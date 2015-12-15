package main


import (
	"text/template"
	"os"
	"flag"
	"strings"
	"fmt"
)

func main() {

	debug := flag.Bool("d", false, "debug mode")
	listSep := flag.String("ls", ",", "list entry separator")
	mapSep := flag.String("ms", ":", "map key-value separator")
	tpl := flag.String("t", "", "go template file")

	flag.Parse()

	if (len(*tpl) == 0) {
		flag.Usage();
		os.Exit(2)
	}

	if _, err := os.Stat(*tpl); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s not found\n", *tpl)
		os.Exit(2)
	}

	envMap := make(map[string]interface{})

	for _, envVar := range os.Environ() {

		envKeyValuePair := strings.Split(envVar, "=")
		envKey, envValue :=  envKeyValuePair[0], envKeyValuePair[1]

		valueList := strings.Split(envValue, *listSep)
		if len(valueList) > 1 {
			valueMapKeyValuePair := strings.Split(valueList[0], *mapSep)
			if len(valueMapKeyValuePair) > 1 {
				valueMap := make(map[string]string)
				for _, valueMapKeyValueString := range valueList {
					valueMapKeyValuePair := strings.Split(valueMapKeyValueString, *mapSep)
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

	var t = template.Must(template.New(*tpl).ParseFiles(*tpl))
	t.Execute(os.Stdout, envMap)
}
