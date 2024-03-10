package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	inputFlag  = flag.String("i", "", "Input file path")
	outputFlag = flag.String("o", "", "Output file path")
)

func encode(inputPath, outputPath string) error {
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	err = ioutil.WriteFile(outputPath, []byte(encoded), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Encoding complete.")
	return nil
}

func decode(inputPath, outputPath string) error {
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputPath, decoded, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Decoding complete.")
	return nil
}

func main() {
	encodeCommand := flag.NewFlagSet("encode", flag.ExitOnError)
	decodeCommand := flag.NewFlagSet("decode", flag.ExitOnError)

	encodeCommand.StringVar(inputFlag, "i", "", "Input file path")
	encodeCommand.StringVar(outputFlag, "o", "", "Output file path")

	decodeCommand.StringVar(inputFlag, "i", "", "Input file path")
	decodeCommand.StringVar(outputFlag, "o", "", "Output file path")

	if len(os.Args) < 3 {
		fmt.Println("Используйте: encoder64 [encode|decode] -i input_file -o output_file")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encode":
		encodeCommand.Parse(os.Args[2:])
	case "decode":
		decodeCommand.Parse(os.Args[2:])
	default:
		fmt.Println("Неизвестная команда:", os.Args[1])
		os.Exit(1)
	}

	if *inputFlag == "" {
		fmt.Println("Необходимо указать файл.")
		os.Exit(1)
	}

	if *outputFlag == "" {
		base := filepath.Base(*inputFlag)
		ext := filepath.Ext(base)
		*outputFlag = filepath.Join(filepath.Dir(*inputFlag), base[:len(base)-len(ext)]+".out")
	}

	switch {
	case encodeCommand.Parsed():
		err := encode(*inputFlag, *outputFlag)
		if err != nil {
			fmt.Println("Error encoding:", err)
			os.Exit(1)
		}
	case decodeCommand.Parsed():
		err := decode(*inputFlag, *outputFlag)
		if err != nil {
			fmt.Println("Error decoding:", err)
			os.Exit(1)
		}
	}
}
