package main

import (
	"fmt"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func main() {
	config := globals.GetJwtConfig()
	fmt.Printf("%+v\n", config)

	temp := nokocore.MapAny{
		"test": &nokocore.MapAny{
			"value": 12,
		},
	}

	nokocore.SetValueReflect(temp.Get("test"), &nokocore.MapAny{
		"value": 24,
	})

	fmt.Println(nokocore.ShikaYamlEncode(temp))

	func() {
		jwtConfig := globals.GetJwtConfig()

		timeUtcNow := nokocore.GetTimeUtcNow()
		expires := timeUtcNow.Add(jwtConfig.GetExpiresIn())

		jwtClaimsDataAccess := nokocore.NewEmptyJwtClaimsDataAccess()

		identity := nokocore.NewUUID()
		sessionId := nokocore.NewUUID()

		nokocore.KeepVoid(expires, identity, sessionId)

		jwtClaimsDataAccess.SetIdentity(identity.String())
		jwtClaimsDataAccess.SetSubject("NokoWebToken")
		jwtClaimsDataAccess.SetAudience(jwtConfig.Audience)
		jwtClaimsDataAccess.SetIssuer(jwtConfig.Issuer)
		jwtClaimsDataAccess.SetIssuedAt(timeUtcNow)
		jwtClaimsDataAccess.SetUser("user")
		jwtClaimsDataAccess.SetSessionId(sessionId.String())
		jwtClaimsDataAccess.SetEmail("user@example.com")
		jwtClaimsDataAccess.SetPhone("+62 8123 4444 5555")
		jwtClaimsDataAccess.SetRole("User")
		jwtClaimsDataAccess.SetExpiresAt(expires)

		jwtClaims := nokocore.CvtJwtClaimsAccessDataToJwtClaims(jwtClaimsDataAccess, jwtConfig.GetSigningMethod())
		fmt.Println(jwtClaims.ToJwtTokenString(jwtConfig.SecretKey))
	}()

	// valid token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieW91ci1hdWRpZW5jZSJdLCJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjE3MzE0MTkwNzgsImlhdCI6MTczMTQxNTQ3OCwiaXNzIjoieW91ci1pc3N1ZXIiLCJqdGkiOiIwMTkzMjA2Ny04Zjc5LTcwYTItOTU1MS0wMmM1NTJlNTMwODEiLCJwaG9uZSI6Iis2MiA4MTIzIDQ0NDQgNTU1NSIsInJvbGUiOiJVc2VyIiwic2lkIjoiMDE5MzIwNjctOGY3OS03MGEzLTk5YTgtNGMxZGI5ZDFhOGE5Iiwic3ViIjoiTm9rb1dlYlRva2VuIiwidXNlciI6InVzZXIifQ.BG1uYn_z9qx022T-SRQR_12P_y76s8eAlyhlt5Ago1I
	// broken token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOm51bGwsImVtYWlsIjoiIiwiZXhwIjpudWxsLCJpYXQiOm51bGwsImlzcyI6IiIsImp0aSI6IiIsInBob25lIjoiIiwicm9sZSI6IiIsInNpZCI6IiIsInN1YiI6IiIsInVzZXIiOiIifQ.lxvWtt5ZIRUtIQyo1JImjLlGA1lWom19Srnom6bNDx0

	var ff = func() string {
		defer nokocore.HandlePanic(nil)
		brokenToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOm51bGwsImVtYWlsIjoiIiwiZXhwIjpudWxsLCJpYXQiOm51bGwsImlzcyI6IiIsImp0aSI6IiIsInBob25lIjoiIiwicm9sZSI6IiIsInNpZCI6IiIsInN1YiI6IiIsInVzZXIiOiIifQ.lxvWtt5ZIRUtIQyo1JImjLlGA1lWom19Srnom6bNDx0"
		jwtToken := nokocore.Unwrap(nokocore.ParseJwtTokenUnverified(brokenToken))
		jwtClaims := nokocore.Unwrap(nokocore.GetJwtClaimsFromJwtToken(jwtToken))
		jwtClaimsDataAccess, jwtSigningMethod := nokocore.CvtJwtClaimsToJwtClaimsAccessData(jwtClaims)
		nokocore.KeepVoid(jwtClaimsDataAccess, jwtSigningMethod)
		fmt.Printf("%+v\n", jwtClaimsDataAccess)
		return "done"
	}()

	fmt.Println("GET OUTPUT", ff)
}
