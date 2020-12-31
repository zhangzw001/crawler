package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)
var (
	cityRe = regexp.MustCompile(`<a href="(/\s+/)">[^<]+</a>`)
)

func Fetch(url string) ([]byte, error ) {

	// 方法1 直接通过http.get
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()

	// 方法2 自定义client
	client := &http.Client{}
	request , err := http.NewRequest(http.MethodGet,url,nil )
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent" ,"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解决字符集的问题
	// 1. 直接转码
	//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	//return ioutil.ReadAll(utf8Reader)
	// 2. 通过获取body的内容知道是什么编码, 但是有可能这里写的也不一定是正确的
	// 3.
	// 3.1 读取bufio, 不直接ReadAll 是因为 ReadAll 完之后, body的数据就没了
	buf := bufio.NewReader(resp.Body)
	// 3.2 然后通过包装的determineEncoding  分析出编码格式
	e := determineEncoding(buf)
	// 3.3 将resp.body转换成对应的编码格式
	utf8Reader := transform.NewReader(buf,e.NewDecoder())
	// 3.4 直接返回转换过的内容
	return ioutil.ReadAll(utf8Reader)

}

// 对编码进行分析
func determineEncoding(buf *bufio.Reader) encoding.Encoding {
	contents,err  := buf.Peek(1024)
	if err != nil {
		return  unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(contents,"")
	return e
}
