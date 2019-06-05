package streams

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	acceptedHeaderType = "application/json"
)

//GetStreamHandler handles GET request
func GetStreamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["streamId"]; !ok {
		writeError(w,"failed to find stream id", http.StatusBadRequest)
		return
	}
	record, err := ReadStream(vars["streamId"])
	if err != nil {
		writeError(w,fmt.Sprintf("failed to find record: %s", err.Error()), http.StatusBadRequest)
		return
	}
	if record == nil {
		writeError(w,fmt.Sprintf("cannot find the record:%s", vars["streamId"]), http.StatusNotFound)
		return
	}
	resBody, err := json.Marshal(record)
	if err != nil {
		fmt.Println(err)
		writeError(w,"failed to process the result", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resBody)
	return
}

func writeError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
