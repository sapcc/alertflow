package handlers

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "AlertFlow handlers: ", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)
)
