package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/getsentry/sentry-go"
)

type DSN struct {
	host      string
	rawurl    string
	key       string
	projectId string
}

func NewDSN(rawurl string) *DSN {
	key := strings.Split(rawurl, "@")[0][7:]

	uri, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	idx := strings.LastIndex(uri.Path, "/")
	if idx == -1 {
		sentry.CaptureException(errors.New("missing projectId in dsn"))
		log.Fatal("missing projectId in dsn")
	}
	projectId := uri.Path[idx+1:]

	var host string
	if strings.Contains(rawurl, "ingest.sentry.io") {
		domain := strings.Split(rawurl, "@")[1]
		prefix := strings.Split(domain, ".ingest")[0]
		host = fmt.Sprintf("%v.ingest.sentry.io", prefix)
	}
	if strings.Contains(rawurl, "@localhost:") {
		host = "localhost:9000"
	}
	if host == "" {
		sentry.CaptureException(errors.New("missing host"))
		log.Fatal("missing host")
	}
	if len(projectId) < 6 {
		sentry.CaptureException(errors.New("bad project Id in dsn" + projectId))
		log.Fatal("bad project Id in dsn")
	}
	if projectId == "" {
		sentry.CaptureException(errors.New("missing project Id"))
		log.Fatal("missing project Id")
	}
	return &DSN{
		host,
		rawurl,
		key,
		projectId,
	}
}

func (d DSN) storeEndpoint() string {
	fullurl := fmt.Sprintf("https://%v/api/%v/store/?sentry_key=%v&sentry_version=7", d.host, d.projectId, d.key[1:])
	if d.host == "localhost:9000" {
		fullurl = strings.Replace(fullurl, "http", "https", 1)
	}
	if fullurl == "" {
		sentry.CaptureException(errors.New("missing fullurl"))
		log.Fatal("missing fullurl")
	}
	return fullurl
}
