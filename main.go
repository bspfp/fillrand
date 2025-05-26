package main

import (
	"crypto/rand"
	"log"
	"os"
)

const (
	targetDrive = `E:`
)

func main() {
	// 랜덤한 바이트 1G 생성
	const size = 1 << 30 // 1GB
	data := make([]byte, size)

	_, err := rand.Read(data)
	if err != nil {
		log.Fatalf("failed to generate random data: %v", err)
	}

	targetfile := targetDrive + `\random_data.bin`
	f, err := os.OpenFile(targetfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	for len(data) > 0 {
		_, err := f.Write(data)
		if err != nil {
			data = data[:len(data)/2]
		}
	}
}
