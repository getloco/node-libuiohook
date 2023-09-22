package main

/*
#cgo windows LDFLAGS: -L. -lcallback

#cgo CFLAGS: -I .
#cgo LDFLAGS: /home/omran/Projects/omranjamal/node-libuiohook/native/callback.o

#include "callback.h"
*/
import "C"

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"reflect"
)

var receiverFn C.callbackFunc

//export SetReceiver
func SetReceiver(fn C.callbackFunc) {
	receiverFn = fn
}

var stopChan chan struct{} = make(chan struct{})

func startListener() {
	evChan := hook.Start()
	defer hook.End()

	for {
		select {
		case ev, ok := <-evChan:
			if !ok {
				return
			}

			t := reflect.TypeOf(ev)
			fmt.Println(t)

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)

				fmt.Printf("Field Name: %s, Field Type: %s, Tag: %s\n",
					field.Name,
					field.Type,
				)
			}
		case <-stopChan:
			return
		}
	}
}

//export Start
func Start() {
	go startListener()
}

//export End
func End() {
	stopChan <- struct{}{}
}

func main() {}
