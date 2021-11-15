package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
)

type Request struct {
	Payload       []byte
	StoreEndpoint string
	Kind          string
	Platform      string
}

func NewRequest(event Event) *Request {
	r := new(Request)

	var bodyBytes []byte
	var err error
	r.Kind = "ERROR"
	bodyBytes, err = json.Marshal(event.Error)

	if err != nil {
		sentry.CaptureException(err)
		fmt.Println(err)
	}

	r.Payload = bodyBytes
	r.StoreEndpoint = event.DSN.storeEndpoint()

	if r.StoreEndpoint == "" || r.Payload == nil {
		sentry.CaptureException(errors.New("missing StoreEndpoint or Payload"))
		fmt.Println("missing StoreEndpoint or Payload")
	}
	return r
}

func (r Request) send() {
	time.Sleep(300 * time.Millisecond)

	var payload []byte
	size := len(r.Payload)
	HUNDRED_KILOBYTES := 100000
	if size > HUNDRED_KILOBYTES {
		var buf bytes.Buffer
		gw := gzip.NewWriter(&buf)
		_, err := gw.Write(r.Payload)
		if err != nil {
			log.Fatal(err)
		}
		err = gw.Close()
		if err != nil {
			log.Fatal(err)
		}
		payload = buf.Bytes()
	} else {
		payload = r.Payload
	}

	request, errNewRequest := http.NewRequest("POST", r.StoreEndpoint, bytes.NewReader(payload)) // &buf
	if errNewRequest != nil {
		sentry.CaptureException(errNewRequest)
		fmt.Println(errNewRequest)
	}

	request.Header.Set("content-type", "application/json")
	if size > HUNDRED_KILOBYTES {
		request.Header.Set("Content-Encoding", "gzip")
	}

	var httpClient = &http.Client{}
	response, requestErr := httpClient.Do(request)

	if response.StatusCode != 200 {
		sentry.CaptureException(errors.New("status code was not 200, it was" + strconv.Itoa(response.StatusCode)))
	}
	if requestErr != nil {
		sentry.CaptureException(requestErr)
		fmt.Println(requestErr)
	}
	responseData, responseDataErr := ioutil.ReadAll(response.Body)

	if responseDataErr != nil {
		sentry.CaptureException(responseDataErr)
		fmt.Println(responseDataErr)
	}
	fmt.Printf("> Kind: %v | %v | Response: %v \n", r.Kind, r.Platform, string(responseData))
}
