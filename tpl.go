package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var environment = make(map[string]interface{})
var templateFile *string

// add custom functions
var customFuctions = template.FuncMap{
	"include":     include,
	"mustInclude": mustInclude,
}

func inputToObject(inputStr string, debug *bool) (result interface{}, err error) {
	if *debug {
		fmt.Fprintf(os.Stderr, "----\ninput is: %v\n", inputStr)
	}

	// try to parse a plain json first
	jsonStr := inputStr
	err = json.Unmarshal([]byte(jsonStr), &result)

	// now try to enrich unquoted json
	if err != nil {
		// insert " after , if next is none of [ { "
		jsonStr = regexp.MustCompile(`,([^[{"])`).ReplaceAllString(jsonStr, ",\"$1")
		// insert " before , if previous is none of ] } "
		jsonStr = regexp.MustCompile(`([^]}"]),`).ReplaceAllString(jsonStr, "$1\",")
		// insert " after [ { if next is none of ] [ } { , "
		jsonStr = regexp.MustCompile(`([\[{])([^][}{,"])`).ReplaceAllString(jsonStr, "$1\"$2")
		// insert " before ] } if previous is none of ] [ } { , "
		jsonStr = regexp.MustCompile(`([^][}{,"])([\]}])`).ReplaceAllString(jsonStr, "$1\"$2")
		// insert " after : if next is none of : [ { "
		jsonStr = regexp.MustCompile(`([^:]):([^:[{"])`).ReplaceAllString(jsonStr, "$1:\"$2")
		// insert " before : if previous is not :
		jsonStr = regexp.MustCompile(`([^:"]):([^:])`).ReplaceAllString(jsonStr, "$1\":$2")
		// replace :: with : (double colons can be used to escape a colon)
		jsonStr = regexp.MustCompile(`::`).ReplaceAllString(jsonStr, ":")
	}
	if *debug {
		fmt.Fprintf(os.Stderr, "json is: %v\n", jsonStr)
	}

	// try parsing json again, if it fails fall back to the plain input value
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil || result == nil || reflect.TypeOf(result).Kind() == reflect.Float64 {
		result = inputStr
	}

	if *debug {
		if err != nil {
			fmt.Fprintf(os.Stderr, "result is: %v, error: %v\n----\n", result, err)
		} else {
			fmt.Fprintf(os.Stderr, "result is: %v\n----\n", result)
		}
	}

	return result, err
}

func renderInclude(fileName string, safeMode bool) string {
	// lookup relative file names in same directory like main template
	lookupDir := ""
	if !strings.HasPrefix(fileName, "/") && !strings.HasPrefix(fileName, "/") {
		lookupDir = path.Dir(*templateFile)
	}

	// ignore non-existing files
	if safeMode {
		if _, err := os.Stat(path.Join(lookupDir, fileName)); os.IsNotExist(err) {
			return ""
		}
	}

	tpl := template.Must(template.New(path.Base(fileName)).Funcs(sprig.TxtFuncMap()).ParseFiles(path.Join(lookupDir, fileName)))

	var result bytes.Buffer
	tpl.Execute(&result, environment)
	return result.String()
}

func include(fileName string) string {
	return renderInclude(fileName, true)
}

func mustInclude(fileName string) string {
	return renderInclude(fileName, false)
}

func main() {
	// set and parse cmd line flags
	debug := flag.Bool("d", false, "enable debug mode")
	prefix := flag.String("p", "", "only consider variables starting with prefix")
	templateFile = flag.String("t", "", "template file")
	version := flag.Bool("v", false, "show version")
	outputFile := flag.String("o", "", "output file")

	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "version %s\n", "0.7.0")
		os.Exit(0)
	}

	if len(*templateFile) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(*templateFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s not found\n", *templateFile)
		os.Exit(2)
	}

	// generate environment map
	for _, envVar := range os.Environ() {
		envKeyValuePair := strings.SplitN(envVar, "=", 2)
		envKey, envValue := envKeyValuePair[0], envKeyValuePair[1]

		if !strings.HasPrefix(envKey, *prefix) {
			continue
		}

		data, err := inputToObject(envValue, debug)
		if err != nil {
			environment[envKey] = envValue
		} else {
			environment[envKey] = data
		}
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "environment map is: %v\n", environment)
	}

	outputWriter := os.Stdout
	if len(*outputFile) > 0 {
		// Create file and truncate it if it already exists
		out, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file: %s", err)
			return
		}
		outputWriter = out
	}

	// render template
	tpl := template.Must(template.New(path.Base(*templateFile)).Funcs(sprig.TxtFuncMap()).Funcs(customFuctions).ParseFiles(*templateFile))
	err := tpl.Execute(outputWriter, environment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template %v: %v\n", *templateFile, err)
		os.Exit(2)
	}
}
