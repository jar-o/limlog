package limlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapImpl struct {
	Zap *zap.Logger
}

func NewLimlogZapImpl() *zapImpl {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &zapImpl{
		Zap: logger,
	}
}

func NewZapConfigWithLevel(level zapcore.Level) zap.Config {
	var cfg zap.Config
	if level == zap.DebugLevel {
		cfg = zap.NewDevelopmentConfig()
	} else { // Otherwise, start with leanest config, user can config from there
		cfg = zap.NewProductionConfig()
	}
	cfg.Level.SetLevel(level)
	return cfg
}

func NewLimlogZapWithConfigImpl(cfg interface{}) *zapImpl {
	conf := cfg.(zap.Config)
	logger, err := conf.Build()
	if err != nil {
		panic(err)
	}
	return &zapImpl{
		Zap: logger,
	}
}

func (o *zapImpl) GetLogger() interface{} {
	return o.Zap
}

func (o *zapImpl) Error(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.Error(msg)
	} else {
		o.Zap.With(fields...).Error(msg)
	}
}

func (o *zapImpl) Warn(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.Warn(msg)
	} else {
		o.Zap.With(fields...).Warn(msg)
	}
}

func (o *zapImpl) Info(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.Info(msg)
	} else {
		o.Zap.With(fields...).Info(msg)
	}
}

func (o *zapImpl) Debug(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.Debug(msg)
	} else {
		o.Zap.With(fields...).Debug(msg)
	}
}

func (o *zapImpl) Trace(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.DPanic(msg)
	} else {
		o.Zap.With(fields...).DPanic(msg)

	}
}

func (o *zapImpl) Panic(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.Panic(msg)
	} else {
		o.Zap.With(fields...).Panic(msg)

	}
}

func (o *zapImpl) Fatal(a ...interface{}) {
	msg, fields := o.convertParams(a...)
	if fields == nil {
		o.Zap.Fatal(msg)
	} else {
		o.Zap.With(fields...).Fatal(msg)
	}
}

func (o *zapImpl) convertParams(a ...interface{}) (string, []zap.Field) {
	if len(a) == 1 {
		msg, ok := a[0].(string)
		if !ok {
			return paramErr, nil
		}
		return msg, nil
	}

	fields := []zap.Field{}
	for _, field := range a[1:] {
		f, ok := field.(zap.Field)
		if !ok {
			return paramErr, nil
		}
		fields = append(fields, f)
	}

	msg, ok := a[0].(string)
	if !ok {
		return paramErr, nil
	}

	return msg, fields
}
