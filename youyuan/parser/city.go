package parser

import (
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/public"
	"regexp"
)


var (
	cityRe = regexp.MustCompile(`<a href="(/\d+-profile/)" target="_blank"><strong title="[^"]+">([^<]+)</strong></a>`)
	cityPageGGRe = regexp.MustCompile(`<a href="([^"]+[g]{2}18[^"]+)">[^<]+</a>`)
	cityPageMMRe = regexp.MustCompile(`<a href="([^"]+[m]{2}18[^"]+)">[^<]+</a>`)
)

func CityParser(contents []byte, gender string,workAddress string ) engine.ParseResult {
	// 根据正则表达式获取城市页面的所有 人的连接
	cityMatches := cityRe.FindAllSubmatch(contents,-1)
	// 首先定义返回的数据
	result := engine.ParseResult{}
	// 对正则表达式获取的所有连接进行存储, ?注意, 这里只查询了第一页
	for _, m := range cityMatches {
		url := public.UrlYouYuan+string(m[1])
		name := string(m[2])
		req := engine.Request{
			Url:        url ,
			// 这里在继续对获取到的用户进行爬取
			ParserFunc: func(contents []byte) engine.ParseResult{
				return ProfileParser(contents, url, name ,gender,workAddress)
			},
		}
		result.Requests = append(result.Requests, req)
		//items = append(items ,m[2])
	}

	// 对翻页进行爬取
	cityPageGGMatches := cityPageGGRe.FindAllSubmatch(contents, -1)
	cityPageMMMatches := cityPageMMRe.FindAllSubmatch(contents, -1)
	for _, m := range cityPageMMMatches {
		//fmt.Printf("Regexp Get >>> cityPage Url: %s ",m[1])

		req := engine.Request{
			Url:        public.UrlYouYuan+string(m[1]),
			// 这里在继续对获取的下一页 进行 CityParser, 因为下一页还是城市信息
			ParserFunc: func(contents []byte) engine.ParseResult{
				return CityParser(contents, public.YouYuanMM,workAddress)
			},
		}
		result.Requests = append(result.Requests, req )
	}

	for _, m := range cityPageGGMatches {
		//fmt.Printf("Regexp Get >>> cityPage Url: %s ",m[1])
		req := engine.Request{
			Url:        public.UrlYouYuan+string(m[1]),
			// 这里在继续对获取的下一页 进行 CityParser, 因为下一页还是城市信息
			ParserFunc: func(contents []byte) engine.ParseResult{
				return CityParser(contents, public.YouYuanGG,workAddress)
			},
		}
		result.Requests = append(result.Requests, req )
	}
	return result
}
