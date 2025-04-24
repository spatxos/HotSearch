package hotSearch

import (
	"github.com/spatxos/HotSearch/model"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Sougou struct {
}

func (*Sougou) GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error) {
	resp, err := http.Get("https://hotlist.imtt.qq.com/Fetch")
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
		if index := gjson.Get(jsonStr, "main."+strconv.Itoa(i)+".title"); !index.Exists() {
			break
		}
		floatValue, _ := strconv.ParseFloat(gjson.Get(jsonStr, "main."+strconv.Itoa(i)+".score").Str, 64)
		popularity := fmt.Sprintf("%.1f万", floatValue)
		hotList = append(hotList, model.HotItem{
			Index:       i + 1,
			Title:       gjson.Get(jsonStr, "main."+strconv.Itoa(i)+".title").Str,
			Description: "",
			Image:       "",
			Popularity:  popularity,
			URL:         gjson.Get(jsonStr, "main."+strconv.Itoa(i)+".url").Str,
		})
	}

	return model.HotSearchData{Source: "搜狗热搜", UpdateTime: updateTime, HotList: hotList}, nil
}
