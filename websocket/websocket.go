package websocket

import (
	"bufio"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"io"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var stripWebsocketHeaders = []string {
	"Upgrade",
	"Connection",
	"Sec-Websocket-Key",
	"Sec-Websocket-Version",
	"Sec-Websocket-Extensions",
}

// IsWebSocketUpgrade checks to see if the request is a WebSocket connection.
func IsWebSocketUpgrade(req *http.Request) bool {
	return websocket.IsWebSocketUpgrade(req)
}

// ClientConnect creates a WebSocket client connection for provided request. Caller is responsible for closing
// the connection. The response body may not contain the entire response and does
// not need to be closed by the application.
func ClientConnect(req *http.Request, tlsClientConfig *tls.Config) (*websocket.Conn, *http.Response, error) {
	req.URL.Scheme = changeRequestScheme(req)
	wsHeaders := websocketHeaders(req)

	d := &websocket.Dialer{TLSClientConfig: tlsClientConfig}
	conn, response, err := d.Dial(req.URL.String(), wsHeaders)
	if err != nil {
		return nil, nil, err
	}
	response.Header.Set("Sec-WebSocket-Accept", generateAcceptKey(req))
	return conn, response, err
}

// HijackConnection takes over an HTTP connection. Caller is responsible for closing connection.
func HijackConnection(w http.ResponseWriter) (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack error")
	}

	conn, brw, err := hj.Hijack()
	if err != nil {
		return nil, nil, err
	}
	return conn, brw, nil
}

// Stream copies copy data to & from provided io.ReadWriters.
func Stream(conn, backendConn io.ReadWriter) {
	proxyDone := make(chan struct{}, 2)

	go func() {
		io.Copy(conn, backendConn)
		proxyDone <- struct{}{}
	}()

	go func() {
		io.Copy(backendConn, conn)
		proxyDone <- struct{}{}
	}()

	// If one side is done, we are done.
	<-proxyDone
}

// the gorilla websocket library sets its own Upgrade, Connection, Sec-WebSocket-Key,
// Sec-WebSocket-Version and Sec-Websocket-Extensions headers.
// https://github.com/gorilla/websocket/blob/master/client.go#L189-L194.
func websocketHeaders(req *http.Request) http.Header {
	wsHeaders := make(http.Header)
	for key, val := range req.Header {
			wsHeaders[key] = val
	}
	// Assume the header keys are in canonical format.
	for _,  header := range stripWebsocketHeaders {
		wsHeaders.Del(header)
	}
	return wsHeaders
}

// sha1Base64 sha1 and then base64 encodes str.
func sha1Base64(str string) string {
	hasher := sha1.New()
	io.WriteString(hasher, str)
	hash := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
}

// generateAcceptKey returns the string needed for the Sec-WebSocket-Accept header.
// https://tools.ietf.org/html/rfc6455#section-1.3 describes this process in more detail.
func generateAcceptKey(req *http.Request) string {
	return sha1Base64(req.Header.Get("Sec-WebSocket-Key") + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
}

// changeRequestScheme is needed as the gorilla websocket library requires the ws scheme.
// (even though it changes it back to http/https, but ¯\_(ツ)_/¯.)
func changeRequestScheme(req *http.Request) string {
	switch req.URL.Scheme {
	case "https":
		return "wss"
	case "http":
		return "ws"
	default:
		return req.URL.Scheme
	}
}
