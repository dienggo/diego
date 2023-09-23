package requester

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// handler ...
type handler struct{}

// New ...
func New() Contract {
	return &handler{}
}

// RAW ...
func (request *handler) RAW(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

// GET request type get
func (request *handler) GET(url string, header map[string]string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	return body, nil
}

// POST request type post
func (request *handler) POST(url string, header map[string]string, payload []byte) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	return body, nil
}

// WithBasicPOST request type post
func (request *handler) WithBasicPOST(url string, header map[string]string, payload []byte, username, password string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	if err != nil {
		return result, err
	}
	if len(header) != 0 {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	return body, nil
}

// PUT request type post
func (request *handler) PUT(url string, header map[string]string, payload []byte) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if (resp.StatusCode != 200) && (resp.StatusCode != 201) {
		return result, errors.New(string(body))
	}
	return body, nil
}

// DELETE request type get
func (request *handler) DELETE(url string, header map[string]string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return result, err
	}
	// set default headers
	req.Header.Add("Content-Type", "application/json")
	// add optional headers
	for content, value := range header {
		req.Header.Set(content, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	return body, nil
}
