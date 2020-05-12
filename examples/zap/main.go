package main

import (
	"fmt"
	"time"

	"github.com/jar-o/limlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// If you don't want to set level or other config options, just use the
	// Production version of Zap:
	// l := limlog.NewLimlogZap()

	// This sets a NewProduction config, but with a custom AtomicLevel. You can
	// also tweak cfg further, as required.
	cfg := limlog.NewZapConfigWithLevel(zap.DebugLevel)
	cfg.Encoding = "console" // By default this is JSON
	l := limlog.NewLimlogZapWithConfig(cfg)
	z := l.L.GetLogger().(*zap.Logger)
	defer z.Sync()

	// Setup some limiter tags
	l.SetLimiter("limiter1", 10, 1*time.Second, 6)
	l.SetLimiter("limiter2", 1, 5*time.Second, 1)

	l.Info("You don't have to limit if you don't want.")
	l.Debug("It's true.")

	for i := 0; i <= 10000000; i++ {
		l.ErrorL("limiter1", fmt.Sprintf("%d", i))
		l.WarnL(
			"limiter1",
			fmt.Sprintf("%d", i),
			zap.Field{Key: "helo", Type: zapcore.StringType, String: "wrld"},
			zap.Field{Key: "ehlo", Type: zapcore.StringType, String: "wlrd"},
		)
		l.InfoL("limiter1", fmt.Sprintf("%d", i))
		l.DebugL("limiter1", fmt.Sprintf("%d", i))
		l.DebugL("limiter2", fmt.Sprintf("limiter2 %d", i))
		// l.Debug(i) // <--- This would spew every iteration
	}
}
