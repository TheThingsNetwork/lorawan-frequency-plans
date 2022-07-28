package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/wcharczuk/go-chart/v2"
	"gopkg.in/yaml.v3"
)

//go:embed "*.tmpl"
var fsys embed.FS

var tmpl = template.Must(template.ParseFS(fsys, "*.tmpl"))

func main() {
	indexBytes, err := os.ReadFile("frequency-plans.yml")
	if err != nil {
		log.Fatal(err)
	}
	var index []struct {
		ID          string `yaml:"id"`
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		BaseID      string `yaml:"base-id"`
		File        string `yaml:"file"`
	}
	if err = yaml.Unmarshal(indexBytes, &index); err != nil {
		log.Fatal(err)
	}
	for _, plan := range index {
		if plan.BaseID != "" {
			log.Printf("Skipping %s: extending a base plan not supported yet", plan.ID)
			continue
		}
		if err := render(plan.ID, plan.File); err != nil {
			log.Fatal(err)
		}
	}

	var buf bytes.Buffer
	if err = tmpl.ExecuteTemplate(&buf, "frequency-plans.md.tmpl", index); err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile("frequency-plans.md", buf.Bytes(), 0o644); err != nil {
		log.Fatal(err)
	}
}

func formatFrequency(f float64) string {
	return strings.TrimRight(fmt.Sprintf("%.3f", f/1_000_000), "0.")
}

func render(id, file string) error {
	planBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var plan struct {
		SubBands []struct {
			MinFrequency uint64 `yaml:"min-frequency"`
			MaxFrequency uint64 `yaml:"max-frequency"`
		} `yaml:"sub-bands"`
		UplinkChannels []struct {
			Frequency uint64 `yaml:"frequency"`
			Radio     uint8  `yaml:"radio"`
		} `yaml:"uplink-channels"`
		DownlinkChannels []struct {
			Frequency uint64 `yaml:"frequency"`
			Radio     uint8  `yaml:"radio"`
		} `yaml:"downlink-channels"`
		LoRaStandardChannel *struct {
			Frequency uint64 `yaml:"frequency"`
			Radio     uint8  `yaml:"radio"`
		} `yaml:"lora-standard-channel"`
		FSKChannel *struct {
			Frequency uint64 `yaml:"frequency"`
			Radio     uint8  `yaml:"radio"`
		} `yaml:"fsk-channel"`
		Radios []struct {
			Frequency uint64 `yaml:"frequency"`
		} `yaml:"radios"`
	}
	if err = yaml.Unmarshal(planBytes, &plan); err != nil {
		return err
	}

	frequencies := make(map[float64]string)

	graph := chart.Chart{
		Title:  id,
		Width:  1920,
		Height: 1080,
		DPI:    150,
	}

	annotations := chart.AnnotationSeries{}

	for _, ch := range plan.UplinkChannels {
		freq := float64(ch.Frequency)
		start, end := freq-62500, freq+62500
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(int(ch.Radio))
		color.A = 128
		graph.Series = append(graph.Series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeColor: color,
				FillColor:   color,
			},
			XValues: []float64{start, end},
			YValues: []float64{float64(1), float64(1)},
		})
		annotations.Annotations = append(annotations.Annotations, chart.Value2{
			XValue: freq,
			YValue: 1,
			Label:  formatFrequency(freq),
		})
	}

	if ch := plan.LoRaStandardChannel; ch != nil {
		freq := float64(ch.Frequency)
		start, end := freq-125000, freq+125000
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(int(ch.Radio))
		color.A = 128
		graph.Series = append(graph.Series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeColor: color,
				FillColor:   color,
			},
			XValues: []float64{start, end},
			YValues: []float64{float64(2), float64(2)},
		})
		annotations.Annotations = append(annotations.Annotations, chart.Value2{
			XValue: freq,
			YValue: 2,
			Label:  formatFrequency(freq) + " (Std)",
		})
	}

	if ch := plan.FSKChannel; ch != nil {
		freq := float64(ch.Frequency)
		start, end := freq-62500, freq+62500
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(int(ch.Radio))
		color.A = 128
		graph.Series = append(graph.Series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeColor: color,
				FillColor:   color,
			},
			XValues: []float64{start, end},
			YValues: []float64{float64(2), float64(2)},
		})
		annotations.Annotations = append(annotations.Annotations, chart.Value2{
			XValue: freq,
			YValue: 2,
			Label:  formatFrequency(freq) + " (FSK)",
		})
	}

	for _, ch := range plan.DownlinkChannels {
		freq := float64(ch.Frequency)
		start, end := freq-62500, freq+62500
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(3)
		color.A = 128
		graph.Series = append(graph.Series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeColor: color,
				FillColor:   color,
			},
			XValues: []float64{start, end},
			YValues: []float64{float64(-1), float64(-1)},
		})
		annotations.Annotations = append(annotations.Annotations, chart.Value2{
			XValue: freq,
			YValue: -1,
			Label:  formatFrequency(freq),
		})
	}

	for i, radio := range plan.Radios {
		freq := float64(radio.Frequency)
		start, end := freq-462500, freq+462500
		frequencies[start] = formatFrequency(start)
		frequencies[freq] = formatFrequency(freq)
		frequencies[end] = formatFrequency(end)
		color := chart.GetDefaultColor(i)
		color.A = 128
		graph.Series = append(graph.Series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeColor: color,
				StrokeWidth: 10,
			},
			XValues: []float64{start, end},
			YValues: []float64{float64(0), float64(0)},
		})
		annotations.Annotations = append(annotations.Annotations, chart.Value2{
			XValue: freq,
			YValue: 0,
			Label:  fmt.Sprintf("Radio %d: %s", i, formatFrequency(freq)),
		})
	}

	var frequencySlice []float64
	for frequency := range frequencies {
		frequencySlice = append(frequencySlice, frequency)
	}
	sort.Float64s(frequencySlice)

	graph.XAxis = chart.XAxis{
		Range: &chart.ContinuousRange{
			Min: frequencySlice[0],
			Max: frequencySlice[len(frequencySlice)-1],
		},
		TickStyle: chart.Style{
			TextRotationDegrees: 45.0,
		},
	}

	for _, freq := range frequencySlice {
		graph.XAxis.Ticks = append(graph.XAxis.Ticks, chart.Tick{
			Value: freq,
			Label: frequencies[freq],
		})
	}

	graph.YAxis = chart.YAxis{
		Range: &chart.ContinuousRange{
			Min: -2,
			Max: 3,
		},
		Ticks: []chart.Tick{
			{Value: -2},
			{Value: -1, Label: "Downlink"},
			{Value: 0, Label: "Radio"},
			{Value: 1, Label: "Uplink"},
			{Value: 2, Label: "Std/FSK"},
			{Value: 3},
		},
	}

	graph.Series = append(graph.Series, annotations)

	var buf bytes.Buffer

	if err = graph.Render(chart.SVG, &buf); err != nil {
		return err
	}
	if err = os.WriteFile(file+".svg", buf.Bytes(), 0o644); err != nil {
		return err
	}

	return nil
}
