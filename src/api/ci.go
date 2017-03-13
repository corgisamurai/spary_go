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
	url := "https://slack.com/api/channels.history?token=xoxp-113726990690-113821933188-153905532965-a2ced34460c2639621b8aba10906496a&channel=C4D8D3XMX&count=1&pretty=1"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	resJson, _ := ioutil.ReadAll(resp.Body)

	result := "0"
	if strings.Contains(string(resJson), "Failed:") {
		result = "1"
	}
	fmt.Fprintf(w, result)
}
