package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	mode := flag.String("mode", "", "Decode or encode")
	input := flag.String("i", "", "Input path")
	output := flag.String("o", "", "Output path")
	flag.Parse()
	if *mode == "" || *input == "" || *output == "" {
		fmt.Println("Choose mode and input/output path.")
		return
	}

	input_data, err := ioutil.ReadFile(*input)
	if err != nil {
		fmt.Println("Error! File is corrupted.")
		return
	}

	var output_data []byte
	switch *mode {
	case "encode":
		output_data = []byte(base64.StdEncoding.EncodeToString(input_data))
	case "decode":
		decoded_data, err := base64.StdEncoding.DecodeString(string(input_data))
		if err != nil {
			fmt.Println("Error! File is not base64 encoded.")
			return
		}
		output_data = decoded_data
	default:
		fmt.Println("Not supported mode.")
		return
	}

	err = ioutil.WriteFile(*output, output_data, 0644)
	if err != nil {
		fmt.Println("Error!", err)
		return
	}
	fmt.Println("Success")
}
