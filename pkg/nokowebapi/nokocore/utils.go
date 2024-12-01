package nokocore

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
)

func HexEncodeToString(data []byte) string {
	return hex.EncodeToString(data)
}

func HexDecodeToBytes(data string) ([]byte, error) {
	return hex.DecodeString(data)
}

func HashSha256(data []byte) []byte {
	hash := sha256.New()
	buff := hash.Sum(data)
	return buff
}

func HashSha256Compare(data []byte, hash []byte) bool {
	temp := HashSha256(data)
	return BytesEquals(temp, hash)
}

func HashSha384(data []byte) []byte {
	hash := sha512.New384()
	buff := hash.Sum(data)
	return buff
}

func HashSha384Compare(data []byte, hash []byte) bool {
	temp := HashSha384(data)
	return BytesEquals(temp, hash)
}

func HashSha512(data []byte) []byte {
	hash := sha512.New()
	buff := hash.Sum(data)
	return buff
}

func HashSha512Compare(data []byte, hash []byte) bool {
	temp := HashSha512(data)
	return BytesEquals(temp, hash)
}

type BytesOrStringImpl interface {
	string | []byte
}

func BytesEquals[V1, V2 BytesOrStringImpl](data V1, buff V2) bool {
	size := len(data)
	if size != len(buff) {
		return false
	}
	for i := 0; i < size; i++ {
		if data[i] != buff[i] {
			return false
		}
	}
	return true
}

func StringEquals(data string, buff string) bool {
	return BytesEquals(data, buff)
}

func CreateEmptyFile(path string) error {
	var err error
	var fileInfo os.FileInfo
	var file *os.File
	KeepVoid(err, fileInfo, file)

	dirPath := filepath.Dir(path)
	filePath := filepath.Base(path)

	// Check if the directory exists, and create it if it doesn't
	if fileInfo, err = os.Stat(dirPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return fmt.Errorf("[OS] Failed to create directory: %s, %w", dirPath, err)
		}

		fmt.Printf("[OS] Directory %s created.\n", dirPath)

	} else {
		fmt.Printf("[OS] Directory %s already exists.\n", dirPath)

	}

	// Check if the file exists, and create it if it doesn't
	if fileInfo, err = os.Stat(path); os.IsNotExist(err) {
		if file, err = os.Create(path); err != nil {
			return fmt.Errorf("[OS] Failed to create file: %s, %w", filePath, err)
		}

		fmt.Printf("[OS] Excel %s has been created.\n", filePath)
		NoErr(file.Close())

	} else {
		fmt.Printf("[OS] Excel %s already exists.\n", filePath)

	}

	return nil
}

func GetNameType(obj any) string {
	var ok bool
	var nameable NameableImpl
	KeepVoid(ok, nameable)

	if IsNone(obj) {
		return "<nil>"
	}

	// try cast nameable and call method
	if nameable, ok = Cast[NameableImpl](obj); ok {
		return nameable.GetNameType()
	}

	//return fmt.Sprintf("%T", obj)
	return GetNameTypeReflect(obj)
}

func ParseEnvToBool(value string) bool {
	if value = strings.TrimSpace(value); value != "" {
		switch strings.ToLower(value) {
		case "1", "y", "yes", "true":
			return true

		default:
			return false
		}
	}

	return false
}

func ParseEnvToInt(value string) int64 {
	var err error
	var val int64
	KeepVoid(err, val)

	if val, err = strconv.ParseInt(value, 10, 64); err != nil {
		return 0
	}

	return val
}

func ParseEnvToUint(value string) uint64 {
	var err error
	var val uint64
	KeepVoid(err, val)

	if val, err = strconv.ParseUint(value, 10, 64); err != nil {
		return 0
	}

	return val
}

func ParseEnvToFloat(value string) float64 {
	var err error
	var val float64
	KeepVoid(err, val)

	if val, err = strconv.ParseFloat(value, 64); err != nil {
		return 0
	}

	return val
}

func ParseEnvToComplex(value string) complex128 {
	var err error
	var val complex128
	KeepVoid(err, val)

	if val, err = strconv.ParseComplex(value, 128); err != nil {
		return 0
	}

	return val
}

func ParseEnvToString(value string) string {
	return strings.TrimSpace(value)
}

func ParseEnvToDuration(value string) time.Duration {
	var err error
	var val time.Duration
	KeepVoid(err, val)

	if val, err = time.ParseDuration(value); err != nil {
		return 0
	}

	return val
}

func StringToUtf16(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

func Utf16ToString(b []uint16) string {
	return string(utf16.Decode(b))
}

func BytesToUtf16(b []byte) []uint16 {
	return StringToUtf16(string(b))
}

func Utf16ToBytes(b []uint16) []byte {
	return []byte(Utf16ToString(b))
}

func ToFileSizeFormat(size int64) string {
	const (
		KB = 1 << 10
		MB = 1 << 20
		GB = 1 << 30
		TB = 1 << 40
	)

	switch {
	case size >= TB:
		return fmt.Sprintf("%.2f TB", float64(size)/TB)

	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)

	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)

	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)

	default:
		return fmt.Sprintf("%d B", size)
	}
}

func Int64ToHexCodes(value int64) string {
	buff := make([]byte, 8)
	for i := 0; i < 8; i++ {
		j := 8 - i - 1
		buff[j] = byte(value & 0xff)
		value >>= 8
	}

	return fmt.Sprintf("%016x", buff)
}

func Int64ToBase64RawURL(value int64) string {
	buff := make([]byte, 8)
	for i := 0; i < 8; i++ {
		j := 8 - i - 1
		buff[j] = byte(value & 0xff)
		value >>= 8
	}

	return base64.RawURLEncoding.EncodeToString(buff)
}
