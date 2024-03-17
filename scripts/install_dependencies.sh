#! bin/bash

#Cgeck for Go

if ! command -v go &> /dev/null
then
    echo "Go could not be found, please install."
    exit 1
fi

if ! command -v lefthook &> /dev/null
then
    echo "Lefthook not found, installing..."
    go install github.com/evilmartians/lefthook@latest
fi

echo "All necessary tools are installed."
 