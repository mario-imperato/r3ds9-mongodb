#!/bin/bash
# tpm-morphia -collection-def-scan-path . -out-dir . -format-code

tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name SysInfo --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name UserRole --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name App --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name FileVariant --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name FileReference --out-dir ../.. --with-format --with-bak

tpm-morphia-cli gen entity --schema-file ./apicms/schema.yml --name entRefStruct --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicms/schema.yml --name file --out-dir ../.. --with-format --with-bak