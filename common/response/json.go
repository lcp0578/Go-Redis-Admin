package response

import "encoding/json"

func init() {

}
type JsonResponse struct {
	Code int
	Msg string
	Data []string
}
func (json *JsonResponse)Encode(code int, msg string, data []string) string {
	encoder := json.NewEncoder()
	var jdata = &JsonResponse{Code:code, Msg:msg, Data:data}
	return encoder.Encode(jdata)
}