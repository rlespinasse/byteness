package gobits

import (
	"testing"
)

func TestContainsBit(t *testing.T) {
	var val bool
	tests := []struct {
		name   string
		data   byte
		bit    uint8
		result bool
	}{
		{"hi in hi", 0xf0, 5, true},
		{"lo in hi", 0xf0, 2, false},
	}
	for _, tt := range tests {

		val = ContainsBit(tt.data, tt.bit)
		if val != tt.result {
			t.Errorf("Test '%v' failed: ContainsBit(0x%x, %v) was %v, should be %v",
				tt.name,
				tt.data, tt.bit,
				val,
				tt.result)
		}
	}
}

func TestSetBit(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		data   byte
		bit    uint8
		result byte
	}{
		{"set low bit", 0xf0, 0, 0xf1},
		{"set already-set bit", 0xf0, 5, 0xf0},
	}
	for _, tt := range tests {

		val = SetBit(tt.data, tt.bit)
		if val != tt.result {
			t.Errorf("Test '%v' failed: SetBit(0x%x, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.bit,
				val,
				tt.result)
		}
	}
}

func TestUnsetBit(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		data   byte
		bit    uint8
		result byte
	}{
		{"unset high bit", 0xf0, 7, 0x70},
		{"unset already-clear bit", 0xf0, 0, 0xf0},
	}
	for _, tt := range tests {

		val = UnsetBit(tt.data, tt.bit)
		if val != tt.result {
			t.Errorf("Test '%v' failed: UnsetBit(0x%x, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.bit,
				val,
				tt.result)
		}
	}
}

func TestGetBit(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		data   byte
		bit    uint8
		result byte
	}{
		{"get low bit of high nibble", 0xf0, 4, 0x10},
		{"get low bit", 0xf0, 0, 0},
	}
	for _, tt := range tests {

		val = GetBit(tt.data, tt.bit)
		if val != tt.result {
			t.Errorf("Test '%v' failed: GetBit(0x%x, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.bit,
				val,
				tt.result)
		}
	}
}
