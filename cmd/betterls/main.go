package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	
	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}
	sort.Strings(fileNames)


}