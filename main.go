package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	filename := "output.txt"
	_,err := os.Stat(filename)
	if os.IsExist(err) {
		fmt.Println(err)
		os.Remove(filename)
		
	}

	var path string
	_, error := fmt.Scanln(&path)

	if error != nil {
		log.Fatal(error)
	}
	os.Create(filename)
	filepath.Walk(path, VisitFile)
}
