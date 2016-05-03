package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	f "github.com/MustWin/gomarathon"
	"os"
	"strings"
)

var (
	url = "https://shipped-tx3-control-01.tx3.shipped-cisco.com/marathon"
	user       = "synthetic-mon"
	pwd        = ""
	_seperator = "===================================================="
)

var (
	Args struct {
		BaseUrl string
		Host    string
		User    string
		PWD    string
		Port    int
		Appid   string
	}
)

func GetArgs() {
	flag.StringVar(&Args.BaseUrl, "url", url, "Provide marathon api url")
	flag.StringVar(&Args.User, "u", user, "Provide Login Username")
	flag.StringVar(&Args.PWD, "p", "", "Provide Login password")
	flag.StringVar(&Args.Host, "host", "", "Provide Host name need to search")
	flag.IntVar(&Args.Port, "port", 0, "Provide Port  Number to search")
	flag.StringVar(&Args.Appid, "id", "", "Provide Application Id to search")
	flag.Parse()

}

func main() {
	GetArgs()
	fmt.Println(_seperator)
	run()
	fmt.Println(_seperator)

}
func run() {
	auth := f.HttpBasicAuth{Args.User, Args.PWD}
	c, err := f.NewClient(Args.BaseUrl, &auth, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("Error MAIN : " + err.Error())
	}
	if len(Args.Host) > 0 {
		if Args.Port == 0 {
			fmt.Printf("\n Error: Please Provide Port number for '%s' Usage s=HOST_ID p=PORT_NUM ", Args)
			return
		}
		//fmt.Println("DEBUG port-- ", port)
		flag := false
		resp, err := c.ListTasks()
		if err != nil {
			fmt.Println("Error MAIN : " + err.Error())
		}
		//fmt.Println(resp.Tasks[0].AppID)
		for _, t := range resp.Tasks {
			//fmt.Println("DEBUG HOST ", t.Host)
			if t.Host == Args.Host {
				for _, p := range t.Ports {
					//fmt.Println("DEBUG port ", p)
					if p == Args.Port {
						resp, err := c.GetApp(t.AppID)
						if err != nil {
							return
						}
						fmt.Println("App Id : " + resp.App.ID)
						for k, v := range resp.App.Labels {
							fmt.Println(k+" : ", v)
						}
						flag = true
					}
				}
			}
		}
		if !flag {
			fmt.Println("No record found.")
		}
	} else if len(Args.Appid) > 0 {

		appid := strings.Replace(Args.Appid, "/", "", -1)
		resp, err := c.GetApp(appid)
		if err != nil {
			fmt.Println("Invalid Details : "+err.Error())
			return
		}

		fmt.Println("App Id : " + resp.App.ID)
		for k, v := range resp.App.Labels {
			fmt.Println(k+" : ", v)
		}
		for i, tsk := range resp.App.Tasks {
			fmt.Println(fmt.Sprintf("Host_%d : %s", i, tsk.Host))
			fmt.Println(fmt.Sprintf("Port_%d : %d", i, tsk.Ports[0]))
		}

	} else {
		os.Exit(0)
	}
}
