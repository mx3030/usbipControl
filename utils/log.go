package utils

import (
	"os"
)

var oldStdout *os.File

func setupLogging() (*os.File, error) {
	logFile, err := os.Create("output.log")
	if err != nil {
		return nil, err
	}
	oldStdout = os.Stdout
	os.Stdout = logFile
	return logFile, nil
}

func cleanupLogging(logFile *os.File) {
	if logFile != nil {
		logFile.Close()
		os.Stdout = oldStdout
	}
}

func WithLogging(enableLogging bool, fn func() error) error {
	var logFile *os.File
	var err error

	if enableLogging {
		logFile, err = setupLogging()
		if err != nil {
			return err
		}
		defer cleanupLogging(logFile)
	}

	return fn()
}
