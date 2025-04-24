package api

import (
	"errors"
	"net/http"

	"github.com/spatxos/HotSearch/hotSearch"
	"github.com/spatxos/HotSearch/model"
	"github.com/spatxos/HotSearch/model/response"
)

func GetHotListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Failed(w, errors.New("only GET requests are allowed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	source := r.URL.Query().Get("source")
	if source == "" {
		response.Failed(w, errors.New("source parameter is required"))
		return
	}
	sourceInstance := hotSearch.NewSource(source)
	if sourceInstance == nil {
		response.Failed(w, errors.New("source not found"))
		return
	}
	hotSearchData, err := sourceInstance.GetHotSearchData(30)
	if err != nil {
		response.Failed(w, errors.New("data cannot be obtained: "+err.Error()))
		return
	}
	response.OkWithData(w, hotSearchData)
}

func GetSourcesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Failed(w, errors.New("only GET requests are allowed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.OkWithSourceList(w, hotSearch.GetAvailableSources())
}

func GetAllHotListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Failed(w, errors.New("only GET requests are allowed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// 获取所有可用的来源
	sources := hotSearch.GetAvailableSources()

	// 存储所有来源的热搜数据
	allHotSearchData := make(map[string]model.HotSearchData)

	// 遍历所有来源，获取热搜数据
	for _, source := range sources {
		sourceInstance := hotSearch.NewSource(source)
		if sourceInstance != nil {
			hotSearchData, err := sourceInstance.GetHotSearchData(30)
			if err == nil {
				allHotSearchData[source] = hotSearchData
			}
		}
	}

	response.OkWithAllHotSearch(w, allHotSearchData)
}
