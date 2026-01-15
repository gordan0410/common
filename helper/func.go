package helper

import "github.com/gordan0410/common/log"

func GoFunc(goFunc func()) {
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				// 這裡禁止使用 Panic
				log.Error().Str(log.LogField_Common, "GoFunc").Msg("goroutine panic")
			}
		}()
		goFunc()
	}()
}
