package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// function returns 3 params, input and output file
	// and err is usage message that will be display on terminal
	inputFileName, outputFileName, err := getFilenamesFromTerminal()

	// we dont use log.Fatal here because we want to show usage message
	// log.fatal display current date, time and err message
	// os.Exit will terminat the programe and returns 1 to show
	// user or application fault
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read and write file from and to the standard input and output
	inFile, outFile := os.Stdin, os.Stdout

	// read file and close it using defer statement
	// defer will make sure that file will be closed
	// once reading and writing is done and when the main function returns
	if inputFileName != "" {
		if inFile, err = os.Open(inputFileName); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	// create file for writing. log.Fatal will print error message
	// and date and time on terminal. It will call osExit with status
	// code 1 to show programmer or code fault.
	if outputFileName != "" {
		if outFile, err = os.Create(outputFileName); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

	fmt.Println("Reading data from ", os.Args[1], " file")

	if err := writeContent(inFile, outFile); err != nil {
		log.Fatal(err)
	}

}

func getFilenamesFromTerminal() (iFileName, oFileName string, err error) {

	// if use aske for help or only type input file show them error message
	// on terminal and returns input and output file as a empty
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err := fmt.Errorf("usage: %s infile.ext outfile.ext", filepath.Base(os.Args[0]))
		return "", "", err
	}

	// if both the input and output files are provided
	// store them in variable to return
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

func writeContent(inFile io.Reader, outFile io.Writer) (err error) {

	// create buffer reader and writer through which
	// there content can be accessed as a byte(or string)
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	// create anonymous defer function that will be called just before
	// writeContent function will return or panic occurs to make sure
	// that the writer buffer is flushed
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	eof := false

	// create infinite loop to read the content of buffer reader
	// set the eof to true when all the content of file is done reading (io.EOF)
	// or err occurs to break the infinite loop
	for !eof {
		var line string

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}

		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}

	fmt.Println("File has been written into ", os.Args[2])

	return nil

}
