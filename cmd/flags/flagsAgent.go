package flags

import (
	"flag"
	"fmt"
	"os"
)

var FlagAddr string
var FlagReportInterval int
var FlagPollInterval int

func ParseFlagsAgent() {

	flag.StringVar(&FlagAddr, "a", ":8080", "address and port to run server")
	flag.IntVar(&FlagReportInterval, "r", 10, "Metric report interval")
	flag.IntVar(&FlagPollInterval, "p", 2, "Metric collection time")
	flag.Parse()
	if len(flag.Args()) > 0 {
		// Вывод сообщения об ошибке и синтаксисе использования
		fmt.Println("Ошибка: неизвестные флаги")
		flag.Usage()
		os.Exit(1)
	}
}
