package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var resquests []Request
	for _,r := range seeds{
		resquests = append(resquests,r)
	}

	for len(resquests) > 0 {
		r := resquests[0]
		resquests = resquests[1:]
		log.Printf("Fetching %s",r.Url)
		body,err := fetcher.Fetcher(r.Url)
		if err != nil {
			log.Printf("Fetcher:error" + "fetching url %s:%v",r.Url,err)
			continue
		}
		ParseResult := r.ParserFunc(body)
		resquests = append(resquests,ParseResult.Requests...)
		for _,item := range ParseResult.Items{
			log.Printf("Got item %v",item)
		}
	}
}