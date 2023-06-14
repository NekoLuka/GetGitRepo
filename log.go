package main

import (
	"io"
	"log"
	"os"
)

type Log struct {
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
}

func (l *Log) init() {
	file, err := os.OpenFile(os.Getenv("LOG_FILE_LOCATION"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)

	l.info = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	l.warning = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	l.error = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

func (l *Log) Info(message string) {
	l.info.Println(message)
}

func (l *Log) Warning(message string) {
	l.warning.Println(message)
}

func (l *Log) Error(message string) {
	l.error.Println(message)
}

func GetLogger() Log {
	var logger Log
	logger.init()
	return logger
}
