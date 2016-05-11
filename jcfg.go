package jcfg

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func completeMissingFields(cfg interface{}) {
	typ := reflect.TypeOf(cfg)

	elemTyp := typ.Elem()
	elemVal := reflect.ValueOf(cfg).Elem()
	for i := 0; i < elemVal.NumField(); i++ {
		if elemVal.Field(i).String() == "" {
			fmt.Printf("Missing field - %s: ", elemTyp.Field(i).Name)
			var newFieldValue string
			fmt.Scanln(&newFieldValue)
			elemVal.Field(i).SetString(newFieldValue)
		}
	}
}

// Parse the given file
func Parse(filePath string, cfg interface{}) interface{} {
	// Check if the file doesn't exist yet
	newMode := false
	if _, err := os.Stat(filePath); err != nil {
		newMode = true
	}

	// open the file
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file %s, got: %s\n", filePath, err))
	}

	// if the file is new, there's no use in parsing it, as it's not valid json
	if !newMode {
		// parse the file
		if err = json.NewDecoder(f).Decode(cfg); err != nil {
			// Failed to parse json, ask user whether to continue
			fmt.Print("Failed to parse configuration file, continue [y/N]? ")
			var x string
			fmt.Scanln(&x)
			if x != "y" {
				panic(fmt.Sprintln("Stopping..."))
			}
		}
	}

	// fill empty fields
	completeMissingFields(cfg)

	// save the new file
	f.Seek(0, 0)
	if err = json.NewEncoder(f).Encode(cfg); err != nil {
		panic(fmt.Sprintf("Failed to save file %s, got: %s\n", filePath, err))
	}

	return cfg
}
