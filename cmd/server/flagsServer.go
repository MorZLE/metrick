package main

import (
	"flag"
	"fmt"
	"os"
)

var flagRunAddr string

func parseFlags() {

	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.Parse()
	if len(flag.Args()) > 0 {
		// Вывод сообщения об ошибке и синтаксисе использования
		fmt.Println("Ошибка: неизвестные флаги")
		flag.Usage()
		os.Exit(1)
	}
}
