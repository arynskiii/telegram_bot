.PHONY:


build:
	go build -o ./home/aryn cmd/main/bot/main.go
run: build
	./.bin/bot