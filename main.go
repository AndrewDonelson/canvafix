package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/xfrr/goffmpeg/media"
	"github.com/xfrr/goffmpeg/transcoder"
)

// outFileName generates a new file name by prepending the given prepend string to the base name of the source file.
// The function expects the source file name to be in the format "#.ext", where # is a number and ext is the file extension.
// The generated file name will be in the format "<prepend> (#).mp4", where # is the base name of the source file.
// If the source file name does not match the expected format, an empty string is returned.
func outFileName(src string, prepend string) string {
	// Split the filename on the "." character
	parts := strings.Split(filepath.Base(src), ".")
	if len(parts) < 2 { // #.ext format expected
		return ""
	}

	return fmt.Sprintf("%s (%s).mp4", prepend, parts[0])
}

// run is a helper function that parses the command-line arguments and prints the value of the Prepend option.
func run(args []string) {
	flags.ParseArgs(&opts, args)
	fmt.Printf("opts.Prepend: %v\n", opts.Prepend)
}

// opts is a struct that holds command-line options for the program.
// The Prepend field is a required string that specifies the prepend file name.
// This field is accessed using the short flag "-p" or the long flag "--prepend".
var opts struct {
	Prepend string `short:"p" long:"prepend" description:"prepend file name" required:"true"`
}

// Saved audio stream to be added to mp4s without audio
var AudioStream *media.Streams

// hasAAC(s media.Stream) bool checks if the given media stream has an AAC audio codec.
func hasAAC(m media.Metadata) bool {
	return m.Streams[1].CodecName == "aac"
}

func main() {
	run(os.Args)

	// read program arguments. should have one argument: the prepend string. if not error and exit
	if len(os.Args) != 4 {
		fmt.Println(
			"This program requires a single argument: the prepend string. It will then find all MP4 files in the current",
			"\ndirectory, prepend the given string to the filename, put the original name (1,2,50,etc) in parenthases and",
			"\nthen use the audio from 1.mp4 on all others.",
			"\nUsage: canvafix --prepend <prepend_file_name>",
		)
	}

	prepend := os.Args[2]
	pattern := "*.mp4"

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var matchingFiles []string
	for _, file := range files {
		match, err := filepath.Match(pattern, file.Name())
		if err != nil {
			log.Fatal("Error:", err)

		}

		if match {
			log.Printf("Processing file: %s...", file.Name())

			// trans is a new instance of the Transcoder type from the transcoder package.
			trans := new(transcoder.Transcoder)
			newFileName := outFileName(file.Name(), prepend)
			err = trans.Initialize(file.Name(), newFileName)
			if err != nil {
				log.Printf(err.Error(), "skipping file", file.Name())
				continue
			}

			hasAAC := hasAAC(trans.MediaFile().Metadata())
			log.Printf("-> AAC: %v, Bitrate: %s, Length: %s...", hasAAC, trans.MediaFile().Metadata().Streams[1].BitRate, trans.MediaFile().Metadata().Format.Duration)

			if hasAAC {
				if file.Name() == "1.mp4" {
					log.Printf("-> Extracting audio...")
					AudioStream = &trans.MediaFile().Metadata().Streams[1]
					continue
				}

				log.Printf("-> Audio channel present, skipping file %s", file.Name())
				if trans.MediaFile().Metadata().Streams[1].BitRate != AudioStream.BitRate {
					log.Printf("-> Audio bitrates mismatch [%s vs %s], updating audio", trans.MediaFile().Metadata().Streams[1].BitRate, AudioStream.BitRate)
					trans.MediaFile().Metadata().Streams[1] = *AudioStream
					done := trans.Run(true)
					progress := trans.Output()
					for p := range progress {
						fmt.Println(p)
					}

					fmt.Println(<-done)
				}
				continue
			}

			matchingFiles = append(matchingFiles, newFileName)
			log.Printf("Processed file %s to %s", file.Name(), newFileName)
		}
	}

	if len(matchingFiles) > 0 {
		fmt.Printf("Completed [%d] files\n", len(matchingFiles))
	}

	log.Println("exited.")
}
