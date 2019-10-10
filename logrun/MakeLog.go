package main

import (
	"flag"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

type resource struct {
	url string
	target string
	start int
	end int
}

func listUa() []string   {
	list := []string {
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	}

	return list
}

func ruleResource() [] resource{
	var res []resource

	r1 := resource{
		url: "http://localhost:8080",
		target:"",
		start : 0,
		end : 0,
	}
	r2 := resource{
		url: "http://localhost:8080/list/{%id}.html",
		target:"",
		start : 0,
		end : 21,
	}
	r3 := resource{
		url: "http://localhost:8080/movie/{%id}.html",
		target:"",
		start : 0,
		end : 12924,
	}

	res = append(append(append(res, r1), r2), r3)
	return res
}
func buildUrl(res []resource) [] string {
	var list []string


	for _,r := range res{
		if len(r.target) == 0{
			list = append(list,r.url)
		}else{
			for i := r.start;i<r.end;i++{
				urlStr := strings.Replace(r.url,r.target, string(i),-1)
				list = append(list,urlStr)
			}
		}

	}

	return list
}

func makeLog(currentUrl , referUrl , ua string) string {
	u := url.Values{}
	u.Set("time","1")
	u.Set("url",currentUrl)
	u.Set("refer",referUrl)
	u.Set("ua",ua)
	paramsStr := u.Encode()

	logTemplate := "127.0.0.1 - - [09/10/2019]   ?{$paramsStr} {$ua}"
	log := strings.Replace(logTemplate,"{$paramsStr}",paramsStr,-1)
	log = strings.Replace(log, "{$ua}", ua, -1)
	return log
}

func randInt (min,max int ) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func main() {

	total := flag.Int("total",100,"how many rows by create?")
	filePath := flag.String("filePath","f://dig.log","路径")
	flag.Parse()
	res := ruleResource()

	list := buildUrl(res)

	for i := 0; i<*total ; i++{
		currentUrl := list[randInt(0,len(list)-1)]
		var referUrl string
		var ua string

		referUrl = list[randInt(0,len(list)-1)]
		ualist := listUa()
		ua = ualist[randInt(0,len(ualist)-1)]
		/*if !((i+1)<*total) {
			referUrl = list[i-1]
		}else {
			referUrl = list[i+1]
		}

		if i > len(ualist){
			ua = ualist[len(ualist)]
		}else{

			ua = ualist[i]
		}*/

		logStr := makeLog(currentUrl, referUrl,ua ) +" \n"

		//ioutil.WriteFile(*filePath, []byte(logStr), 0644)

		fd,_ := os.OpenFile(*filePath,os.O_APPEND|os.O_RDWR|os.O_CREATE,0644)
		fd.Write([]byte(logStr))
		fd.Close()

	}


}


