package main

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

const defaultWait = 1000

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	wait, _ := strconv.Atoi(os.Getenv(`WAIT_MILLISECONDS`))
	if wait == 0 {
		wait = defaultWait
	}
	container := os.Getenv(`CONTAINER`)

	stdFields := log.Fields{
		`container`: container,
		`wait`:      wait,
	}

	cntr := 1

	for {
		log.WithFields(stdFields).WithFields(log.Fields{`time`: time.Now(), `cntr`: cntr}).Info(`l`)
		cntr++
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}
