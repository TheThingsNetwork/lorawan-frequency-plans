package schema

import (
	"bytes"
	"embed"
	"os"
	"sort"
	"text/template"

	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/utils"
)

//go:embed "*.tmpl"
var fsys embed.FS

var (
	tmpl   = template.Must(template.ParseFS(fsys, "*.tmpl"))
	schema = Schema{
		BandIDs:     getBandIDs(),
		PhyVersions: getPhyVersions(),
	}
)

type Schema struct {
	BandIDs            string
	EndDeviceBaseIDs   string
	GatewayBaseIDs     string
	EndDeviceModifiers string
	GatewayModifiers   string
	EndDeviceFiles     string
	GatewayFiles       string
	PhyVersions        string
}

func Generate() error {
	var err error
	schema.EndDeviceBaseIDs, schema.GatewayBaseIDs, err = getBaseIDs()
	if err != nil {
		return err
	}
	schema.EndDeviceModifiers, err = getYamlFileNames("end-device/modifiers")
	if err != nil {
		return err
	}
	schema.GatewayModifiers, err = getYamlFileNames("gateway/modifiers")
	if err != nil {
		return err
	}
	schema.EndDeviceFiles, err = getYamlFileNames("end-device")
	if err != nil {
		return err
	}
	schema.GatewayFiles, err = getYamlFileNames("gateway")
	if err != nil {
		return err
	}

	executeTemplate("frequency-plans-description-schema.json.tmpl", "schema.json")
	executeTemplate("frequency-plans-end-device-schema.json.tmpl", "end-device/schema.json")
	executeTemplate("frequency-plans-end-device-modifiers-schema.json.tmpl", "end-device/modifiers/schema.json")
	executeTemplate("frequency-plans-gateway-schema.json.tmpl", "gateway/schema.json")
	executeTemplate("frequency-plans-gateway-modifiers-schema.json.tmpl", "gateway/modifiers/schema.json")

	return nil
}

func executeTemplate(file, output string) error {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, file, schema); err != nil {
		return err
	}
	if err := os.WriteFile(output, buf.Bytes(), 0o644); err != nil {
		return err
	}
	return nil
}

func getBaseIDs() (string, string, error) {
	endDevice, gateway, err := utils.GetBaseIDs()
	if err != nil {
		return "", "", err
	}
	return buildCSV(endDevice), buildCSV(gateway), nil
}

func getBandIDs() string {
	return buildCSV(utils.GetBandIDs())
}

func getYamlFileNames(folder string) (string, error) {
	filenames, err := utils.GetYamlFileNames(folder)
	if err != nil {
		return "", err
	}
	return buildCSV(filenames), nil
}

func getPhyVersions() string {
	return buildCSV(utils.GetPhyVersions())
}

func buildCSV(inputs map[string]bool) (output string) {
	keys := []string{}
	for input := range inputs {
		keys = append(keys, input)
	}
	sort.Strings(keys)
	first := true
	for _, key := range keys {
		if !first {
			output += ","
		}
		output += "\"" + key + "\""
		first = false
	}
	return output
}
