package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

// see https://datatracker.ietf.org/doc/html/rfc9113
// see https://chatgpt.com/c/688ccbe2-1ecc-8324-bcf2-218a1b38982c

func runMain() error {
	logger := slog.Default()

	addr := "curl.se:443"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", addr, err)
	}

	defer conn.Close()

	tlsConn := tls.Client(conn, &tls.Config{
		ServerName: "curl.se",
		NextProtos: []string{"h2"},
	})

	if err := tlsConn.Handshake(); err != nil {
		return fmt.Errorf("TLS handshake failed: %w", err)
	}

	state := tlsConn.ConnectionState()

	if state.NegotiatedProtocol != "h2" {
		return fmt.Errorf("expected h2, got %s", state.NegotiatedProtocol)
	}

	logger.Info("TLS handshake completed", "protocol", state.NegotiatedProtocol)

	if _, err = tlsConn.Write([]byte(http2.ClientPreface)); err != nil {
		return fmt.Errorf("failed to write client preface: %w", err)
	}

	framer := http2.NewFramer(tlsConn, tlsConn)

	if err := framer.WriteSettings(); err != nil {
		return fmt.Errorf("failed to write settings: %w", err)
	}

	logger.Info("HTTP/2 settings written")

	headers := []hpack.HeaderField{
		{Name: ":method", Value: http.MethodGet},
		{Name: ":scheme", Value: "https"},
		{Name: ":authority", Value: "curl.se"},
		{Name: ":path", Value: "/"},
		{Name: "user-agent", Value: "go-http2-client"},
	}

	var hpackBuf bytes.Buffer
	enc := hpack.NewEncoder(&hpackBuf)
	for _, hf := range headers {
		if err := enc.WriteField(hf); err != nil {
			return fmt.Errorf("failed to encode header field %s: %w", hf.Name, err)
		}
	}

	headerBlock := hpackBuf.Bytes()
	if err := framer.WriteHeaders(http2.HeadersFrameParam{
		StreamID:      1,
		BlockFragment: headerBlock,
		EndStream:     true,
		EndHeaders:    true,
	}); err != nil {
		return fmt.Errorf("failed to send headers: %w", err)
	}

	logger.Info("Headers sent", "headers", headers)

	for {
		frame, err := framer.ReadFrame()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read frame: %w", err)
		}

		switch f := frame.(type) {
		case *http2.HeadersFrame:
			logger.Info("Received HeadersFrame")
		case *http2.DataFrame:
			logger.Info("Received DataFrame", "length", f.Length, "data", f.Data())
		case *http2.SettingsFrame:
			logger.Info("Received SettingFrame")
		case *http2.GoAwayFrame:
			logger.Info("Received GoAwayFrame")
			return nil
		default:
			logger.Warn("Received unknown frame type", "type", f.Header().Type)

		}
	}

	return nil
}

func main() {
	if err := runMain(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
