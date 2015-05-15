/*
 * file   : test_channel.go
 * author : ning
 * date   : 2015-05-11 10:27:53
 */

package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/idning/golog"
	"github.com/ngaut/logging"

	"time"
)

func BenchmarkDummy1(N int64) {
	var i int64
	logging.SetLevel(logging.LOG_LEVEL_WARN)
	logging.SetOutputByName("1.log")
	for i = 0; i < N; i++ {

	}
}

func BenchmarkDummy2(N int64) {
	var i int64
	var j int64
	logging.SetLevel(logging.LOG_LEVEL_WARN)
	logging.SetOutputByName("1.log")
	for i = 0; i < N; i++ {
		if i%2 == 1 {
			j = i
			i = j
		}
	}
}

func variadic(format string, v ...interface{}) {
	return
}

func BenchmarkVariadic(N int64) {
	var i int64
	for i = 0; i < N; i++ {
		variadic("hello %v %v", "abc", "def")
	}
}

func BenchmarkLogging(N int64) {
	var i int64
	//logging.SetLevel(logging.LOG_LEVEL_WARN)
	//logging.SetOutputByName("1.log")
	for i = 0; i < N; i++ {
		logging.Debug("hello %v %v", "abc", "def")
	}
}

func BenchmarkGoLog(N int64) {
	var i int64
	//golog.SetLevel(golog.LEVEL_ERROR)
	for i = 0; i < N; i++ {
		golog.Debug("hello %v %v", "abc", "def")
	}
}

func BenchmarkBeeLog(N int64) {
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"test.log", "level": 3}`)

	var i int64
	for i = 0; i < N; i++ {
		log.Debug("hello %v %v", "abc", "def")
	}
}

func TimeIt(name string, testFunc func(N int64)) {
	var n int64
	n = 1000

	for {
		t0 := time.Now()
		testFunc(n)
		t1 := time.Now()
		diff := t1.Sub(t0)
		if diff.Seconds() < 1 {
			n *= 10
			continue
		}

		fmt.Printf("qps of %10s: %10.0f, diff:%v\n",
			name, 1.0*float64(n)/diff.Seconds(), diff)
		break
	}
}

func main() {
	TimeIt("dummy1", BenchmarkDummy1)
	TimeIt("dummy2", BenchmarkDummy2)
	TimeIt("variadic", BenchmarkVariadic)
	TimeIt("logging", BenchmarkLogging)
	TimeIt("golog", BenchmarkGoLog)
	TimeIt("beelog", BenchmarkBeeLog)
}
