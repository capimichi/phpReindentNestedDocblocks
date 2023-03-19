#!/bin/bash

#read version from version.txt
version=$(cat version.txt)
mkdir "dist/${version}"

zip_path="dist/${version}/php_reindent_nested_docblocks.macos.${version}.zip"
rm "${zip_path}"
env GOOS=darwin GOARCH=amd64 go build .
mv phpReindentNestedDocblocks php_reindent_nested_docblocks
zip -r "${zip_path}" php_reindent_nested_docblocks
rm php_reindent_nested_docblocks

zip_path="dist/${version}/php_reindent_nested_docblocks.linux32.${version}.zip"
rm "${zip_path}"
env GOOS=linux GOARCH=386 go build .
mv phpReindentNestedDocblocks php_reindent_nested_docblocks
zip -r "${zip_path}" php_reindent_nested_docblocks
rm php_reindent_nested_docblocks

zip_path="dist/${version}/php_reindent_nested_docblocks.linux64.${version}.zip"
rm "${zip_path}"
env GOOS=linux GOARCH=amd64 go build .
mv phpReindentNestedDocblocks php_reindent_nested_docblocks
zip -r "${zip_path}" php_reindent_nested_docblocks
rm php_reindent_nested_docblocks

zip_path="dist/${version}/php_reindent_nested_docblocks.windows32.${version}.zip"
rm "${zip_path}"
env GOOS=windows GOARCH=386 go build .
mv phpReindentNestedDocblocks.exe php_reindent_nested_docblocks.exe
zip -r "${zip_path}" php_reindent_nested_docblocks.exe
rm php_reindent_nested_docblocks.exe

zip_path="dist/${version}/php_reindent_nested_docblocks.windows64.${version}.zip"
rm "${zip_path}"
env GOOS=windows GOARCH=amd64 go build .
mv phpReindentNestedDocblocks.exe php_reindent_nested_docblocks.exe
zip -r "${zip_path}" php_reindent_nested_docblocks.exe
rm php_reindent_nested_docblocks.exe