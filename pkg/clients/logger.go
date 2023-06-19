package clients

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "AlertFlow Clients: ", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)
)
