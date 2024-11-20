package nokocore

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

var ErrPassInvalidLength = errors.New("password invalid length")

type PassConfig struct {
	MinLength int
	MaxLength int
}

func NewPassConfig() *PassConfig {
	return &PassConfig{
		MinLength: 8,
		MaxLength: 512,
	}
}

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

type PassImpl interface {
	Hash() (string, error)
	Equals(hash string) bool
}

type Password struct {
	Value string

	passConfig  *PassConfig
	pdkf2Config *Pdkf2Config
}

func NewPassword(password string) PassImpl {
	return &Password{
		Value:       password,
		passConfig:  NewPassConfig(),
		pdkf2Config: NewPdkf2Config(),
	}
}

func (p *Password) Hash() (string, error) {
	var n int
	var err error
	KeepVoid(n, err)

	if len(p.Value) < p.passConfig.MinLength || len(p.Value) > p.passConfig.MaxLength {
		return "", ErrPassInvalidLength
	}

	salt := make([]byte, p.pdkf2Config.SaltSize)
	if n, err = rand.Read(salt); err != nil {
		return "", err
	}

	buff := []byte(p.Value)
	key := pbkdf2.Key(buff, salt, p.pdkf2Config.N, p.pdkf2Config.KeyLength, sha256.New)

	temp := p.pdkf2Config.Prefix + Base64EncodeURLSafe(salt) + p.pdkf2Config.SepChar + Base64EncodeURLSafe(key)
	return temp, nil
}

func (p *Password) Equals(hash string) bool {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if len(p.Value) < p.passConfig.MinLength || len(p.Value) > p.passConfig.MaxLength {
		return false
	}

	if hash, ok = strings.CutPrefix(hash, p.pdkf2Config.Prefix); !ok {
		return false
	}

	if !strings.Contains(hash, p.pdkf2Config.SepChar) {
		return false
	}

	tokens := strings.Split(hash, p.pdkf2Config.SepChar)
	if len(tokens) != 2 {
		return false
	}

	salt := Unwrap(Base64DecodeURLSafe(tokens[0]))
	pass := Unwrap(Base64DecodeURLSafe(tokens[1]))

	if len(salt) != p.pdkf2Config.SaltSize {
		return false
	}

	buff := []byte(p.Value)
	key := pbkdf2.Key(buff, salt, p.pdkf2Config.N, p.pdkf2Config.KeyLength, sha256.New)

	return BytesEquals(pass, key)
}
