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
	time.Sleep(100 * time.Millisecond)

	lcd.Home()
	lcd.Write([]byte("2"))
	time.Sleep(100 * time.Millisecond)

	lcd.Home()
	lcd.Write([]byte("3"))
	time.Sleep(100 * time.Millisecond)

	lcd.Clear()
	lcd.ChangeRow(1)
	lcd.Write([]byte("Go"))
	lcd.SetupDisplay(true, true, true)
}
