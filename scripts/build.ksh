
export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  env GOOS=darwin go build -o bin/authtoken-ws.darwin authtokenws
  res=$?
fi

if [ $res -eq 0 ]; then
  env GOOS=linux go build -o bin/authtoken-ws.linux authtokenws
  res=$?
fi

exit $res
