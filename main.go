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
	"os/user"
	"path/filepath"
	"strings"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

// expandTilde replaces the tilde (~) in a path with the user's home directory
func expandTilde(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	if path == "~" {
		return usr.HomeDir, nil
	}

	return filepath.Join(usr.HomeDir, path[2:]), nil
}

// Russian messages
var (
	msgErrorRequiredPaths      = "Ошибка: Требуются пути к файлам XML и XSD"
	msgUsage                   = "Использование: validate_xml_by_xsd -xml <путь_к_xml_файлу> -xsd <путь_к_xsd_схеме>"
	msgErrorProcessingXMLPath  = "Ошибка при обработке пути к XML файлу: %v"
	msgErrorProcessingXSDPath  = "Ошибка при обработке пути к XSD файлу: %v"
	msgErrorGettingCurrentDir  = "Ошибка при получении текущей директории: %v"
	msgErrorChangingDir        = "Ошибка при смене директории на %s: %v"
	msgErrorReadingXSDFile     = "Ошибка чтения XSD файла: %v"
	msgErrorParsingXSDSchema   = "Ошибка при разборе XSD схемы: %v"
	msgErrorRestoringDir       = "Ошибка при восстановлении исходной директории: %v"
	msgErrorReadingXMLFile     = "Ошибка чтения XML файла: %v"
	msgErrorParsingXMLDoc      = "Ошибка при разборе XML документа: %v"
	msgValidationFailed        = "Валидация XML не пройдена. Ошибки:"
	msgValidationSuccessful    = "Валидация XML успешно завершена! XML файл соответствует XSD схеме."
)

func main() {
	// Define command line flags
	xmlFile := flag.String("xml", "", "Path to XML file (required)")
	xsdFile := flag.String("xsd", "", "Path to XSD schema file (required)")
	flag.Parse()

	// Check if required flags are provided
	if *xmlFile == "" || *xsdFile == "" {
		fmt.Println(msgErrorRequiredPaths)
		fmt.Println(msgUsage)
		os.Exit(1)
	}

	// Expand tilde in file paths if present
	expandedXmlPath, err := expandTilde(*xmlFile)
	if err != nil {
		fmt.Printf(msgErrorProcessingXMLPath+"\n", err)
		os.Exit(1)
	}
	*xmlFile = expandedXmlPath

	expandedXsdPath, err := expandTilde(*xsdFile)
	if err != nil {
		fmt.Printf(msgErrorProcessingXSDPath+"\n", err)
		os.Exit(1)
	}
	*xsdFile = expandedXsdPath

	// Get the directory of the XSD file to resolve relative includes
	xsdDir := filepath.Dir(*xsdFile)
	
	// Save current working directory
	originalWd, err := os.Getwd()
	if err != nil {
		fmt.Printf(msgErrorGettingCurrentDir+"\n", err)
		os.Exit(1)
	}
	
	// Change to XSD directory to resolve relative includes
	err = os.Chdir(xsdDir)
	if err != nil {
		fmt.Printf(msgErrorChangingDir+"\n", xsdDir, err)
		os.Exit(1)
	}
	
	// Read XSD schema file
	xsdContent, err := os.ReadFile(*xsdFile)
	if err != nil {
		// Restore original working directory before exiting
		_ = os.Chdir(originalWd)
		fmt.Printf(msgErrorReadingXSDFile+"\n", err)
		os.Exit(1)
	}
	
	// Parse XSD schema
	schema, err := xsd.Parse(xsdContent)
	if err != nil {
		// Restore original working directory before exiting
		_ = os.Chdir(originalWd)
		fmt.Printf(msgErrorParsingXSDSchema+"\n", err)
		os.Exit(1)
	}
	defer schema.Free()
	
	// Restore original working directory
	err = os.Chdir(originalWd)
	if err != nil {
		fmt.Printf(msgErrorRestoringDir+"\n", err)
		os.Exit(1)
	}

	// Read XML file
	xmlContent, err := os.ReadFile(*xmlFile)
	if err != nil {
		fmt.Printf(msgErrorReadingXMLFile+"\n", err)
		os.Exit(1)
	}

	// Parse XML document
	doc, err := libxml2.Parse(xmlContent)
	if err != nil {
		fmt.Printf(msgErrorParsingXMLDoc+"\n", err)
		os.Exit(1)
	}
	defer doc.Free()

	// Validate XML against XSD schema
	if err := schema.Validate(doc); err != nil {
		fmt.Println(msgValidationFailed)
		for i, validationErr := range err.(xsd.SchemaValidationError).Errors() {
			fmt.Printf("%d. %s\n", i+1, validationErr.Error())
		}
		os.Exit(1)
	}

	fmt.Println(msgValidationSuccessful)
}
