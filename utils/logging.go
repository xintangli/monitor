package utils

import (
	"github.com/op/go-logging"
	"os"
	"fmt"
)

var Log *logging.Logger

func InitLogging()  {
	Log = logging.MustGetLogger("example")

	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	logFile, err := os.OpenFile("D:\\log.txt", os.O_APPEND, 0666)
	fmt.Println("open log File error :", err)

	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stdout, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)
}
