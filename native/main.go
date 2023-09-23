package main

/*
// #cgo windows LDFLAGS: -L. -lcallback

#cgo CFLAGS: -I .
#cgo LDFLAGS: /Users/samiha/Projects/node-libuiohook/native/callback.o

#include "callback.h"
*/
import "C"

import (
	hook "github.com/robotn/gohook"
)

var stopChan chan struct{} = make(chan struct{})
var globalEv hook.Event

func startListener(fn C.callbackFunc) {
	evChan := hook.Start()
	defer hook.End()

	for {
		select {
		case ev, ok := <-evChan:
			if ok {
				globalEv = ev
				// C.callCallback(
				// 	fn,

				// 	C.uint8_t(ev.Kind),
				// 	C.int64_t(ev.When.UnixMilli()),
				// 	C.uint16_t(ev.Mask),
				// 	C.uint16_t(ev.Reserved),
				// 	C.uint16_t(ev.Keycode),
				// 	C.uint16_t(ev.Rawcode),
				// 	C.int32_t(ev.Keychar),
				// 	C.uint16_t(ev.Button),
				// 	C.uint16_t(ev.Clicks),
				// 	C.int16_t(ev.X),
				// 	C.int16_t(ev.Y),
				// 	C.uint16_t(ev.Amount),
				// 	C.int32_t(ev.Rotation),
				// 	C.uint8_t(ev.Direction),
				// )
			}
		case <-stopChan:
			return
		}
	}
}

//export Join
func Join(fn C.callbackFunc) {
	// startListener(fn)
}

//export Kill
func Kill() {
	stopChan <- struct{}{}
}

func main() {}
