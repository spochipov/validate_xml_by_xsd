#!/bin/bash

# Set environment variables for libxml2
export CGO_CFLAGS="-I/opt/homebrew/opt/libxml2/include/libxml2"
export CGO_LDFLAGS="-L/opt/homebrew/opt/libxml2/lib -lxml2"

# Build the application
echo "Building validate_xml_by_xsd..."
go build -o validate_xml_by_xsd

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "Build successful!"
    echo "Usage: ./validate_xml_by_xsd -xml <xml_file_path> -xsd <xsd_file_path>"
else
    echo "Build failed!"
fi
