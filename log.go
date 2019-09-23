package limlog

// Implementation for the standard Golang logger

import (
	"log"
)

type limlogImpl struct{}

func NewLimlogImpl() *limlogImpl {
	return &limlogImpl{}
}

func (o *limlogImpl) GetLogger() interface{} {
	return nil
}

func (o *limlogImpl) Error(v ...interface{}) {
	v = append(levelAsInterface("ERROR"), v...)
	log.Println(v...)
}

func (o *limlogImpl) Warn(v ...interface{}) {
	v = append(levelAsInterface("WARN"), v...)
	log.Println(v...)
}

func (o *limlogImpl) Info(v ...interface{}) {
	v = append(levelAsInterface("INFO"), v...)
	log.Println(v...)
}

func (o *limlogImpl) Debug(v ...interface{}) {
	v = append(levelAsInterface("DEBUG"), v...)
	log.Println(v...)
}

func (o *limlogImpl) Trace(v ...interface{}) {
	v = append(levelAsInterface("TRACE"), v...)
	log.Println(v...)
}

func (o *limlogImpl) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (o *limlogImpl) Panic(v ...interface{}) {
	log.Panic(v...)
}

func levelAsInterface(level string) []interface{} {
	r := make([]interface{}, 1)
	r[0] = level
	return r
}
