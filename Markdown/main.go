package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8">
			<title>Markdown Preview Tool</title>
		</head>
		<body>
	`
	footer = `
		</body>
	</html>
	`
)

func main() {
	// parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	// if user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	// read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	// create temporary file and check for errors
	temp, err := ioutil.TempFile("", "mdp*.html")
	if err != nil {
		return err
	}

	outName := temp.Name()

	fmt.Fprintln(out, outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {

	// parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	// write the byts to the file
	return ioutil.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	// define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	// apppend filename to parameters slice
	cParams = append(cParams, fname)

	// locate executable in PATH
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	// open the file using default program
	return exec.Command(cPath, cParams...).Run()
}
