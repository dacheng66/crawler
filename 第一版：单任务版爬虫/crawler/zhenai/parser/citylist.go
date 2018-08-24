package parser

import (
	"crawler/engine"
	"regexp"
	"fmt"
)

const  cityListRe  = `<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte)engine.ParseResult  {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	limit := 3
	result := engine.ParseResult{}
	for _,m := range matches{
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
		limit --
		if limit == 0{
			break
		}
		//fmt.Printf("URL:%s , ADDRESS:%s\n",m[1],m[2])
	}
	fmt.Printf("Matches found:%d\n",len(matches))
	return result
}
