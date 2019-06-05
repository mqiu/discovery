package streams

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
)


const (
	acceptedHeaderType = "application/json"
)

//getMessageHandler handles GET request
func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if v, ok := vars["message"]; ok {
		if str, err := FindSaved(v); err == nil {
			resp := responseData{
				Message: str,
			}
			resBody, err := json.Marshal(resp)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte("unable to process data"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(resBody)
			return
		}
	}
	w.WriteHeader(404)
	resp := responseData{
		ErrMsg: "Message not found",
	}
	resBody, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("unable to process data"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(resBody)
	return
}