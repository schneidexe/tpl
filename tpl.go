package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Masterminds/sprig"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

var (
	// BuildVersion is used to pass version during build
    BuildVersion string = ""
)

func inputToObject(inputStr string, debug *bool) (result interface{}, err error) {
	jsonStr := inputStr
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

	if *debug {
		fmt.Fprintf(os.Stderr, "json is: %v\n", jsonStr)
	}

	err = json.Unmarshal([]byte(jsonStr), &result)

	if *debug {
		fmt.Fprintf(os.Stderr, "result is: %v, error: %v\n", result, err)
	}

	if err != nil || result == nil || reflect.TypeOf(result).Kind() == reflect.Float64 {
		result = inputStr
		if *debug {
			fmt.Fprintf(os.Stderr, "result is: %v\n", result)
		}
	}

	return result, err
}

func main() {
	// set and parse cmd line flags
	debug := flag.Bool("d", false, "enable debug mode")
	templateFile := flag.String("t", "", "template file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "version %s\n", BuildVersion)
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
	environment := make(map[string]interface{})
	for _, envVar := range os.Environ() {
		envKeyValuePair := strings.SplitN(envVar, "=", 2)
		envKey, envValue := envKeyValuePair[0], envKeyValuePair[1]

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

	// render template
	tpl := template.Must(template.New(path.Base(*templateFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(*templateFile))
	err := tpl.Execute(os.Stdout, environment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template %v: %v", *templateFile, err)
		os.Exit(2)
	}
}
