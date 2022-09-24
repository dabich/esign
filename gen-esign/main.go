// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// gen-esign creates the esign subpackages based upon DocuSign's
// esignature.rest.swagger.json definition file.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/jfcote87/esign"
)

const (
	openAPIdefinition = "https://github.com/docusign/eSign-OpenAPI-Specification"
	generatorVersion  = "20190720"
	pkgBaseName       = "github.com/jfcote87/esign"
)

var (
	definitionFileMap = map[string]esign.APIVersion{
		esign.APIv2.Name():     esign.APIv2,
		esign.APIv21.Name():    esign.APIv2,
		esign.AdminV2.Name():   esign.APIv2,
		esign.ClickV1.Name():   esign.APIv2,
		esign.MonitorV2.Name(): esign.MonitorV2,
		esign.RoomsV2.Name():   esign.APIv2,
	}
	apiParametersMap = map[esign.APIVersion]APIGenerateCfg{
		esign.APIv2: {
			DocPrefix:      "esign-api/v2/",
			CallVersion:    "esign.APIv2",
			ResourceMap:    v2ResourceMap,
			PackagePath:    "v2",
			ModelFile:      "v2/model/model.go",
			ModelPackage:   "model",
			ModelIsPackage: true,
			fldOverrides:   GetFieldOverrides(),
			paramOverrides: GetParameterOverrides(),
		},
		esign.APIv21: {
			DocPrefix:      "esign-api/",
			CallVersion:    "esign.APIv21",
			ResourceMap:    v21ResourceMap,
			PackagePath:    "v2.1",
			ModelFile:      "v2.1/model/model.go",
			ModelPackage:   "model",
			ModelIsPackage: true,
			fldOverrides:   GetFieldOverrides(),
			paramOverrides: GetParameterOverrides(),
		},
		esign.AdminV2: {
			DocPrefix:      "admin-api/",
			CallVersion:    "esign.AdminV2",
			ResourceMap:    adminResourceMap,
			PackagePath:    "admin/",
			ModelFile:      "admin/admin.go",
			ModelPackage:   "admin",
			ModelIsPackage: true,
			fldOverrides:   make(map[string]map[string]string),
			paramOverrides: make(map[string]map[string]string),
		},
		esign.RoomsV2: {
			DocPrefix:      "rooms-api/",
			CallVersion:    "esign.RoomsV2",
			ResourceMap:    roomsResourceMap,
			PackagePath:    "rooms/",
			ModelFile:      "rooms/rooms.go",
			ModelPackage:   "rooms",
			ModelIsPackage: true,
			fldOverrides:   make(map[string]map[string]string),
			paramOverrides: make(map[string]map[string]string),
		},
		esign.ClickV1: {
			DocPrefix:      "click-api/",
			CallVersion:    "esign.ClickV1",
			ResourceMap:    clickResourceMap,
			PackagePath:    "click/",
			ModelFile:      "click/model.go",
			ModelPackage:   "click",
			ModelIsPackage: false,
			fldOverrides:   make(map[string]map[string]string),
			paramOverrides: make(map[string]map[string]string),
		},
		esign.MonitorV2: {
			DocPrefix:      "monitor-api/",
			CallVersion:    "esign.MonitorV2",
			ResourceMap:    monitorResourceMap,
			PackagePath:    "monitor/",
			ModelFile:      "monitor/model.go",
			ModelPackage:   "monitor",
			ModelIsPackage: false,
			fldOverrides:   make(map[string]map[string]string),
			paramOverrides: make(map[string]map[string]string),
		},
	}

	baseDir     = flag.String("src", "../.", "source directory")
	templDir    = flag.String("templates", "gen-esign/templates", "directory containing output templates.")
	specsFolder = flag.String("swaggerfiles", "gen-esign/specs", "directory containing swagger specification files")
	skipFormat  = flag.Bool("skip_format", false, "skip gofmt command")
)

// APIGenerateCfg contains parameters for generating an eSignature version
type APIGenerateCfg struct {
	esign.APIVersion
	Templates      *template.Template // templates
	BaseDir        string             // source directory
	BasePkg        string
	SkipFormat     bool
	Name           string
	Version        string
	DocPrefix      string
	CallVersion    string
	PackagePath    string
	ModelFile      string
	ModelPackage   string
	ModelIsPackage bool
	ResourceMap    map[string]string
	fldOverrides   map[string]map[string]string
	paramOverrides map[string]map[string]string
}

func main() {
	flag.Parse()
	log.Printf("baseDir: %s", *baseDir)
	log.Printf("tmpl: %s", *templDir)
	log.Printf("specs: %s", *specsFolder)
	curdir, _ := os.Getwd()
	log.Printf("path: %s", curdir)
	pkgBaseDir, pkgSwaggerDir, pkgTemplates, skipFormatting := *baseDir, *specsFolder, *templDir, *skipFormat
	genTemplates, err := template.ParseFiles(path.Join(pkgTemplates, "service.tmpl"), path.Join(pkgTemplates, "model.tmpl"))
	if err != nil {
		log.Fatalf("Templates: %v", err)
	}
	docmap, err := decodeSwaggerDocs(pkgSwaggerDir)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := os.Chdir(pkgBaseDir); err != nil {
		log.Fatalf("unable to set directory to %s: %v", pkgBaseDir, err)
	}
	if pkgBaseDir, err = os.Getwd(); err != nil {
		log.Fatalf("unable to retrieve working diretory: %v", err)
	}

	for v, doc := range docmap {
		cfg, ok := apiParametersMap[v]
		if !ok {
			log.Printf("skipping %s has no parameters entry", v.Name())
			continue
		}
		cfg.APIVersion = v
		cfg.Name = v.Name()
		cfg.Version = doc.Info.Version
		cfg.BaseDir = pkgBaseDir
		cfg.BasePkg = pkgBaseName
		cfg.Templates = genTemplates
		cfg.SkipFormat = skipFormatting
		if err := cfg.genVersion(&doc); err != nil {
			log.Printf("%s %v", cfg.Name, err)
			return
		}
	}
}

func decodeSwaggerDocs(folderName string) (map[esign.APIVersion]Document, error) {
	fis, err := ioutil.ReadDir(folderName)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if !strings.HasSuffix(folderName, "/") {
		folderName += "/"
	}
	var results = make(map[esign.APIVersion]Document)
	for _, f := range fis {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".json") {
			continue
		}
		b, err := ioutil.ReadFile(folderName + f.Name())
		if err != nil {
			return nil, err
		}
		var doc *Document
		if err = json.Unmarshal(b, &doc); err != nil {
			return nil, fmt.Errorf("%s decode %w", f.Name(), err)
		}
		apikey := doc.Info.Title + ":" + doc.Info.Version
		apiVersion, ok := definitionFileMap[apikey]
		if !ok {
			return nil, fmt.Errorf("no matching api version for %s", apikey)
		}
		results[apiVersion] = *doc
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no definition files specified in %s", folderName)
	}
	return results, nil
}

func (api *APIGenerateCfg) genVersion(doc *Document) error {
	// Put the Definitions (structs) in order
	sort.Sort(doc.Definitions)

	defMap := make(map[string]Definition)
	structName := ""
	defList := make([]Definition, 0, len(doc.Definitions))
	// create defMap for lookup later field and parameter
	// lookups.  Make certain defList has only unique names.
	for _, def := range doc.Definitions {
		defMap["#/definitions/"+def.ID] = def
		if structName != def.Name {
			defList = append(defList, def)
		}
		structName = def.Name
	}

	// generate model.go first
	if err := api.doModel(defList, defMap); err != nil {
		return fmt.Errorf("%v Generating model.go failed: %v", api.Version, err)
	}

	sort.Sort(doc.Operations)
	ops := make(map[string][]Operation, 0)
	for _, op := range doc.Operations {
		if op.Status == "restricted" {
			log.Printf("Skipping: %s %s", op.Status, op.OperationID)
			continue
		}
		if op.Service == "" {
			log.Printf("No service specified: %s", op.OperationID)
			continue
		}
		fullService := api.Version + ":" + op.Service
		if newServiceName, ok := ServiceNameOverride[fullService]; ok {
			op.Service = newServiceName
		}
		fullOpName := api.Version + ":" + op.OperationID
		if !OperationSkipList[fullOpName] {
			serviceName, ok := api.ResourceMap[op.Tags[0]]
			if ok {
				op.Service = serviceName
			} else {
				log.Printf("No service found for %s", op.Tags[0])
			}
			ops[op.Service] = append(ops[op.Service], op)
		}
	}
	tagDescMap := make(map[string]string)
	for _, tag := range doc.DSTags {
		tagDescMap[tag.Name] = tag.Description
	}

	serviceTmpl := api.Templates.Lookup("service.tmpl")
	for k, v := range ops {
		log.Printf("Generating %s", k)
		descrip, _ := tagDescMap[k]

		if err := api.doPackage(serviceTmpl, k, descrip, v, defMap); err != nil {
			return fmt.Errorf("%s generate %s.go failed: %v", api.Version, k, err)
		}
	}
	return nil
}

// getDocument loads the swagger def file and applies overrides
func getDocument(fn string) (*Document, error) {
	// Open swagger file and parse
	f, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("Unable to open: %v", err)
	}
	defer f.Close()
	var doc *Document
	if err = json.NewDecoder(f).Decode(&doc); err != nil {
		return nil, fmt.Errorf("Unable to parse: %v", err)
	}
	var opAdditions OpList
	// Add additional operations from overrides package
	for i, op := range doc.Operations {
		if strings.Contains(op.Description, "**Deprecated**") {
			doc.Operations[i].Deprecated = true
		}
		// add media download when empty get response
		if op.HTTPMethod == "GET" && op.Responses["200"].Schema == nil {
			newResp := op.Responses["200"]
			newResp.Schema = &SchemaRef{
				Type: "file",
			}
			doc.Operations[i].Responses["200"] = newResp
		}
	}
	doc.Operations = append(doc.Operations, opAdditions...)

	return doc, nil
}

func (op *Operation) hasJSONResponse() bool {
	if okResponse, ok := op.Responses["200"]; ok {
		return okResponse.Schema != nil && okResponse.Schema.Ref > ""
	}
	return false
}

// doModel generates the model.go in the model package
func (api *APIGenerateCfg) doModel(defList []Definition, defMap map[string]Definition) error {
	modelTempl := api.Templates.Lookup("model.tmpl")
	// create model.go
	// get field overrides and tab overrides
	tabDefs := TabDefs(api.Name, defMap, api.fldOverrides)

	var data = struct {
		Definitions  []Definition
		DefMap       map[string]Definition
		FldOverrides map[string]map[string]string
		CustomCode   string
		DocPrefix    string
		VersionID    string
		IsPackage    bool
	}{
		Definitions:  append(tabDefs, defList...), // Prepend tab definitions
		DefMap:       defMap,
		FldOverrides: api.fldOverrides,
		CustomCode:   CustomCode,
		DocPrefix:    api.DocPrefix,
		VersionID:    api.Version,
		IsPackage:    api.ModelIsPackage,
	}
	modelBuffer := &bytes.Buffer{}
	if err := modelTempl.Execute(modelBuffer, data); err != nil {
		return err
	}

	if *skipFormat {
		return api.makePackageFile("model", modelBuffer.Bytes())
	}
	fmtBytes, err := format.Source(modelBuffer.Bytes())
	if err != nil {
		log.Printf("Source Error: %v", err)
		return err
	}
	return api.makePackageFile("model", fmtBytes)
}

// ExtOperation contains all needed info
// for the template merge
type ExtOperation struct {
	Operation
	OpPayload         *Payload
	HasUploads        bool
	IsMediaUpload     bool
	PathParams        []PathParam
	FuncName          string
	QueryOptions      []QueryOpt
	Result            string
	DownloadAdditions []DownloadAddition
	JSONResponse      bool
}

// doPackage creates a subpackage go file
func (api *APIGenerateCfg) doPackage(resTempl *template.Template, serviceName string, description string, ops []Operation, defMap map[string]Definition) error {
	packageName := strings.ToLower(serviceName)
	comments := strings.Split(strings.TrimRight(description, "\n"), "\n")
	if packageName == "uncategorized" {
		comments = append(comments, "Uncategorized calls may change or move to other packages.")
	}

	extOps := make([]ExtOperation, 0, len(ops))
	for _, op := range ops {
		payload := op.Payload(defMap, api.ModelPackage)
		extOps = append(extOps, ExtOperation{
			Operation:         op,
			OpPayload:         payload,
			HasUploads:        IsUploadFilesOperation(api.Version + ":" + op.OperationID),
			IsMediaUpload:     payload != nil && payload.Type == "*esign.UploadFile",
			PathParams:        op.PathParameters(),
			FuncName:          op.GoFuncName(GetServicePrefixes(op.Service)),
			QueryOptions:      op.QueryOpts(api.paramOverrides),
			Result:            op.Result(defMap, api.ModelPackage),
			DownloadAdditions: GetDownloadAdditions(api.Version + ":" + op.OperationID),
			JSONResponse:      op.hasJSONResponse(),
		})
	}
	var data = struct {
		Service      string
		Package      string
		Directory    string
		Operations   []ExtOperation
		Comments     []string
		Packages     []string
		PackagePath  string
		ModelPackage string
		DocPrefix    string
		VersionID    string
		CallVersion  string
		AddDocLinks  bool
		Accept       string
	}{
		Service:      serviceName,
		Package:      packageName,
		Directory:    api.BasePkg,
		Operations:   extOps,
		Comments:     comments,
		Packages:     []string{`"context"`, `"net/url"`},
		PackagePath:  api.PackagePath,
		ModelPackage: api.ModelPackage,
		DocPrefix:    api.DocPrefix,
		VersionID:    api.Version,
		CallVersion:  api.CallVersion,
		AddDocLinks:  (serviceName != "Uncategorized"),
	}
	importMap := make(map[string]bool)
	for _, op := range extOps {
		if len(op.PathParameters()) > 0 {
			importMap[`"strings"`] = true
			break
		}
	}
	for _, o := range extOps {
		for _, q := range o.QueryOptions {
			switch q.Type {
			case "...string":
				importMap[`"strings"`] = true
			case "int":
				importMap[`"fmt"`] = true
			case "float64":
				importMap[`"fmt"`] = true
			case "time.Time":
				importMap[`"time"`] = true
			}
		}
		if o.IsMediaUpload {
			importMap[`"io"`] = true
		}
	}

	for k, v := range importMap {
		if v {
			data.Packages = append(data.Packages, k)
		}
	}
	sort.Strings(data.Packages)
	data.Packages = append(data.Packages,
		"",
		"\""+api.BasePkg+"\"")
	if api.ModelPackage > "" {
		data.Packages = append(data.Packages,
			"\""+api.BasePkg+"/"+api.PackagePath+api.ModelPackage+"\"")
	}

	pkgBuffer := &bytes.Buffer{}
	if err := resTempl.Execute(pkgBuffer, data); err != nil {
		return err
	}
	if !*skipFormat {
		pkgBytes, err := format.Source(pkgBuffer.Bytes())
		if err == nil {
			pkgBuffer = bytes.NewBuffer(pkgBytes)
		}
	}
	return api.makePackageFile(packageName, pkgBuffer.Bytes())

}

func (api *APIGenerateCfg) getEsignDir() string {
	p := path.Join(api.BaseDir, api.Version)
	if strings.HasPrefix(p, api.Version) {
		p = "./" + p
	}
	return p
}

func (api *APIGenerateCfg) pkgdir(packageName string) (string, string) {
	//if packageName == "Uncategorized" && api.UncategorizedToTop {
	//	return api.getEsignDir(), api.Version
	//}
	return api.getEsignDir() + "/" + packageName, packageName
}

func (api *APIGenerateCfg) makePackageFile(packageName string, content []byte) error {
	pkgDir, fileName := api.pkgdir(packageName)
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(pkgDir+"/"+fileName+".go", content, 0755)
}
