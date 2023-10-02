package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

func NewConfigAgent() *ConfigAgent {
	cnf := &ConfigAgent{}
	return ParseFlagsAgent(cnf)
}

type ConfigAgent struct {
	FlagAddr           string
	FlagReportInterval int
	FlagPollInterval   int
}

func ParseFlagsAgent(p *ConfigAgent) *ConfigAgent {

	flag.StringVar(&p.FlagAddr, "a", ":8080", "address and port to run server")
	flag.IntVar(&p.FlagReportInterval, "r", 10, "metric report interval")
	flag.IntVar(&p.FlagPollInterval, "p", 2, "metric collection time")
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		p.FlagAddr = envRunAddr
	}

	if ReportInterval := os.Getenv("REPORT_INTERVAL "); ReportInterval != "" {
		ReportInterval, err := strconv.Atoi(ReportInterval)
		if err != nil {
			log.Fatalln("Ошибка преобразования строки в число:", err)

		}
		p.FlagReportInterval = ReportInterval
	}

	if PollInterval := os.Getenv("POLL_INTERVAL "); PollInterval != "" {
		PollInterval, err := strconv.Atoi(PollInterval)
		if err != nil {
			log.Fatalln("Ошибка преобразования строки в число:", err)
		}
		p.FlagPollInterval = PollInterval
	}
	return p
}
