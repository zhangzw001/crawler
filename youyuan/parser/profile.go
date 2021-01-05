package parser

import (
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/model"
	"github.com/zhangzw001/crawler/public"
	"regexp"
	"strconv"
)

var (
	//nickNameRe   = regexp.MustCompile(`<strong>([^<]+)</strong>`)
	hokouRe      = regexp.MustCompile(`<li>籍 贯：<span class="black">([^<]+)</span></li>`)
	ageRe        = regexp.MustCompile(`<p class="local">[^\d]+([\d]+)岁[^<]+</p><ol class="hoby">`)
	heightRe     = regexp.MustCompile(`<li>身高：<span class="black">([0-9]+)cm</span></li>`)
	weightRe     = regexp.MustCompile(`<li>体 重：<span class="black">([^<]+)斤</span></li>`)
	incomeRe     = regexp.MustCompile(`<li>月 薪：<span class="black">([^<]+)</span></li>`)
	marriageRe   = regexp.MustCompile(`<li>婚姻：<span class="black">([^<]+)</span></li>`)
	educationRe  = regexp.MustCompile(`<li>学 历：<span class="black">([^<]+)</span></li>`)
	occupationRe = regexp.MustCompile(`<li>职业：<span class="black">([^<]+)</span></li>`)
	houseRe      = regexp.MustCompile(`<li>住房：<span class="black">([^<]+)</span></li>`)
	sexRe        = regexp.MustCompile(`<li>能否接受婚前性行为：<span class="black">([^<]+)</span></li>`)
	idUrlRe      = regexp.MustCompile(`http://www.youyuan.com/(\d+)-profile/`)

	// 这里只能取到id
	p = regexp.MustCompile(`<li class="inPerson" data-kd="(\d+)"><a href="/login.html"`)
)



func ProfileParser(contents []byte, url string ,  name string, gender string  ,workAddress string ) engine.ParseResult {
	//fmt.Printf("gender : %s , url : %s\n",gender,url)
	profile := model.Profile{}
	// 整型字段处理
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	//

	// 昵称
	//profile.Name = extractString(contents, nickNameRe)
	profile.Name = name
	profile.Gender = gender
	profile.WorkAddress = workAddress

	profile.Hokou = extractString(contents, hokouRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.House = extractString(contents, houseRe)
	profile.Sex = extractString(contents, sexRe)


	//空
	profile.Car = ""

	//fmt.Printf("url: %s, id :%s\n",url, extractString([]byte(url),idUrlRe))
	result := engine.ParseResult{
		//Items: []interface{}{profile},
		Items: []engine.Item{
			{
				Url:     url,
				Type:    public.EsType,
				Id:      extractString([]byte(url),idUrlRe),
				Payload: profile,
			},
		},
	}
	//fmt.Printf("%v\n",result)
	return result
}


func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
