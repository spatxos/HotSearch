package response

import (
	"github.com/spatxos/HotSearch/model"
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

type SourceListResponse struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
	Msg  string   `json:"msg"`
}

type AllHotSearchResponse struct {
	Code int                            `json:"code"`
	Data map[string]model.HotSearchData `json:"data"`
	Msg  string                         `json:"msg"`
}

func OkWithData(w http.ResponseWriter, data model.HotSearchData) {
	_ = json.NewEncoder(w).Encode(Response{
		Code: SUCCESS,
		Data: data,
		Msg:  "success",
	})
}

func OkWithSourceList(w http.ResponseWriter, sources []string) {
	_ = json.NewEncoder(w).Encode(SourceListResponse{
		Code: SUCCESS,
		Data: sources,
		Msg:  "success",
	})
}

func OkWithAllHotSearch(w http.ResponseWriter, data map[string]model.HotSearchData) {
	_ = json.NewEncoder(w).Encode(AllHotSearchResponse{
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
