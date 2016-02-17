package main

import (
	"log"
	"time"

	"github.com/davecheney/i2c"
	"github.com/kmtr/aqm0802"
)

func main() {
	i, err := i2c.New(0x3e, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()
	lcd, err := aqm0802.NewLCD(i)
	if err != nil {
		log.Fatal(err)
	}
	lcd.ChangeRow(0)
	lcd.Write([]byte("1"))
	time.Sleep(500 * time.Millisecond)

	lcd.Home()
	lcd.Write([]byte("2"))
	time.Sleep(500 * time.Millisecond)

	lcd.Home()
	lcd.Write([]byte("3"))
	time.Sleep(500 * time.Millisecond)

	lcd.Clear()
	lcd.ChangeRow(1)
	lcd.Write([]byte("Go"))
	time.Sleep(time.Second)
	lcd.SetupDisplay(true, true, true)

	lcd.Clear()
	lcd.ShiftCursorRight()
	lcd.ShiftCursorRight()
	lcd.Write([]byte("G"))
	time.Sleep(500 * time.Millisecond)
	lcd.ShiftCursorRight()
	time.Sleep(500 * time.Millisecond)
	lcd.Write([]byte("o"))
	time.Sleep(500 * time.Millisecond)

	lcd.ShiftCursorLeft()
	time.Sleep(time.Second)
	lcd.ShiftCursorLeft()
	time.Sleep(time.Second)
	lcd.Write([]byte{byte(7)})
	time.Sleep(time.Second)

	lcd.Home()
	time.Sleep(time.Second)
	lcd.ShiftDisplayRight()
	time.Sleep(time.Second)
	lcd.ShiftDisplayRight()
	time.Sleep(time.Second)
	lcd.ShiftDisplayLeft()
	time.Sleep(time.Second)
	lcd.ShiftDisplayLeft()
	time.Sleep(time.Second)
	lcd.Clear()
}
