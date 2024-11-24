package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"nokowebapi/nokocore"
	"regexp"
	"strconv"
	"strings"
)

type CORSConfig struct {
	Origins       []string `mapstructure:"origins" json:"origins" yaml:"origins"`
	Methods       []string `mapstructure:"methods" json:"methods" yaml:"methods"`
	Headers       []string `mapstructure:"headers" json:"headers" yaml:"headers"`
	ExposeHeaders []string `mapstructure:"expose_headers" json:"exposeHeaders" yaml:"expose_headers"`
	Credentials   bool     `mapstructure:"credentials" json:"credentials" yaml:"credentials"`
	MaxAge        int      `mapstructure:"max_age" json:"maxAge" yaml:"max_age"`
}

var defaultCORSConfig = &CORSConfig{
	Origins:     []string{"*"},
	Methods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
	Headers:     []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
	Credentials: true,
	MaxAge:      86400,
}

func matchScheme(domain, pattern string) bool {
	domainIdx := strings.Index(domain, ":")
	patternIdx := strings.Index(pattern, ":")

	return domainIdx != -1 && patternIdx != -1 && domain[:domainIdx] == pattern[:patternIdx]
}

func matchSubdomain(domain, pattern string) bool {
	if !matchScheme(domain, pattern) {
		return false
	}

	domainIdx := strings.Index(domain, "://")
	patternIdx := strings.Index(pattern, "://")

	if domainIdx == -1 || patternIdx == -1 {
		return false
	}
	domainAuth := domain[domainIdx+3:]

	if len(domainAuth) > 253 {
		return false
	}
	patAuth := pattern[patternIdx+3:]

	domainComp := strings.Split(domainAuth, ".")
	patternComp := strings.Split(patAuth, ".")

	for i := len(domainComp)/2 - 1; i >= 0; i-- {
		opp := len(domainComp) - 1 - i
		domainComp[i], domainComp[opp] = domainComp[opp], domainComp[i]
	}

	for i := len(patternComp)/2 - 1; i >= 0; i-- {
		opp := len(patternComp) - 1 - i
		patternComp[i], patternComp[opp] = patternComp[opp], patternComp[i]
	}

	for i, v := range domainComp {
		if len(patternComp) <= i {
			return false
		}
		p := patternComp[i]
		if p == "*" {
			return true
		}
		if p != v {
			return false
		}
	}

	return false
}

func CORS() echo.MiddlewareFunc {
	return CORSWithConfig(defaultCORSConfig)
}

func CORSWithConfig(config *CORSConfig) echo.MiddlewareFunc {
	if len(config.Origins) == 0 {
		config.Origins = defaultCORSConfig.Origins
	}

	hasCustomAllowMethods := true
	if len(config.Methods) == 0 {
		hasCustomAllowMethods = false
		config.Methods = defaultCORSConfig.Methods
	}

	var allowOriginPatterns []string
	for i, origin := range config.Origins {
		nokocore.KeepVoid(i)

		pattern := regexp.QuoteMeta(origin)
		pattern = strings.ReplaceAll(pattern, "\\*", ".*")
		pattern = strings.ReplaceAll(pattern, "\\?", ".")
		pattern = "^" + pattern + "$"
		allowOriginPatterns = append(allowOriginPatterns, pattern)
	}

	allowMethods := strings.Join(config.Methods, ",")
	allowHeaders := strings.Join(config.Headers, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")

	maxAge := "0"
	if config.MaxAge > 0 {
		maxAge = strconv.Itoa(config.MaxAge)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()
			res := ctx.Response()

			allowOrigin := ""
			origin := req.Header.Get(echo.HeaderOrigin)

			res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
			preflight := req.Method == http.MethodOptions

			routerAllowMethods := ""
			if preflight {
				tmpAllowMethods, ok := ctx.Get(echo.ContextKeyHeaderAllow).(string)
				if ok && tmpAllowMethods != "" {
					routerAllowMethods = tmpAllowMethods
					res.Header().Set(echo.HeaderAllow, routerAllowMethods)
				}
			}

			if origin == "" {
				if !preflight {
					return next(ctx)
				}
				return ctx.NoContent(http.StatusNoContent)
			}

			for i, o := range config.Origins {
				nokocore.KeepVoid(i)

				if o == "*" || o == origin {
					allowOrigin = o
					break
				}

				if matchSubdomain(origin, o) {
					allowOrigin = origin
					break
				}
			}

			checkPatterns := false
			if allowOrigin == "" {
				if len(origin) <= (253+3+5) && strings.Contains(origin, "://") {
					checkPatterns = true
				}
			}
			if checkPatterns {
				for i, re := range allowOriginPatterns {
					nokocore.KeepVoid(i)

					if match, _ := regexp.MatchString(re, origin); match {
						allowOrigin = origin
						break
					}
				}
			}

			if allowOrigin == "" {
				if !preflight {
					return next(ctx)
				}
				return ctx.NoContent(http.StatusNoContent)
			}

			res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
			if config.Credentials {
				res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			}

			if !preflight {
				if exposeHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return next(ctx)
			}

			res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestMethod)
			res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestHeaders)

			if !hasCustomAllowMethods && routerAllowMethods != "" {
				res.Header().Set(echo.HeaderAccessControlAllowMethods, routerAllowMethods)
			} else {
				res.Header().Set(echo.HeaderAccessControlAllowMethods, allowMethods)
			}

			if allowHeaders != "" {
				res.Header().Set(echo.HeaderAccessControlAllowHeaders, allowHeaders)

			} else {
				h := req.Header.Get(echo.HeaderAccessControlRequestHeaders)
				if h != "" {
					res.Header().Set(echo.HeaderAccessControlAllowHeaders, h)
				}
			}

			if config.MaxAge != 0 {
				res.Header().Set(echo.HeaderAccessControlMaxAge, maxAge)
			}

			return ctx.NoContent(http.StatusNoContent)
		}
	}
}
