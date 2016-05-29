// +build !linux

package main

import "github.com/davecheney/i2c"

func newI2C() (*i2c.I2C, error) {
	return nil, nil
}
