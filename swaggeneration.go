package swaggeneration

import (
	"bytes"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Builder struct {
	manager        *template.Template
	data           *data
	outputFileName string
}

type data struct {
	collection map[string]interface{}
	api        []map[string]interface{}
}

func Init(baseTemplatePath, dataPath, outputFileName string) *Builder {
	return &Builder{
		manager:        loadBaseTemplate(baseTemplatePath),
		data:           loadData(dataPath, outputFileName),
		outputFileName: outputFileName,
	}
}

func (b *Builder) Generate() {
	outputFile, err := os.Create(b.outputFileName)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	var bufferList []*bytes.Buffer

	// generate api collection doc
	var master bytes.Buffer
	err = b.manager.ExecuteTemplate(&master, "_api-collection.yaml", b.data.collection)
	if err != nil {
		panic(err)
	}
	bufferList = append(bufferList, &master)

	// generate api doc
	for _, m := range b.data.api {
		var apiBuffer bytes.Buffer
		err = b.manager.ExecuteTemplate(&apiBuffer, "_api.yaml", m)
		if err != nil {
			panic(err)
		}

		bufferList = append(bufferList, &apiBuffer)
	}

	if err != nil {
		panic(err)
	}

	for _, buffer := range bufferList {
		_, errWrite := outputFile.Write(buffer.Bytes())
		if errWrite != nil {
			panic(errWrite)
		}
	}
}

func loadBaseTemplate(baseTemplatePath string) *template.Template {
	tmp, err := template.New("templates").Funcs(templateFunc()).ParseGlob(filepath.Join(baseTemplatePath, "*.yaml"))
	if err != nil {
		panic(err)
	}

	return template.Must(tmp, err)
}

func loadData(dataPath, outputFileName string) *data {
	fileList, errReadDir := os.ReadDir(dataPath)
	if errReadDir != nil {
		panic(errReadDir)
	}

	d := &data{}

	for _, entry := range fileList {
		if path.Ext(entry.Name()) != ".yaml" {
			continue
		}

		if entry.Name() == "collection.yaml" {
			d.collection = loadYAML(dataPath, entry)
			continue
		}

		if entry.Name() == filepath.Base(outputFileName) {
			continue
		}

		d.api = append(d.api, loadYAML(dataPath, entry))
	}

	return d
}

func loadYAML(basePath string, entry os.DirEntry) map[string]interface{} {
	f, err := os.ReadFile(filepath.Join(basePath, entry.Name()))
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal YAML content into struct
	var data map[string]interface{}
	err = yaml.Unmarshal(f, &data)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	return data
}
