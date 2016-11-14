export GOPATH=$(pwd)
go get -u github.com/gorilla/mux
go get -u github.com/patrickmn/go-cache
go get -u github.com/go-sql-driver/mysql

# for tests
go get -u gopkg.in/yaml.v2
go get -u github.com/parnurzeal/gorequest
