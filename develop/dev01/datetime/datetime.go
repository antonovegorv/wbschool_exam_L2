package datetime

import (
	"time"

	"github.com/beevik/ntp"
)

// Datetime — struct to hold fields and methods of getting local and exact datetime.
type Datetime struct {
	hostname string
}

// New — constructor for a Datetime struct.
func New(hostname string) *Datetime {
	return &Datetime{hostname: hostname}
}

// SetHostname — setter for a hostname field.
func (dt *Datetime) SetHostname(hostname string) {
	dt.hostname = hostname
}

// GetLocal — method to get local time from machine.
func (dt *Datetime) GetLocal() time.Time {
	return time.Now()
}

// GetExact — method to get exact time from ntp server.
func (dt *Datetime) GetExact() (time.Time, error) {
	resp, err := ntp.Query(dt.hostname)
	if err != nil {
		return time.Time{}, err
	}

	err = resp.Validate()
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().Add(resp.ClockOffset), nil
}
