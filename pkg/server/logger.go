package server

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "AlertFlow server: ", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)
)
