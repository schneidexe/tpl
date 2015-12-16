package main


import (
	"text/template"
	"os"
	"flag"
	"strings"
	"fmt"
	"encoding/json"
	"io"
//	"log"
	"regexp"
	"errors"
)

func incr(val int) int {
	return val + 1
}

// add quotes to json
func createValidJson(shellson string) (data interface{}, err error) {
	validJson := ""
	last := ""

	for position, rune := range shellson {
		current := string(rune)

		isOpeningBrace, _ := regexp.MatchString("[{\\[]", current)
		isColonOrComma, _ := regexp.MatchString("[:,]", current)
		isNotSpecial, _ := regexp.MatchString("[^{\\[:,]", current)
		lastWasSpecial, _ := regexp.MatchString("[{\\[:,]", last)
		isClosingBrace, _ := regexp.MatchString("[}\\]]", current)
		lastWasClosingBrace, _ := regexp.MatchString("[^}\\]]", last)

		if position > 0 && isOpeningBrace && !lastWasSpecial {
			validJson += "\""
		}

		if isNotSpecial && lastWasSpecial {
			validJson += "\""
		}

		if isColonOrComma && lastWasClosingBrace {
			validJson += "\""
		}

		if (isClosingBrace && lastWasClosingBrace) {
			validJson += "\""
		}

		validJson += current
		last = current
	}

	decoder := json.NewDecoder(strings.NewReader(validJson))
	for {
		var data interface{}
		if err := decoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			return shellson, errors.New("json conversion validation failed for " + validJson)
		}
		return data, nil
	}

	return shellson, errors.New("json conversion validation failed for " + validJson)
}

func main() {

	// set and parse cmd line flags
	debug := flag.Bool("d", false, "debug mode")
	templateFile := flag.String("t", "", "go template file")

	flag.Parse()

	if (len(*templateFile) == 0) {
		flag.Usage();
		os.Exit(2)
	}

	if _, err := os.Stat(*templateFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s not found\n", *templateFile)
		os.Exit(2)
	}

	// generate data map
	environment := make(map[string]interface{})
	for _, envVar := range os.Environ() {

		envKeyValuePair := strings.Split(envVar, "=")
		envKey, envValue :=  envKeyValuePair[0], envKeyValuePair[1]

		data, err := createValidJson(envValue);
		if (err != nil) {
			environment[envKey] = envValue;
		} else {
			environment[envKey] = data;
		}
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "environment map is: %v\n", environment)
	}

	// render template
	tmpl := template.Must(template.ParseGlob(*templateFile))
	err := tmpl.Execute(os.Stdout, environment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during template execution: %s", err)
		os.Exit(1)
	}

}
