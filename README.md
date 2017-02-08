# XFE
Various IBM X-Force Exchange modules
https://exchange.xforce.ibmcloud.com/new

## python

```In [8]: print(xfexchange.info())

Copyright 2016,2017 - Joerg Stephan (@johest)
X-Force Exchange Python module
xforceexchange.apicall(Indicator Type, Indicator, Apikey)

Indicator can be any : IP / hash(md5,sha) / stdcode (cve)

Please create your Apikey on https://exchange.xforce.ibmcloud.com
	
Indicator Type: (%s = Object)
	 url_mal - url/malware/%s
	 hash - malware/%s
	 ip2 - ipr/malware/%s
	 dns - resolve/%s
	 url - url/%s
	 whois - whois/%s
	 vuln - vulnerabilities/search/%s
	 ip1 - ipr/%s
```

## golang

```$ go run xforceexcchange.go 
Copyright 2016,2017 - Joerg Stephan (@johest)
XforceExchange script needs at least 3 arguments
	Indicator Type - see below
	OBJECT - hash/ip/stdcode
	APIKEY - your apkikey

Indicator types - XFE url extension
	 url   -  url/%s
	 url_mal   -  url/malware/%s
	 whois   -  whois/%s
	 ip1   -  ipr/%s
	 ip2   -  ipr/malware/%s
	 hash   -  malware/%s
	 vuln   -  /vulnerabilities/search/%s
	 dns   -  resolve/%s
```
