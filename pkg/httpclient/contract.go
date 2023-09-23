package requester

import (
	"io"
	"net/http"
)

type Contract interface {
	GET(url string, header map[string]string) ([]byte, error)
	DELETE(url string, header map[string]string) ([]byte, error)
	POST(url string, header map[string]string, payload []byte) ([]byte, error)
	PUT(url string, header map[string]string, payload []byte) ([]byte, error)
	RAW(method, url string, body io.Reader) (*http.Request, error)
	WithBasicPOST(url string, header map[string]string, payload []byte, username, password string) ([]byte, error)
}
