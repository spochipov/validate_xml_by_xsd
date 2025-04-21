// Package main provides a simple XML validator against XSD schemas
package main

// CGO_CFLAGS and CGO_LDFLAGS are required for libxml2
// To build this application, you need to set the following environment variables:
// export CGO_CFLAGS="-I/opt/homebrew/opt/libxml2/include/libxml2"
// export CGO_LDFLAGS="-L/opt/homebrew/opt/libxml2/lib -lxml2"

import (
	"flag"
	"fmt"
	"os"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

func main() {
	// Define command line flags
	xmlFile := flag.String("xml", "", "Path to XML file (required)")
	xsdFile := flag.String("xsd", "", "Path to XSD schema file (required)")
	flag.Parse()

	// Check if required flags are provided
	if *xmlFile == "" || *xsdFile == "" {
		fmt.Println("Error: Both XML and XSD file paths are required")
		fmt.Println("Usage: validate_xml_by_xsd -xml <xml_file_path> -xsd <xsd_file_path>")
		os.Exit(1)
	}

	// Read XSD schema file
	xsdContent, err := os.ReadFile(*xsdFile)
	if err != nil {
		fmt.Printf("Error reading XSD file: %v\n", err)
		os.Exit(1)
	}

	// Parse XSD schema
	schema, err := xsd.Parse(xsdContent)
	if err != nil {
		fmt.Printf("Error parsing XSD schema: %v\n", err)
		os.Exit(1)
	}
	defer schema.Free()

	// Read XML file
	xmlContent, err := os.ReadFile(*xmlFile)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		os.Exit(1)
	}

	// Parse XML document
	doc, err := libxml2.Parse(xmlContent)
	if err != nil {
		fmt.Printf("Error parsing XML document: %v\n", err)
		os.Exit(1)
	}
	defer doc.Free()

	// Validate XML against XSD schema
	if err := schema.Validate(doc); err != nil {
		fmt.Println("XML validation failed. Errors:")
		for i, validationErr := range err.(xsd.SchemaValidationError).Errors() {
			fmt.Printf("%d. %s\n", i+1, validationErr.Error())
		}
		os.Exit(1)
	}

	fmt.Println("XML validation successful! The XML file conforms to the XSD schema.")
}
