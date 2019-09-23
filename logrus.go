package limlog

import (
	"github.com/sirupsen/logrus"
)

type limlogrusImpl struct {
	Logrus *logrus.Logger
}

const paramErr = "Problem with params. Needs to be (string[, logrus.Fields])"

// Using an instance allows for more advanced features as well as multiple,
// fine-grain loggers running within a single process. But for simple stuff, you
// can just use a package level logger. NewLimlogrusWithInstance facilitates
// setting up either.
func NewLimlogrusWithInstance(createInstance bool) *limlogrusImpl {
	if createInstance {
		return &limlogrusImpl{
			Logrus: logrus.New(),
		}
	} else {
		return &limlogrusImpl{}
	}
}

func (o *limlogrusImpl) GetLogger() interface{} {
	return o.Logrus
}

func (o *limlogrusImpl) Error(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Error(msg)
	} else {
		o.Logrus.WithFields(fields).Error(msg)
	}
}

func (o *limlogrusImpl) Warn(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Warn(msg)
	} else {
		o.Logrus.WithFields(fields).Warn(msg)
	}
}

func (o *limlogrusImpl) Info(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Info(msg)
	} else {
		o.Logrus.WithFields(fields).Info(msg)
	}
}

func (o *limlogrusImpl) Debug(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Debug(msg)
	} else {
		o.Logrus.WithFields(fields).Debug(msg)
	}
}

func (o *limlogrusImpl) Trace(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Trace(msg)
	} else {
		o.Logrus.WithFields(fields).Trace(msg)
	}
}

func (o *limlogrusImpl) Fatal(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Fatal(msg)
	} else {
		o.Logrus.WithFields(fields).Fatal(msg)
	}
}

func (o *limlogrusImpl) Panic(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if o.Logrus == nil {
		logrus.WithFields(fields).Panic(msg)
	} else {
		o.Logrus.WithFields(fields).Panic(msg)
	}
}

func (o *limlogrusImpl) convertParams(a ...interface{}) (string, logrus.Fields) {
	if len(a) == 1 {
		msg, ok := a[0].(string)
		if !ok {
			return paramErr, nil
		}
		return msg, nil
	}

	if len(a) == 2 {
		msg, ok := a[0].(string)
		if !ok {
			return paramErr, nil
		}
		fields, ok := a[1].(logrus.Fields)
		if !ok {
			return paramErr, nil
		}
		return msg, fields

	}
	return "", nil
}
