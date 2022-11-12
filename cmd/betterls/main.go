package main

import (
	"log"
	"os"
)

func main(){
	files, err:= os.ReadDir(".");

	if err != nil {
		log.Fatal(err)
	}
	
	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	
}