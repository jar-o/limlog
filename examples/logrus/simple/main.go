package main

import (
	"fmt"
	"time"

	"github.com/jar-o/limlog"
	"github.com/sirupsen/logrus"
)

func main() {
	// You can make adjustments if you'd like
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	l := limlog.NewLimlogrus()
	l.SetLimiter("limiter1", 4, 1*time.Second, 6)

	l.Info("You don't have to limit if you don't want.")
	l.Debug("It's true.")

	for i := 0; i <= 10000000; i++ {
		l.ErrorL("limiter1", fmt.Sprintf("%d", i))
		l.WarnL("limiter1", fmt.Sprintf("%d", i))
		l.TraceL("limiter1", fmt.Sprintf("%d", i))
		l.InfoL("limiter1", fmt.Sprintf("%d", i))
		l.DebugL("limiter1", fmt.Sprintf("%d", i))
		// l.Debug(i) // <--- This will spew every i
	}
}
