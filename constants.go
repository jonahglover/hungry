package main

import (
	"fmt"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("hungry")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Critical("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

/*
log.Info("info")
log.Notice("notice")
log.Warning("warning")
log.Error("err")
log.Critical("crit")
*/
