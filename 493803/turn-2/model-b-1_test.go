package main

import (
	"runtime"
	"testing"
)

var globalVar int = 0

func TestGlobalVarWithKeepAlive(t *testing.T) {
	// Simulate work using globalVar
	globalVar++
	runtime.KeepAlive(globalVar) // Pin globalVar
}
