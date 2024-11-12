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
	Origins       []string `mapstructure:"origins" yaml:"origins" json:"origins"`
	Methods       []string `mapstructure:"methods" yaml:"methods" json:"methods"`
	Headers       []string `mapstructure:"headers" yaml:"headers" json:"headers"`
	ExposeHeaders []string `mapstructure:"expose_headers" yaml:"expose_headers" json:"exposeHeaders"`
	Credentials   bool     `mapstructure:"credentials" yaml:"credentials" json:"credentials"`
	MaxAge        int      `mapstructure:"max_age" yaml:"max_age" json:"maxAge"`
}

var DefaultCORSConfig = CORSConfig{
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
	return CORSWithConfig(DefaultCORSConfig)
}

func CORSWithConfig(config CORSConfig) echo.MiddlewareFunc {
	if len(config.Origins) == 0 {
		config.Origins = DefaultCORSConfig.Origins
	}

	hasCustomAllowMethods := true
	if len(config.Methods) == 0 {
		hasCustomAllowMethods = false
		config.Methods = DefaultCORSConfig.Methods
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
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			allowOrigin := ""
			origin := req.Header.Get(echo.HeaderOrigin)

			res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
			preflight := req.Method == http.MethodOptions

			routerAllowMethods := ""
			if preflight {
				tmpAllowMethods, ok := c.Get(echo.ContextKeyHeaderAllow).(string)
				if ok && tmpAllowMethods != "" {
					routerAllowMethods = tmpAllowMethods
					c.Response().Header().Set(echo.HeaderAllow, routerAllowMethods)
				}
			}

			if origin == "" {
				if !preflight {
					return next(c)
				}
				return c.NoContent(http.StatusNoContent)
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
					return next(c)
				}
				return c.NoContent(http.StatusNoContent)
			}

			res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
			if config.Credentials {
				res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			}

			if !preflight {
				if exposeHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return next(c)
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

			return c.NoContent(http.StatusNoContent)
		}
	}
}
