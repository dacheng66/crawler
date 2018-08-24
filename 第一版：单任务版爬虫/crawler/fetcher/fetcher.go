package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
		"golang.org/x/text/encoding"
		"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"bufio"
)

func Fetcher(url string)([]byte,error){

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//resp,err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func  determineEncoding(r *bufio.Reader)  encoding.Encoding{
	//bytes,err := bufio.NewReader(r).Peek(1024) //不使用这个原因是去除了1024个字节
	bytes,err := r.Peek(1024)
	if err !=nil{
		log.Printf("Fetcher error：%v",err)
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}