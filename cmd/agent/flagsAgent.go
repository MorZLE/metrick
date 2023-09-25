package main

import (
	"flag"
	"fmt"
	"os"
)

var flagAddr string
var flagReportInterval int
var flagPollInterval int

func parseFlagsAgent() {

	flag.StringVar(&flagAddr, "a", ":8080", "address and port to run server")
	flag.IntVar(&flagReportInterval, "r", 10, "Metric report interval")
	flag.IntVar(&flagPollInterval, "p", 2, "Metric collection time")
	flag.Parse()
	if len(flag.Args()) > 0 {
		// Вывод сообщения об ошибке и синтаксисе использования
		fmt.Println("Ошибка: неизвестные флаги")
		flag.Usage()
		os.Exit(1)
	}
}
