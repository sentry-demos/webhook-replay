# Webhook Replay

VIDEO DEMO  
https://www.loom.com/share/d847771caeac40ac8e46e94d9a5f0bbd

## Setup & Run

You need 2 orgs in Sentry. Let's call them the origin and destination org. The webserver described here is something you're going to run, and its code is in main.go.

1. In the origin org, setup an Internal Integration with a URL that is the webhook-replay webserver hosted by ngrok. `ngrok http 8000`
2. Put the DSN of the destination org's project in .env. This is for the webserver (main.go)
3. Run the webserver by running `go build -o main *.go && ./main --local`
4. Send an error event to a project in the origin org.
5. Check your destination org's project to see the error event. It will also be in the origin org's project.

Error event -> origin org project -> Internal Integration -> webhook (ngrok) -> destination org project

## Serverless Function
This is a serverless function and ideally you'd want something that runs 24/7, like App Engine.
1. Create a lambda function in AWS and set your DSN's as environment variables in the lambda runtime environment.
2. `GOOS=linux GOARCH=amd64 go build -o ./main *.go`
3. `zip function.zip main`
4. Upload `function.zip` to your lambda
5. Configure Internal Integration URL to be the lambda.
6. Send error to origin org's project.
7. Check destination org's project.