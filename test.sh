#!/bin/bash

echo "===== Testing XML Validator ====="
echo ""

echo "1. Testing with valid XML file (valid.xml):"
./validate.sh -xml valid.xml -xsd schema.xsd
echo ""

echo "2. Testing with invalid XML file (invalid.xml):"
./validate.sh -xml invalid.xml -xsd schema.xsd
echo ""

echo "===== Test Complete ====="
