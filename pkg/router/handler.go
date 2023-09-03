package router

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

// statusRecorder implement and extends ResponseWriter
type statusRecorder struct {
	http.ResponseWriter
	Status int
}

// WriteHeader sends an HTTP response header with the provided
// status code.
//
// If WriteHeader is not called explicitly, the first call to Write
// will trigger an implicit WriteHeader(http.StatusOK).
// Thus explicit calls to WriteHeader are mainly used to
// send error codes or 1xx informational responses.
//
// The provided code must be a valid HTTP 1xx-5xx status code.
// Any number of 1xx headers may be written, followed by at most
// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
// headers may be buffered. Use the Flusher interface to send
// buffered data. The header map is cleared when 2xx-5xx headers are
// sent, but not with 1xx headers.
//
// The server will automatically send a 100 (Continue) header
// on the first read from the request body if the request has
// an "Expect: 100-continue" header.
func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func httpHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewReader(body))

		recorder := &statusRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		h.ServeHTTP(recorder, r)

		// Calculate response time
		timeEnd := time.Now()
		responseTime := timeEnd.Sub(timeStart)
		wrapped := log.Fields{
			"status":            recorder.Status,
			"method":            r.Method,
			"url":               r.URL.Path,
			"body":              string(body),
			"header":            r.Header,
			"client_ip":         r.RemoteAddr,
			"remote_ip":         r.RemoteAddr,
			"response_time_mcs": responseTime.Microseconds(),
		}

		if recorder.Status >= 400 {
			log.WithFields(wrapped).Error("Error HTTP Observer " + r.URL.Path)
		} else {
			log.WithFields(wrapped).Info("Info HTTP Observer " + r.URL.Path)
		}
	})
}
