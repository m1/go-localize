package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	jsonFileExt = ".json"
	yamlFileExt = ".yaml"
)

type localizationFile map[string]string

const (
	defaultOutputDir = "localizations"
)

var (
	input  = flag.String("input", "", "input localizations folder")
	output = flag.String("output", "", "where to output the generated package")

	errFlagInputNotSet = errors.New("the flag -input must be set")
)

func main() {
	flag.Parse()

	if err := run(input, output); err != nil {
		log.Fatal(err.Error())
	}
}

func run(in, out *string) error {
	inputDir, outputDir, err := parseFlags(in, out)
	if err != nil {
		return err
	}

	files, err := getLocalizationFiles(inputDir)
	if err != nil {
		return err
	}

	localizations, err := generateLocalizations(files)
	if err != nil {
		return err
	}

	return generateFile(outputDir, localizations)
}

func generateLocalizations(files []string) (map[string]string, error) {
	localizations := map[string]string{}
	for _, file := range files {
		newLocalizations, err := getLocalizationsFromFile(file)
		if err != nil {
			return nil, err
		}
		for key, value := range newLocalizations {
			localizations[key] = value
		}
	}
	return localizations, nil
}

func getLocalizationFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if !info.IsDir() && (ext == jsonFileExt || ext == yamlFileExt) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func generateFile(output string, localizations map[string]string) error {
	dir := output
	parent := output
	if strings.Contains(output, string(filepath.Separator)) {
		parent = filepath.Base(dir)
	}

	err := os.MkdirAll(output, 0700)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("%v/%v.go", dir, parent))
	if err != nil {
		return err
	}

	return packageTemplate.Execute(f, struct {
		Timestamp     time.Time
		Localizations map[string]string
		Package       string
	}{
		Timestamp:     time.Now(),
		Localizations: localizations,
		Package:       parent,
	})
}

func getLocalizationsFromFile(file string) (map[string]string, error) {
	newLocalizations := map[string]string{}

	openFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(openFile)
	if err != nil {
		return nil, err
	}

	localizationFile := localizationFile{}
	ext := filepath.Ext(file)
	if ext == jsonFileExt {
		err = json.Unmarshal(byteValue, &localizationFile)
	} else {
		err = yaml.Unmarshal(byteValue, &localizationFile)
	}

	if err != nil {
		return nil, err
	}

	slicePath := getSlicePath(file)
	for key, value := range localizationFile {
		newLocalizations[strings.Join(append(slicePath, key), ".")] = value
	}

	return newLocalizations, nil
}

func getSlicePath(file string) []string {
	dir, file := filepath.Split(file)

	paths := strings.Replace(dir, *input, "", -1)
	pathSlice := strings.Split(paths, string(filepath.Separator))

	var strs []string
	for _, part := range pathSlice {
		part := strings.TrimSpace(part)
		part = strings.Trim(part, "/")
		if part != "" {
			strs = append(strs, part)
		}
	}

	strs = append(strs, strings.Replace(file, filepath.Ext(file), "", -1))
	return strs
}

func parseFlags(input *string, output *string) (string, string, error) {
	var inputDir, outputDir string

	if *input == "" {
		return "", "", errFlagInputNotSet
	}
	if *output == "" {
		outputDir = defaultOutputDir
	} else {
		outputDir = *output
	}

	inputDir = *input

	return inputDir, outputDir, nil
}
