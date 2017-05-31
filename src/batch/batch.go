package batch

import (
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"lib"
	"net/http"
)

type Results struct {
	NumberOfResults string  `xml:"NumberOfResults"`
	DisplayPerPage  string  `xml:"DisplayPerPage"`
	DisplayFrom     string  `xml:"DisplayFrom"`
	APIVersion      string  `xml:"APIVersion"`
	Onsen           []Onsen `xml:"Onsen"`
}

type Onsen struct {
	OnsenName         string `xml:"OnsenName"`
	OnsenNameKana     string `xml:"OnsenNameKana"`
	OnsenID           string `xml:"OnsenID"`
	OnsenAddress      string `xml:"OnsenAddress"`
	Area              Area   `xml:"Area"`
	NatureOfOnsen     string `xml:"NatureOfOnsen"`
	OnsenAreaName     string `xml:"OnsenAreaName"`
	OnsenAreaNameKana string `xml:"OnsenAreaNameKana"`
	OnsenAreaID       string `xml:"OnsenAreaID"`
	OnsenAreaURL      string `xml:"OnsenAreaURL"`
	OnsenAreaCaption  string `xml:"OnsenAreaCaption"`
}

type Area struct {
	Region     string `xml:"Region"`
	Prefecture string `xml:"Prefecture"`
	LargeArea  string `xml:"LargeArea"`
	SmallArea  string `xml:"SmallArea"`
}

func ImportOnsenList() {
	for i := 1; i < 48; i++ {

		url := "http://jws.jalan.net/APICommon/OnsenSearch/V1/?key=aqr15a41839ced&pref="
		url += fmt.Sprintf("%02d", i)
		url += "0000&xml_ptn=1"

		req, _ := http.NewRequest("GET", url, nil)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		results := Results{}
		err := xml.Unmarshal(byteArray, &results)
		if err != nil {
			fmt.Println("ERROR:")
			fmt.Println(string(byteArray))
			return
		}

		db := lib.DbOpen()

		query := "INSERT INTO spa"
		query += "(name, address, url, effect)"
		query += "VALUES (? ,? ,? ,?)"

		for i := range results.Onsen {
			db.Query(query,
				results.Onsen[i].OnsenName,
				results.Onsen[i].OnsenAddress,
				results.Onsen[i].OnsenAreaURL,
				results.Onsen[i].NatureOfOnsen)
		}
		db.Close()
	}
}
