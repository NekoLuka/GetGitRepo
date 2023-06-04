package main

import (
	"io"
	"log"
	"os"
)

type Log struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func (l *Log) init() {
	file, err := os.OpenFile(os.Getenv("LOG_FILE_LOCATION"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)

	l.Info = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	l.Warning = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	l.Error = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

func GetLogger() Log {
	var logger Log
	logger.init()
	return logger
}
