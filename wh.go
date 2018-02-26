package wh

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithValidationsError(w http.ResponseWriter, code int, message string, validations interface{}) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message, "validations": validations})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithEmptySuccess(w http.ResponseWriter) {
	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func BodyToModel(w http.ResponseWriter, body io.ReadCloser, model interface{}) bool {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&model); err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return true
	}
	defer body.Close()

	return false
}
