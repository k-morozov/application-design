package handlers

import (
	"applicationDesign/internal/log"
	"net/http"
)

func Ping(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logFromContext := log.FromContext(ctx)

	logFromContext.Debug().Msg("ping is OK")
	rw.WriteHeader(http.StatusOK)
}
