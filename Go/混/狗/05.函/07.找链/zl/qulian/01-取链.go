/* 从网页提取网址
 */
package qulian // 取链

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func H取链(网址 string) ([]string, error) {
	回, 障 := http.Get(网址)

	if 障 != nil {
		return nil, 障
	}

	if 回.StatusCode != http.StatusOK {
		回.Body.Close()
		return nil, fmt.Errorf("取%s障%s", 网址, 回.Status)
	}

	网页, 障 := html.Parse(回.Body)
	回.Body.Close()

	if 障 != nil {
		return nil, fmt.Errorf("析%s障%v", 网址, 障)
	}

	var 链集 []string
	访节 := func(节 *html.Node) { // 匿函
		if 节.Type == html.ElementNode && 节.Data == "a" {
			for _, 锚 := range 节.Attr {
				if 锚.Key != "href" {
					continue
				}
				链, 障 := 回.Request.URL.Parse(锚.Val) // 绝对链址
				if 障 != nil {
					continue
				}
				链集 = append(链集, 链.String())
			}
		}
	}

	逐节(网页, 访节, nil)

	return 链集, nil
}

func 逐节(节点 *html.Node, 预备, 善后 func(节点 *html.Node)) {
	if 预备 != nil {
		预备(节点)
	}
	for 子 := 节点.FirstChild; 子 != nil; 子 = 子.NextSibling {
		逐节(子, 预备, 善后)
	}
	if 善后 != nil {
		善后(节点)
	}
}
