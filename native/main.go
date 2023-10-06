package main

import (
	"bufio"
	"fmt"
	hook "github.com/robotn/gohook"
	"os"
)

// KIND 11 == mouse wheel
// KIND 10 == mouse drag
// KIND 9 == mouse move
// KIND 8 = mouse down
// KIND 7 = mouse hold
// KIND 6 = mouse up
// KIND 5 = key up
// KIND 4 = key hold
// KIND 3 = key down

func start() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("", fmt.Sprintf("RE: %x K: %x R: %x C: %x", ev.Reserved, ev.Keycode, ev.Rawcode, ev.Keychar))
	}
}

func main() {
	go start()

	reader := bufio.NewReader(os.Stdin)

	for {
		command, _ := reader.ReadByte()

		if command == 255 { // simple-listen
			listenFor, _ := reader.ReadByte()

			if listenFor == 0xFF { // key down

			} else if listenFor == 0xFE { // key up

			} else if listenFor == 0xFD { // mouse click

			} else if listenFor == 0xFC { // mouse down

			} else if listenFor == 0xFB { // mouse up

			} else if listenFor == 0xFA { // mouse move

			} else if listenFor == 0xF9 { // mouse drag

			} else if listenFor == 0xF8 { // mouse wheel

			}
		} else if command == 254 { // unregister simple-listen

		} else if command == 253 { // compound-listen

		} else if command == 252 { // unregister compound-listen

		}
	}
}
