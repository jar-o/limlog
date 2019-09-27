# limlog
Golang logger with rate limiting

## Overview

Logging often gets in that uncomfortable place where you spam too much, or omit
too much, and it's hard to always get it just right -- where "just right" means
enough logging to tell you what's going on, and not so much that you're drowning
in log lines.

Enter rate limiting and `limlog`. This package uses simple token bucket rate
limiting -- usually the same algorithm generating `429` on your requests -- and
applies that to log lines. It's simple to use, and works out of the box with the
standard [log](https://golang.org/pkg/log/) package, as well as the popular
[logrus](https://github.com/sirupsen/logrus) package.

## Install

Just run

```
go get -u github.com/jar-o/limlog
```

to install.

## Get a feel for it

You can run the included
[examples](https://github.com/jar-o/limlog/blob/master/examples)
pretty easily, e.g.

```
go run examples/log/main.go
```


## Details

`limlog` is a simple and lightweight level logger at its core. There are
several [examples](https://github.com/jar-o/limlog/blob/master/examples) in this
repo that you can refer to. However, a quick example of using the standard `log`
package would look a little like

```
package main

import (
	"log"
	"github.com/jar-o/limlog"
)

func main() {
	l := limlog.NewLimlog()

	// All entirely optional
	log.SetPrefix("HELOWRLD: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	// Total of 4 log lines per second, with a burst of 6
	l.SetLimiter("limiter1", 4, 6)

	l.Info("You don't have to limit your log lines if you don't want.")

	for i := 0; i <= 10000000; i++ {
		l.DebugL("limiter1", i)
	}
}
```

Loggers that want to integrate into `limlog` must implement the following interface:

```
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
```

As you can see `limlog` provides all the common logging levels. To do a basic
logging call (no limiting) you would do something like

```
l.Error(...)
l.Debug(...)
... etc ...
```

However, for limiting you first setup a limiter, then use the `L` version of the
log level calls. E.g.

```
// This limiter allows only one log line per second, with a burst of one (i.e.
// no burst)
l.SetLimiter("mylimiterkey", 1, 1)

// In a loop somwhere:
l.InfoL("mylimiterkey", ...)
```

Note that we're using `WarnL()` instead of `Warn()`.

The rate limiter is designed around keys, which are just strings that should
make sense for whatever context in which you find yourself logging. For example,
say you have organizations using a service, and at some point they reach the
max for their subscription level. However, they are using scripts and so the
log line for reaching their subscription level begins to spam your logs.

```
my-cool-service time="2019-09-23T08:54:15Z" level=warning msg="Org 1679 has reached their subscription limit for item 111" metric=warning

my-cool-service time="2019-09-23T08:54:15Z" level=warning msg="Org 1679 has reached their subscription limit for item 123" metric=warning

... etc ...
```

Because of the scripts these log lines fill up your logs. However, it's a useful
signal, knowing that the org is frequently bumping up against their subscription
limits, so omitting them entirely is not helpful either. With `limlog` you'd
setup a key for this and use `WarnL()`:

```
lkey := fmt.Sprintf("sub%d", orgID)
l.SetLimiter(lkey, 1, 1)
l.WarnL(lkey, fmt.Sprintf("Org %d has reached their subscription limit at item %d", orgID, itemID)
```

Now you'll see a log line for when the customer's subscription is hit, but no more than once a second.

# Logrus

The [logrus](https://github.com/sirupsen/logrus) package provides more advanced
behavior than the standard `log` package. For starters, it can have logger
instances in addition to a single package logger. The advantage of logger
instances is that you can have multiple custom loggers within a single process.

To set that up, just call `NewLimlogrusInstance()`.

```
l := limlog.NewLimlogrusInstance()
inst := l.L.GetLogger().(*logrus.Logger)
// ... Adjust settings on inst ...
```

See the
[advanced](https://github.com/jar-o/limlog/blob/master/examples/logrus/advanced/main.go)
example for details on how this works.
