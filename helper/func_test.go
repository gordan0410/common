package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestGoFunc(t *testing.T) {
	GoFunc(func() {
		panic("test panic")
	})

	time.Sleep(time.Second * 1)
	fmt.Println("start")
}
