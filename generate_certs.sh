#!/bin/bash

CERTS_DIR="certs"
OPENSSL_CONFIG="$CERTS_DIR/openssl.cnf"

# Create certs directory if it does not exist
mkdir -p $CERTS_DIR

# Generate a single private key and certificate with SAN
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $CERTS_DIR/server.key -out $CERTS_DIR/server.crt -config $OPENSSL_CONFIG

echo "server certificate and key generated successfully."
