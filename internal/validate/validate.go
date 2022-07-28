package validate

import (
	"fmt"

	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/model"
	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/utils"
)

// Validate validates the defined frequency-plans.
func Validate() error {
	if err := validateFolder(model.FrequencyPlanEndDevice{}, "end-device"); err != nil {
		return err
	}
	if err := validateFolder(model.FrequencyPlanEndDeviceModifier{}, "end-device/modifiers"); err != nil {
		return err
	}
	if err := validateFolder(model.FrequencyPlanGateway{}, "gateway"); err != nil {
		return err
	}
	if err := validateFolder(model.FrequencyPlanGatewayModifier{}, "gateway/modifiers"); err != nil {
		return err
	}
	if err := validateFile(model.FrequencyPlanDescriptions{}, "frequency-plans.yml"); err != nil {
		return err
	}
	return nil
}

func validateFile(def model.Definition, file string) error {
	definition, err := def.Parse(file)
	if err != nil {
		return err
	}
	return definition.Validate()
}

func validateFolder(def model.Definition, folder string) error {
	files, err := utils.GetYamlFileNames(folder)
	if err != nil {
		return err
	}
	for file := range files {
		if file == "" {
			continue
		}
		definition, err := def.Parse(folder + "/" + file)
		if err != nil {
			return err
		}
		err = definition.Validate()
		if err != nil {
			return fmt.Errorf("%s/%s: %s", folder, file, err)
		}
	}
	return nil
}
