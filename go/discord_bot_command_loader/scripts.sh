#!/bin/bash

case $1 in

	"build")
		cd bot/ && go build main.go -o bot
	;;

	"dev")
		cd bot/ && air -c .air.toml
	;;

	"loader")
		cd loader/ && cargo run --verbose --color always
	;;

esac