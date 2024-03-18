package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)



func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.WriteHeader(*status)
	w.Header().Set("content-type", "application/json")
	jsonBytes, _ := json.Marshal(i)
	w.Write(jsonBytes)
	return nil
}

func DecodeJSONRequest(w http.ResponseWriter, r *http.Request, body interface{}) bool {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&body)
    if err != nil {
        errorMsg := ""
        var errorCode int = http.StatusBadRequest // По умолчанию большинство ошибок будет 400

        switch {
        case errors.As(err, &json.SyntaxError{}):
            errorMsg = fmt.Sprintf("Request body contains badly-formed JSON")

        case errors.Is(err, io.ErrUnexpectedEOF):
            errorMsg = "Request body contains badly-formed JSON"

        case errors.As(err, &json.UnmarshalTypeError{}):
            errorMsg = "Request body contains an incorrect value for the type"

        case strings.HasPrefix(err.Error(), "json: unknown field"):
            errorMsg = "Request body contains unknown field"
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            errorMsg += ": " + fieldName

        case errors.Is(err, io.EOF):
            errorMsg = "Request body must not be empty"

        case err.Error() == "http: request body too large":
            errorMsg = "Request body must not exceed 1MB"
            errorCode = http.StatusRequestEntityTooLarge

        default:
            log.Println(err)
            errorMsg = http.StatusText(http.StatusInternalServerError)
            errorCode = http.StatusInternalServerError
        }


        response := map[string]string{
            "error": errorMsg,
        }

        EncodeJSONResponse(response, &errorCode, w)
        return false
    }
    return true
}