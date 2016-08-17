# aqm0802

AQM0802 LCD i2c(ST7032i) controller library

## building example

### target Raspberry Pi

```
$ env GOOS=linux GOARCH=arm GOARM=7 go build -o example.raspi ./example
```

### target Mac (mock mode)

```
$ go build -o mock ./example
```