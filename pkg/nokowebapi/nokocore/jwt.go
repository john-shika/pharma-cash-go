package nokocore

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math"
	"strings"
	"time"
)

var ErrJwtTokenNotFound = errors.New("jwt token not found")
var ErrJwtClaimsInvalid = errors.New("invalid JWT claims")
var ErrJwtIdentityNotFound = errors.New("jwt identity not found")
var ErrJwtIssuedAtNotFound = errors.New("jwt issued at not found")
var ErrJwtIssuerNotFound = errors.New("jwt issuer not found")
var ErrJwtSubjectNotFound = errors.New("jwt subject not found")
var ErrJwtExpiresNotFound = errors.New("jwt expires not found")
var ErrJwtSessionIdNotFound = errors.New("jwt session id not found")
var ErrJwtUserInvalid = errors.New("jwt user not found")
var ErrJwtEmailInvalid = errors.New("jwt email not found")
var ErrJwtSecretKeyNotFound = errors.New("jwt secret key not found")

type JwtConfig struct {
	Algorithm string `mapstructure:"algorithm,required" json:"algorithm,required"`
	SecretKey string `mapstructure:"secret_key,required" json:"secretKey,required"`
	Audience  string `mapstructure:"audience,required" json:"audience,required"`
	Issuer    string `mapstructure:"issuer,required" json:"issuer,required"`
	ExpiresIn string `mapstructure:"expires_in,required" json:"expiresIn,required"`
}

func NewJwtConfig() *JwtConfig {
	return new(JwtConfig)
}

func (j *JwtConfig) GetNameType() string {
	return "Jwt"
}

func (j *JwtConfig) GetSigningMethod() jwt.SigningMethod {
	switch strings.ToUpper(j.Algorithm) {
	case "ES256":
		return jwt.SigningMethodES256
	case "ES384":
		return jwt.SigningMethodES384
	case "ES512":
		return jwt.SigningMethodES512
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	case "PS256":
		return jwt.SigningMethodPS256
	case "PS384":
		return jwt.SigningMethodPS384
	case "PS512":
		return jwt.SigningMethodPS512
	case "RS256":
		return jwt.SigningMethodRS256
	case "RS384":
		return jwt.SigningMethodRS384
	case "RS512":
		return jwt.SigningMethodRS512
	default:
		return jwt.SigningMethodHS256
	}
}

type JwtClaimNamed string

const (
	JwtClaimIdentity  JwtClaimNamed = "jti"
	JwtClaimSubject   JwtClaimNamed = "sub"
	JwtClaimSessionId JwtClaimNamed = "sid"
	JwtClaimIssuedAt  JwtClaimNamed = "iat"
	JwtClaimIssuer    JwtClaimNamed = "iss"
	JwtClaimAudience  JwtClaimNamed = "aud"
	JwtClaimExpiresAt JwtClaimNamed = "exp"
	JwtClaimUser      JwtClaimNamed = "user"
	JwtClaimEmail     JwtClaimNamed = "email"
	JwtClaimPhone     JwtClaimNamed = "phone"
	JwtClaimRole      JwtClaimNamed = "role"
)

type JwtTokenImpl interface {
	SignedString(key interface{}) (string, error)
	SigningString() (string, error)
	EncodeSegment(seg []byte) string
}

type TimeOrJwtNumericDateImpl interface {
	*jwt.NumericDate | time.Time | string | int64 | int
}

func GetTimeUtcFromAny(value any) time.Time {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return Default[time.Time]()
	}
	value = val.Interface()
	switch value.(type) {
	case jwt.NumericDate:
		return value.(jwt.NumericDate).Time.UTC()
	case time.Time:
		return value.(time.Time).UTC()
	case string:
		return Unwrap(ParseTimeUtcByStringISO8601(value.(string)))
	case int64:
		return GetTimeUtcByTimestamp(value.(int64))
	case int:
		return GetTimeUtcByTimestamp(int64(value.(int)))
	default:
		return Default[time.Time]()
	}
}

func GetTimeUtcFromStrict[V TimeOrJwtNumericDateImpl](value V) time.Time {
	return GetTimeUtcFromAny(value)
}

func GetJwtNumericDateFromAny(value any) *jwt.NumericDate {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return nil
	}
	value = val.Interface()
	switch value.(type) {
	case jwt.NumericDate:
		t := value.(jwt.NumericDate)
		return &t
	case time.Time:
		return jwt.NewNumericDate(value.(time.Time))
	case string:
		t := Unwrap(ParseTimeUtcByStringISO8601(value.(string)))
		return jwt.NewNumericDate(t)
	case int64:
		t := GetTimeUtcByTimestamp(value.(int64))
		return jwt.NewNumericDate(t)
	case int:
		t := GetTimeUtcByTimestamp(int64(value.(int)))
		return jwt.NewNumericDate(t)
	default:
		return nil
	}
}

func GetJwtNumericDateFromStrict[V TimeOrJwtNumericDateImpl](value V) *jwt.NumericDate {
	return GetJwtNumericDateFromAny(value)
}

type JwtClaimsImpl interface {
	GetDataAccess() *JwtClaimsDataAccess
	Get(key string) (any, bool)
	Set(key string, value any) bool
	ToJwtToken() JwtTokenImpl
	ToJwtTokenString(secretKey string) (string, error)
	ParseNumericDate(key string) (*jwt.NumericDate, error)
	ParseString(key string) (string, error)
	ParseStringMany(key string) ([]string, error)
	GetIdentity() (string, error)
	SetIdentity(identity string) bool
	GetSubject() (string, error)
	SetSubject(subject string) bool
	GetIssued() (*jwt.NumericDate, error)
	SetIssuedAt(date any) bool
	GetIssuer() (string, error)
	SetIssuer(issuer string) bool
	GetAudience() ([]string, error)
	SetAudience(audience []string) bool
	GetExpires() (*jwt.NumericDate, error)
	SetExpiresAt(date any) bool
	GetSessionId() (string, error)
	SetSessionId(sessionId string) bool
	GetUser() (string, error)
	SetUser(user string) bool
	GetRole() (string, error)
	SetRole(role string) bool
	GetEmail() (string, error)
	SetEmail(email string) bool
}

type JwtClaimsDataAccessImpl interface {
	GetIdentity() string
	SetIdentity(identity string)
	GetSubject() string
	SetSubject(subject string)
	GetIssuer() string
	SetIssuer(issuer string)
	GetAudience() []string
	SetAudience(audience []string)
	GetIssued() *jwt.NumericDate
	SetIssuedAt(date any)
	GetExpires() *jwt.NumericDate
	SetExpiresAt(date any)
	GetSessionId() string
	SetSessionId(sessionId string)
	GetUser() string
	SetUser(user string)
	GetRole() string
	SetRole(role string)
	GetEmail() string
	SetEmail(email string)
}

type JwtClaimsDataAccess struct {
	ID        string           `json:"jti,omitempty"`
	Issuer    string           `json:"iss,omitempty"`
	Subject   string           `json:"sub,omitempty"`
	Audience  []string         `json:"aud,omitempty"`
	NotBefore *jwt.NumericDate `json:"nbf,omitempty"`
	IssuedAt  *jwt.NumericDate `json:"iat,omitempty"`
	ExpiresAt *jwt.NumericDate `json:"exp,omitempty"`
	SessionId string           `json:"sid,omitempty"`
	User      string           `json:"name,omitempty"`
	Email     string           `json:"email,omitempty"`
	Role      string           `json:"role,omitempty"`
}

func NewJwtClaimsDataAccess(claims *jwt.RegisteredClaims) JwtClaimsDataAccessImpl {
	return &JwtClaimsDataAccess{
		ID:        claims.ID,
		Subject:   claims.Subject,
		Issuer:    claims.Issuer,
		Audience:  claims.Audience,
		NotBefore: claims.NotBefore,
		IssuedAt:  claims.IssuedAt,
		ExpiresAt: claims.ExpiresAt,
	}
}

func NewEmptyJwtClaimsDataAccess() JwtClaimsDataAccessImpl {
	return new(JwtClaimsDataAccess)
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIdentity() string {
	return claimsDataAccess.ID
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIdentity(identity string) {
	claimsDataAccess.ID = identity
}

func (claimsDataAccess *JwtClaimsDataAccess) GetSubject() string {
	return claimsDataAccess.Subject
}

func (claimsDataAccess *JwtClaimsDataAccess) SetSubject(subject string) {
	claimsDataAccess.Subject = subject
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIssuer() string {
	return claimsDataAccess.Issuer
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIssuer(issuer string) {
	claimsDataAccess.Issuer = issuer
}

func (claimsDataAccess *JwtClaimsDataAccess) GetAudience() []string {
	return claimsDataAccess.Audience
}

func (claimsDataAccess *JwtClaimsDataAccess) SetAudience(audience []string) {
	claimsDataAccess.Audience = audience
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIssued() *jwt.NumericDate {
	return claimsDataAccess.IssuedAt
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIssuedAt(date any) {
	claimsDataAccess.IssuedAt = GetJwtNumericDateFromAny(date)
}

func (claimsDataAccess *JwtClaimsDataAccess) GetExpires() *jwt.NumericDate {
	return claimsDataAccess.ExpiresAt
}

func (claimsDataAccess *JwtClaimsDataAccess) SetExpiresAt(date any) {
	claimsDataAccess.ExpiresAt = GetJwtNumericDateFromAny(date)
}

func (claimsDataAccess *JwtClaimsDataAccess) GetSessionId() string {
	return claimsDataAccess.SessionId
}

func (claimsDataAccess *JwtClaimsDataAccess) SetSessionId(sessionId string) {
	claimsDataAccess.SessionId = sessionId
}

func (claimsDataAccess *JwtClaimsDataAccess) GetUser() string {
	return claimsDataAccess.User
}

func (claimsDataAccess *JwtClaimsDataAccess) SetUser(user string) {
	claimsDataAccess.User = user
}

func (claimsDataAccess *JwtClaimsDataAccess) GetRole() string {
	return claimsDataAccess.Role
}

func (claimsDataAccess *JwtClaimsDataAccess) SetRole(role string) {
	claimsDataAccess.Role = role
}

func (claimsDataAccess *JwtClaimsDataAccess) GetEmail() string {
	return claimsDataAccess.Email
}

func (claimsDataAccess *JwtClaimsDataAccess) SetEmail(email string) {
	claimsDataAccess.Email = email
}

func encodeSecretKey(secretKey string, jwtSigningMethod jwt.SigningMethod) []byte {
	data := []byte(secretKey)
	switch jwtSigningMethod.Alg() {
	case jwt.SigningMethodHS256.Alg():
		return HashSha256(data)
	case jwt.SigningMethodHS384.Alg():
		return HashSha384(data)
	case jwt.SigningMethodHS512.Alg():
		return HashSha512(data)
	default:
		return data
	}
}

type JwtClaims struct {
	method jwt.SigningMethod
	claims jwt.MapClaims
}

func NewJwtClaims(claims jwt.MapClaims, jwtSigningMethod jwt.SigningMethod) JwtClaimsImpl {
	return &JwtClaims{
		method: jwtSigningMethod,
		claims: claims,
	}
}

func JwtClaimsEmpty(jwtSigningMethod jwt.SigningMethod) JwtClaimsImpl {
	return &JwtClaims{
		method: jwtSigningMethod,
		claims: jwt.MapClaims{},
	}
}

func GetJwtClaimsFromJwtToken(token JwtTokenImpl, jwtSigningMethod jwt.SigningMethod) (JwtClaimsImpl, error) {
	var ok bool
	var claims jwt.MapClaims
	KeepVoid(ok, claims)

	jwtToken := token.(*jwt.Token)
	if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, ErrJwtClaimsInvalid
	}

	return NewJwtClaims(claims, jwtSigningMethod), nil
}

func (j *JwtClaims) GetSigningMethod() jwt.SigningMethod {
	return j.method
}

func (j *JwtClaims) SetSigningMethod(method jwt.SigningMethod) {
	j.method = method
}

func (j *JwtClaims) GetDataAccess() *JwtClaimsDataAccess {
	return &JwtClaimsDataAccess{
		ID:        Unwrap(j.GetIdentity()),
		Subject:   Unwrap(j.GetSubject()),
		Issuer:    Unwrap(j.GetIssuer()),
		Audience:  Unwrap(j.GetAudience()),
		IssuedAt:  Unwrap(j.GetIssued()),
		ExpiresAt: Unwrap(j.GetExpires()),
		SessionId: Unwrap(j.GetIdentity()),
		User:      Unwrap(j.GetUser()),
		Role:      Unwrap(j.GetRole()),
		Email:     Unwrap(j.GetEmail()),
	}
}

func (j *JwtClaims) Get(key string) (any, bool) {
	var ok bool
	var value any
	KeepVoid(ok, value)

	value, ok = j.claims[key]
	return value, ok
}

func (j *JwtClaims) Set(key string, value any) bool {
	var ok bool
	var temp any
	KeepVoid(ok, temp)

	if temp, ok = j.claims[key]; !ok {
		return false
	}

	j.claims[key] = value
	return true
}

func (j *JwtClaims) ToJwtToken() JwtTokenImpl {
	return jwt.NewWithClaims(j.method, j.claims)
}

func (j *JwtClaims) ToJwtTokenString(secretKey string) (string, error) {
	var err error
	var tokenString string
	KeepVoid(err, tokenString)

	token := j.ToJwtToken()
	dataSecretKey := encodeSecretKey(secretKey, j.method)
	if tokenString, err = token.SignedString(dataSecretKey); err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewNumericDateFromSeconds(f float64) *jwt.NumericDate {
	round, frac := math.Modf(f)
	return jwt.NewNumericDate(time.Unix(int64(round), int64(frac*1e9)))
}

func (j *JwtClaims) ParseNumericDate(key string) (*jwt.NumericDate, error) {
	var ok bool
	var err error
	var value any
	KeepVoid(ok, err, value)

	if value, ok = j.Get(key); !ok {
		return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
	}

	switch exp := value.(type) {
	case float64:
		if exp == 0 {
			return nil, nil
		}

		return NewNumericDateFromSeconds(exp), nil
	case json.Number:
		if value, err = exp.Float64(); err != nil {
			return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
		}

		return NewNumericDateFromSeconds(value.(float64)), nil
	}

	return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
}

func (j *JwtClaims) ParseString(key string) (string, error) {
	var ok bool
	var value any
	var temp string
	KeepVoid(ok, value, temp)

	if value, ok = j.Get(key); !ok {
		return EmptyString, ErrDataTypeInvalid
	}

	if temp, ok = value.(string); !ok {
		return EmptyString, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
	}

	return temp, nil
}

func (j *JwtClaims) ParseStringMany(key string) ([]string, error) {
	var ok bool
	var value any
	var temp []string
	KeepVoid(ok, value, temp)

	if value, ok = j.Get(key); !ok {
		return nil, ErrDataTypeInvalid
	}

	switch value.(type) {
	case string:
		temp = []string{value.(string)}
	case []string:
		temp = value.([]string)
	case []any:
		values := value.([]any)
		temp = make([]string, len(values))
		for i, v := range values {
			if temp[i], ok = v.(string); !ok {
				return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
			}
		}
	default:
		return nil, NewThrow(fmt.Sprintf("%s is invalid", key), ErrDataTypeInvalid)
	}

	return temp, nil
}

func (j *JwtClaims) GetIdentity() (string, error) {
	return j.ParseString(string(JwtClaimIdentity))
}

func (j *JwtClaims) SetIdentity(identity string) bool {
	return j.Set(string(JwtClaimIdentity), identity)
}

func (j *JwtClaims) GetSubject() (string, error) {
	return j.ParseString(string(JwtClaimSubject))
}

func (j *JwtClaims) SetSubject(subject string) bool {
	return j.Set(string(JwtClaimSubject), subject)
}

func (j *JwtClaims) GetIssued() (*jwt.NumericDate, error) {
	var err error
	var temp *jwt.NumericDate
	KeepVoid(err, temp)

	temp, err = j.ParseNumericDate(string(JwtClaimIssuedAt))
	return temp, err
}

func (j *JwtClaims) SetIssuedAt(date any) bool {
	return j.Set(string(JwtClaimIssuedAt), GetJwtNumericDateFromAny(date))
}

func (j *JwtClaims) GetIssuer() (string, error) {
	return j.ParseString(string(JwtClaimIssuer))
}

func (j *JwtClaims) SetIssuer(issuer string) bool {
	return j.Set(string(JwtClaimIssuer), issuer)
}

func (j *JwtClaims) GetAudience() ([]string, error) {
	return j.ParseStringMany(string(JwtClaimAudience))
}

func (j *JwtClaims) SetAudience(audience []string) bool {
	return j.Set(string(JwtClaimAudience), audience)
}

func (j *JwtClaims) GetExpires() (*jwt.NumericDate, error) {
	var err error
	var temp *jwt.NumericDate
	KeepVoid(err, temp)

	temp, err = j.ParseNumericDate(string(JwtClaimExpiresAt))
	return temp, err
}

func (j *JwtClaims) SetExpiresAt(date any) bool {
	return j.Set(string(JwtClaimExpiresAt), GetJwtNumericDateFromAny(date))
}

func (j *JwtClaims) GetSessionId() (string, error) {
	return j.ParseString(string(JwtClaimSessionId))
}

func (j *JwtClaims) SetSessionId(sessionId string) bool {
	return j.Set(string(JwtClaimSessionId), sessionId)
}

func (j *JwtClaims) GetUser() (string, error) {
	return j.ParseString(string(JwtClaimUser))
}

func (j *JwtClaims) SetUser(user string) bool {
	return j.Set(string(JwtClaimUser), user)
}

func (j *JwtClaims) GetRole() (string, error) {
	return j.ParseString(string(JwtClaimRole))
}

func (j *JwtClaims) SetRole(role string) bool {
	return j.Set(string(JwtClaimRole), role)
}

func (j *JwtClaims) GetEmail() (string, error) {
	return j.ParseString(string(JwtClaimEmail))
}

func (j *JwtClaims) SetEmail(email string) bool {
	return j.Set(string(JwtClaimEmail), email)
}

func (j *JwtClaims) GetPhone() (string, error) {
	return j.ParseString(string(JwtClaimPhone))
}

func (j *JwtClaims) SetPhone(phone string) bool {
	return j.Set(string(JwtClaimPhone), phone)
}

func ParseJwtTokenUnverified(token string) (JwtTokenImpl, error) {
	var err error
	var parts []string
	var jwtToken JwtTokenImpl
	KeepVoid(err, parts, jwtToken)

	parser := jwt.NewParser()
	claims := make(jwt.MapClaims)

	if jwtToken, parts, err = parser.ParseUnverified(token, claims); err != nil {
		return nil, ErrJwtTokenNotFound
	}

	return jwtToken, nil
}

func ParseJwtToken(token string, secretKey string, jwtSigningMethod jwt.SigningMethod) (JwtTokenImpl, error) {
	var err error
	var jwtToken JwtTokenImpl
	KeepVoid(err, jwtToken)

	parser := jwt.NewParser()
	claims := jwt.MapClaims{}
	keyFunc := func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwtSigningMethod.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		dataSecretKey := encodeSecretKey(secretKey, token.Method)
		return dataSecretKey, nil
	}

	if jwtToken, err = parser.ParseWithClaims(token, claims, keyFunc); err != nil {
		return nil, ErrJwtTokenNotFound
	}

	return jwtToken, nil
}

func CvtJwtClaimsAccessDataToJwtClaims(claimsDataAccess *JwtClaimsDataAccess, jwtSigningMethod jwt.SigningMethod) JwtClaimsImpl {
	claims := NewJwtClaims(jwt.MapClaims{}, jwtSigningMethod)
	claims.SetIdentity(claimsDataAccess.GetIdentity())
	claims.SetSubject(claimsDataAccess.GetSubject())
	claims.SetIssuer(claimsDataAccess.GetIssuer())
	claims.SetAudience(claimsDataAccess.GetAudience())
	claims.SetIssuedAt(claimsDataAccess.GetIssued())
	claims.SetExpiresAt(claimsDataAccess.GetExpires())
	claims.SetSessionId(claimsDataAccess.GetSessionId())
	claims.SetUser(claimsDataAccess.GetUser())
	claims.SetRole(claimsDataAccess.GetRole())
	claims.SetEmail(claimsDataAccess.GetEmail())
	return claims
}

func CvtJwtClaimsToJwtClaimsAccessData(claims JwtClaimsImpl) *JwtClaimsDataAccess {
	claimsDataAccess := new(JwtClaimsDataAccess)
	claimsDataAccess.SetIdentity(Unwrap(claims.GetIdentity()))
	claimsDataAccess.SetSubject(Unwrap(claims.GetSubject()))
	claimsDataAccess.SetIssuer(Unwrap(claims.GetIssuer()))
	claimsDataAccess.SetAudience(Unwrap(claims.GetAudience()))
	claimsDataAccess.SetIssuedAt(Unwrap(claims.GetIssued()))
	claimsDataAccess.SetExpiresAt(Unwrap(claims.GetExpires()))
	claimsDataAccess.SetSessionId(Unwrap(claims.GetSessionId()))
	claimsDataAccess.SetUser(Unwrap(claims.GetUser()))
	claimsDataAccess.SetRole(Unwrap(claims.GetRole()))
	claimsDataAccess.SetEmail(Unwrap(claims.GetEmail()))
	return claimsDataAccess
}
