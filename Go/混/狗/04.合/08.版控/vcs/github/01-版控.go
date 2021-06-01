/* 很多网站的网页服务提供了故障跟踪
* 如Gitee， Github等
* 本狗宝采用网页服务：https://docs.github.com/en/rest/reference/search#search-issues
* 服务请求和返回结果以JSON表示
 */
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const Z障址 = "https://api.github.com/search/issues"

type Z障果 struct {
	Z障数 int   `json:"total_count"` //此total_count来自Github文档，在此确定名字对应，不分大小写
	Z障集 []*Z障 `json:"items"`
}

type Z障 struct {
	S数   int       `json:"number"`
	W网址  string    `json:"html_url"`
	B标题  string    `json:"title"`
	Z状态  string    `json:"state"`
	Y用户  *Z账户      `json:"User"`
	C创建日 time.Time `json:"create_at"`
	Z正文  string    `json:"body"` // Markdown 格式
}

type Z账户 struct {
	D登录名 string `json:"login"`
	W网址  string `json:"html_url"`
}

func X寻障(项集 []string) (*Z障果, error) {
	求 := url.QueryEscape(strings.Join(项集, " "))
	回, 障 := http.Get(Z障址 + "?q=" + 求)
	if 障 != nil {
		return nil, 障
	}

	if 回.StatusCode != http.StatusOK {
		回.Body.Close()
		return nil, fmt.Errorf("查询败:%s", 回.Status)
	}

	var 果 Z障果
	if 障 := json.NewDecoder(回.Body).Decode(&果); 障 != nil {
		回.Body.Close()
		return nil, 障
	}

	回.Body.Close()
	return &果, nil
}
