package model

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type FrequencyPlanGateway struct {
	BandID              string             `yaml:"band-id"`
	SubBands            []SubBand          `yaml:"sub-bands"`
	Channels            []Channel          `yaml:"channels"`
	LoRaStandardChannel *FrequencyDataRate `yaml:"lora-standard-channel,omitempty"`
	FSKChannel          *FrequencyDataRate `yaml:"fsk-channel,omitempty"`
	TimeOffAir          *TimeOffAir        `yaml:"time-off-air,omitempty"`
	DwellTime           *DwellTime         `yaml:"dwell-time,omitempty"`
	Radios              []Radio            `yaml:"radios"`
	ClockSource         int                `yaml:"clock-source"`
	MaxEIRP             *float32           `yaml:"max-eirp,omitempty"`
}

func (f FrequencyPlanGateway) Parse(file string) (Definition, error) {
	indexBytes, err := os.ReadFile(file)
	if err != nil {
		return FrequencyPlanGateway{}, err
	}
	err = yaml.Unmarshal(indexBytes, &f)
	return f, err
}

func (f FrequencyPlanGateway) Validate() error {
	for i, subBand := range f.SubBands {
		if err := subBand.Validate(); err != nil {
			return fmt.Errorf("SubBand %d: %w", i, err)
		}
	}
	for i, channel := range f.Channels {
		if err := channel.Validate(); err != nil {
			return fmt.Errorf("Channel %d: %w", i, err)
		}
	}
	if f.LoRaStandardChannel != nil {
		if err := f.LoRaStandardChannel.Validate(); err != nil {
			return fmt.Errorf("LoRaStandardChannel: %w", err)
		}
	}
	if f.FSKChannel != nil {
		if err := f.FSKChannel.Validate(); err != nil {
			return fmt.Errorf("FSKChannel: %w", err)
		}
	}
	if f.TimeOffAir != nil {
		if err := f.TimeOffAir.Validate(); err != nil {
			return fmt.Errorf("TimeOffAir: %w", err)
		}
	}
	if f.DwellTime != nil {
		if err := f.DwellTime.Validate(); err != nil {
			return fmt.Errorf("DwellTime: %w", err)
		}
	}
	for i, radio := range f.Radios {
		if err := radio.Validate(); err != nil {
			return fmt.Errorf("Radio %d: %w", i, err)
		}
	}
	return nil
}

func (f FrequencyPlanGateway) Modify(modifier FrequencyPlanGatewayModifier) FrequencyPlanGateway {
	modified := f
	set(modifier.SubBands, modified.SubBands)
	set(modifier.Channels, modified.Channels)
	set(modifier.Radios, modified.Radios)
	set(modifier.ClockSource, modified.ClockSource)
	setPointer(modifier.LoRaStandardChannel, modified.LoRaStandardChannel)
	setPointer(modifier.FSKChannel, modified.FSKChannel)
	setPointer(modifier.TimeOffAir, modified.TimeOffAir)
	setPointer(modifier.DwellTime, modified.DwellTime)
	setPointer(modifier.MaxEIRP, modified.MaxEIRP)
	return modified
}

type FrequencyPlanGatewayModifier struct {
	SubBands            *[]SubBand         `yaml:"sub-bands,omitempty"`
	Channels            *[]Channel         `yaml:"channels,omitempty"`
	LoRaStandardChannel *FrequencyDataRate `yaml:"lora-standard-channel,omitempty"`
	FSKChannel          *FrequencyDataRate `yaml:"fsk-channel,omitempty"`
	TimeOffAir          *TimeOffAir        `yaml:"time-off-air,omitempty"`
	DwellTime           *DwellTime         `yaml:"dwell-time,omitempty"`
	Radios              *[]Radio           `yaml:"radios,omitempty"`
	ClockSource         *int               `yaml:"clock-source,omitempty"`
	MaxEIRP             *float32           `yaml:"max-eirp,omitempty"`
}

func (f FrequencyPlanGatewayModifier) Parse(file string) (Definition, error) {
	indexBytes, err := os.ReadFile(file)
	if err != nil {
		return FrequencyPlanGatewayModifier{}, err
	}
	err = yaml.Unmarshal(indexBytes, &f)
	return f, err
}

func (f FrequencyPlanGatewayModifier) Validate() error {
	if f.SubBands != nil {
		for i, subBand := range *f.SubBands {
			if err := subBand.Validate(); err != nil {
				return fmt.Errorf("SubBand %d: %w", i, err)
			}
		}
	}
	if f.Channels != nil {
		for i, channel := range *f.Channels {
			if err := channel.Validate(); err != nil {
				return fmt.Errorf("Channel %d: %w", i, err)
			}
		}
	}
	if f.LoRaStandardChannel != nil {
		if err := f.LoRaStandardChannel.Validate(); err != nil {
			return fmt.Errorf("LoRaStandardChannel: %w", err)
		}
	}
	if f.FSKChannel != nil {
		if err := f.FSKChannel.Validate(); err != nil {
			return fmt.Errorf("FSKChannel: %w", err)
		}
	}
	if f.TimeOffAir != nil {
		if err := f.TimeOffAir.Validate(); err != nil {
			return fmt.Errorf("TimeOffAir: %w", err)
		}
	}
	if f.DwellTime != nil {
		if err := f.DwellTime.Validate(); err != nil {
			return fmt.Errorf("DwellTime: %w", err)
		}
	}
	if f.Radios != nil {
		for i, radio := range *f.Radios {
			if err := radio.Validate(); err != nil {
				return fmt.Errorf("Radio %d: %w", i, err)
			}
		}
	}
	return nil
}
