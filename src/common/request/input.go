package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Input struct {
	Request *http.Request
}

func (input *Input) InputBody() (string, error) {
	contentBody, err := ioutil.ReadAll(input.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(contentBody), nil
}

func (input *Input) IsPost() bool {
	return input.Request.Method == "POST"
}

func (input *Input) IsGet() bool {
	return input.Request.Method == "GET"
}

func (input *Input) GetValue(key string) string {
	input.Request.ParseForm()
	return input.Request.FormValue(key)
}
