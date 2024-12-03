package hotSearch

import (
	"HotSearch/model"
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Baidu struct {
}

func (*Baidu) GetHotList(maxNum int) (HotSearchData model.HotSearchData, err error) {
	resp, err := http.Get("https://top.baidu.com/board?tab=realtime")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	var jsonStr string
	reg := regexp.MustCompile(`<!--s-data:({.*?})-->`)
	result := reg.FindAllStringSubmatch(string(body), -1)
	if len(result) > 0 && len(result[0]) > 1 {
		jsonStr = result[0][1]
	} else {
		return model.HotSearchData{}, errors.New("failed to get data")
	}

	updateTime := time.Unix(gjson.Get(jsonStr, "data.cards.0.updateTime").Int(), 0).Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "data.cards.0.content."+strconv.Itoa(i)+".index"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Index:       i + 1,
			Title:       gjson.Get(jsonStr, "data.cards.0.content."+strconv.Itoa(i)+".word").Str,
			Description: gjson.Get(jsonStr, "data.cards.0.content."+strconv.Itoa(i)+".desc").Str,
			Image:       gjson.Get(jsonStr, "data.cards.0.content."+strconv.Itoa(i)+".img").Str,
			Popularity:  gjson.Get(jsonStr, "data.cards.0.content."+strconv.Itoa(i)+".hotScore").Str,
			URL:         gjson.Get(jsonStr, "data.cards.0.content."+strconv.Itoa(i)+".rawUrl").Str,
		})
	}

	return model.HotSearchData{Source: "百度热搜", UpdateTime: updateTime, HotList: hotList}, nil
}
