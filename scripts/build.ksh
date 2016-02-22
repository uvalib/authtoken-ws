export GOPATH=$(pwd)

env GOOS=linux go build -o bin/authtoken-ws.linux authtokenws
env GOOS=darwin go build -o bin/authtoken-ws.darwin authtokenws
