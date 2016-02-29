package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	inputFileName, outputFileName, err := getFilenamesFromTerminal()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout

	if inputFileName != "" {
		if inFile, err = os.Open(inputFileName); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	if outputFileName != "" {
		if outFile, err = os.Create(outputFileName); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

}

func getFilenamesFromTerminal() (iFileName, oFileName string, err error) {

	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err := fmt.Errorf("usage: %s infile.ext outfile.ext", filepath.Base(os.Args[0]))
		return "", "", err
	}

	if len(os.Args) > 1 {
		iFileName = os.Args[1]
		if len(os.Args) > 2 {
			oFileName = os.Args[2]
		}
	}

	if iFileName != "" && (iFileName == oFileName) {
		log.Fatal("Can not overwrite input file")
	}

	// Empty return because we already specify variable name and there type
	// in function return syntax (iFilename, oFileName string, err error)
	// Do not use this syntax because its called poor go style
	return
}
