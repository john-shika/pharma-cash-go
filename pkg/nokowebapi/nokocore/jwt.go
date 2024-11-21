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

var ErrJwtTokenInvalid = errors.New("invalid jwt token")
var ErrJwtClaimsInvalid = errors.New("invalid jwt claims")
var ErrJwtIdentityNotFound = errors.New("jwt identity not found")
var ErrJwtIssuedAtNotFound = errors.New("jwt issued at not found")
var ErrJwtIssuerNotFound = errors.New("jwt issuer not found")
var ErrJwtSubjectNotFound = errors.New("jwt subject not found")
var ErrJwtExpiresNotFound = errors.New("jwt expires not found")
var ErrJwtSessionIdNotFound = errors.New("jwt session id not found")
var ErrJwtUserNotFound = errors.New("jwt user not found")
var ErrJwtEmailNotFound = errors.New("jwt email not found")
var ErrJwtSecretKeyNotFound = errors.New("jwt secret key not found")

type JwtConfig struct {
	Algorithm string   `mapstructure:"algorithm,required" json:"algorithm,required"`
	SecretKey string   `mapstructure:"secret_key,required" json:"secretKey,required"`
	Audience  []string `mapstructure:"audience,required" json:"audience,required"`
	Issuer    string   `mapstructure:"issuer,required" json:"issuer,required"`
	ExpiresIn string   `mapstructure:"expires_in,required" json:"expiresIn,required"`
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

func (j *JwtConfig) GetExpiresIn() time.Duration {
	return Unwrap(time.ParseDuration(j.ExpiresIn))
}

type JwtClaimNamed string

const (
	JwtClaimIdentity  JwtClaimNamed = "jti"
	JwtClaimSubject   JwtClaimNamed = "sub"
	JwtClaimSessionId JwtClaimNamed = "sid"
	JwtClaimIssuedAt  JwtClaimNamed = "iat"
	JwtClaimIssuer    JwtClaimNamed = "iss"
	JwtClaimAudience  JwtClaimNamed = "aud"
	JwtClaimExpires   JwtClaimNamed = "exp"
	JwtClaimUser      JwtClaimNamed = "user"
	JwtClaimEmail     JwtClaimNamed = "email"
	JwtClaimPhone     JwtClaimNamed = "phone"
	JwtClaimRole      JwtClaimNamed = "role"
	JwtClaimAdmin     JwtClaimNamed = "admin"
	JwtClaimLevel     JwtClaimNamed = "level"
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
	GetSigningMethod() jwt.SigningMethod
	SetSigningMethod(method jwt.SigningMethod)
	GetDataAccess() *JwtClaimsDataAccess
	Get(key string) (any, bool)
	Set(key string, value any) bool
	ToJwtToken() JwtTokenImpl
	ToJwtTokenString(secretKey string) string
	ParseNumericDate(key JwtClaimNamed) *jwt.NumericDate
	ParseString(key JwtClaimNamed) string
	ParseManyString(key JwtClaimNamed) []string
	GetIdentity() string
	SetIdentity(identity string)
	GetSubject() string
	SetSubject(subject string)
	GetIssued() *jwt.NumericDate
	SetIssuedAt(date any)
	GetIssuer() string
	SetIssuer(issuer string)
	GetAudience() []string
	SetAudience(audience []string)
	GetExpires() *jwt.NumericDate
	SetExpires(date any)
	GetSessionId() string
	SetSessionId(sessionId string)
	GetUser() string
	SetUser(user string)
	GetEmail() string
	SetEmail(email string)
	GetPhone() string
	SetPhone(phone string)
	GetRole() string
	SetRole(role string)
	GetAdmin() bool
	SetAdmin(admin bool)
	GetLevel() int
	SetLevel(level int)
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
	GetEmail() string
	SetEmail(email string)
	GetPhone() string
	SetPhone(phone string)
	GetRole() string
	SetRole(role string)
	GetAdmin() bool
	SetAdmin(admin bool)
	GetLevel() int
	SetLevel(level int)
}

type JwtClaimsDataAccess struct {
	Identity  string           `mapstructure:"jti" json:"jti,omitempty" yaml:"jti,omitempty"`
	Issuer    string           `mapstructure:"iss" json:"iss,omitempty" yaml:"iss,omitempty"`
	Subject   string           `mapstructure:"sub" json:"sub,omitempty" yaml:"sub,omitempty"`
	Audience  []string         `mapstructure:"aud" json:"aud,omitempty" yaml:"aud,omitempty"`
	NotBefore *jwt.NumericDate `mapstructure:"nbf" json:"nbf,omitempty" yaml:"nbf,omitempty"`
	IssuedAt  *jwt.NumericDate `mapstructure:"iat" json:"iat,omitempty" yaml:"iat,omitempty"`
	ExpiresAt *jwt.NumericDate `mapstructure:"exp" json:"exp,omitempty" yaml:"exp,omitempty"`
	SessionId string           `mapstructure:"sid" json:"sid,omitempty" yaml:"sid,omitempty"`
	User      string           `mapstructure:"name" json:"name,omitempty" yaml:"name,omitempty"`
	Email     string           `mapstructure:"email" json:"email,omitempty" yaml:"email,omitempty"`
	Phone     string           `mapstructure:"phone" json:"phone,omitempty" yaml:"phone,omitempty"`
	Role      string           `mapstructure:"role" json:"role,omitempty" yaml:"role,omitempty"`
	Admin     bool             `mapstructure:"admin" json:"admin,omitempty" yaml:"admin,omitempty"`
	Level     int              `mapstructure:"level" json:"level,omitempty" yaml:"level,omitempty"`
}

func NewJwtClaimsDataAccess(claims *jwt.RegisteredClaims) JwtClaimsDataAccessImpl {
	return &JwtClaimsDataAccess{
		Identity:  claims.ID,
		Subject:   claims.Subject,
		Issuer:    claims.Issuer,
		Audience:  claims.Audience,
		NotBefore: claims.NotBefore,
		IssuedAt:  claims.IssuedAt,
		ExpiresAt: claims.ExpiresAt,
	}
}

func NewEmptyJwtClaimsDataAccess() JwtClaimsDataAccessImpl {
	identity := NewUUID().String()
	return &JwtClaimsDataAccess{
		Identity: identity,
	}
}

func (claimsDataAccess *JwtClaimsDataAccess) GetIdentity() string {
	return claimsDataAccess.Identity
}

func (claimsDataAccess *JwtClaimsDataAccess) SetIdentity(identity string) {
	claimsDataAccess.Identity = identity
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

func (claimsDataAccess *JwtClaimsDataAccess) GetEmail() string {
	return claimsDataAccess.Email
}

func (claimsDataAccess *JwtClaimsDataAccess) SetEmail(email string) {
	claimsDataAccess.Email = email
}

func (claimsDataAccess *JwtClaimsDataAccess) GetPhone() string {
	return claimsDataAccess.Phone
}

func (claimsDataAccess *JwtClaimsDataAccess) SetPhone(phone string) {
	claimsDataAccess.Phone = phone
}

func (claimsDataAccess *JwtClaimsDataAccess) GetRole() string {
	return claimsDataAccess.Role
}

func (claimsDataAccess *JwtClaimsDataAccess) SetRole(role string) {
	claimsDataAccess.Role = role
}

func (claimsDataAccess *JwtClaimsDataAccess) GetAdmin() bool {
	return claimsDataAccess.Admin
}

func (claimsDataAccess *JwtClaimsDataAccess) SetAdmin(admin bool) {
	claimsDataAccess.Admin = admin
}

func (claimsDataAccess *JwtClaimsDataAccess) GetLevel() int {
	return claimsDataAccess.Level
}

func (claimsDataAccess *JwtClaimsDataAccess) SetLevel(level int) {
	claimsDataAccess.Level = level
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

func GetJwtClaimsFromJwtToken(token JwtTokenImpl) (JwtClaimsImpl, error) {
	var ok bool
	var claims jwt.MapClaims
	KeepVoid(ok, claims)

	jwtToken := token.(*jwt.Token)
	if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, ErrJwtClaimsInvalid
	}

	return NewJwtClaims(claims, jwtToken.Method), nil
}

func (j *JwtClaims) GetSigningMethod() jwt.SigningMethod {
	return j.method
}

func (j *JwtClaims) SetSigningMethod(method jwt.SigningMethod) {
	j.method = method
}

func (j *JwtClaims) GetDataAccess() *JwtClaimsDataAccess {
	return &JwtClaimsDataAccess{
		Identity:  j.GetIdentity(),
		Subject:   j.GetSubject(),
		Issuer:    j.GetIssuer(),
		Audience:  j.GetAudience(),
		IssuedAt:  j.GetIssued(),
		ExpiresAt: j.GetExpires(),
		SessionId: j.GetIdentity(),
		User:      j.GetUser(),
		Email:     j.GetEmail(),
		Phone:     j.GetPhone(),
		Role:      j.GetRole(),
		Admin:     j.GetAdmin(),
		Level:     j.GetLevel(),
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
	j.claims[key] = value
	return true
}

func (j *JwtClaims) ToJwtToken() JwtTokenImpl {
	return jwt.NewWithClaims(j.method, j.claims)
}

func (j *JwtClaims) ToJwtTokenString(secretKey string) string {
	token := j.ToJwtToken()
	dataSecretKey := encodeSecretKey(secretKey, j.method)
	return Unwrap(token.SignedString(dataSecretKey))
}

func NewNumericDateFromSeconds(f float64) *jwt.NumericDate {
	round, frac := math.Modf(f)
	return jwt.NewNumericDate(time.Unix(int64(round), int64(frac*1e9)))
}

func (j *JwtClaims) ParseNumericDate(key JwtClaimNamed) *jwt.NumericDate {
	var ok bool
	var err error
	KeepVoid(ok, err)

	value := Unwrap(j.Get(string(key)))

	switch exp := value.(type) {
	case float64:
		if exp != 0 {
			return NewNumericDateFromSeconds(exp)
		}

		panic("zero value")
	case json.Number:
		val := Unwrap(exp.Float64())
		return NewNumericDateFromSeconds(val)
	default:
		panic("invalid data type")
	}
}

func (j *JwtClaims) ParseString(key JwtClaimNamed) string {
	temp := Unwrap(j.Get(string(key)))
	return ToString(temp)
}

func (j *JwtClaims) ParseBool(key JwtClaimNamed) bool {
	temp := Unwrap(j.Get(string(key)))
	return GetBoolValueReflect(temp)
}

func (j *JwtClaims) ParseInt(key JwtClaimNamed) int {
	temp := Unwrap(j.Get(string(key)))
	return int(GetIntValueReflect(temp))
}

func (j *JwtClaims) ParseManyString(key JwtClaimNamed) []string {
	var ok bool
	var temp []string
	KeepVoid(ok, temp)

	value := Unwrap(j.Get(string(key)))

	switch value.(type) {
	case string:
		temp = []string{value.(string)}
	case []string:
		temp = value.([]string)
	case []any:
		values := value.([]any)
		temp = make([]string, len(values))
		for i, v := range values {
			temp[i] = v.(string)
		}
	default:
		panic("invalid data type")
	}

	return temp
}

func (j *JwtClaims) GetIdentity() string {
	return j.ParseString(JwtClaimIdentity)
}

func (j *JwtClaims) SetIdentity(identity string) {
	j.Set(string(JwtClaimIdentity), identity)
}

func (j *JwtClaims) GetSubject() string {
	return j.ParseString(JwtClaimSubject)
}

func (j *JwtClaims) SetSubject(subject string) {
	j.Set(string(JwtClaimSubject), subject)
}

func (j *JwtClaims) GetIssued() *jwt.NumericDate {
	return j.ParseNumericDate(JwtClaimIssuedAt)
}

func (j *JwtClaims) SetIssuedAt(date any) {
	j.Set(string(JwtClaimIssuedAt), GetJwtNumericDateFromAny(date))
}

func (j *JwtClaims) GetIssuer() string {
	return j.ParseString(JwtClaimIssuer)
}

func (j *JwtClaims) SetIssuer(issuer string) {
	j.Set(string(JwtClaimIssuer), issuer)
}

func (j *JwtClaims) GetAudience() []string {
	return j.ParseManyString(JwtClaimAudience)
}

func (j *JwtClaims) SetAudience(audience []string) {
	j.Set(string(JwtClaimAudience), audience)
}

func (j *JwtClaims) GetExpires() *jwt.NumericDate {
	return j.ParseNumericDate(JwtClaimExpires)
}

func (j *JwtClaims) SetExpires(date any) {
	j.Set(string(JwtClaimExpires), GetJwtNumericDateFromAny(date))
}

func (j *JwtClaims) GetSessionId() string {
	return j.ParseString(JwtClaimSessionId)
}

func (j *JwtClaims) SetSessionId(sessionId string) {
	j.Set(string(JwtClaimSessionId), sessionId)
}

func (j *JwtClaims) GetUser() string {
	return j.ParseString(JwtClaimUser)
}

func (j *JwtClaims) SetUser(user string) {
	j.Set(string(JwtClaimUser), user)
}

func (j *JwtClaims) GetEmail() string {
	return j.ParseString(JwtClaimEmail)
}

func (j *JwtClaims) SetEmail(email string) {
	j.Set(string(JwtClaimEmail), email)
}

func (j *JwtClaims) GetPhone() string {
	return j.ParseString(JwtClaimPhone)
}

func (j *JwtClaims) SetPhone(email string) {
	j.Set(string(JwtClaimPhone), email)
}

func (j *JwtClaims) GetRole() string {
	return j.ParseString(JwtClaimRole)
}

func (j *JwtClaims) SetRole(role string) {
	j.Set(string(JwtClaimRole), role)
}

func (j *JwtClaims) GetAdmin() bool {
	return j.ParseBool(JwtClaimAdmin)
}

func (j *JwtClaims) SetAdmin(admin bool) {
	j.Set(string(JwtClaimAdmin), admin)
}

func (j *JwtClaims) GetLevel() int {
	return j.ParseInt(JwtClaimLevel)
}

func (j *JwtClaims) SetLevel(level int) {
	j.Set(string(JwtClaimLevel), level)
}

func ParseJwtTokenUnverified(token string) (JwtTokenImpl, error) {
	var err error
	var parts []string
	var jwtToken JwtTokenImpl
	KeepVoid(err, parts, jwtToken)

	parser := jwt.NewParser()
	claims := make(jwt.MapClaims)

	if jwtToken, parts, err = parser.ParseUnverified(token, claims); err != nil {
		return nil, ErrJwtTokenInvalid
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
		return nil, ErrJwtTokenInvalid
	}

	return jwtToken, nil
}

func CvtJwtClaimsDataAccessToJwtClaims(claimsDataAccess JwtClaimsDataAccessImpl, jwtSigningMethod jwt.SigningMethod) JwtClaimsImpl {
	claims := NewJwtClaims(jwt.MapClaims{}, jwtSigningMethod)
	claims.SetIdentity(claimsDataAccess.GetIdentity())
	claims.SetSubject(claimsDataAccess.GetSubject())
	claims.SetIssuer(claimsDataAccess.GetIssuer())
	claims.SetAudience(claimsDataAccess.GetAudience())
	claims.SetIssuedAt(claimsDataAccess.GetIssued())
	claims.SetExpires(claimsDataAccess.GetExpires())
	claims.SetSessionId(claimsDataAccess.GetSessionId())
	claims.SetUser(claimsDataAccess.GetUser())
	claims.SetEmail(claimsDataAccess.GetEmail())
	claims.SetPhone(claimsDataAccess.GetPhone())
	claims.SetRole(claimsDataAccess.GetRole())
	claims.SetAdmin(claimsDataAccess.GetAdmin())
	claims.SetLevel(claimsDataAccess.GetLevel())
	return claims
}

func CvtJwtClaimsToJwtClaimsDataAccess(claims JwtClaimsImpl) (JwtClaimsDataAccessImpl, jwt.SigningMethod) {
	claimsDataAccess := new(JwtClaimsDataAccess)
	claimsDataAccess.SetIdentity(claims.GetIdentity())
	claimsDataAccess.SetSubject(claims.GetSubject())
	claimsDataAccess.SetIssuer(claims.GetIssuer())
	claimsDataAccess.SetAudience(claims.GetAudience())
	claimsDataAccess.SetIssuedAt(claims.GetIssued())
	claimsDataAccess.SetExpiresAt(claims.GetExpires())
	claimsDataAccess.SetSessionId(claims.GetSessionId())
	claimsDataAccess.SetUser(claims.GetUser())
	claimsDataAccess.SetEmail(claims.GetEmail())
	claimsDataAccess.SetPhone(claims.GetPhone())
	claimsDataAccess.SetRole(claims.GetRole())
	claimsDataAccess.SetAdmin(claims.GetAdmin())
	claimsDataAccess.SetLevel(claims.GetLevel())
	return claimsDataAccess, claims.GetSigningMethod()
}

func GenerateJwtToken(jwtClaims JwtClaimsImpl, secretKey string) string {
	return jwtClaims.ToJwtTokenString(secretKey)
}
