package test

import (
	"log"
	"testing"
)

//命令行：go test -v helloworld_test.go helloworld.go
func TestDefertest(t *testing.T)  {
	Timeout()
	log.Println("<<<test test")
}
