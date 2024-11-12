package nokocore

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

func TestJwtClaims_ToJwtClaimsAccessData(t *testing.T) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		t.Error(err)
	//	}
	//}()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJ0aGlzIGlzIGp3dCBpZGVudGl0eSIsInNpZCI6InRoaXMgaXMgc2Vzc2lvbiBpZCIsImV4cCI6NDc2OTI0MDkzMSwiaXNzIjoidGhpcyBpcyBpc3N1ZXIiLCJhdWQiOiJ0aGlzIGlzIGF1ZGllbmNlIiwic3ViIjoiMTIzNDU2Nzg5MCIsIm5hbWUiOiJKb2huIERvZSIsImVtYWlsIjoiam9obmRvZUBleGFtcGxlLmNvbSIsInJvbGUiOiJVc2VyIiwiaWF0IjoxNzI5MjQwOTMxfQ.9tuR0WlIBiyay9554ELUL-9dBjn9C-Ba5UN7TqffP-Y"
	jwtToken := Unwrap(ParseJwtToken(token, "your-256-bit-secret", jwt.SigningMethodHS256))
	jwtClaims := Unwrap(GetJwtClaimsFromJwtToken(jwtToken, jwt.SigningMethodHS256))
	jwtClaimsAccessData := CvtJwtClaimsToJwtClaimsAccessData(jwtClaims)

	fmt.Println(ShikaYamlEncode(jwtClaimsAccessData))
}
