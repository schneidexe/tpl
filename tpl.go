package main

import (
	"text/template"
	"os"
	"flag"
	"strings"
	"fmt"
	"encoding/json"
	"regexp"
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
		isNotSpecial, _ := regexp.MatchString("[^{\\[:,]", currentChar)
		lastWasSpecial, _ := regexp.MatchString("[{\\[:,]", lastChar)
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

		if (isClosingBrace && lastWasClosingBrace) {
			jsonStr += "\""
		}

		jsonStr += currentChar
		lastChar = currentChar
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "json is: %v\n", jsonStr)
	}

	err = json.Unmarshal([]byte(jsonStr), &result);
	if err != nil {
		result = inputStr
		if *debug {
			fmt.Fprintf(os.Stderr, "result is: %v (error: %v)\n", result, err)
		}
	} else if *debug {
		fmt.Fprintf(os.Stderr, "result is: %v\n", result)
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
		fmt.Fprintf(os.Stdout, "version %s\n", "0.4-alpha")
		os.Exit(0)
	}

	if len(*templateFile) == 0 {
		flag.Usage();
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
		envKey, envValue :=  envKeyValuePair[0], envKeyValuePair[1]

		data, err := inputToObject(envValue, debug)
		if (err != nil) {
			environment[envKey] = envValue
		} else {
			environment[envKey] = data
		}
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "environment map is: %v\n", environment)
	}

	// render template
	tpl := template.Must(template.ParseGlob(*templateFile))
	err := tpl.Execute(os.Stdout, environment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template %v: %v", *templateFile, err)
		os.Exit(2)
	}

}
