package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishchenko-gv/go-example-app/api/apierr"
)

type Endpoint func(http.ResponseWriter, *http.Request) (any, error)

func ep(endpoint Endpoint) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := endpoint(w, r)
		if err != nil {
			writeHttpErr(w, err)
			return
		}

		writeHttpJson(w, resp)
	})
}

func writeHttpErr(w http.ResponseWriter, err error) {
	errType := ""
	message := "internal error"
	status := http.StatusInternalServerError

	switch t := err.(type) {
	case *apierr.Error:
		errType = t.Type
		message = t.Message
		status = t.Status
	}

	fmt.Println(message)

	resp := fmt.Sprintf(`{"error":"%s","message":"%s"}`, errType, message)

	w.WriteHeader(status)
	w.Write([]byte(resp))
}

func writeHttpJson(w http.ResponseWriter, resp any) {
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		writeHttpErr(w, err)
	}
}
