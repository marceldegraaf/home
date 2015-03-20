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

var (
	config *serial.Config
	usb    io.ReadWriteCloser
	reader *bufio.Reader
	C      chan *Sample
)

type Sample struct {
	CurrentlyIn  float64   `json:"currently_in"`   // Current kW into meter
	CurrentlyOut float64   `json:"currently_out"`  // Current kW out of meter
	TotalInLow   float64   `json:"total_in_low"`   // Total kWh in (low tariff)
	TotalInHigh  float64   `json:"total_in_high"`  // Total kWh in (high tariff)
	TotalOutLow  float64   `json:"total_out_low"`  // Total kWh out (low tariff)
	TotalOutHigh float64   `json:"total_out_high"` // Total kWh out (high tariff)
	Tariff       int64     `json:"current_tariff"` // Tariff (1 = low, 2 = high)
	Timestamp    time.Time `json:"timestamp"`
}

func Initialize() chan *Sample {
	C = make(chan *Sample, 6)
	var err error

	config = &serial.Config{Name: device, Baud: rate}
	usb, err = serial.OpenPort(config)
	if err != nil {
		log.Fatalf("kaifa/Initialize: %s", err)
	}

	reader = bufio.NewReader(usb)

	return C
}

func Poll() {
	var m string
	var s *Sample
	var err error

	for {
		if m, err = reader.ReadString(delimiter); err != nil {
			log.Errorf("kaifa/Poll: reading from serial interface: %s", err)
			continue
		}

		if s, err = parse(m); err != nil {
			log.Errorf("kaifa/Poll: parsing message: %s", err)
			continue
		}

		C <- s
	}
}

func parse(m string) (*Sample, error) {
	s := &Sample{Tariff: 1}

	return s, nil
}
