package batch

import (
  "encoding/xml"
  "fmt"
  "lib"
  "io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Results struct {
  NumberOfResults string `xml:"NumberOfResults"`
  DisplayPerPage string `xml:"DisplayPerPage"`
  DisplayFrom string `xml:"DisplayFrom"`
  APIVersion string `xml:"APIVersion"`
  Onsen []Onsen `xml:"Onsen"`
}

type Onsen struct {
  OnsenName string `xml:"OnsenName"`
  OnsenNameKana string `xml:"OnsenNameKana"`
  OnsenID string `xml:"OnsenID"`
  OnsenAddress string `xml:"OnsenAddress"`
  Area Area `xml:"Area"`
  NatureOfOnsen string `xml:"NatureOfOnsen"`
  OnsenAreaName string `xml:"OnsenAreaName"`
  OnsenAreaNameKana string `xml:"OnsenAreaNameKana"`
  OnsenAreaID string `xml:"OnsenAreaID"`
  OnsenAreaURL string `xml:"OnsenAreaURL"`
  OnsenAreaCaption string `xml:"OnsenAreaCaption"`
}

type Area struct {
  Region string `xml:"Region"`
  Prefecture string `xml:"Prefecture"`
  LargeArea string `xml:"LargeArea"`
  SmallArea string `xml:"SmallArea"`
}

func ImportOnsenList() {
  url := "http://jws.jalan.net/APICommon/OnsenSearch/V1/?key=aqr15a41839ced&l_area=010300&count=1&xml_ptn=1"
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
  defer db.Close()

  query := "INSERT INTO spa ("
  query += "name,"
  query += "address,"
  query += "url,"
  query += "effect"
  query += ") VALUES ("
  query += "?,"
  query += "?,"
  query += "?,"
  query += "?"
  query += ")"

  db.Query(query,
    results.Onsen[0].OnsenName,
    results.Onsen[0].OnsenAddress,
    results.Onsen[0].OnsenAreaURL,
    results.Onsen[0].NatureOfOnsen)
}
