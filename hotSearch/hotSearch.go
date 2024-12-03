package hotSearch

import "HotSearch/model"

type Source interface {
	GetHotList(maxNum int) (HotSearchData model.HotSearchData, err error)
}

func NewSource(source string) Source {
	switch source {
	case "baidu":
		return &Baidu{}
	case "bilibili":
		return &Bilibili{}
	case "douyin":
		return &Douyin{}
	case "kuaishou":
		return &Kuaishou{}
	case "pengpai": // 20
		return &Pengpai{}
	case "qqnews":
		return &QQnews{}
	case "sina":
		return &Sina{}
	case "sougou":
		return &Sougou{}
	case "sspai":
		return &Sspai{}
	case "tieba":
		return &Tieba{}
	case "toutiao":
		return &Toutiao{}
	case "weibo":
		return &Weibo{}
	case "zhihu":
		return &Zhihu{}
	default:
		return nil
	}
}
