#!/bin/bash
# tpm-morphia -collection-def-scan-path . -out-dir . -format-code

tpm-morphia -collection-def-file ./user-tpmm.json -out-dir . -format-code
tpm-morphia -collection-def-file ./kv-tpmm.json -out-dir . -format-code
tpm-morphia -collection-def-file ./file-tpmm.json -out-dir . -format-code
tpm-morphia -collection-def-file ./site-tpmm.json -out-dir . -format-code