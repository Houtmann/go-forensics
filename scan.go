package main

import (
	"os"
)

func VisitFile(fp string, fi os.FileInfo, err error) error {

	filename := "output.txt"

	if err != nil {
		// fmt.Println(err) // can't walk here,
		return nil // but continue walking elsewhere
	}
	if !!fi.IsDir() {
		// fmt.Println("directory:"+fp)
		return nil // not a file.
	}
	//fmt.Println("file:"+fp)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)

	f.WriteString("\n")
	f.WriteString(fp)
	f.Close()

	return nil

}
