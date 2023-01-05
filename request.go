package wechat

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func ClientPost(url string, in interface{}) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		var res *http.Response

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		jsonBody, _ := json.Marshal(in)
		res, err = client.Post(url, "application/json", strings.NewReader(string(jsonBody)))

		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
	}()
	return func() ([]byte, error) {
		<-c
		return body, err
	}
}

func ClientGet(url string) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		var res *http.Response

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		res, err = client.Get(url)

		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
	}()
	return func() ([]byte, error) {
		<-c
		return body, err
	}
}
