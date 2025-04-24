# HotSearch

## 文档导航

- [功能概述](#功能概述)
- [运行命令](#运行命令)
- [接口说明](#接口说明)
  - [获取单个来源的热搜数据](#1-获取单个来源的热搜数据)
  - [获取所有可用的热搜来源](#2-获取所有可用的热搜来源)
  - [获取所有来源的热搜数据](#3-获取所有来源的热搜数据)
- [数据结构](#数据结构)
  - [响应封装结构](#1-响应封装结构)
  - [数据模型](#2-数据模型)
- [注意事项](#注意事项)
- [支持的热榜来源](#支持的热榜来源)
- [API 文档](#api-文档)
- [项目引用说明](#项目引用说明)
- [文档导航](#文档导航)

### 功能概述

这是一个基于Go语言实现的API接口，用于获取多个新闻平台的热搜数据。该接口能够实时抓取各大新闻网站的热搜内容，并返回相关的详细信息，提供统一的API服务。

### 运行命令

```bash
go run main.go
```

服务将在 `http://127.0.0.1:7490` 启动。

### 接口说明

#### 1. 获取单个来源的热搜数据

```
GET http://127.0.0.1:7490/api/news?source=:source
```

##### 请求参数

| 参数   | 类型   | 说明                                                         | 是否必填 |
| ------ | ------ | ------------------------------------------------------------ | -------- |
| source | string | 热榜来源平台标识，支持以下值：<br>`baidu`、`bilibili`、`douyin`、`kuaishou`、`pengpai`、`qqnews`、`sina`、`sougou`、`sspai`、`tieba`、`toutiao`、`weibo`、`zhihu` | 是       |

##### 响应示例

成功响应：
```json
{
    "code": 0,
    "data": {
        "source": "贴吧热议榜", 
        "update_time": "2024-12-04 00:31:39",  
        "hot_list": [
            {
                "index": 1,
                "title": "新闻标题",
                "description": "新闻描述",
                "image": "图片URL",
                "popularity": "2824170",
                "url": "新闻链接"
            },
            ...
            {
                "index": 30,
                "title": "新闻标题",
                "description": "新闻描述",
                "image": "图片URL",
                "popularity": "8428",
                "url": "新闻链接"
            }
        ]
    },
    "msg": "success"
}
```

失败响应：
```json
{
    "code": 7,
    "data": {
        "source": "",
        "update_time": "",
        "hot_list": null
    },
    "msg": "source not found"
}
```

#### 2. 获取所有可用的热搜来源

```
GET http://127.0.0.1:7490/api/sources
```

##### 响应示例
```json
{
    "code": 0,
    "data": [
        "baidu",
        "bilibili",
        "douyin",
        "kuaishou",
        "pengpai",
        "qqnews",
        "sina",
        "sougou",
        "sspai",
        "tieba",
        "toutiao",
        "weibo",
        "zhihu"
    ],
    "msg": "success"
}
```

#### 3. 获取所有来源的热搜数据

```
GET http://127.0.0.1:7490/api/all
```

##### 响应示例
```json
{
    "code": 0,
    "data": {
        "baidu": {
            "source": "百度热搜",
            "update_time": "2024-03-21 10:00:00",
            "hot_list": [...]
        },
        "bilibili": {
            "source": "B站排行榜",
            "update_time": "2024-03-21 10:00:00",
            "hot_list": [...]
        },
        // ... 其他来源的数据
    },
    "msg": "success"
}
```

### 数据结构

#### 1. 响应封装结构

```go
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

// Response 通用响应结构
type Response struct {
	Code int                 `json:"code"`   // 响应状态码
	Data model.HotSearchData `json:"data"`   // 热榜数据
	Msg  string              `json:"msg"`    // 响应消息
}

// SourceListResponse 来源列表响应结构
type SourceListResponse struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
	Msg  string   `json:"msg"`
}

// AllHotSearchResponse 所有热搜数据响应结构
type AllHotSearchResponse struct {
	Code int                     `json:"code"`
	Data map[string]model.HotSearchData `json:"data"`
	Msg  string                 `json:"msg"`
}

// OkWithData 成功响应
func OkWithData(w http.ResponseWriter, data model.HotSearchData) {
	_ = json.NewEncoder(w).Encode(Response{
		Code: SUCCESS,
		Data: data,
		Msg:  "success",
	})
}

// OkWithSourceList 成功响应来源列表
func OkWithSourceList(w http.ResponseWriter, sources []string) {
	_ = json.NewEncoder(w).Encode(SourceListResponse{
		Code: SUCCESS,
		Data: sources,
		Msg:  "success",
	})
}

// OkWithAllHotSearch 成功响应所有热搜数据
func OkWithAllHotSearch(w http.ResponseWriter, data map[string]model.HotSearchData) {
	_ = json.NewEncoder(w).Encode(AllHotSearchResponse{
		Code: SUCCESS,
		Data: data,
		Msg:  "success",
	})
}

// Failed 错误响应
func Failed(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(Response{
		Code: ERROR,
		Data: model.HotSearchData{},
		Msg:  err.Error(),
	})
}
```

#### 2. 数据模型

```go
package model

// HotItem 热点新闻项
type HotItem struct {
	Index       int    `json:"index"`       // 排名
	Title       string `json:"title"`       // 标题
	Description string `json:"description"` // 描述
	Image       string `json:"image"`       // 图片
	Popularity  string `json:"popularity"`  // 热度
	URL         string `json:"url"`         // 链接
}

// HotSearchData 热榜数据结构
type HotSearchData struct {
	Source     string    `json:"source"`      // 数据源
	UpdateTime string    `json:"update_time"` // 更新时间
	HotList    []HotItem `json:"hot_list"`    // 热搜列表
}
```

### 注意事项

1. 除澎湃热榜返回20条数据外，其他平台返回30条数据。
2. 不是所有平台的返回数据都包含 `description`、`image` 和 `popularity` 字段。

### 支持的热榜来源

下面是支持的 `source` 参数的具体列表，用户可以通过这些值来查询不同平台的热榜：

- `baidu`：百度热搜
- `bilibili`：Bilibili 排行榜
- `douyin`：抖音热榜
- `kuaishou`：快手热榜
- `pengpai`：澎湃热榜
- `qqnews`：腾讯热点榜
- `sina`：新浪热榜
- `sougou`：搜狗热搜
- `sspai`：少数派最热
- `tieba`：贴吧热议榜
- `toutiao`：头条热榜
- `weibo`：微博热搜
- `zhihu`：知乎热榜

## API 文档

### 1. 获取热搜列表

- 请求方式：GET
- 请求地址：`/api/news?source={source}`
- 请求参数：
  - source: 数据源名称，可选值：`douyin`, `bilibili`, `tieba`, `toutiao`, `pengpai`, `kuaishou`, `zhihu`, `sina`, `sspai`, `weibo`, `qqnews`, `sougou`
- 返回数据：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "source": "数据源名称",
      "updateTime": "更新时间",
      "list": [
        {
          "title": "热搜标题",
          "url": "热搜链接",
          "hot": "热度值"
        }
      ]
    }
  }
  ```

### 2. 获取所有数据源

- 请求方式：GET
- 请求地址：`/api/sources`
- 返回数据：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": [
      "douyin",
      "bilibili",
      "tieba",
      "toutiao",
      "pengpai",
      "kuaishou",
      "zhihu",
      "sina",
      "sspai",
      "weibo",
      "qqnews",
      "sougou"
    ]
  }
  ```

### 3. 获取所有数据源的热搜数据

- 请求方式：GET
- 请求地址：`/api/all`
- 返回数据：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "douyin": {
        "source": "抖音热榜",
        "updateTime": "更新时间",
        "list": [...]
      },
      "bilibili": {
        "source": "B站排行榜",
        "updateTime": "更新时间",
        "list": [...]
      },
      // ... 其他数据源
    }
  }
  ```

## 项目引用说明

### Go 项目引用

1. 添加依赖：
```bash
go get github.com/spatxos/HotSearch
```

2. 在代码中使用：
```go
import (
	"github.com/spatxos/HotSearch/api"
)

// 使用示例
func main() {
	var httpServer http.Server
    http.HandleFunc("/api/hotsearch/list", api.GetHotListHandler)
	http.HandleFunc("/api/hotsearch/sources", api.GetSourcesHandler)
	http.HandleFunc("/api/hotsearch/all", api.GetAllHotListHandler)
    ...
}
```


