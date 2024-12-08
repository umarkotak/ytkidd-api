run:
	go run .

migrate-up:
	go run . migrate up

build:
	go build -o ytkidd-api cmd/web/main.go

rund:
	nohup ./ytkidd-api &

stopd:
	pkill ytkidd-api

statusd:
	ps aux | grep ytkidd-api
