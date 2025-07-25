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

Send request with HTTP/2 support securely:
```shell
curl -v --http2 https://localhost:8443/hello --cacert certs/cert.pem
```

Send request with HTTP/2 support insecurely:
```shell
curl -v --http2 https://localhost:8443/hello --insecure
```


or the HTTP/2 Over Cleartext (h2c) protocol:
```shell
curl -v --http2-prior-knowledge http://localhost:8080/hello
```

## Run nginx

```shell
podman run --rm -it --name nginx \
  -p 8443:8443 \
  -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf:ro \
  -v $(pwd)/certs:/etc/nginx/certs:ro \
  nginx:1.29.0-alpine
```

## References

- [How HTTP/2 Works and How to Enable It in Go](https://victoriametrics.com/blog/go-http2/)