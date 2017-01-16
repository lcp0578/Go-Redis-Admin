package response

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Code int
	Msg  string
	Data map[string]string
}

func OuputJson(w http.ResponseWriter, jr *JsonResponse) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(jr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
