package main

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfile} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)

type Password string

func (p Password) Redacted() interface{}{
	return logging.Redact(string(p))
}
func main(){
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend2Leveled := logging.AddModuleLevel(backend2Formatter)
	backend2Leveled.SetLevel(logging.INFO, "")

	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Leveled, backend2Leveled)

	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
	doSomething()
}

func doSomething(){
	log.Info("HAI I'M INSIDE ANOTHER FUNCTION, YOU MIGHT BE ABLE TO SEE ME ON THE LEFT SIDE")
}
