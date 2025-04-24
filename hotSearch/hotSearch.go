package hotSearch

import "github.com/spatxos/HotSearch/model"

type Source interface {
	GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error)
}

// SourceInfo 存储来源信息
type SourceInfo struct {
	Name   string
	Source Source
}

// 存储所有可用的来源
var availableSources = []SourceInfo{
	{Name: "baidu", Source: &Baidu{}},       // 百度热搜
	{Name: "bilibili", Source: &Bilibili{}}, // B站排行榜
	{Name: "douyin", Source: &Douyin{}},     // 抖音热榜
	{Name: "kuaishou", Source: &Kuaishou{}}, // 快手热榜
	{Name: "pengpai", Source: &Pengpai{}},   // 澎湃热榜
	{Name: "qqnews", Source: &QQnews{}},     // 腾讯热点榜
	{Name: "sina", Source: &Sina{}},         // 新浪热榜
	{Name: "sougou", Source: &Sougou{}},     // 搜狗热搜
	{Name: "sspai", Source: &Sspai{}},       // 少数派最热
	{Name: "tieba", Source: &Tieba{}},       // 贴吧热议榜
	{Name: "toutiao", Source: &Toutiao{}},   // 头条热榜
	{Name: "weibo", Source: &Weibo{}},       // 微博热搜
	{Name: "zhihu", Source: &Zhihu{}},       // 知乎热榜
}

func NewSource(source string) Source {
	for _, info := range availableSources {
		if info.Name == source {
			return info.Source
		}
	}
	return nil
}

// GetAvailableSources 返回所有可用的热搜来源名称
func GetAvailableSources() []string {
	sources := make([]string, len(availableSources))
	for i, info := range availableSources {
		sources[i] = info.Name
	}
	return sources
}
