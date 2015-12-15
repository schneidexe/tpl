package main


import (
	"text/template"
	"os"
	"flag"
	"strings"
	"fmt"
)

func incr(val int) int {
	return val+1
}

func main() {

	debug := flag.Bool("d", false, "debug mode")
	listSeparator := flag.String("ls", ",", "list entry separator")
	mapSeparator := flag.String("ms", ":", "map key-value separator")
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

	envMap := make(map[string]interface{})

	for _, envVar := range os.Environ() {

		envKeyValuePair := strings.Split(envVar, "=")
		envKey, envValue :=  envKeyValuePair[0], envKeyValuePair[1]

		valueList := strings.Split(envValue, *listSeparator)
		if len(valueList) > 1 {
			valueMapKeyValuePair := strings.Split(valueList[0], *mapSeparator)
			if len(valueMapKeyValuePair) > 1 {
				valueMap := make(map[string]string)
				for _, valueMapKeyValueString := range valueList {
					valueMapKeyValuePair := strings.Split(valueMapKeyValueString, *mapSeparator)
					mapKey, mapValue :=  valueMapKeyValuePair[0], valueMapKeyValuePair[1]
					valueMap[mapKey] = mapValue
				}
				envMap[envKey] = valueMap;
				if *debug {
					fmt.Printf("%s is a map: %s\n", envKey, valueMap);
				}
			} else {
				envMap[envKey] = valueList;
				if *debug {
					fmt.Printf("%s is a list: %s\n", envKey, valueList);
				}
			}
		} else {
			envMap[envKey] = envValue;
			if *debug {
				fmt.Printf("%s is a simple string: %s\n", envKey, envValue);
			}
		}
	}

	if *debug {
		fmt.Printf("environment map is: %v\n", envMap)
	}

	tmpl := template.Must(template.ParseGlob(*templateFile))
	err := tmpl.Execute(os.Stdout, envMap)
	if err != nil {
		fmt.Printf("error during template execution: %s", err);
		os.Exit(1)
	}

}
