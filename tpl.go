package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"
)

func inputToObject(inputStr string, debug *bool) (result interface{}, err error) {
	jsonStr := ""
	lastChar := ""

	if *debug {
		fmt.Fprintf(os.Stderr, "input is: %v\n", inputStr)
	}

	for position, rune := range inputStr {
		currentChar := string(rune)

		isOpeningBrace, _ := regexp.MatchString("[{\\[]", currentChar)
		isColonOrComma, _ := regexp.MatchString("[:,]", currentChar)
		isNotSpecial, _ := regexp.MatchString("[^{}\\[\\]:,]", currentChar)
		lastWasSpecial, _ := regexp.MatchString("[{}\\[\\]:,]", lastChar)
		isClosingBrace, _ := regexp.MatchString("[}\\]]", currentChar)
		lastWasClosingBrace, _ := regexp.MatchString("[^}\\]]", lastChar)

		if position > 0 && isOpeningBrace && !lastWasSpecial {
			jsonStr += "\""
		}

		if isNotSpecial && lastWasSpecial {
			jsonStr += "\""
		}

		if isColonOrComma && lastWasClosingBrace {
			jsonStr += "\""
		}

		if isClosingBrace && !lastWasSpecial {
			jsonStr += "\""
		}

		jsonStr += currentChar
		lastChar = currentChar
	}

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
			fmt.Fprintf(os.Stderr, "result is: %v\n")
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
		fmt.Fprintf(os.Stdout, "version %s\n", "0.4.2")
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

		envKeyValuePair := strings.Split(envVar, "=")
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
	tpl := template.Must(template.New(path.Base(*templateFile)).Funcs(sprig.FuncMap()).ParseFiles(*templateFile))
	err := tpl.Execute(os.Stdout, environment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template %v: %v", *templateFile, err)
		os.Exit(2)
	}
}
