package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"hash/crc64"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var PrivateKey *rsa.PrivateKey

const expectedChecksum uint64 = 0x5E6D6D40FB2D7468

func findFile(rootDir, targetFile string, maxDepth int, stopDirs []string) (string, error) {
	var foundFile string
	var currentDir = rootDir

	var searchDir func(string, int) bool

	searchDir = func(dir string, depth int) bool {
		if depth > maxDepth {
			return false
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Ошибка при чтении директории:", err)
			return false
		}

		for _, file := range files {
			filePath := filepath.Join(dir, file.Name())

			if file.Name() == targetFile {
				foundFile = filePath
				return true
			}

			if file.Name() == "main.go" && depth == 0 {
				continue
			}

			for _, stopDir := range stopDirs {
				if strings.Contains(filePath, stopDir) {
					return false
				}
			}

			if file.IsDir() {
				if searchDir(filePath, depth+1) {
					return true
				}
			}
		}
		return false
	}

	for {
		if searchDir(currentDir, 0) {
			return foundFile, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir || parentDir == "/" {
			break
		}
		currentDir = parentDir
	}

	return "", fmt.Errorf("файл %s не найден", targetFile)
}

func init() {
	startDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}

	keyPath, err := findFile(startDir, "prk.der", 3, nil)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		// fmt.Println("Файл найден по пути:", keyPath)
	}

	keyData, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		log.Fatalf("Ошибка чтения приватного ключа: %v", err)
	}

	table := crc64.MakeTable(crc64.ISO)
	checksum := crc64.Checksum(keyData, table)
	if checksum != expectedChecksum {
		log.Fatalf("Контрольная сумма файла не совпадает: %v", checksum)
	}

	PrivateKey, err = x509.ParsePKCS1PrivateKey(keyData)
	if err != nil {
		log.Fatalf("Ошибка при парсинге приватного ключа: %v", err)
	}
}

func DecryptData(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	blockSize := privateKey.Size()

	if len(data) <= blockSize {
		decrypted, err := rsa.DecryptPKCS1v15(nil, privateKey, data)
		if err != nil {
			return nil, fmt.Errorf("не удалось расшифровать данные: %v", err)
		}
		return decrypted, nil
	}

	var decryptedData []byte
	for i := 0; i < len(data); i += blockSize {
		end := i + blockSize
		if end > len(data) {
			end = len(data)
		}

		block := data[i:end]
		decryptedBlock, err := rsa.DecryptPKCS1v15(nil, privateKey, block)
		if err != nil {
			return nil, fmt.Errorf("не удалось расшифровать блок данных: %v", err)
		}
		decryptedData = append(decryptedData, decryptedBlock...)
	}

	return decryptedData, nil
}
