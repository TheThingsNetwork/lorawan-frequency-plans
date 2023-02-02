package model

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type FrequencyPlanEndDevice struct {
	BandID                  string                   `yaml:"band-id"`
	SubBands                []SubBand                `yaml:"sub-bands"`
	Channels                []Channel                `yaml:"channels"`
	LoRaStandardChannel     *FrequencyDataRate       `yaml:"lora-special-channel,omitempty"`
	FSKChannel              *FrequencyDataRate       `yaml:"fsk-channel,omitempty"`
	TimeOffAir              *TimeOffAir              `yaml:"time-off-air,omitempty"`
	DwellTime               *DwellTime               `yaml:"dwell-time,omitempty"`
	ListenBeforeTalk        *ListenBeforeTalk        `yaml:"listen-before-talk,omitempty"`
	PingSlot                *FrequencyMinMaxDataRate `yaml:"ping-slot,omitempty"`
	PingSlotDefaultDataRate *string                  `yaml:"ping-slot-default-data-rate,omitempty"`
	RX2Channel              *FrequencyMinMaxDataRate `yaml:"rx2-channel,omitempty"`
	RX2DefaultDataRate      *string                  `yaml:"rx2-default-data-rate,omitempty"`
	MaxEIRP                 *float32                 `yaml:"max-eirp,omitempty"`
}

func (f FrequencyPlanEndDevice) Parse(file string) (Definition, error) {
	indexBytes, err := os.ReadFile(file)
	if err != nil {
		return FrequencyPlanEndDevice{}, err
	}
	err = yaml.Unmarshal(indexBytes, &f)
	return f, err
}

func (f FrequencyPlanEndDevice) Validate() error {
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
	if f.ListenBeforeTalk != nil {
		if err := f.ListenBeforeTalk.Validate(); err != nil {
			return fmt.Errorf("ListenBeforeTalk: %w", err)
		}
	}
	if f.PingSlot != nil {
		if err := f.PingSlot.Validate(); err != nil {
			return fmt.Errorf("PingSlot: %w", err)
		}
	}
	if f.PingSlotDefaultDataRate != nil {
		if err := validateDataRate(*f.PingSlotDefaultDataRate); err != nil {
			return fmt.Errorf("PingSlotDefaultDataRate: %w", err)
		}
	}
	if f.RX2Channel != nil {
		if err := f.RX2Channel.Validate(); err != nil {
			return fmt.Errorf("RX2Channel: %w", err)
		}
	}
	if f.RX2DefaultDataRate != nil {
		if err := validateDataRate(*f.RX2DefaultDataRate); err != nil {
			return fmt.Errorf("RX2DefaultDataRate: %w", err)
		}
	}
	return nil
}

func (f FrequencyPlanEndDevice) Modify(modifier FrequencyPlanEndDeviceModifier) FrequencyPlanEndDevice {
	modified := f
	set(modifier.SubBands, modified.SubBands)
	set(modifier.Channels, modified.Channels)
	setPointer(modifier.LoRaStandardChannel, modified.LoRaStandardChannel)
	setPointer(modifier.FSKChannel, modified.FSKChannel)
	setPointer(modifier.TimeOffAir, modified.TimeOffAir)
	setPointer(modifier.DwellTime, modified.DwellTime)
	setPointer(modifier.ListenBeforeTalk, modified.ListenBeforeTalk)
	setPointer(modifier.PingSlot, modified.PingSlot)
	setPointer(modifier.PingSlotDefaultDataRate, modified.PingSlotDefaultDataRate)
	setPointer(modifier.RX2Channel, modified.RX2Channel)
	setPointer(modifier.RX2DefaultDataRate, modified.RX2DefaultDataRate)
	setPointer(modifier.MaxEIRP, modified.MaxEIRP)
	return modified
}

type FrequencyPlanEndDeviceModifier struct {
	SubBands                *[]SubBand               `yaml:"sub-bands,omitempty"`
	Channels                *[]Channel               `yaml:"channels,omitempty"`
	LoRaStandardChannel     *FrequencyDataRate       `yaml:"lora-special-channel,omitempty"`
	FSKChannel              *FrequencyDataRate       `yaml:"fsk-channel,omitempty"`
	TimeOffAir              *TimeOffAir              `yaml:"time-off-air,omitempty"`
	DwellTime               *DwellTime               `yaml:"dwell-time,omitempty"`
	ListenBeforeTalk        *ListenBeforeTalk        `yaml:"listen-before-talk,omitempty"`
	PingSlot                *FrequencyMinMaxDataRate `yaml:"ping-slot,omitempty"`
	PingSlotDefaultDataRate *string                  `yaml:"ping-slot-default-data-rate,omitempty"`
	RX2Channel              *FrequencyMinMaxDataRate `yaml:"rx2-channel,omitempty"`
	RX2DefaultDataRate      *string                  `yaml:"rx2-default-data-rate,omitempty"`
	MaxEIRP                 *float32                 `yaml:"max-eirp,omitempty"`
}

func (f FrequencyPlanEndDeviceModifier) Parse(file string) (Definition, error) {
	indexBytes, err := os.ReadFile(file)
	if err != nil {
		return FrequencyPlanEndDeviceModifier{}, err
	}
	err = yaml.Unmarshal(indexBytes, &f)
	return f, err
}

func (f FrequencyPlanEndDeviceModifier) Validate() error {
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
	if f.ListenBeforeTalk != nil {
		if err := f.ListenBeforeTalk.Validate(); err != nil {
			return fmt.Errorf("ListenBeforeTalk: %w", err)
		}
	}
	if f.PingSlot != nil {
		if err := f.PingSlot.Validate(); err != nil {
			return fmt.Errorf("PingSlot: %w", err)
		}
	}
	if f.PingSlotDefaultDataRate != nil {
		if err := validateDataRate(*f.PingSlotDefaultDataRate); err != nil {
			return fmt.Errorf("PingSlotDefaultDataRate: %w", err)
		}
	}
	if f.RX2Channel != nil {
		if err := f.RX2Channel.Validate(); err != nil {
			return fmt.Errorf("RX2Channel: %w", err)
		}
	}
	if f.RX2DefaultDataRate != nil {
		if err := validateDataRate(*f.RX2DefaultDataRate); err != nil {
			return fmt.Errorf("RX2DefaultDataRate: %w", err)
		}
	}
	return nil
}
