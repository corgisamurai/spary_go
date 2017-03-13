package api

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

type Slacks struct {
  Slacks []Slack `json:"spa"`
}

type Slack struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func Ci(w http.ResponseWriter, r *http.Request) {
	url := "https://slack.com/api/channels.history?token=xoxp-113726990690-113821933188-153829293618-d5bb3d12cc6635ecbfd7e62c230b30f9&channel=C4D8D3XMX&count=1&pretty=1"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	resJson, _ := ioutil.ReadAll(resp.Body)

	text := "Build Success"
	color := "green"
	if strings.Contains(string(resJson), "Failed") {
		text = "Build Failed"
  	color = "red"
	}
	fmt.Println(string(resJson)) // debug
	html := "<html><body style='font-size: 100px;text-align: center; margin-top: 100px;background-color:" + color + ";color:white;'>" + text + "</body></html>"
	fmt.Fprintf(w, html)
}
