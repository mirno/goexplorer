#! /bin/bash

#ToDo: Verify if the CA is existing since we should not be overriding this the whole time. Using Make this would be an ideal resolve.

# Detect the OS
OS="unknown"

case "$(uname -s)" in  
    Darwin)
        OS="macos"
        ;;
    Linux)
        OS="linux"
        ;;
    CYGWIN*|MINGW32*|MSYS*|MINGW*)
        OS="windows"
        ;;
esac

# Use openssl command based on the OS
OPENSSL="openssl"
if [ "$OS" = "windows"]; then
    OPENSSL="openssl.exe"
fi
# Generate CA and sign the cert
$OPENSSL req -x509 -new -nodes -keyout .private/ca.key -sha256 -days 1024 -out .private/ca.crt -subj "//CN=Example CA"
$OPENSSL req -new -nodes -keyout .private/server.key -out .private/server.csr -config req.conf
$OPENSSL x509 -req -in .private/server.csr -CA .private/ca.crt -CAkey .private/ca.key -CAcreateserial -out .private/server.crt -days 365 -sha256 -extfile signing.conf -extensions v3_req