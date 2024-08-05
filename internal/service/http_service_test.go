package service

import (
	"applicationDesign/internal/config"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/zerolog"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceWork(t *testing.T) {
	tests := []struct {
		name         string
		pingRequest  *http.Request
		pingResponse *http.Response
	}{
		{
			name:        "ping",
			pingRequest: httptest.NewRequest(http.MethodGet, "/ping", strings.NewReader("")),
			pingResponse: &http.Response{
				Status:        "200 OK",
				StatusCode:    http.StatusOK,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
	}

	s, _ := NewServiceHTTP(zerolog.Nop(), config.NewServiceConfigForDebug())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			RunCheck(t, s.Ping, test.pingResponse, test.pingRequest)
		})
	}
}

func RunCheck(t *testing.T, handler func(rw http.ResponseWriter, req *http.Request),
	responseExpected *http.Response, request *http.Request) {
	w := httptest.NewRecorder()
	handler(w, request)
	responseAdd := w.Result()
	defer responseAdd.Body.Close()

	assert.Equal(t, responseExpected, responseAdd)
}
