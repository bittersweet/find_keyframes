package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func calculateSeconds(input string) float64 {
	splitted_duration := strings.Split(input, ":")
	hours, err := strconv.ParseFloat(splitted_duration[0], 32)
	if err != nil {
		panic("Wrong input")
	}
	minutes, err := strconv.ParseFloat(splitted_duration[1], 32)
	if err != nil {
		panic("Wrong input")
	}
	seconds, err := strconv.ParseFloat(splitted_duration[2], 32)
	if err != nil {
		panic("Wrong input")
	}

	return (hours * 60 * 60) + (minutes * 60) + seconds
}

func extractKeyframes(file string) []float64 {
	command := exec.Command("ffprobe", "-show_frames", "-select_streams", "v", "-print_format", "json=c=1", file)
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Fatal("Failure with command:", err)
	}
	if err := command.Start(); err != nil {
    log.Fatal("Failure on starting command:", err)
	}
	type Frame struct {
		KeyFrame int    `json:"key_frame"`
		Time     string `json:"pkt_dts_time"`
	}
	type Frames struct {
		Frames []Frame
	}
	var f Frames
	if err := json.NewDecoder(stdout).Decode(&f); err != nil {
    log.Fatal("Decoding error:", err)
	}
	if err := command.Wait(); err != nil {
    log.Fatal("Failure on command wait:", err)
	}
	keyframes := make([]float64, 0)
	for _, v := range f.Frames {
		if v.KeyFrame == 1 {
			time, _ := strconv.ParseFloat(v.Time, 32)
			if err != nil {
				panic("Wrong input")
			}
			keyframes = append(keyframes, time)
		}
	}
	fmt.Println("Keyframes:", len(keyframes))

	return keyframes
}

func main() {
	filename := os.Args[1]
	fmt.Println("Processing", filename)
	start := calculateSeconds(os.Args[2])
	fmt.Printf("Input HH:MM:SS: %s\n", os.Args[2])
	fmt.Printf("Start in seconds: %f\n", start)

	frames := extractKeyframes(filename)
	var closest_keyframe float64
	var next_keyframe float64
	for index, value := range frames {
		if value > start {
			closest_keyframe = frames[index-1]
			next_keyframe = value
			break
		}
	}

	fmt.Printf("Closest: %f Next: %f\n", closest_keyframe, next_keyframe)

}
