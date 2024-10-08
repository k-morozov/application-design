package service

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic/rental"
	"applicationDesign/internal/logic/rental/accommodation"
	"applicationDesign/internal/logic/rental/rental_manager"
	"bytes"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceValidation(t *testing.T) {
	// lg := zerolog.Nop()
	// lg := log.NewLogger("debug")
	// s, _ := NewServiceHTTP(lg, config.NewServiceConfigForDebug(),
	// 	OptLogger(lg))

	rentalManager := rental_manager.NewHotelManager(zerolog.Nop())
	s, _ := NewServiceHTTP(rentalManager, config.NewServiceConfigForDebug())

	tests := []struct {
		name     string
		handler  func(rw http.ResponseWriter, req *http.Request)
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "orders_fail_unknown_filed",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"hotel_id": "reddison",
				"room_id": "lux",
				"email": "guest@mail.ru",
				"from": "2024-01-02T00:00:00Z",
				"to": "2024-01-04T00:00:00Z",
				"new_field": "some_value"
			}`)),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
		{
			name:    "orders_fail_empty",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader("")),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
		{
			name:    "orders_fail_parse_to",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"hotel_id": "reddison",
				"room_id": "lux",
				"email": "guest@mail.ru",
				"from": "2024-01-02T00:00:00Z"
			}`)),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
		{
			name:    "orders_fail_parse_from",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"hotel_id": "reddison",
				"room_id": "lux",
				"email": "guest@mail.ru",
				"to": "2024-01-04T00:00:00Z"
			}`)),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
		{
			name:    "orders_fail_parse_email",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"hotel_id": "reddison",
				"room_id": "lux",
				"from": "2024-01-02T00:00:00Z",
				"to": "2024-01-04T00:00:00Z"
			}`)),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
		{
			name:    "orders_fail_parse_room_id",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"hotel_id": "reddison",
				"email": "guest@mail.ru",
				"from": "2024-01-02T00:00:00Z",
				"to": "2024-01-04T00:00:00Z"
			}`)),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
		{
			name:    "orders_fail_parse_hotel_id",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"room_id": "lux",
				"email": "guest@mail.ru",
				"from": "2024-01-02T00:00:00Z",
				"to": "2024-01-04T00:00:00Z"
			}`)),
			response: &http.Response{
				Status:        "400 Bad Request",
				StatusCode:    http.StatusBadRequest,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			RunCheck(t, test.handler, test.response, test.request)
		})
	}
}

func TestSimpleServiceWork(t *testing.T) {
	lg := zerolog.Nop()
	//lg := log.NewLogger("debug")

	rentalManager := rental_manager.NewHotelManager(lg)

	var hotelID rental.TRentalID = "reddison"
	var roomID accommodation.TAccommodationID = "lux"

	testHotel := rental.NewHotel(hotelID, lg)
	testHotel.AddAccommodation(roomID)

	rentalManager.AddRental(testHotel)

	s, _ := NewServiceHTTP(rentalManager, config.NewServiceConfigForDebug())

	tests := []struct {
		name     string
		handler  func(rw http.ResponseWriter, req *http.Request)
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "ping",
			handler: s.Ping,
			request: httptest.NewRequest(http.MethodGet, "/ping", strings.NewReader("")),
			response: &http.Response{
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
		{
			name:    "orders",
			handler: s.Orders,
			request: httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{
				"hotel_id": "reddison",
				"room_id": "lux",
				"email": "guest@mail.ru",
				"from": "2030-01-02T00:00:00Z",
				"to": "2030-01-04T00:00:00Z"
			}`)),
			response: &http.Response{
				Status:        "201 Created",
				StatusCode:    http.StatusCreated,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          io.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			RunCheck(t, test.handler, test.response, test.request)
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
