#!/usr/bin/make -ef

NAME = go-myspace-libp2p

default: tidy

init: go.mod go.sum

go.mod:
	go mod init $(NAME)

go.sum:
	go mod tidy

tidy:
	go mod tidy

.PHONY: tidy