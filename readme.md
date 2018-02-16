## Simple Library for Validating Oauth Tokens from Okta

Simple Library for Validating Okta Oauth 2.0 (JWT) tokens from:
* Default Okta OIDC Service (*Every Okta Instance has this*)
* Okta Authorization Server


This package is very basic, and pulls the keys from Okta On-Demand. It doesn't cache the key,
I will likely extend this so you can Parse with a preloaded key.

### Run it like this
~~~~
package main

import (
	"oktajwt"
	"fmt"
)

func main () {

	token := `eyJraWQiOiJ0TnRuRGdlWGVwYmYyTlpsVmp6S0dkdFBVOW1uT2lkdHcwcHVvM3MtXzhZIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIwMHUxOGVlaHUzNDlhUzJ5WDFkOCIsIm5hbWUiOiJva3RhcHJveHkgb2t0YXByb3h5IiwidmVyIjoxLCJpc3MiOiJodHRwczovL2NvbXBhbnl4Lm9rdGEuY29tIiwiYXVkIjoidlpWNkNwOHJuNWx4ck45YVo2ODgiLCJpYXQiOjE1MTg4MDk4MjcsImV4cCI6MTUxODgxMzQyNywianRpIjoiSUQueElpTzJWMThJQm5GU0lXVHRBN0dQOUd2MWdCdHR3QlhPOERKcHlnM3RoUSIsImFtciI6WyJwd2QiXSwiaWRwIjoiMDBveTc0YzBnd0hOWE1SSkJGUkkiLCJub25jZSI6Im4tMFM2X1d6QTJNaiIsInByZWZlcnJlZF91c2VybmFtZSI6Im9rdGFwcm94eUBva3RhLmNvbSIsImF1dGhfdGltZSI6MTUxODgwOTgyMSwiYXRfaGFzaCI6ImdGRTJ6V3NYdXNiUmRZZF85S2dIZFEifQ.DnXnMQLY2mCRCifu1FD4wH6JEYweuVTEZso0G6893JlQfxCwlnv9t4c94DIoROMvX_HsWSLhMRoQMutVLUUJfgER-3uoFTcyWyVgDUm-SRFk6jHEe7nioMdEx-ScIpbxK4GOo-Z7XXY3qtOixsDB8omAt90ipafLxzSZj7d2FJndFgfoqCvcHOWvWATVK74AyRjkv3OLfxQ4Cjn7M8c7q0BK3W-71dhxqqU5FXQ8LHotsD1757IZF4MHixySCJGpaq4zSwB9unBrkSmVB7Pxdpy9boNTw6c1pxyIW_aJFCLiREaoHizVXZeBkoCriQuQMcOYOJ7C9UulVeMmz0-2rw`

	result, err:=oktajwt.Oktaparsejwt ( token )
	if result!=nil {
		fmt.Println(err)

	} else {
		fmt.Println("Token is good")

	}
}

~~~~