package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/fogleman/gg"
	"github.com/go-audio/wav"
)

const (
	imgSize         = 1024
	radius          = 300
	spikeLen        = 1000
	spikesPerCircle = 360
	audioPath       = "output.wav" // Must be PCM
	sliceSeconds    = 1            // duration of each frame in seconds
)

func main() {
	file, err := os.Open(audioPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	if !decoder.IsValidFile() {
		log.Fatal("Invalid WAV file")
	}

	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		log.Fatal(err)
	}

	sampleRate := decoder.SampleRate
	channels := buf.Format.NumChannels
	rawData := buf.Data
	data := make([]int, len(rawData)/channels)
	for i := 0; i < len(data); i++ {
		sum := 0
		for c := 0; c < channels; c++ {
			sum += rawData[i*channels+c]
		}
		data[i] = sum / channels
	}

	samplesPerSlice := int64(sampleRate * sliceSeconds)
	totalSamples := int64(len(data))
	totalSlices := int(totalSamples / samplesPerSlice)
	fmt.Printf("Generating %d frames...\n", totalSlices)

	for i := 0; i < totalSlices; i++ {
		start := i * int(samplesPerSlice)
		end := start + int(samplesPerSlice)
		if end > len(data) {
			end = len(data)
		}
		frameData := data[start:end]

		img := drawWaveformCircle(frameData, spikesPerCircle)
		filename := fmt.Sprintf("images/frame_%04d.png", i+1)
		img.SavePNG(filename)
	}
}

func drawWaveformCircle(data []int, points int) *gg.Context {
	dc := gg.NewContext(imgSize, imgSize)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(1)

	cx, cy := float64(imgSize/2), float64(imgSize/2)

	// Normalize audio data to [-1, 1]
	norm := normalize(data, points)

	dc.NewSubPath()
	for i := 0; i < points; i++ {
		angle := 2 * math.Pi * float64(i) / float64(points)
		amp := norm[i]
		r := float64(radius) + amp*spikeLen // now spikeLen is more like "waveHeight"

		x := cx + r*math.Cos(angle)
		y := cy + r*math.Sin(angle)

		if i == 0 {
			dc.MoveTo(x, y)
		} else {
			dc.LineTo(x, y)
		}
	}
	dc.ClosePath()
	dc.Stroke()

	return dc
}

func normalize(data []int, n int) []float64 {
	step := len(data) / n
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < step; j++ {
			idx := i*step + j
			if idx >= len(data) {
				break
			}
			sum += int(math.Abs(float64(data[idx])))
		}
		avg := float64(sum) / float64(step) / float64(1<<15)
		out[i] = avg
	}
	return out
}
