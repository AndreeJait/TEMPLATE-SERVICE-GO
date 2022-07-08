package common

import (
	"fmt"
	"strings"

	"github.com/AndreeJait/GO-ANDREE-UTILITIES/util/andreerror"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func JwtMiddleWare(config *config.EnvConfiguration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorizationHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authorizationHeader, "Bearer") {
				return SystemResponse(c, nil, andreerror.New(andreerror.UNAUTHORIZE, errors.New("invalid Token.")))
			}

			tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("signing method invalid")
				} else if method != JWT_SIGNING_METHOD {
					return nil, fmt.Errorf("signing method invalid")
				}

				return []byte(config.SecretKey), nil
			})

			if err != nil {
				return SystemResponse(c, nil, andreerror.New(andreerror.UNAUTHORIZE, errors.New("invalid Token.")))
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				return SystemResponse(c, nil, andreerror.New(andreerror.UNAUTHORIZE, errors.New("invalid Token.")))
			}

			c.Set("userInfo", claims)

			return next(c)
		}
	}
}
