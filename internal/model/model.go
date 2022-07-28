package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	dataRates = map[string]bool{
		"SF7BW125":      true,
		"SF7BW250":      true,
		"SF8BW125":      true,
		"SF8BW500":      true,
		"SF9BW125":      true,
		"SF9BW500":      true,
		"SF10BW125":     true,
		"SF10BW500":     true,
		"SF11BW125":     true,
		"SF11BW500":     true,
		"SF12BW125":     true,
		"SF12BW500":     true,
		"FSK50":         true,
		"M0CW137CR1/3":  true,
		"M0CW137CR2/3":  true,
		"M0CW336CR1/3":  true,
		"M0CW336CR2/3":  true,
		"M0CW1523CR1/3": true,
		"M0CW1523CR2/3": true,
	}
	chipTypes = map[string]bool{
		"SX1255": true,
		"SX1257": true,
	}
	mandatoryFieldsError = errors.New("all mandatory fields should be set")
)

type Definition interface {
	Parse(file string) (Definition, error)
	Validate() error
}

type SubBand struct {
	MinFrequency *int     `yaml:"min-frequency,omitempty"`
	MaxFrequency *int     `yaml:"max-frequency,omitempty"`
	DutyCyle     *float32 `yaml:"duty-cycle,omitempty"`
	MaxEIRP      *float32 `yaml:"max-eirp,omitempty"`
}

func (s SubBand) Validate() error {
	if s.MinFrequency == nil || s.MaxFrequency == nil {
		return mandatoryFieldsError
	}
	if err := validateFrequency(*s.MinFrequency); err != nil {
		return err
	}
	if err := validateFrequency(*s.MaxFrequency); err != nil {
		return err
	}
	if *s.MinFrequency >= *s.MaxFrequency {
		return errors.New("min-frequency should not be the same as or bigger than max-frequency")
	}
	return nil
}

type Channel struct {
	UplinkFrequency   *int    `yaml:"uplink-frequency"`
	DownlinkFrequency *int    `yaml:"downlink-frequency,omitempty"`
	MinDataRate       *string `yaml:"min-data-rate"`
	MaxDataRate       *string `yaml:"max-data-rate"`
	Radio             *int    `yaml:"radio,omitempty"`
	Default           *bool   `yaml:"default"`
}

func (c Channel) Validate() error {
	if c.UplinkFrequency == nil || c.MinDataRate == nil || c.MaxDataRate == nil {
		return mandatoryFieldsError
	}
	if err := validateFrequency(*c.UplinkFrequency); err != nil {
		return err
	}
	if c.DownlinkFrequency != nil {
		if err := validateFrequency(*c.DownlinkFrequency); err != nil {
			return err
		}
	}
	if err := validateDataRate(*c.MinDataRate); err != nil {
		return err
	}
	if err := validateDataRate(*c.MaxDataRate); err != nil {
		return err
	}
	return nil
}

type FrequencyDataRate struct {
	Frequency *int    `yaml:"frequency,omitempty"`
	DataRate  *string `yaml:"data-rate,omitempty"`
	Radio     *int    `yaml:"radio,omitempty"`
}

func (f FrequencyDataRate) Validate() error {
	if f.Frequency == nil || f.DataRate == nil {
		return mandatoryFieldsError
	}
	if err := validateFrequency(*f.Frequency); err != nil {
		return err
	}
	if err := validateDataRate(*f.DataRate); err != nil {
		return err
	}
	return nil
}

type TimeOffAir struct {
	Fraction *float32       `yaml:"fraction,omitempty"`
	Duration *time.Duration `yaml:"duration,omitempty"`
}

func (t TimeOffAir) Validate() error {
	if t.Fraction == nil || t.Duration == nil {
		return mandatoryFieldsError
	}
	return nil
}

type DwellTime struct {
	Uplinks   *bool          `yaml:"uplinks,omitempty"`
	Downlinks *bool          `yaml:"downlinks,omitempty"`
	Duration  *time.Duration `yaml:"duration,omitempty"`
}

func (d DwellTime) Validate() error {
	if d.Uplinks == nil || d.Downlinks == nil || d.Duration == nil {
		return mandatoryFieldsError
	}
	return nil
}

type ListenBeforeTalk struct {
	RSSIOffset *int `yaml:"rssi-offset,omitempty"`
	RSSITarget *int `yaml:"rssi-targer,omitempty"`
	ScanTime   *int `yaml:"scan-time,omitempty"`
}

func (l ListenBeforeTalk) Validate() error {
	if l.RSSIOffset == nil || l.RSSITarget == nil || l.ScanTime == nil {
		return mandatoryFieldsError
	}
	return nil
}

type Radio struct {
	Enable     *bool   `yaml:"enable,omitempty"`
	ChipType   *string `yaml:"chip-type,omitempty"`
	Frequency  *int    `yaml:"frequency,omitempty"`
	RSSIOffset *int    `yaml:"rssi-offset,omitempty"`
	TX         *TX     `yaml:"tx,omitempty"`
}

func (r Radio) Validate() error {
	if r.Enable == nil || r.ChipType == nil || r.Frequency == nil || r.RSSIOffset == nil {
		return mandatoryFieldsError
	}
	if !chipTypes[*r.ChipType] {
		return errors.New("Unknown chip type")
	}
	if err := validateFrequency(*r.Frequency); err != nil {
		return err
	}
	return nil
}

type TX struct {
	MinFrequency   *int `yaml:"min-frequency,omitempty"`
	MaxFrequency   *int `yaml:"max-frequency,omitempty"`
	NotchFrequency *int `yaml:"notch-frequency,omitempty"`
}

func (tx TX) Validate() error {
	if tx.MinFrequency == nil || tx.MaxFrequency == nil {
		return mandatoryFieldsError
	}
	if err := validateFrequency(*tx.MinFrequency); err != nil {
		return err
	}
	if err := validateFrequency(*tx.MaxFrequency); err != nil {
		return err
	}
	if err := validateFrequencyRange(*tx.NotchFrequency, 126000, 250000); err != nil {
		return err
	}
	return nil
}

type FrequencyMinMaxDataRate struct {
	Frequency   *int    `yaml:"frequency,omitempty"`
	MinDataRate *string `yaml:"min-data-rate,omitempty"`
	MaxDataRate *string `yaml:"max-data-rate,omitempty"`
}

func (f FrequencyMinMaxDataRate) Validate() error {
	if f.Frequency == nil || f.MaxDataRate == nil || f.MinDataRate == nil {
		return mandatoryFieldsError
	}
	if err := validateFrequency(*f.Frequency); err != nil {
		return err
	}
	if err := validateDataRate(*f.MinDataRate); err != nil {
		return err
	}
	if err := validateDataRate(*f.MaxDataRate); err != nil {
		return err
	}
	return nil
}

func validateDataRate(dataRate string) error {
	if _, err := strconv.Atoi(dataRate); err == nil || dataRates[dataRate] {
		return nil
	}
	return fmt.Errorf("data rate %s is neither a known datarate nor an integer", dataRate)
}

func validateFrequency(frequency int) error {
	if frequency < 0 {
		return errors.New("frequencies can't be negative")
	}
	return nil
}

func validateFrequencyRange(frequency int, min, max int) error {
	if err := validateFrequency(frequency); err != nil {
		return err
	}
	if frequency < min || frequency > max {
		return fmt.Errorf("frequency should be between %d and %d Hz", min, max)
	}
	return nil
}

func set[T any](modifier *T, base T) {
	if modifier != nil {
		base = *modifier
	}
}

func setPointer[T any](modifier *T, base *T) {
	if modifier != nil {
		if base == nil {
			base = new(T)
		}
		*base = *modifier
	}
}
