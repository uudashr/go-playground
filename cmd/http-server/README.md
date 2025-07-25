## Generate Certificate

```shell
openssl req -x509 -newkey rsa:2048 -nodes \
    -keyout certs/key.pem \
    -out certs/cert.pem \
    -days 365 \
    -config certs/localhost-openssl.cnf \
    -extensions v3_req
```

## Send request

```shell
http --verify=certs/cert.pem https://localhost/hello
```

