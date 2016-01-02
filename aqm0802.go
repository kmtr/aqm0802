// Package aqm0802 provides AQM0802 LCD control API
package aqm0802

import "github.com/davecheney/i2c"

import "time"

const contrast = 0x20
const defaultWait = 26300 * time.Nanosecond

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
	cmds := []byte{
		0x38,
		0x39,
		0x14,
		contrast & 15,
		0x5F,
		0x6A,
		0x38,
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
