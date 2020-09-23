package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

//ResponseJSON func common
func ResponseJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

//ResponseError func common
func ResponseError(w http.ResponseWriter, err resterrors.RestErr) {
	ResponseJSON(w, err.Status(), err)
}
