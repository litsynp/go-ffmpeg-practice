package main

import (
	"fmt"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("filename is required")
	}

	filename := args[0]

	fmt.Println("filename: ", filename)

	compressVideo(filename)
}

func compressVideo(videoFilePath string) {
	err := ffmpeg.Input("./media/test_video.mp4").
		Output("./media/test_output.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		log.Fatal(err)
	}
}
