package main

import (
	"flag"
	"fmt"
	"gwhois/lib"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

var shouldExportResponse bool
var exportFileName string

func init() {
	flag.BoolVar(&shouldExportResponse, "save", false, "Save the API response to disk")
	flag.StringVar(&exportFileName, "filename", "api_response.json", "The filename the API response should be saved as")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Google Whois 0.0.1 by Doctor Chaos\n\nUsage of %s: <document URL>\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "Error: No URL has been provided. Please provide a URL")
		os.Exit(1)
	}

	parsedUrl, err := url.Parse(flag.Arg(0))

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	if parsedUrl.Path == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Error: Invalid URL (URL does not contain path)")
		os.Exit(1)
	}

	var extractedIdentifier string
	for _, v := range strings.Split(parsedUrl.Path, "/") {
		if len(v) != 44 {
			continue
		}
		extractedIdentifier = v
	}

	if extractedIdentifier == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Error: Invalid URL (URL does not contain identifier)")
		os.Exit(1)
	}

	info, infoResponse, err := lib.GetFileInformation(extractedIdentifier)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to retrieve document information from Google! (%s)\n", err)
		os.Exit(1)
	}

	if shouldExportResponse {
		err := ioutil.WriteFile(exportFileName, infoResponse, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Unable to save response! (%s)\n", err)
		}
	}

	fmt.Fprintf(os.Stdout, "Document ID: %s\nDocument created at: %s\nDocument last modified at: %s\nDocument title: \"%s\"\n", extractedIdentifier, info.Createddate, info.Modifieddate, info.Title)

	for _, v := range info.Owners {
		fmt.Fprintf(os.Stdout, "Document Owner: %s (%s) (%s)\n", v.Displayname, v.Emailaddress, v.ID)
	}

	fmt.Fprintf(os.Stdout, "Last modified by: %s\nPublicly Editable: %t\n", info.Lastmodifyingusername, info.Editable)
}
