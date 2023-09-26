package flags

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var FlagRunAddr string

func ParseFlags() {

	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")

	flag.Parse()

	if len(flag.Args()) > 0 {
		// Вывод сообщения об ошибке и синтаксисе использования
		fmt.Println("Ошибка: неизвестные флаги")
		flag.Usage()
		os.Exit(1)
	}
	log.Printf("Starting UpServer on %s\n", FlagRunAddr)
}
