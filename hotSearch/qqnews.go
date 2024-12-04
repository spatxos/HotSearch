package hotSearch

import (
	"HotSearch/model"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
	"time"
)

type QQnews struct {
}

func (*QQnews) GetHotSearchData(maxNum int) (model.HotSearchData, error) {
	resp, err := http.Get("https://i.news.qq.com/gw/event/pc_hot_ranking_list?offset=0&page_size=51")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	jsonStr := string(body)

	updateTime := time.Now().Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i+1)+".id"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Index:       i + 1,
			Title:       gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i+1)+".title").Str,
			Description: gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i+1)+".abstract").Str,
			Image:       "",
			Popularity:  "",
			URL:         gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i+1)+".url").Str,
		})
	}

	return model.HotSearchData{Source: "腾讯热点榜", UpdateTime: updateTime, HotList: hotList}, nil
}
