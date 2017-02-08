package main

import ("fmt"; "net/http"; "strings"; "io/ioutil"; "os")

var extension = map[string]string{
	"ip1": "ipr/%s",
	"ip2": "ipr/malware/%s",
	"hash": "malware/%s",
	"vuln": "/vulnerabilities/search/%s",
	"dns": "resolve/%s",
	"url": "url/%s",
	"url_mal": "url/malware/%s",
	"whois": "whois/%s"}

var BaseUrl = "https://api.xforce.ibmcloud.com/"

func main(){
	
	if len(os.Args) < 3 {
		fmt.Println("Copyright 2016,2017 - Joerg Stephan (@johest)\n"+
			"XforceExchange script needs at least 3 arguments\n"+
		"\tIndicator Type - see below\n" +
		"\tOBJECT - hash/ip/stdcode\n"+
		"\tAPIKEY - your apkikey\n"+
		"\n"+
		"Indicator types - XFE url extension")
		
		for k, v := range extension{
			fmt.Println("\t", k, "  - ", v)
		} 
	}else{
			fmt.Printf(apicall(os.Args[1],os.Args[2], os.Args[3]))
}
}

func apicall(inditype string, query string, apikey string) string{
	client := &http.Client{}

req, err := http.NewRequest("GET", BaseUrl+strings.Replace(extension[inditype], "%s", query, -1), nil)
req.Header.Add("Authorization","Basic "+apikey)
req.Header.Add("Accept", "application/json")
req.Header.Add("User-Agent", "Mozilla 5.0")
	
resp, err := client.Do(req)
if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}
bodyText, err := ioutil.ReadAll(resp.Body)
if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}
s := string(bodyText)
return s
}