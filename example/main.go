package main

import (
	"log"

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
	lcd.Write([]byte("123"))
	lcd.ChangeRow(1)
	lcd.Write([]byte("Go"))
}
