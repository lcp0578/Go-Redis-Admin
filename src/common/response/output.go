package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func init() {

}

type Output struct {
	ResponseWriter http.ResponseWriter
	Status         int
}

func (output *Output) Json(data interface{}, hasIndent bool, coding bool) error {
	output.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	var content []byte
	var err error
	if hasIndent {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(output.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	if coding {
		content = []byte(stringsToJson(string(content)))
	}
	_, err = output.ResponseWriter.Write(content)
	return err
}

func (output *Output) WriteString(content string) {
	output.ResponseWriter.Write([]byte(content))
}

func stringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}

func (output *Output) Xml(data interface{}) error {
	return nil
}
