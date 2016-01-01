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
	if _, err := lcd.i.Write([]byte{0x38, 0x39, 0x14, contrast & 15, 0x78, 0x5f, 0x6a}); err != nil {
		return err
	}
	time.Sleep(300 * time.Millisecond)
	if _, err := lcd.i.Write([]byte{0x38, 0x0c, 0x01}); err != nil {
		return err
	}
	time.Sleep(300 * time.Millisecond)
	return nil
}

// ChangeRow changes cursor line.
// LCD has two lines.
// Line number is zero-based.
func (lcd *LCD) ChangeRow(n int) {
	switch n {
	case 0:
		lcd.i.Write([]byte{0, 0x80})
	case 1:
		lcd.i.Write([]byte{0, 0xc0})
	}
}

func (lcd *LCD) Write(buf []byte) (int, error) {
	i, err := lcd.i.Write(append([]byte{0x40}, buf...))
	time.Sleep(defaultWait)
	return i - 1, err
}
