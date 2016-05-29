// +build linux

package main

import "github.com/davecheney/i2c"

func newI2C() (*i2c.I2C, error) {
	return i2c.New(0x3e, 1)
}
