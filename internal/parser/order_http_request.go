package parser

import (
	"applicationDesign/internal/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog"
)

func ParseBodyOrderRequest(req *http.Request, lg zerolog.Logger) (*models.Order, error) {
	// rawRequestBody, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	lg.Error().Err(err).Msg("failed read body")
	// 	return nil, errors.New("broken body")
	// }

	// if !json.Valid(rawRequestBody) {
	// 	lg.Error().Err(err).Msg("failed convert body to json")
	// 	return nil, errors.New("broken body")
	// }

	// bodyRequest := models.Order{}

	// if err = json.Unmarshal(rawRequestBody, &bodyRequest); err != nil {
	// 	lg.Error().Err(err).Msg("failed unmarshal body")
	// 	return nil, errors.New("broken body")
	// }

	// lg.Debug().
	// 	Any("bodyRequest", bodyRequest).
	// 	Msg("unmarshal")

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	bodyRequest := models.Order{}

	if err := dec.Decode(&bodyRequest); err != nil {
		lg.Error().Err(err).Msg("failed parse body request")
		return nil, errors.New("broken body")
	}

	if !bodyRequest.Validate() {
		lg.Error().Msg("failed validate")
		return nil, errors.New("failed validate")
	}

	return &bodyRequest, nil
}
