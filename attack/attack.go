package attack

import (
	"github.com/op/go-logging"
)

var logger *logging.Logger

// SetLogger ...
func SetLogger(l *logging.Logger) {
	logger = l
}

func Attack(target string) {
	logger.Noticef("Attacking: %s", target)
}
