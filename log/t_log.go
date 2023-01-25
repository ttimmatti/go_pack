package t_log

import (
	"log"
	"os"
)

func SetLogger(logFilePath string) error {
	log_file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.Default().SetFlags(log.Lshortfile)
	log.Default().SetOutput(log_file)

	return nil
}
