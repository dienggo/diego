package app

import (
	"bytes"
	"encoding/json"
	"github.com/dienggo/diego/pkg/render"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// NewUseCase is method to handle UseCase implemented
func NewUseCase(uc UseCase) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		uc.Handle(UseCaseHandler{Request: request, Writer: writer})
	})
}

// TestUseCase is method to handle UseCase implemented on test purpose
func TestUseCase(uc UseCase, method string, url string, data any) (w *httptest.ResponseRecorder, err error) {
	marshal, _ := json.Marshal(data)

	r, err := http.NewRequest(method, url, bytes.NewBuffer(marshal))
	r.Header.Set("Content-Type", "application/json")

	if err == nil {
		// Create a response recorder to capture the response
		w = httptest.NewRecorder()
		uc.Handle(UseCaseHandler{Request: r, Writer: w})
	}
	if err != nil {
		println("Err", err.Error())
	}
	return w, err
}

type UseCase interface {
	Handle(uch UseCaseHandler)
}

type UseCaseHandler struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

// CastAndValidate is casting Request http into dto struct owned and make validate Request
func (uc UseCaseHandler) CastAndValidate(target interface{}) error {
	return NewHttpProcessor(uc.Request).Cast(target)
}

// GetParam is get data by key on request parameter or json data or body request
func (uc UseCaseHandler) GetParam(key string) string {
	value := mux.Vars(uc.Request)[key]
	if value == "" {
		value = uc.Request.URL.Query().Get(key)
	}
	return value
}

// GetLimit is method to get param limit from request
func (uc UseCaseHandler) GetLimit() int {
	i, err := strconv.Atoi(uc.GetParam("limit"))
	if err != nil {
		log.Error("error get limit ", err.Error(), mux.Vars(uc.Request)["limit"])
	}
	return i
}

// GetPage is method to get param page from request
func (uc UseCaseHandler) GetPage() int {
	i, err := strconv.Atoi(uc.GetParam("page"))
	if err != nil {
		log.Error("error get page ", err.Error())
	}
	return i
}

// JsonResponse is writes the response headers and calls JSON to render data.
func (uc UseCaseHandler) JsonResponse(code int, message string, data any) {
	render.JsonFormatted(uc.Writer, render.Data{
		Code:    code,
		Name:    http.StatusText(code),
		Message: message,
		Data:    data,
	})
}

// JsonResponsePaginate is method to build template json response with pagination format
func (uc UseCaseHandler) JsonResponsePaginate(code int, message string, data any, count int) {
	render.JsonFormattedPaginate(uc.Writer, code, message, data, count, uc.GetLimit(), uc.GetPage())
}

// Response is writes the response headers and calls JSON to render data.
func (uc UseCaseHandler) Response() render.FormattedJson {
	return render.FormattedJson{Writer: uc.Writer}
}
