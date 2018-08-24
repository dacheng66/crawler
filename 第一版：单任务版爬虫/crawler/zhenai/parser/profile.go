package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/model"
	"strconv"
	)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([0-9]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([0-9]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte,name string) engine.ParseResult {
	 profile := model.Profile{}
	 age,err := strconv.Atoi(extractString(contents,ageRe))
	 if err == nil{
	 	profile.Age = age
	 }

	 height,err := strconv.Atoi(extractString(contents,heightRe))
	 if err == nil{
		profile.Height = height
	 }

	 weight,err := strconv.Atoi(extractString(contents,weightRe))
	 if err == nil{
		profile.Weight = weight
	 }

	 profile.Name = name
	 profile.Gender = extractString(contents,genderRe)
	 profile.Income = extractString(contents,incomeRe)
	 profile.Marriage = extractString(contents,marriageRe)
  	 profile.Education = extractString(contents,educationRe)
	 profile.Occupation = extractString(contents,occupationRe)
	 profile.Hokou = extractString(contents,hokouRe)
	 profile.Xinzuo = extractString(contents,xinzuoRe)
	 profile.House = extractString(contents,houseRe)
	 profile.Car = extractString(contents,carRe)

	 result := engine.ParseResult{
	 	Items:[]interface{}{profile},
	 }
	 return result
}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match)>=2 {
		return string(match[1])
	}else {
		return ""
	}
}