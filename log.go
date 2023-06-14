package main

import (
	"io"
	"log"
	"os"
	"strconv"
)

var _logLevel int8 = 1

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

	level, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatal(err)
	}
	_logLevel = int8(level)

	mw := io.MultiWriter(os.Stdout, file)

	l.info = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	l.warning = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	l.error = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

func (l *Log) Info(message any) {
	if _logLevel < 2 {
		return
	}
	l.info.Println(message)
}

func (l *Log) Warning(message any) {
	if _logLevel < 1 {
		return
	}
	l.warning.Println(message)
}

func (l *Log) Error(message any, fatal bool) {
	if _logLevel < 0 {
		return
	}
	if fatal {
		l.error.Fatal(message)
	} else {
		l.error.Println(message)
	}
}

func SetLogLevel(level int8) {
	_logLevel = level
}

func GetLogger() Log {
	var logger Log
	logger.init()
	return logger
}
