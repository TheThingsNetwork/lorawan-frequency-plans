package model

import (
	"fmt"
	"os"

	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/utils"
	"gopkg.in/yaml.v2"
)

type FrequencyPlanDescriptions struct {
	EndDeviceDescriptions []FrequencyPlanDescription `yaml:"end-device-descriptions"`
	GatewayDescriptions   []FrequencyPlanDescription `yaml:"gateway-descriptions"`
}

func (f FrequencyPlanDescriptions) Parse(file string) (Definition, error) {
	indexBytes, err := os.ReadFile(file)
	if err != nil {
		return FrequencyPlanDescriptions{}, err
	}
	err = yaml.Unmarshal(indexBytes, &f)
	if err != nil {
		return nil, err
	}
	return f, err
}

func (f FrequencyPlanDescriptions) Validate() error {
	for _, description := range f.EndDeviceDescriptions {
		if err := description.Validate(EndDeviceFrequencyPlan); err != nil {
			return err
		}
	}
	for _, description := range f.GatewayDescriptions {
		if err := description.Validate(GatewayFrequencyPlan); err != nil {
			return err
		}
	}
	return nil
}

type FrequencyPlanType string

var (
	GatewayFrequencyPlan   FrequencyPlanType = "gateway"
	EndDeviceFrequencyPlan FrequencyPlanType = "end-device"
)

// EndDeviceFrequencyPlanDescription describes an end device frequency plan in the YAML format.
type FrequencyPlanDescription struct {
	ID            string    `yaml:"id"`
	BandID        string    `yaml:"band-id"`
	BaseID        *string   `yaml:"base-id,omitempty"`
	Name          string    `yaml:"name"`
	Description   string    `yaml:"description"`
	BaseFrequency uint16    `yaml:"base-frequency"`
	CountryCodes  []string  `yaml:"country-codes"`
	File          *string   `yaml:"file,omitempty"`
	Modifiers     *[]string `yaml:"modifiers,omitempty"`
	Endorsed      bool      `yaml:"endorsed"`
}

func (f FrequencyPlanDescription) Validate(source FrequencyPlanType) error {
	bandIDs := utils.GetBandIDs()
	if !bandIDs[f.BandID] {
		return fmt.Errorf("Frequency plan %s: BandID is invalid", f.ID)
	}
	if f.BaseID != nil && f.File != nil {
		return fmt.Errorf("Frequency plan %s: BaseID may only be defined with Modifiers and not files", f.ID)
	}
	if f.File != nil && f.Modifiers != nil {
		return fmt.Errorf("Frequency plan %s: Either define files or modifiers", f.ID)
	}

	EndDeviceBaseIDs, GatewayBaseIDs, err := utils.GetBaseIDs()
	if err != nil {
		return err
	}
	switch source {
	case EndDeviceFrequencyPlan:
		if f.BaseID != nil && !EndDeviceBaseIDs[*f.BaseID] {
			return fmt.Errorf("Frequency plan %s: Set BaseID doesn't exist", f.ID)
		}
	case GatewayFrequencyPlan:
		if f.BaseID != nil && !GatewayBaseIDs[*f.BaseID] {
			return fmt.Errorf("Frequency plan %s: Set BaseID doesn't exist", f.ID)
		}
	}

	if f.Modifiers != nil {
		modifiers, err := utils.GetYamlFileNames(string(source) + "/modifiers")
		if err != nil {
			return err
		}
		for _, modifier := range *f.Modifiers {
			if !modifiers[modifier] {
				return fmt.Errorf("Frequency plan %s: %s modifier %s doesn't exist", source, f.ID, modifier)
			}
		}
	}

	if f.File != nil {
		files, err := utils.GetYamlFileNames(string(source))
		if err != nil {
			return err
		}
		if !files[*f.File] {
			return fmt.Errorf("Frequency plan %s: %s file %s doesn't exist", source, f.ID, *f.File)
		}
	}

	return nil
}

func (f FrequencyPlanDescription) HasModifiers() bool {
	return f.BaseID != nil && f.Modifiers != nil
}
