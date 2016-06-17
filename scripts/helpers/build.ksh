export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  GOOS=darwin go build -o bin/authtoken-ws.darwin authtokenws
  res=$?
fi

if [ $res -eq 0 ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/authtoken-ws.linux authtokenws
  res=$?
fi

exit $res
