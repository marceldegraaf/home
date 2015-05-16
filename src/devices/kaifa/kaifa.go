package kaifa

import (
	"bufio"
	"io"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/tarm/goserial"
)

const (
	device        = "/dev/ttyUSB0"
	rate          = 115200
	delimiter     = '\x21' // "!" character
	crcDelimiter  = '\x0a' // newline
	crcPolynomial = 0xA001 // IBM CRC16
)

type Kaifa struct {
	C      chan *Point
	config *serial.Config
	usb    io.ReadWriteCloser
	reader *bufio.Reader
}

type Point struct {
	CurrentlyIn  float64   `json:"currently_in"`   // Current kW into meter
	CurrentlyOut float64   `json:"currently_out"`  // Current kW out of meter
	TotalInLow   float64   `json:"total_in_low"`   // Total kWh in (low tariff)
	TotalInHigh  float64   `json:"total_in_high"`  // Total kWh in (high tariff)
	TotalOutLow  float64   `json:"total_out_low"`  // Total kWh out (low tariff)
	TotalOutHigh float64   `json:"total_out_high"` // Total kWh out (high tariff)
	Tariff       int64     `json:"current_tariff"` // Tariff (1 = low, 2 = high)
	Timestamp    time.Time `json:"timestamp"`
}

func Initialize() *Kaifa {
	var err error

	k := Kaifa{C: make(chan *Point, 10)}

	k.config = &serial.Config{Name: device, Baud: rate}
	k.usb, err = serial.OpenPort(k.config)
	if err != nil {
		log.Fatalf("kaifa/Initialize: %s", err)
	}

	k.reader = bufio.NewReader(k.usb)

	return &k
}

func (k *Kaifa) Poll() {
	var m string
	var err error

	for {
		if m, err = k.reader.ReadString(delimiter); err != nil {
			log.Errorf("kaifa/Poll: reading from serial interface: %s", err)
			continue
		}

		k.C <- parse(m)
	}
}
