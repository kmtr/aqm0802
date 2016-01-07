// Package aqm0802 provides AQM0802 LCD control API
package aqm0802

import (
	"time"

	"github.com/davecheney/i2c"
)

const defaultContrast = 0x32
const defaultWait = 26300 * time.Nanosecond

const (
	// FunctionSetIS0 represents command that is Function Set IS=0
	FunctionSetIS0 = 0x38
	// FunctionSetIS1 represents command that is Function Set IS=1
	// You must send this before to use extra commands
	// and send FunctionSetIS0 after using extra commands
	FunctionSetIS1 = 0x39
)

// LCD is struct of AQM0802 LCD
type LCD struct {
	i *i2c.I2C
}

// NewLCD creates an LCD instance.
func NewLCD(i *i2c.I2C) (*LCD, error) {
	lcd := LCD{
		i: i,
	}
	if err := lcd.Reset(); err != nil {
		return nil, err
	}
	return &lcd, nil
}

// Reset sends a reset command.
func (lcd *LCD) Reset() error {
	cu, cl := parseContrastValue(defaultContrast)
	cmds := []byte{
		FunctionSetIS0,
		FunctionSetIS1,
		0x14,
		cu,
		0x5C | cl,
		0x6A,
		FunctionSetIS0,
		0x0C,
		0x01,
	}
	for _, cmd := range cmds {
		if err := lcd.Cmd(cmd); err != nil {
			return err
		}
	}
	return nil
}

// ChangeRow changes cursor line.
// LCD has two lines.
// Line number is zero-based.
func (lcd *LCD) ChangeRow(n int) {
	switch n {
	case 0:
		lcd.Cmd(0x80)
	case 1:
		lcd.Cmd(0xC0)
	}
}

func (lcd *LCD) Write(buf []byte) (int, error) {
	i, err := lcd.i.Write(append([]byte{0x40}, buf...))
	time.Sleep(defaultWait)
	return i - 1, err
}

// Cmd sends a control command.
func (lcd *LCD) Cmd(cmd byte) error {
	_, err := lcd.i.Write(append([]byte{0x0}, cmd))
	time.Sleep(defaultWait)
	return err
}

// Clear clears LCD
func (lcd *LCD) Clear() error {
	err := lcd.Cmd(0x01)
	time.Sleep((1080 - defaultWait) * time.Microsecond)
	return err
}

// Home sets cursor to home
func (lcd *LCD) Home() error {
	err := lcd.Cmd(byte(0x02))
	time.Sleep((1080 - defaultWait) * time.Microsecond)
	return err
}

// SetupDisplay setup display.
// 1st arg: Display ON/OFF.
// Even when the display is turned off, the data is remained in RAM.
// 2nd arg: Cursor ON/OFF.
// Even when the cursor is disappeared, the index register remains its data.
// 3rd arg" Cursor Blink ON/OFF
func (lcd *LCD) SetupDisplay(on bool, cur bool, blink bool) error {
	cmd := 0x01
	cmd = cmd << 1
	if on {
		cmd++
	}
	cmd = cmd << 1
	if cur {
		cmd++
	}
	cmd = cmd << 1
	if blink {
		cmd++
	}
	return lcd.Cmd(byte(cmd))
}

// SetContrast sets contrast
func (lcd *LCD) SetContrast(c int) error {
	cu, cl := parseContrastValue(c)
	if err := lcd.Cmd(FunctionSetIS1); err != nil {
		return err
	}
	defer lcd.Cmd(FunctionSetIS0)
	cmds := []byte{
		cu,
		0x5C | cl,
	}
	for _, cmd := range cmds {
		if err := lcd.Cmd(cmd); err != nil {
			return err
		}
	}
	return nil
}

func parseContrastValue(c int) (byte, byte) {
	u := byte(c & 15)
	l := byte(c >> 4 & 3)
	return byte(u), byte(l)
}

// RegisterCG registers graphinc data with the CGRAM
// GraphicData consist of 5bit(0x1F == 31) x 8 lines bit array.
// First line is data[0].
// Last line(data[7]) is cursor line.
//
// Character Pattern
// -------------------
// 0 0 0 0 0 | data[0]
// 0 0 0 0 0 | data[1]
// ...
// 0 0 0 0 0 | data[7]
//
func (lcd *LCD) RegisterCG(p int, data [8]byte) error {
	addr := byte(0x40 + p*8)
	// Set CGRAM
	err := lcd.Cmd(addr)
	if err != nil {
		return err
	}
	for _, line := range data {
		// Write Data
		_, err = lcd.i.Write(append([]byte{addr}, line))
		if err != nil {
			return err
		}
	}
	// Set DDRAM
	err = lcd.Cmd(byte(0x80))
	if err != nil {
		return err
	}
	return nil
}

// ShiftCursorLeft shifts cursor to the left
func (lcd *LCD) ShiftCursorLeft() error {
	return lcd.Cmd(byte(0x10))
}

// ShiftCursorRight shifts cursor to the right
func (lcd *LCD) ShiftCursorRight() error {
	return lcd.Cmd(byte(0x14))
}

// ShiftDisplayLeft shifts cursor to the left
func (lcd *LCD) ShiftDisplayLeft() error {
	return lcd.Cmd(byte(0x18))
}

// ShiftDisplayRight shifts cursor to the right
func (lcd *LCD) ShiftDisplayRight() error {
	return lcd.Cmd(byte(0x1C))
}
