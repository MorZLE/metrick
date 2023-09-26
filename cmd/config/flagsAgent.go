package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var FlagAddr string
var FlagReportInterval int
var FlagPollInterval int

func ParseFlagsAgent() {

	flag.StringVar(&FlagAddr, "a", ":8080", "address and port to run server")
	flag.IntVar(&FlagReportInterval, "r", 10, "Metric report interval")
	flag.IntVar(&FlagPollInterval, "p", 2, "Metric collection time")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		FlagAddr = envRunAddr
	}

	if ReportInterval := os.Getenv("ReportInterval"); ReportInterval != "" {
		ReportInterval, err := strconv.Atoi(ReportInterval)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			return
		}
		FlagReportInterval = ReportInterval
	}

	if PollInterval := os.Getenv("PollInterval"); PollInterval != "" {
		PollInterval, err := strconv.Atoi(PollInterval)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			return
		}
		FlagPollInterval = PollInterval
	}
}
