package fetcher

import (
	"testing"
)

func TestFetch(t *testing.T) {
	const url = "http://www.youyuan.com/find/beijing/gg18-0/advance-0-0-0-0-0-0-0/p1/"

	contents, err := Fetch(url)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%s",contents)
}
