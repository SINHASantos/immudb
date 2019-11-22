/*
Copyright 2019 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logger

import (
	"log"
	"os"
	"strings"
)

const (
	logDebug = iota
	logInfo
	logWarn
	logError
)

type simpleLogger struct {
	Logger   *log.Logger
	LogLevel LogLevel
}

var SimpleLogger = &simpleLogger{
	Logger:   log.New(os.Stderr, "immudb ", log.LstdFlags),
	LogLevel: logLevelFromEnvironment(),
}

func (l *simpleLogger) Errorf(f string, v ...interface{}) {
	if l.LogLevel <= logError {
		l.Logger.Printf("ERROR: "+f, v...)
	}
}

func (l *simpleLogger) Warningf(f string, v ...interface{}) {
	if l.LogLevel <= logWarn {
		l.Logger.Printf("WARNING: "+f, v...)
	}
}

func (l *simpleLogger) Infof(f string, v ...interface{}) {
	if l.LogLevel <= logInfo {
		l.Logger.Printf("INFO: "+f, v...)
	}
}

func (l *simpleLogger) Debugf(f string, v ...interface{}) {
	if l.LogLevel <= logDebug {
		l.Logger.Printf("DEBUG: "+f, v...)
	}
}

func logLevelFromEnvironment() LogLevel {
	logLevel, _ := os.LookupEnv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "error":
		return logError
	case "warn":
		return logWarn
	case "info":
		return logInfo
	case "debug":
		return logDebug
	}
	return logInfo
}
