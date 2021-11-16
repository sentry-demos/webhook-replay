# Webhook Replay

ngrok and the Internation Integration are missing from this readme.

## Development
1. set your Sentry DSN's in .env
2. `go build -o main *.go`
3. `./main --local`
4. Send error to origin Sentry org
5. Check Sentry.io to see your events.

## Production
1. Create a lambda function in AWS and set your DSN's as environment variables in the lambda runtime environment.
2. `GOOS=linux GOARCH=amd64 go build -o ./main *.go`
3. `zip function.zip main`
4. Upload `function.zip` to your lambda
5. Send error to origin Sentry org.
6. Check Sentry.io to see your events.


## Test
1. ./ngrok http 8000  
2. go build -o main *.go && ./main --local