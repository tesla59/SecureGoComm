# Project Shepherd (WIP)

Refer to `shepherd_features.yaml` for planned features.

## How to generate certificates
1. Generate Certificate Authority (CA)
Execute the commands inside certificates/CA
```
# Generate CA private key (protect with passphrase)
openssl genrsa -aes256 -out ca.key 4096
# Create CA certificate
openssl req -new -x509 -days 365 -key ca.key -out ca.crt
```

2. Generate Server Certificate
Execute the commands inside certificates/server
```
# Generate server private key
openssl genrsa -out server.key 2048
# Create Certificate Signing Request (CSR)
openssl req -new -key server.key -out server.csr
# Sign server certificate with CA
openssl x509 -req -days 365 -in server.csr -CA ../CA/ca.crt -CAkey ../CA/ca.key -CAcreateserial -out server.crt
```

3. Generate Client Certificate
Execute the commands inside certificates/client
```
# Generate client private key
openssl genrsa -out client.key 2048
# Create Certificate Signing Request (CSR)
openssl req -new -key client.key -out client.csr
# Sign client certificate with CA
openssl x509 -req -days 365 -in client.csr -CA ../CA/ca.crt -CAkey ../CA/ca.key -CAcreateserial -out client.crt
```
