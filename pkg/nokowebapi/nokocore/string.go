package nokocore

const (
	Digits      = "0123456789"
	OctDigits   = "01234567"
	HexDigits   = "0123456789ABCDEF"
	AlphaUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaLower  = "abcdefghijklmnopqrstuvwxyz"
	AlphaNum    = AlphaLower + AlphaUpper + Digits
	Punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	WhiteSpace  = " \t\r\n\v\f\xc2\xa0"
	Printable   = Digits + AlphaUpper + AlphaLower + Punctuation + WhiteSpace
)
