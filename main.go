package main

import (
	"encoding/json"
	"flag"
	"os"
)

// Version 9e6ded44a17c7d5e04ae4c536eab5350f83c2bd8

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "instance_settings.json", "path to a TBot Settings file.")

	flag.Parse()

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var settings InstanceSetting
	err = json.Unmarshal(fileBytes, &settings)
	if err != nil {
		panic(err)
	}
	fileBytes, err = json.MarshalIndent(settings, "", "   ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		panic(err)
	}
}
