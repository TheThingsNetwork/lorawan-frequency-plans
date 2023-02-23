package utils

import (
	"os"
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"gopkg.in/yaml.v2"
)

// GetBaseIDs returns the base ids for end-devices and gateways
func GetBaseIDs() (map[string]bool, map[string]bool, error) {
	plans := struct {
		EndDevices []struct {
			ID string `yaml:"id"`
		} `yaml:"end-device-descriptions"`
		Gateways []struct {
			ID string `yaml:"id"`
		} `yaml:"gateway-descriptions"`
	}{}
	indexBytes, err := os.ReadFile("frequency-plans.yml")
	if err != nil {
		return nil, nil, err
	}
	err = yaml.Unmarshal(indexBytes, &plans)
	if err != nil {
		return nil, nil, err
	}

	EndDeviceBaseIDs := make(map[string]bool)
	for _, description := range plans.EndDevices {
		if description.ID == "" {
			continue
		}
		EndDeviceBaseIDs[description.ID] = true
	}
	GatewayBaseIDs := make(map[string]bool)
	for _, description := range plans.EndDevices {
		if description.ID == "" {
			continue
		}
		GatewayBaseIDs[description.ID] = true
	}
	return EndDeviceBaseIDs, GatewayBaseIDs, nil
}

func GetBandIDs() map[string]bool {
	bandIDs := make(map[string]bool)
	for band := range band.All {
		bandIDs[band] = true
	}
	return bandIDs
}

func GetYamlFileNames(folder string) (map[string]bool, error) {
	fileNames := make(map[string]bool)
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".yml") {
			continue
		}
		fileNames[file.Name()] = true
	}
	return fileNames, nil
}

func GetPhyVersions() map[string]bool {
	phyVersions := make(map[string]bool)
	for _, phyVersion := range ttnpb.PHYVersion_name {
		phyVersions[phyVersion] = true
	}
	return phyVersions
}
