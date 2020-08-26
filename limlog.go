package limlog

import (
	"time"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Logger interface {
	GetLogger() interface{}
	Error(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
	Trace(v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})
}

type Limlog struct {
	L            Logger
	rateLimiters map[string]*rate.Limiter
}

func NewLimlog() *Limlog {
	return &Limlog{
		L:            NewLimlogImpl(),
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

func NewLimlogrus() *Limlog {
	return &Limlog{
		L:            NewLimlogrusWithInstance(false),
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

func NewLimlogrusInstance() *Limlog {
	return &Limlog{
		L:            NewLimlogrusWithInstance(true),
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

func NewLimlogZap() *Limlog {
	return &Limlog{
		L:            NewLimlogZapImpl(),
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

func NewLimlogZapWithConfig(cfg interface{}) *Limlog {
	return &Limlog{
		L:            NewLimlogZapWithConfigImpl(cfg),
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

func NewLimlogWithZap(zap *zap.Logger) *Limlog {
	return &Limlog{
		L:            &zapImpl{Zap: zap},
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

// Use SetLimiter to set how many logLines are emitted in a given interval for
// a specific identifier.
func (o *Limlog) SetLimiter(limiter string, logLines float64, interval time.Duration, burst int) {
	if logLines == float64(burst) {
		burst++
	}
	// Check first, and only add the limiter if it doesn't exist already.
	if _, ok := o.rateLimiters[limiter]; !ok {
		o.rateLimiters[limiter] = rate.NewLimiter(rate.Limit(logLines/interval.Seconds()), burst)
	}
}

// No limiting, just log
func (o *Limlog) Error(v ...interface{}) {
	o.L.Error(v...)
}

func (o *Limlog) Warn(v ...interface{}) {
	o.L.Warn(v...)
}

func (o *Limlog) Info(v ...interface{}) {
	o.L.Info(v...)
}

func (o *Limlog) Debug(v ...interface{}) {
	o.L.Debug(v...)
}

func (o *Limlog) Trace(v ...interface{}) {
	o.L.Trace(v...)
}

func (o *Limlog) Fatal(v ...interface{}) {
	o.L.Fatal(v...)
}

func (o *Limlog) Panic(v ...interface{}) {
	o.L.Panic(v...)
}

// Use limiting
func (o *Limlog) ErrorL(limiter string, v ...interface{}) {
	if o.allowLog(limiter) {
		o.L.Error(v...)
	}
}

func (o *Limlog) WarnL(limiter string, v ...interface{}) {
	if o.allowLog(limiter) {
		o.L.Warn(v...)
	}
}

func (o *Limlog) InfoL(limiter string, v ...interface{}) {
	if o.allowLog(limiter) {
		o.L.Info(v...)
	}
}

func (o *Limlog) DebugL(limiter string, v ...interface{}) {
	if o.allowLog(limiter) {
		o.L.Debug(v...)
	}
}

func (o *Limlog) TraceL(limiter string, v ...interface{}) {
	if o.allowLog(limiter) {
		o.L.Trace(v...)
	}
}

func (o *Limlog) allowLog(limiter string) bool {
	if lim, ok := o.rateLimiters[limiter]; ok {
		return lim.Allow()
	}
	// If there is no rate limiter set up, don't log
	return false
}
