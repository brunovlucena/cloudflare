package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/brunovlucena/cloudflare/src/handler"
	"github.com/brunovlucena/cloudflare/src/util"
)

const (
	baseUrl = "https://api.cloudflare.com/client/v4/zones"
)

var (
	authEmail = os.Getenv("CLOUDFLARE_AUTH_EMAIL")
	authKey   = os.Getenv("CLOUDFLARE_AUTH_KEY")
	domains   = map[string]string{
		"lesara.be":    "",
		"lesara.se":    "",
		"lesara.dk":    "",
		"lesara.co.uk": "",
		"lesara.de":    "",
		"lesara.it":    "",
		"lesara.at":    "",
		"lesara.ch":    "",
		"lesara.fr":    "",
		"lesara.nl":    "",
	}
)

func main() {
	getListZoneIds()
}

func makeRequest(method, url string, reader io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, reader)
	req.Header.Add("X-Auth-Email", authEmail)
	req.Header.Add("X-Auth-key", authKey)
	req.Header.Add("Content-Type", "application/json")
	return req
}

func getListZoneIds() {
	client := &http.Client{}
	method := "GET"
	url := baseUrl
	// make request
	req := makeRequest(method, url, nil)
	resp, err := client.Do(req)
	// check
	util.ParseErr(err)
	// read
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	var response handler.Response
	json.Unmarshal(buf.Bytes(), &response)
	if response.Success {
		i := 0
		for key, _ := range domains {
			domains[key] = response.Result[i].Id
			fmt.Println(domains[key])
			i++
		}
	} else {
		fmt.Println("Bang! Wrong credentials.")
	}
}
