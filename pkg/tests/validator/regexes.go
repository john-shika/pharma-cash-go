package nokocore

import (
	"nokowebapi/nokocore"
	"regexp"
	"sync"
)

// ref: https://raw.githubusercontent.com/go-playground/validator/refs/heads/master/regexes.go

const (
	alphaRegex                 = "^[a-zA-Z]+$"
	alphaNumericRegex          = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegex          = "^[\\p{L}]+$"
	alphaUnicodeNumericRegex   = "^[\\p{L}\\p{N}]+$"
	numericRegex               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegex                = "^[0-9]+$"
	hexadecimalRegex           = "^(0[xX])?[0-9a-fA-F]+$"
	hexColorRegex              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	rgbRegex                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegex                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegex                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegex                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailRegex                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	e164Regex                  = "^\\+[1-9]?[0-9]{7,14}$"
	base32Regex                = "^(?:[A-Z2-7]{8})*(?:[A-Z2-7]{2}={6}|[A-Z2-7]{4}={4}|[A-Z2-7]{5}={3}|[A-Z2-7]{7}=|[A-Z2-7]{8})$"
	base64Regex                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegex             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	base64RawURLRegex          = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2,4})$"
	iSBN10Regex                = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13Regex                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	iSSNRegex                  = "^(?:[0-9]{4}-[0-9]{3}[0-9X])$"
	uUID3Regex                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID4Regex                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID5Regex                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegex                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID3RFC4122Regex          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uUID4RFC4122Regex          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUID5RFC4122Regex          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUIDRFC4122Regex           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uLIDRegex                  = "^(?i)[A-HJKMNP-TV-Z0-9]{26}$"
	md4Regex                   = "^[0-9a-f]{32}$"
	md5Regex                   = "^[0-9a-f]{32}$"
	sha256Regex                = "^[0-9a-f]{64}$"
	sha384Regex                = "^[0-9a-f]{96}$"
	sha512Regex                = "^[0-9a-f]{128}$"
	ripemd128Regex             = "^[0-9a-f]{32}$"
	ripemd160Regex             = "^[0-9a-f]{40}$"
	tiger128Regex              = "^[0-9a-f]{32}$"
	tiger160Regex              = "^[0-9a-f]{40}$"
	tiger192Regex              = "^[0-9a-f]{48}$"
	aSCIIRegex                 = "^[\x00-\x7F]*$"
	printableASCIIRegex        = "^[\x20-\x7E]*$"
	multibyteRegex             = "[^\x00-\x7F]"
	dataURIRegex               = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudeRegex              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegex             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	sSNRegex                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnameRegexRFC952        = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                                                   // https://tools.ietf.org/html/rfc952
	hostnameRegexRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?$`                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	fqdnRegexStringRFC1123     = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as hostnameRegexRFC1123 but must contain a non-numerical TLD (possibly ending with '.')
	btcAddressRegex            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`                                                                             // bitcoin address
	btcAddressUpperRegexBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerRegexBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressRegex            = `^0x[0-9a-fA-F]{40}$`
	ethAddressUpperRegex       = `^0x[0-9A-F]{40}$`
	ethAddressLowerRegex       = `^0x[0-9a-f]{40}$`
	uRLEncodedRegex            = `^(?:[^%]|%[0-9A-Fa-f]{2})*$`
	hTMLEncodedRegex           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	hTMLRegex                  = `<[/]?([a-zA-Z]+).*?>`
	jWTRegex                   = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
	splitParamsRegex           = `'[^']*'|\S+`
	bicRegex                   = `^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`
	semverRegex                = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$` // numbered capture groups https://semver.org/
	dnsLabelRegexRFC1035       = "^[a-z]([-a-z0-9]*[a-z0-9]){0,62}$"
	cveRegex                   = `^CVE-(1999|2\d{3})-(0[^0]\d{2}|0\d[^0]\d{1}|0\d{2}[^0]|[1-9]{1}\d{3,})$` // CVE Format Id https://cve.mitre.org/cve/identifiers/syntaxchange.html
	mongodbIdRegex             = "^[a-f\\d]{24}$"
	mongodbConnStringRegex     = "^mongodb(\\+srv)?:\\/\\/(([a-zA-Z\\d]+):([a-zA-Z\\d$:\\/?#\\[\\]@]+)@)?(([a-z\\d.-]+)(:[\\d]+)?)((,(([a-z\\d.-]+)(:(\\d+))?))*)?(\\/[a-zA-Z-_]{1,64})?(\\?(([a-zA-Z]+)=([a-zA-Z\\d]+))(&(([a-zA-Z\\d]+)=([a-zA-Z\\d]+))?)*)?$"
	cronRegex                  = `(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|((\*|\d+)(\/|-)\d+)|\d+|\*) ?){5,7})`
	spicedbIDRegex             = `^(([a-zA-Z0-9/_|\-=+]{1,})|\*)$`
	spicedbPermissionRegex     = "^([a-z][a-z0-9_]{1,62}[a-z0-9])?$"
	spicedbTypeRegex           = "^([a-z][a-z0-9_]{1,61}[a-z0-9]/)?[a-z][a-z0-9_]{1,62}[a-z0-9]$"

	emailxRegex = "^(([a-zA-Z0-9]+)(([\\.\\-\\_]([a-zA-Z0-9]+))+)?)\\@(([a-zA-Z0-9]+)(([\\.\\-\\_]([a-zA-Z0-9]+))+)?)\\.([a-zA-Z0-9]+)$"
	phoneRegex  = "^(\\+[0-9]{1,2}|0)\\s?([0-9]{3,4}(-|\\s)?){2,3}[0-9]{3,4}$"
)

var (
	AlphaRegex                 = LazyRegexCompile(alphaRegex)
	AlphaNumericRegex          = LazyRegexCompile(alphaNumericRegex)
	AlphaUnicodeRegex          = LazyRegexCompile(alphaUnicodeRegex)
	AlphaUnicodeNumericRegex   = LazyRegexCompile(alphaUnicodeNumericRegex)
	NumericRegex               = LazyRegexCompile(numericRegex)
	NumberRegex                = LazyRegexCompile(numberRegex)
	HexadecimalRegex           = LazyRegexCompile(hexadecimalRegex)
	HexColorRegex              = LazyRegexCompile(hexColorRegex)
	RgbRegex                   = LazyRegexCompile(rgbRegex)
	RgbaRegex                  = LazyRegexCompile(rgbaRegex)
	HslRegex                   = LazyRegexCompile(hslRegex)
	HslaRegex                  = LazyRegexCompile(hslaRegex)
	E164Regex                  = LazyRegexCompile(e164Regex)
	EmailRegex                 = LazyRegexCompile(emailRegex)
	Base32Regex                = LazyRegexCompile(base32Regex)
	Base64Regex                = LazyRegexCompile(base64Regex)
	Base64URLRegex             = LazyRegexCompile(base64URLRegex)
	Base64RawURLRegex          = LazyRegexCompile(base64RawURLRegex)
	ISBN10Regex                = LazyRegexCompile(iSBN10Regex)
	ISBN13Regex                = LazyRegexCompile(iSBN13Regex)
	ISSNRegex                  = LazyRegexCompile(iSSNRegex)
	UUID3Regex                 = LazyRegexCompile(uUID3Regex)
	UUID4Regex                 = LazyRegexCompile(uUID4Regex)
	UUID5Regex                 = LazyRegexCompile(uUID5Regex)
	UUIDRegex                  = LazyRegexCompile(uUIDRegex)
	UUID3RFC4122Regex          = LazyRegexCompile(uUID3RFC4122Regex)
	UUID4RFC4122Regex          = LazyRegexCompile(uUID4RFC4122Regex)
	UUID5RFC4122Regex          = LazyRegexCompile(uUID5RFC4122Regex)
	UUIDRFC4122Regex           = LazyRegexCompile(uUIDRFC4122Regex)
	ULIDRegex                  = LazyRegexCompile(uLIDRegex)
	Md4Regex                   = LazyRegexCompile(md4Regex)
	Md5Regex                   = LazyRegexCompile(md5Regex)
	Sha256Regex                = LazyRegexCompile(sha256Regex)
	Sha384Regex                = LazyRegexCompile(sha384Regex)
	Sha512Regex                = LazyRegexCompile(sha512Regex)
	Ripemd128Regex             = LazyRegexCompile(ripemd128Regex)
	Ripemd160Regex             = LazyRegexCompile(ripemd160Regex)
	Tiger128Regex              = LazyRegexCompile(tiger128Regex)
	Tiger160Regex              = LazyRegexCompile(tiger160Regex)
	Tiger192Regex              = LazyRegexCompile(tiger192Regex)
	ASCIIRegex                 = LazyRegexCompile(aSCIIRegex)
	PrintableASCIIRegex        = LazyRegexCompile(printableASCIIRegex)
	MultibyteRegex             = LazyRegexCompile(multibyteRegex)
	DataURIRegex               = LazyRegexCompile(dataURIRegex)
	LatitudeRegex              = LazyRegexCompile(latitudeRegex)
	LongitudeRegex             = LazyRegexCompile(longitudeRegex)
	SSNRegex                   = LazyRegexCompile(sSNRegex)
	HostnameRegexRFC952        = LazyRegexCompile(hostnameRegexRFC952)
	HostnameRegexRFC1123       = LazyRegexCompile(hostnameRegexRFC1123)
	FqdnRegexRFC1123           = LazyRegexCompile(fqdnRegexStringRFC1123)
	BtcAddressRegex            = LazyRegexCompile(btcAddressRegex)
	BtcUpperAddressRegexBech32 = LazyRegexCompile(btcAddressUpperRegexBech32)
	BtcLowerAddressRegexBech32 = LazyRegexCompile(btcAddressLowerRegexBech32)
	EthAddressRegex            = LazyRegexCompile(ethAddressRegex)
	EthAddressUpperRegex       = LazyRegexCompile(ethAddressUpperRegex)
	EthAddressLowerRegex       = LazyRegexCompile(ethAddressLowerRegex)
	URLEncodedRegex            = LazyRegexCompile(uRLEncodedRegex)
	HTMLEncodedRegex           = LazyRegexCompile(hTMLEncodedRegex)
	HTMLRegex                  = LazyRegexCompile(hTMLRegex)
	JWTRegex                   = LazyRegexCompile(jWTRegex)
	SplitParamsRegex           = LazyRegexCompile(splitParamsRegex)
	BicRegex                   = LazyRegexCompile(bicRegex)
	SemverRegex                = LazyRegexCompile(semverRegex)
	DnsLabelRegexRFC1035       = LazyRegexCompile(dnsLabelRegexRFC1035)
	CveRegex                   = LazyRegexCompile(cveRegex)
	MongodbIdRegex             = LazyRegexCompile(mongodbIdRegex)
	MongodbConnectionRegex     = LazyRegexCompile(mongodbConnStringRegex)
	CronRegex                  = LazyRegexCompile(cronRegex)
	SpicedbIDRegex             = LazyRegexCompile(spicedbIDRegex)
	SpicedbPermissionRegex     = LazyRegexCompile(spicedbPermissionRegex)
	SpicedbTypeRegex           = LazyRegexCompile(spicedbTypeRegex)
)

var cachesRegex = make(map[string]*regexp.Regexp)

type ReOrStrImpl interface {
	string | *regexp.Regexp
}

func GetRegexPattern[T ReOrStrImpl](pattern T) *regexp.Regexp {
	var ok bool
	var re *regexp.Regexp
	var str string
	nokocore.KeepVoid(ok, re, str)

	// check pattern is regex pointer or string
	if re, ok = nokocore.CastPtr[regexp.Regexp](pattern); !ok {
		if str, ok = nokocore.Cast[string](pattern); !ok {
			panic("pattern must be 'regexp.Regexp' or string type")
		}

		// register new regex and store in cachesRegex
		if re, ok = cachesRegex[str]; !ok {
			re = regexp.MustCompile(str)
			cachesRegex[str] = re
			return re
		}

		return re
	}

	return nil
}

func LazyRegexCompile[T ReOrStrImpl](pattern T) func() *regexp.Regexp {
	var regex *regexp.Regexp
	var once sync.Once
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = GetRegexPattern(pattern)
		})
		return regex
	}
}

func RegexMatch[T ReOrStrImpl](pattern T, value []byte) bool {
	re := GetRegexPattern(pattern)
	return re.Match(value)
}

func RegexMatchString[T ReOrStrImpl](pattern T, value string) bool {
	re := GetRegexPattern(pattern)
	return re.MatchString(value)
}

func RegexFind[T ReOrStrImpl](pattern T, value []byte) []byte {
	re := GetRegexPattern(pattern)
	return re.Find(value)
}

func RegexFindString[T ReOrStrImpl](pattern T, value string) string {
	re := GetRegexPattern(pattern)
	return re.FindString(value)
}

func RegexFindAll[T ReOrStrImpl](pattern T, value []byte, n int) [][]byte {
	re := GetRegexPattern(pattern)
	return re.FindAll(value, n)
}

func RegexFindAllString[T ReOrStrImpl](pattern T, value string, n int) []string {
	re := GetRegexPattern(pattern)
	return re.FindAllString(value, n)
}

func RegexReplaceAll[T ReOrStrImpl](pattern T, value []byte, replace []byte) []byte {
	re := GetRegexPattern(pattern)
	return re.ReplaceAll(value, replace)
}

func RegexReplaceAllString[T ReOrStrImpl](pattern T, value string, replace string) string {
	re := GetRegexPattern(pattern)
	return re.ReplaceAllString(value, replace)
}

func RegexSplit[T ReOrStrImpl](pattern T, value string, n int) []string {
	re := GetRegexPattern(pattern)
	return re.Split(value, n)
}
