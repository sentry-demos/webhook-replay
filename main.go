package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	all        *bool
	ignore     *bool
	httpClient *http.Client

	local           *bool
	DSN_PROJECT     string
	DSN_JOB_MONITOR string
)

func init() {
	// parseYamlConfig()

	// initializeSentry()
	// sentry.CaptureMessage("job started")

	// ignore = flag.Bool("i", false, "ignore sending the event to Sentry.io")
	// n = flag.Int("n", 25, "default number of events to read from a source")

	// defaultPrefix := "error"
	// filePrefix = flag.String("prefix", defaultPrefix, "file prefix")
	// flag.Parse()

	// httpClient = &http.Client{}
}

func main() {
	errEnv := godotenv.Load()
	// This is where your webhook info gets sent to in Sentry as events
	DSN_PROJECT = os.Getenv("DSN_PROJECT") // ido-native in testorg-az

	local = flag.Bool("local", false, "local development web server")
	flag.Parse()
	if errEnv != nil && *local == true {
		log.Print("no .env or environment")
	}
	if *local == true {
		fmt.Println("local development web server localhost:8000")
		http.HandleFunc("/", HandleLocalRequest)
		http.ListenAndServe(":8000", nil)
	} else {
		fmt.Println("aws lambda environment")
		lambda.Start(HandleLambdaRequest)
	}

}

func HandleLocalRequest(writer http.ResponseWriter, req *http.Request) {
	// TODO
	// prettyPrint(req.Body)
	decoder := json.NewDecoder(req.Body)
	var webhook Webhook
	err := decoder.Decode(&webhook)
	if err != nil {
		sentry.CaptureException(errors.New("unable to decode webhook"))
		log.Fatal("unable to decode webhook")
	}
	App(webhook)
	fmt.Fprintf(writer, "%v \n", "done")
}

func HandleLambdaRequest(ctx context.Context, webhook Webhook) (string, error) {
	defer sentry.Flush(2 * time.Second)
	App(webhook)
	return fmt.Sprintf("App() complete"), nil
}

func App(webhook Webhook) {
	// TODO
	prettyPrint(webhook.Data.Error["title"].(string))

	// data := webhook["Data"].(map[string]interface{}) // does not support indexing

	// WORKS, this printed
	// data := webhook.Data
	// fmt.Println("\n> App done", data)
}

func prettyPrint(v interface{}) {
	pp, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(pp))
}
