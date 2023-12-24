package main

import "log"

func main() {
	svr := initServer()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
