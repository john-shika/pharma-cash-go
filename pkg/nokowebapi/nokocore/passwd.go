package nokocore

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

var ErrPassInvalidLength = errors.New("password invalid length")

type PassConfig struct {
	MinLength int `mapstructure:"min_length" json:"minLength" yaml:"min_length"`
	MaxLength int `mapstructure:"max_length" json:"maxLength" yaml:"max_length"`
}

func NewPassConfig() *PassConfig {
	return &PassConfig{
		MinLength: 8,
		MaxLength: 512,
	}
}

type Pdkf2Config struct {
	N         int    `mapstructure:"n" json:"n" yaml:"n"`
	SaltSize  int    `mapstructure:"salt_size" json:"saltSize" yaml:"salt_size"`
	KeyLength int    `mapstructure:"key_length" json:"keyLength" yaml:"key_length"`
	Separator string `mapstructure:"separator" json:"separator" yaml:"separator"`
	Prefix    string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}

func NewPdkf2Config() *Pdkf2Config {
	return &Pdkf2Config{
		N:         10000,
		SaltSize:  16,
		KeyLength: 32,
		Separator: "$",
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
	var buffer bytes.Buffer
	KeepVoid(n, err, buffer)

	if len(p.Value) < p.passConfig.MinLength || len(p.Value) > p.passConfig.MaxLength {
		return "", ErrPassInvalidLength
	}

	salt := make([]byte, p.pdkf2Config.SaltSize)
	if n, err = rand.Read(salt); err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(p.Value), salt, p.pdkf2Config.N, p.pdkf2Config.KeyLength, sha256.New)

	buffer.WriteString(p.pdkf2Config.Prefix)
	buffer.WriteString(Base64EncodeURLSafe(salt))
	buffer.WriteString(p.pdkf2Config.Separator)
	buffer.WriteString(Base64EncodeURLSafe(key))

	temp := buffer.String()
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

	if !strings.Contains(hash, p.pdkf2Config.Separator) {
		return false
	}

	tokens := strings.Split(hash, p.pdkf2Config.Separator)
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
