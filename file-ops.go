package main

import (
	"fmt"
	"os"
	"strings"
)

var BASEPATH = os.Getenv("temp") + "\\flexins\\"

func CreateFile(name string, payload []byte) {
	fi, err := os.Create(os.Getenv("temp") + "\\flexins\\" + name)
	if err != nil {
		fmt.Printf("Error while creating file on disk: %s\n", err)
		os.Exit(1)
	}
	defer fi.Close()

	_, errX := fi.Write(payload)
	if errX != nil {
		fmt.Printf("Error writing file to disk: %s\n", errX)
		os.Exit(1)
	}
	fi.Close()
}

func createSubDir(name string) {
	err := os.Mkdir(os.Getenv("temp")+"\\flexins\\"+name, 0755)
	if err != nil && !strings.Contains(err.Error(), "file already exists") {
		panic(err)
	}
}

func CreateTemp() {
	fmt.Println(os.Getenv("temp"))
	err := os.Mkdir(os.Getenv("temp")+"\\flexins", 0755)
	if err != nil && !strings.Contains(err.Error(), "file already exists") {
		panic(err)
	}
}
