// Copyright (c) 2019 InfraCloud Technologies
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package logging

import (
	// "github.com/infracloudio/botkube/pkg/logging"
	"github.com/hmharsh/botkube/pkg/logging"
)

func Info(message interface{}) {
	logging.Logger.Info(message)
}

func Trace(message interface{}) {
	logging.Logger.Trace(message)
}

func Debug(message interface{}) {
	logging.Logger.Debug(message)
}

func Warn(message interface{}) {
	logging.Logger.Warn(message)
}

func Error(message interface{}) {
	logging.Logger.Error(message)
}

func Fatal(message interface{}) {
	logging.Logger.Fatal(message)
}

func Panic(message interface{}) {
	logging.Logger.Panic(message)
}

func Infof(format string, v ...interface{}) {
	logging.Logger.Infof(format, v...)
}

func Tracef(format string, v ...interface{}) {
	logging.Logger.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logging.Logger.Debugf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logging.Logger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logging.Logger.Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logging.Logger.Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logging.Logger.Panicf(format, v...)
}
