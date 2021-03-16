package public



const (
	UrlYouYuan = "http://www.youyuan.com"
	YouYuanGG = "男"
	YouYuanMM = "女"

	//elastc search
	EsUrl ="http://172.16.76.220:9200"
	EsIndex = "crawler"
	EsType = "youyuan"
)

var (
	Duplicated = make(map[string]bool)
)
