package api

import (
	"encoding/json"
	"net/http"
)

func wrap(endpoint Endpoint) http.HandlerFunc {
	middlewaresChain := chainMiddlewares(
		loggingMiddleware,
		jsonResponseMiddleware,
		authMiddleware,
	)

	jsonWriter := jsonWriterWrap(endpoint)

	return middlewaresChain(jsonWriter).ServeHTTP
}

func jsonWriterWrap(endpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := endpoint(w, r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}
