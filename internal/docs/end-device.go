package docs

import (
	"bytes"
	"os"
	"sort"

	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/model"
	"github.com/wcharczuk/go-chart/v2"
)

func renderEndDevice(id string, plan model.FrequencyPlanEndDevice) error {
	frequencies := make(map[float64]string)

	graph := chart.Chart{
		Title:  id,
		Width:  1920,
		Height: 1080,
		DPI:    150,
	}

	annotations := chart.AnnotationSeries{}

	for _, ch := range plan.Channels {
		freq := float64(*ch.UplinkFrequency)
		start, end := freq-62500, freq+62500
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(0)
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

		if ch.DownlinkFrequency == nil {
			continue
		}
		freq = float64(*ch.DownlinkFrequency)
		start, end = freq-62500, freq+62500
		frequencies[freq] = formatFrequency(freq)
		color = chart.GetDefaultColor(3)
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

	if ch := plan.LoRaStandardChannel; ch != nil {
		freq := float64(*ch.Frequency)
		start, end := freq-125000, freq+125000
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(1)
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
		freq := float64(*ch.Frequency)
		start, end := freq-62500, freq+62500
		frequencies[freq] = formatFrequency(freq)
		color := chart.GetDefaultColor(2)
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

	if err := graph.Render(chart.SVG, &buf); err != nil {
		return err
	}
	if err := os.WriteFile("./docs/images/end-device/"+id+".svg", buf.Bytes(), 0o644); err != nil {
		return err
	}

	return nil
}
