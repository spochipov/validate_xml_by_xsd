#!/bin/bash

# Set environment variables for libxml2
export CGO_CFLAGS="-I/opt/homebrew/opt/libxml2/include/libxml2"
export CGO_LDFLAGS="-L/opt/homebrew/opt/libxml2/lib -lxml2"

# Run the validator
./validate_xml_by_xsd "$@"

# If the exit code is 0, the validation was successful
# The message is already printed by the validator
