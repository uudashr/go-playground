

```mermaid
sequenceDiagram
    Client ->> Server: Connection Preface + SETTINGS
    Server ->> Client: SETTINGS
    Server ->> Client: SETTINGS ACK (optional)
    Client ->> Server: SETTINGS ACK (optional)
    Client ->> Server: HEADERS (GET /)
    Client ->> Server: DATA (optional)
    Server ->> Client: HEADERS
    Server ->> Client: DATA
    Server ->> Client: TRAILERS (optional)
```