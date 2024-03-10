package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	// Создание временных файлов
	inputData := []byte("Hello, World!")
	inputFile, err := ioutil.TempFile("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(inputFile.Name())
	defer inputFile.Close()
	if _, err := inputFile.Write(inputData); err != nil {
		t.Fatal(err)
	}

	outputFile, err := ioutil.TempFile("", "test_output_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	// Запуск кодирования
	encodeCmd := exec.Command("./encoder64", "encode", "-i", inputFile.Name(), "-o", outputFile.Name())
	if err := encodeCmd.Run(); err != nil {
		t.Fatal(err)
	}

	// Запуск декодирования
	decodeCmd := exec.Command("./encoder64", "decode", "-i", outputFile.Name(), "-o", outputFile.Name()+".decoded.txt")
	if err := decodeCmd.Run(); err != nil {
		t.Fatal(err)
	}

	// Проверка корректности декодированного файла
	decodedData, err := ioutil.ReadFile(outputFile.Name() + ".decoded.txt")
	if err != nil {
		t.Fatal(err)
	}

	decodedStr := string(decodedData)
	if decodedStr != string(inputData) {
		t.Errorf("Expected decoded data: %s, got: %s", string(inputData), decodedStr)
	}
}

func TestMain(m *testing.M) {
	// Компиляция кода перед запуском тестов
	buildCmd := exec.Command("go", "build", "encoder64.go")
	if err := buildCmd.Run(); err != nil {
		panic(err)
	}
	defer os.Remove("./encoder64")
	defer os.Remove("./test_output_*.txt.decoded.txt")

	// Запуск тестов
	os.Exit(m.Run())
}
