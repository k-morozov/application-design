package handlers

import (
	"applicationDesign/internal/log"
	"applicationDesign/internal/storage"
	"net/http"
)

func Ping(rw http.ResponseWriter, req *http.Request, store storage.Storage) {
	ctx := req.Context()
	logFromContext := log.FromContext(ctx)

	if err := PingImpl(ctx, store); err != nil {
		logFromContext.Error().Msg("ping failed")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logFromContext.Debug().Msg("ping is OK")
	rw.WriteHeader(http.StatusOK)
}
