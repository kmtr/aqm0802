// +build !linux

package aqm0802

import (
	"bytes"
	"log"

	"github.com/davecheney/i2c"
)

// LCD is struct of AQM0802 LCD
type LCD struct {
	w *bytes.Buffer
}

// NewLCD creates an LCD instance.
func NewLCD(i *i2c.I2C) (*LCD, error) {
	lcd := new(LCD)
	lcd.w = bytes.NewBuffer([]byte{})
	return lcd, nil
}

// Reset sends a reset command.
func (lcd *LCD) Reset() error {
	log.Print("debug: reset")
	return nil
}

// ChangeRow changes cursor line.
func (lcd *LCD) ChangeRow(n int) {
	log.Printf("debug: change row to %d", n)
}

func (lcd *LCD) Write(buf []byte) (int, error) {
	i, err := lcd.w.Write(buf)
	log.Printf("debug: lcd\t|%s", lcd.w.String())
	return i, err
}

// Cmd sends a control command.
func (lcd *LCD) Cmd(cmd byte) error {
	log.Printf("debug: send cmd %s", string(cmd))
	return nil
}

// Clear clears LCD
func (lcd *LCD) Clear() error {
	log.Printf("debug: Clear")
	lcd.w.Reset()
	return nil
}

// Home sets cursor to home
func (lcd *LCD) Home() error {
	log.Printf("debug: Home")
	lcd.w.Reset()
	return nil
}

// SetupDisplay setup display.
func (lcd *LCD) SetupDisplay(on bool, cur bool, blink bool) error {
	log.Printf("debug: SetupDisplay on=%v, cur=%v, blink=%v", on, cur, blink)
	return nil
}

// SetContrast sets contrast
func (lcd *LCD) SetContrast(c int) error {
	log.Printf("SetContrast c=%d", c)
	return nil
}

// RegisterCG registers graphinc data with the CGRAM
func (lcd *LCD) RegisterCG(p int, data [8]byte) error {
	log.Printf("RegisterCG p=%d, data=%v", p, data)
	return nil
}

// ShiftCursorLeft shifts cursor to the left
func (lcd *LCD) ShiftCursorLeft() error {
	return nil
}

// ShiftCursorRight shifts cursor to the right
func (lcd *LCD) ShiftCursorRight() error {
	return nil
}

// ShiftDisplayLeft shifts cursor to the left
func (lcd *LCD) ShiftDisplayLeft() error {
	return nil
}

// ShiftDisplayRight shifts cursor to the right
func (lcd *LCD) ShiftDisplayRight() error {
	return nil
}
