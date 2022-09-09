package main

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File

func setupLogger(logPath string) {
	logFile, _ := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.Default().SetOutput(logFile)
}

func lprintPanic(err error, message string) {
	log.Default().Println(err.Error())
	log.Default().Println(message)
	closeLogFile()
	panic(err)
}

func lprintPanicIdentical(message string) {
	log.Default().Println(fmt.Errorf(message))
	log.Default().Println(message)
	closeLogFile()
	panic(fmt.Errorf(message))
}

func closeLogFile() {
	logFile.Close()
}
