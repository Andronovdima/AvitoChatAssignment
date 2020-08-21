package respond

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Println(err)
	Respond(w, r, code, map[string]string{"message": errors.Cause(err).Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}
