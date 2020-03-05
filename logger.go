package main

import (
	"os"
	"log"
	"time"
	"fmt"
)

func LogInit(prefix string) (*os.File, *log.Logger) {
	fileName := GenFileLogName(prefix)
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	return f, log.New(f, prefix, log.LstdFlags)
}

func GenFileLogName(prefix string) (string) { return fmt.Sprintf("%s-%s.log", prefix, time.Now().Format("2006-01-02T15-04-05"))}
