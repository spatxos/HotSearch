package response

import (
	"HotSearch/model"
	"encoding/json"
	"net/http"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type Response struct {
	Code int                 `json:"code"`
	Data model.HotSearchData `json:"data"`
	Msg  string              `json:"msg"`
}

func OkWithData(w http.ResponseWriter, data model.HotSearchData) {
	_ = json.NewEncoder(w).Encode(Response{
		Code: SUCCESS,
		Data: data,
		Msg:  "success",
	})
}

func Failed(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(Response{
		Code: ERROR,
		Data: model.HotSearchData{},
		Msg:  err.Error(),
	})
}
