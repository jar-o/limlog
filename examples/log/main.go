package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jar-o/limlog"
)

func main() {
	l := limlog.NewLimlog()
	log.SetPrefix("WUT: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	// You can write to a file too:
	// f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// wrt := io.MultiWriter(os.Stdout, f)
	// log.SetOutput(wrt)

	// Logs 1 line every 3 seconds, with a burst of 6
	l.SetLimiter("limiter1", 1, 3*time.Second, 6)
	l.Info("You don't have to limit if you don't want.")
	for i := 0; i <= 10000000; i++ {
		// Each of these share the same limiter identifier. So it will be
		// indeterminate which ones will emit during a given run.
		l.ErrorL("limiter1", fmt.Sprintf("%d", i))
		l.WarnL("limiter1", fmt.Sprintf("%d", i))
		l.TraceL("limiter1", fmt.Sprintf("%d", i))
		l.InfoL("limiter1", fmt.Sprintf("%d", i))
		l.DebugL("limiter1", fmt.Sprintf("%d", i))
		// l.Debug(i) // <--- This would spew every i
	}
}
