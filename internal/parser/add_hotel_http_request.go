package parser

import (
	"applicationDesign/internal/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func ParseBodyAddHotelRequest(req *http.Request, lg zerolog.Logger) (*models.AddHotel, error) {
	rawRequestBody, err := io.ReadAll(req.Body)
	if err != nil {
		lg.Error().Err(err).Msg("failed read body")
		return nil, errors.New("broken body")
	}

	if !json.Valid(rawRequestBody) {
		lg.Error().Err(err).Msg("failed convert body to json")
		return nil, errors.New("broken body")
	}

	bodyRequest := models.AddHotel{}

	if err = json.Unmarshal(rawRequestBody, &bodyRequest); err != nil {
		lg.Error().Err(err).Msg("failed unmarshal body")
		return nil, errors.New("broken body")
	}

	lg.Debug().
		Any("bodyRequest", bodyRequest).
		Msg("unmarshal")

	return &bodyRequest, nil
}
