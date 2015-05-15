/*
 * file   : log_test.go
 * author : ning
 * date   : 2015-05-14 10:19:36
 */

package golog

import (
	"testing"
)

func logs() {
	Error("error msg")
	Notice("notice msg")
	Info("info msg")
	Debug("debug msg")
	Verbose("debug msg")
}

func TestBasic(t *testing.T) {
	logs()

	SetLevel(6)
	logs()

	SetFile("test.log")
	SetLevel(5)
	logs()

	SetLevel(6)
	logs()
}

func TestBench(t *testing.T) {
	var i int64
	var N int64

	N = 1000 * 1000 * 10
	for i = 0; i < N; i++ {
		Debug("hello %v %v", "abc", "def")
	}
}
