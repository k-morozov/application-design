package handlers

import (
	"applicationDesign/internal/log"
	"applicationDesign/internal/provider"
	"net/http"
)

func Ping(rw http.ResponseWriter, req *http.Request, serviceProvider provider.Provider) {
	ctx := req.Context()
	logFromContext := log.FromContext(ctx)

	if err := serviceProvider.Ping(); err != nil {
		logFromContext.Error().Msg("ping failed")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logFromContext.Debug().Msg("ping is OK")
	rw.WriteHeader(http.StatusOK)
}
