#!/usr/bin/env bash
# This script converts gocov coverage to cobertura format.
#
# Usage:
# ./convert_gocov_cobertura.sh <input_file> <output_file>
# Source: https://stackoverflow.com/a/31440740/688954

# Receive arg 1 as input file and arg 2 as outputfile
INPUT_FILE=$1
OUTPUT_FILE=$2

if [ -z "$INPUT_FILE" ] || [ -z "$OUTPUT_FILE" ]; then
    echo "Usage: $0 <input_file> <output_file>"
    exit 1
fi

gocov convert $INPUT_FILE | gocov-xml > $OUTPUT_FILE
