package apikeys

import (
	"crypto/rand"
	"errors"
	"fmt"
	"hash/crc32"
	"math/big"
	"strings"

	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/rs/zerolog"
)

type ApiKeyGenerator struct {
	prefix         string
	entropyLength  int
	checksumLength int
	logger         zerolog.Logger
}

const base62Characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var generator *ApiKeyGenerator

func Init(config config.ApiKeyConfig) {
	generator = &ApiKeyGenerator{
		prefix:         config.Prefix,
		entropyLength:  config.EntropyLength,
		checksumLength: config.ChecksumLength,
		logger:         logger.NewLogger("api-key-gen"),
	}
}

func New() (string, error) {
	entropy, err := generateRandomBase62(generator.entropyLength)
	if err != nil {
		return "", err
	}

	checksum := getCRC32Checksum(entropy)
	encodedChecksum := base62Encode(checksum, generator.checksumLength)

	return fmt.Sprintf("%s_%s%s", generator.prefix, entropy, encodedChecksum), nil
}

func ValidateApiKey(apiKey string) (bool, error) {
	if !strings.HasPrefix(apiKey, generator.prefix) {
		return false, errors.New("Invalid API Key prefix")
	}

	expectedLegnth := len(generator.prefix) + 1 + generator.entropyLength + generator.checksumLength
	if len(apiKey) != expectedLegnth {
		return false, errors.New("Invalid API Key length")
	}

	entropy := apiKey[len(generator.prefix) : len(generator.prefix)+generator.entropyLength]
	checksumEncoded := apiKey[len(apiKey)-generator.checksumLength:]

	expectedChecksum := getCRC32Checksum(entropy)
	actualChecksum := base62Decode(checksumEncoded)

	if actualChecksum != expectedChecksum {
		return false, errors.New("Invalid checksum")
	}

	return true, nil
}

func generateRandomBase62(length int) (string, error) {
	var result strings.Builder
	max := big.NewInt(int64(len(base62Characters)))

	for range length {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		result.WriteByte(base62Characters[num.Int64()])
	}

	return result.String(), nil
}

func getCRC32Checksum(str string) uint32 {
	crc32q := crc32.MakeTable(crc32.Castagnoli)
	return crc32.Checksum([]byte(str), crc32q)
}

func base62Encode(checksum uint32, length int) string {
	var result []byte
	num := checksum

	for num > 0 {
		remainder := num % 62
		result = append([]byte{base62Characters[remainder]}, result...)
		num /= 62
	}

	for len(result) < length {
		result = append([]byte{'0'}, result...)
	}

	return string(result)
}

func base62Decode(input string) uint32 {
	var result uint32
	for _, char := range input {
		index := strings.IndexRune(base62Characters, char)
		if index == -1 {
			return 0
		}
		result = result*62 + uint32(index)
	}
	return result
}
