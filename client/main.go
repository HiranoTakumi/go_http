package main

import (
	"log"
	"net/http"
	// "io/ioutil"
	"encoding/json"
	"os"
	"time"
)

var (
	timeout_second = 10
	client         = newClient(Timeout(timeout_second))
)

type httpClientOpts struct {
	Timeout time.Duration
}

type option func(*httpClientOpts)

func Timeout(t int) option {
	return func(h *httpClientOpts) {
		h.Timeout = time.Duration(t) * time.Second
	}
}

type Response struct {
	SleepTime int `json:"sleepTime"`
}

func newClient(opts ...option) *http.Client {
	h := &httpClientOpts{
		Timeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(h)
	}
	return &http.Client{
		Timeout: h.Timeout,
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("ERROR: 引数を指定してください。")
	}
	url := "http://localhost:3000/"
	second := os.Args[1]
	url = url + "?second=" + second

	client := &http.Client{}
	// client := newClient(Timeout(timeout_second))

	req, e := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if e != nil {
		log.Println(e)
		return
	}
	req.SetBasicAuth("id", "password")

	resp, e := client.Do(req)
	if e != nil {
		log.Println(e)
		log.Println(resp)
		return
	}
	defer resp.Body.Close()

	var output Response
	// e := json.Unmarshal(resp.Body, &output)
  e = json.NewDecoder(resp.Body).Decode(&output)
	if e != nil {
		log.Println(e.Error())
	} else {
		log.Println(output)
	}
}
