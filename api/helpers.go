package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func wrap(endpoint Endpoint) http.HandlerFunc {
	return jsonWriterWrap(endpoint)
}

func jsonWriterWrap(endpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := endpoint(w, r)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "internal error", 500)
			return
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "internal error", 500)
			return
		}
	}
}
