#!/bin/bash
# tpm-morphia -collection-def-scan-path . -out-dir . -format-code

tpm-morphia -collection-def-file ./user-card-tpmm.json -out-dir . -format-code
tpm-morphia -collection-def-file ./file-tpmm.json -out-dir . -format-code
