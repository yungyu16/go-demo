package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/karalabe/hid"
)

var lockPacket = []uint8{0x10, 0xff, 0x0b, 0x1e, 0x00, 0x00, 0x00}
var unlockPacket = []uint8{0x10, 0xff, 0x0b, 0x1e, 0x01, 0x00, 0x00}

func main() {
	args := os.Args
	cmd := args[1:]
	if len(cmd) != 1 {
		panic("cmd数量不等于1")
	}
	keys := findKeyDev()
	switch cmd[0] {
	case "lock":
		writeHidPacket(keys, lockPacket)
	case "unlock":
		writeHidPacket(keys, unlockPacket)
	}
	fmt.Println("done")
}

func writeHidPacket(keys hid.DeviceInfo, data []uint8) {
	if data == nil {
		panic("data is nil")
	}
	if len(data) != 7 {
		panic("data length error")
	}
	open, err := keys.Open()
	if err != nil {
		panic(err)
	}
	defer func(open *hid.Device) {
		err := open.Close()
		if err != nil {
			panic(err)
		}
	}(open)
	write, err := open.Write(data)
	if err != nil {
		panic(err)
	}
	if write != 7 {
		panic("write error")
	}
}

func findKeyDev() (k hid.DeviceInfo) {
	supported := hid.Supported()
	if !supported {
		panic("unsupported hid")
	}
	enumerate := hid.Enumerate(0, 0)
	for _, dev := range enumerate {
		fmt.Printf("%+v \n", dev)
		if strings.Contains(dev.Product, "羊羽Keys") {
			k = dev
		}
	}
	return
}
