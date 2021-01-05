package parser

import (
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/public"
	"regexp"
)



var (
	cityListRe = regexp.MustCompile(`<a href="(/\S+/)">([^<]+)</a>`)
)

func CityListParser(contents []byte) engine.ParseResult {
	// 对内容进行正则表达式匹配
	// 这里获取的城市列表页, 都是每个城市默认的 女
	// 这里只有在城市的首页才能取到例如  北京男士征婚/北京女士征婚, 所以必须从这里开始传 给ProfileParser
	cityListMatches := cityListRe.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	//var items []engine.Item
	for _, m := range cityListMatches{
		// 这里必须定义遍历, 如果直接传string(m[2]) 给后面, 等到函数执行的时候, 传入的值都不会变了, 都是同一个地址
		workAddress := string(m[2])

		//fmt.Printf("Regexp Get >>> city Url: %s, city Name: %s \n",m[1],m[2])
		// 添加正则匹配到的url到 Request列表
		req := engine.Request{
			Url:        public.UrlYouYuan+string(m[1]),
			// 这里在继续对获取到的每个城市继续爬取, 爬取到用户的连接
			ParserFunc: func(contents []byte) engine.ParseResult{
				return CityParser(contents, public.YouYuanMM,workAddress)
			},
		}
		result.Requests = append(result.Requests, req)
		//items = append(items, string(m[2]))
	}
	return result
}
