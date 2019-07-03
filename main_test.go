package gotools

import (
	"testing"
)

func TestDebugPrint(t *testing.T) {
	DebugPrint("Test1", "Test2")
	DebugPrint("=============")
	DebugPrint("Test1", "\r\n", "Test2")
	DebugPrint("=============")
	DebugPrint("Test1", "\r\n", "Test2","\r\n")
	DebugPrint("=============")
	DebugPrint("Test1\r\n", "Test2\r\n")
	DebugPrint("=============")
	DebugPrint(123, 456)
	DebugPrint("=============")
	DebugPrint("", "t")
}
