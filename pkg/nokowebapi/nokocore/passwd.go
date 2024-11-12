package nokocore

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

var ErrPasswordInvalidLength = errors.New("password invalid length")

type PasswordConfig struct {
	MinLength int
	MaxLength int
}

func NewPasswordConfig() *PasswordConfig {
	return &PasswordConfig{
		MinLength: 8,
		MaxLength: 32,
	}
}

var passConfig = NewPasswordConfig()

type Pdkf2Config struct {
	N         int
	SaltSize  int
	KeyLength int
	SepChar   string
	Prefix    string
}

func NewPdkf2Config() *Pdkf2Config {
	return &Pdkf2Config{
		N:         10000,
		SaltSize:  16,
		KeyLength: 32,
		SepChar:   "$",
		Prefix:    "PDKF2_",
	}
}

var pdkf2Config = NewPdkf2Config()

func HashPassword(password string) (string, error) {
	var n int
	var err error
	KeepVoid(n, err)

	if len(password) < passConfig.MinLength || len(password) > passConfig.MaxLength {
		return EmptyString, ErrPasswordInvalidLength
	}

	salt := make([]byte, pdkf2Config.SaltSize)
	if n, err = rand.Read(salt); err != nil {
		return EmptyString, err
	}

	buff := []byte(password)
	key := pbkdf2.Key(buff, salt, pdkf2Config.N, pdkf2Config.KeyLength, sha256.New)
	temp := pdkf2Config.Prefix + Base64EncodeURLSafe(salt) + pdkf2Config.SepChar + Base64EncodeURLSafe(key)
	return temp, nil
}

func CompareHashPassword(hash string, password string) bool {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if len(hash) < passConfig.MinLength || len(hash) > passConfig.MaxLength {
		return false
	}

	if hash, ok = strings.CutPrefix(hash, pdkf2Config.Prefix); !ok {
		return false
	}

	if !strings.Contains(hash, pdkf2Config.SepChar) {
		return false
	}

	tokens := strings.Split(hash, pdkf2Config.SepChar)
	if len(tokens) != 2 {
		return false
	}

	salt := Unwrap(Base64DecodeURLSafe(tokens[0]))
	pass := Unwrap(Base64DecodeURLSafe(tokens[1]))

	buff := []byte(password)
	key := pbkdf2.Key(buff, salt, 10000, 32, sha256.New)
	return BytesEquals(pass, key)
}
