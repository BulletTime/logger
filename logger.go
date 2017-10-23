// The MIT License (MIT)
//
// Copyright © 2017 Sven Agneessens <sven.agneessens@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// These flags define which loggers should be enabled.
const (
	LDebug = 1 << iota
	LInfo
	LWarning
	LError
)

// These variables define which loggers are available.
// They indicate a different log level and will print an appropriate prefix.
var (
	Debug   = log.New(ioutil.Discard, "Discard: ", log.LstdFlags)
	Info    = log.New(ioutil.Discard, "Discard: ", log.LstdFlags)
	Warning = log.New(ioutil.Discard, "Discard: ", log.LstdFlags)
	Error   = log.New(ioutil.Discard, "Discard: ", log.LstdFlags)
)

// Init will initilize all the loggers (Debug, Info, Warning, Error).
// Filename is the file name to which the loggers will write.
// This file name is a string and can include an extension (.log) if wanted.
// Multi is a boolean that will ensure, if set, that the log is also written to
// the standard output and standard error.
// If not set, it will only write to the file.
// The logFlag argument defines the logging properties.
// (see https://golang.org/pkg/log/#pkg-constants)
// The levelFlag argument defines which levels should be enabled.
// The flags are or'ed together, for example to enable all:
// LDebug|LInfo|LWarning|LError
func Init(fileName string, multi bool, logFlag int, levelFlag int) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open logfile:", err)
	}

	var stdHandle io.Writer
	var errHandle io.Writer

	if multi {
		stdHandle = io.MultiWriter(file, os.Stdout)
		errHandle = io.MultiWriter(file, os.Stderr)
	} else {
		stdHandle = file
		errHandle = file
	}

	if levelFlag&LDebug != 0 {
		Debug = log.New(stdHandle, "Debug: ", logFlag)
	}
	if levelFlag&LInfo != 0 {
		Info = log.New(stdHandle, "Info: ", logFlag)
	}
	if levelFlag&LWarning != 0 {
		Warning = log.New(stdHandle, "Warning: ", logFlag)
	}
	if levelFlag&LError != 0 {
		Error = log.New(errHandle, "Error: ", logFlag)
	}

}
