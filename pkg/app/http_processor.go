package app

import (
	"encoding/json"
	"fmt"
	"github.com/dienggo/diego/pkg/validates"
	"github.com/gorilla/schema"
	"io"
	"net/http"
	"reflect"
)

type HttpProcessor interface {
	Cast(target interface{}) error
}

// NewHttpProcessor : instance httpProcessor
func NewHttpProcessor(request *http.Request) HttpProcessor {
	return &httpProcessor{Request: request}
}

// httpProcessor context untuk http UseCase
type httpProcessor struct {
	Request *http.Request
}

// Cast data berdasarkan dari http.Request atau MessageProcessor
func (d *httpProcessor) Cast(target interface{}) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target %T cannot addressable, must pointer target", target)
	}

	return d.httpCast(target)
}

func (d *httpProcessor) httpCast(target interface{}) error {
	if d.Request == nil {
		return fmt.Errorf("unable to cast http data, null request")
	}

	// httpCast transform request payload data
	// GET -> params-query-string
	// POST -> json-body
	if err := d.grabMethod(target); err != nil {
		return err
	}
	// validate payload request or params
	return validates.ValidateStructFormatted(target)
}

// Transform query-string into json struct
func (d *httpProcessor) transform(target interface{}, src map[string][]string) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.SetAliasTag("url")
	if err := decoder.Decode(target, src); err != nil {
		return fmt.Errorf("unable to decode query string:%s", err.Error())
	}
	return nil
}

// Grab request method
// Take a destination source of struct
func (d *httpProcessor) grabMethod(target interface{}) error {
	switch d.Request.Method {
	case http.MethodPost, http.MethodPut:
		cType := d.Request.Header.Get("Content-Type")
		if !d.isJSON(cType) {
			return fmt.Errorf("unsupported http content-type=%s", cType)
		}
		return d.decodeJSON(d.Request.Body, target)

	case http.MethodGet:
		return d.transform(target, d.Request.URL.Query())
	default:
		return fmt.Errorf("unsupported method or content-type")
	}
}

func (d *httpProcessor) isJSON(cType string) bool {
	return cType == "application/json"
}

func (d *httpProcessor) decodeJSON(body io.ReadCloser, dst interface{}) error {
	if body == nil {
		return nil
	}
	err := json.NewDecoder(body).Decode(dst)
	if err != nil {
		return fmt.Errorf("unable decode request body, err:%s", err.Error())
	}

	return nil
}
