package utils

import (
	"encoding/json"

	"net/http"

	"time"
)



func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.WriteHeader(*status)
	w.Header().Set("content-type", "application/json")
	jsonBytes, _ := json.Marshal(i)
	w.Write(jsonBytes)
	return nil
}

func DecodeJSONRequest(w http.ResponseWriter, r *http.Request, body interface{}) error {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&body)
    if err != nil {
		return err
    }
    return nil
}

func SetCookie(w http.ResponseWriter, name, value string) {
	if name == "userid"{
		cookie := &http.Cookie{
			Name:    name,
			Value:   value,
		}
		http.SetCookie(w, cookie)
		return
	}

	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(24 * time.Hour), 
	}


	http.SetCookie(w, cookie)
}