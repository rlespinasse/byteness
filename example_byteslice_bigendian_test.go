package byteslice

import "fmt"

func ExampleLSet() {
	data := []byte{0xAA, 0xCA, 0x55}
	setData := []byte{0x10, 0x12}

	fmt.Printf("%x\n", LSet(data, setData))
	// Output: bada55
}

func ExampleLToogle() {
	data := []byte{0xAB, 0xCB, 0x44}
	setData := []byte{0x11, 0x11, 0x11}

	fmt.Printf("%x\n", LToogle(data, setData))
	// Output: bada55
}
