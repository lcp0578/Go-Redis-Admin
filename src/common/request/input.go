package request

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

type Input struct {
	Request *http.Request
}

func (input *Input)InputBody()(string,error){
	contentBody, err := ioutil.ReadAll(input.Request.Body)
	if err!=nil{
		fmt.Println(err.Error())
		return "",err
	}
	return string(contentBody),nil
}
