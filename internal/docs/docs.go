package docs

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/model"
)

//go:embed "*.tmpl"
var fsys embed.FS

var tmpl = template.Must(template.ParseFS(fsys, "*.tmpl"))

// Generate generates the documentation for the frequency-plans.
func Generate(sourceFile, destinationFolder string) error {
	output, err := model.FrequencyPlanDescriptions{}.Parse(sourceFile)
	if err != nil {
		return err
	}
	descriptions := output.(model.FrequencyPlanDescriptions)

	if err := renderPlans("gateway", descriptions.GatewayDescriptions, model.FrequencyPlanGateway{}); err != nil {
		return err
	}

	if err := renderPlans("end-device", descriptions.EndDeviceDescriptions, model.FrequencyPlanEndDevice{}); err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "frequency-plans.md.tmpl", output); err != nil {
		return err
	}
	if err := os.WriteFile(destinationFolder+"/frequency-plans.md", buf.Bytes(), 0o644); err != nil {
		return err
	}
	return nil
}

func formatFrequency(f float64) string {
	return strings.TrimRight(fmt.Sprintf("%.3f", f/1_000_000), "0.")
}

func renderPlans(folder string, descriptions []model.FrequencyPlanDescription, definition model.Definition) error {
	for _, plan := range descriptions {
		fileName := ""
		if plan.HasModifiers() {
			for _, description := range descriptions {
				if description.ID == *plan.BaseID {
					fileName = *description.File
					break
				}
			}
		} else {
			fileName = *plan.File
		}
		basePlan, err := definition.Parse(folder + "/" + fileName)
		if err != nil {
			return err
		}
		if plan.HasModifiers() {
			for _, modifierName := range *plan.Modifiers {
				switch definition.(type) {
				case model.FrequencyPlanEndDevice:
					modifier, err := model.FrequencyPlanEndDeviceModifier{}.Parse(folder + "/modifiers/" + modifierName)
					if err != nil {
						return err
					}
					basePlan = basePlan.(model.FrequencyPlanEndDevice).Modify(modifier.(model.FrequencyPlanEndDeviceModifier))
				case model.FrequencyPlanGateway:
					modifier, err := model.FrequencyPlanGatewayModifier{}.Parse(folder + "/modifiers/" + modifierName)
					if err != nil {
						return err
					}
					basePlan = basePlan.(model.FrequencyPlanGateway).Modify(modifier.(model.FrequencyPlanGatewayModifier))
				}
			}
		}
		if err := render(plan.ID, basePlan); err != nil {
			return err
		}
	}
	return nil
}

func render(id string, definition model.Definition) error {
	switch mod := definition.(type) {
	case model.FrequencyPlanGateway:
		return renderGateway(id, mod)
	case model.FrequencyPlanEndDevice:
		return renderEndDevice(id, mod)
	default:
		return errors.New("unsupported type")
	}
}
