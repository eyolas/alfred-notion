package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	NotionUrl = "https://www.notion.so"
	url       = "https://www.notion.so/api/v3/search"
)

func callNotion(search string, spaceID string, cookie string) ([]*AlfredLink, error) {
	query := NewSearchQuery(search, spaceID)
	q, _ := json.Marshal(query)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(q))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)

	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(requestDump))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("eror: %#v", err)
		return nil, err
	}
	defer response.Body.Close()

	// responseDump, err := httputil.DumpResponse(response, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(responseDump))

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject SearchResult
	json.Unmarshal(responseData, &responseObject)

	return responseObject.AlfredLinks(NotionUrl), nil
}
