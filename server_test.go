/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-22T02:10:22+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T02:36:15+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewServer(NewRequestLog())
	testServer := httptest.NewServer(http.HandlerFunc(server.handleRequest))
	defer testServer.Close()

	http.Get(testServer.URL)
	http.Get(testServer.URL)
	response, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	jsonResponse, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	responseObj := serverResponse{}
	json.Unmarshal(jsonResponse, &responseObj)
	fmt.Printf("Response JSON: %s", string(jsonResponse))
	fmt.Printf("Response object: %+v", responseObj)

	if responseObj.MinuteHitsCount != 3 {
		t.Errorf("Expected to see 3 hits but was %d", responseObj.MinuteHitsCount)
	}
}
