package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	)

func main() {
	engine.Run(engine.Request{
		Url:"http://city.zhenai.com/",
		ParserFunc:parser.ParseCityList,
	})
}

