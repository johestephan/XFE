import requests


BASEurl = "https://api.xforce.ibmcloud.com/"

extensions = {"ip1": "ipr/%s",
			  "ip2": "ipr/malware/%s",
			  "hash": "malware/%s",
			  "vuln": "vulnerabilities/search/%s",
			  "dns": "resolve/%s",
			  "url": "url/%s",
			  "url_mal": "url/malware/%s",
			  "whois": "whois/%s"}


def MyHeader(key=False):
	global limit
	if key is False:
		return None
	
	return {"Authorization": "Basic %s " % key,
		   "Accept": "application/json",
		   'User-Agent': 'Mozilla 5.0'}


	
def apicall(indicator_type, indicator, key=False):
	try:
		myURL = BASEurl + (extensions[str(indicator_type)])%indicator
		jsondata = requests.get(myURL, headers=MyHeader(key)).json()
	except:
		return None
	return jsondata

def info():
	print('''
Copyright 2016,2017 - Joerg Stephan (@johest)
X-Force Exchange Python module
xforceexchange.apicall(Indicator Type, Indicator, Apikey)

Indicator can be any : IP / hash(md5,sha) / stdcode (cve)

Please create your Apikey on https://exchange.xforce.ibmcloud.com
	
Indicator Type: (%s = Object)''')
	for item in extensions:
		print("\t %s - %s" % (item ,extensions[item]))


