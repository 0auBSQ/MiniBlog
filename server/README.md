# Golang server

## Required packages

- Pq (Postgresql driver) `go get github.com/lib/pq`
- Gorilla mux `go get -u github.com/gorilla/mux`
- Whirlpool (far better than sha256) `go get github.com/jzelinskie/whirlpool`

## Usage

`go run server.go`

## Postgresql setup

Launch :
```
sudo service postgresql-13 start
```

Init :
```
sudo -u postgres psql
ALTER USER postgres WITH PASSWORD 'test';
```
