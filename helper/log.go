package helper

import (
	"fmt"
	"time"
)

func Log(path string) {
	logMessage := time.Now().Format("2006-02-1") + " " + path
	fmt.Println(logMessage)
}
