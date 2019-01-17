package fetcher

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)

/**
fetcher：根据url获取对应的数据
 */

func Fetch(url string) ([] byte, error) {
	/*
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
*/
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp, err := http.DefaultClient.Do(request)
	//fmt.Println("resp:", resp,",err:",err)
	if resp == nil {
		fmt.Println("resp:",resp)
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp:", resp)
		return nil, fmt.Errorf("error:status code:%d", resp.StatusCode)
	}

	//如果页面传来的不是utf8，我们需要转为utf8格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Ftcher error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
