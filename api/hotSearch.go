package api

import (
	"HotSearch/hotSearch"
	"HotSearch/model/response"
	"errors"
	"net/http"
	"strings"
)

func GetHotListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Failed(w, errors.New("only GET requests are allowed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	segments := strings.Split(path, "/")
	if len(segments) != 4 {
		response.Failed(w, errors.New("URL parameter error"))
		return
	}
	source := hotSearch.NewSource(segments[3])
	if source == nil {
		response.Failed(w, errors.New("source not found"))
		return
	}
	hotSearchData, err := source.GetHotList(30)
	if err != nil {
		response.Failed(w, errors.New("data cannot be obtained: "+err.Error()))
		return
	}
	response.OkWithData(w, hotSearchData)
}
