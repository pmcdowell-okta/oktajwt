package oktajwt

import (
	//"fmt"
	//"os"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"encoding/json"
	"github.com/lestrrat/go-jwx/jwk"
	b64 "encoding/base64"
	"errors"

)

type MyError struct {
	What string
}

type OauthToken struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	Ver int `json:"ver"`
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Iat int `json:"iat"`
	Exp int `json:"exp"`
	Jti string `json:"jti"`
	Amr []string `json:"amr"`
	Idp string `json:"idp"`
	Nonce string `json:"nonce"`
	PreferredUsername string `json:"preferred_username"`
	AuthTime int `json:"auth_time"`
	AtHash string `json:"at_hash"`
}


func Oktaparsejwt ( token string ) (interface{} ,  error) {

	token2, err := jwt.Parse(token, getKey)
	_ = token2
	if err != nil {
		return 1, err
		} else {
		return nil, err

		}
	}


func getKey(token *jwt.Token) (interface{}, error) {

	// TODO: cache response so we don't have to make a request every time
	// we want to verify a JWT

	keyUrl:=""
	rawToken:=strings.Split(token.Raw,".")

	jwtFragment, err := b64.RawURLEncoding.DecodeString(rawToken[1])

	if (err!=nil) {
		return nil, errors.New("Could not 64Decode Token ?")
	}

	res:= OauthToken{}
	err = (json.Unmarshal(jwtFragment, &res))
	if err!=nil {
		return nil, errors.New("Could not Unmashal token ?")
	}

	if ( IsOktaAuthorizationServer(res.Iss)) {
		keyUrl=res.Iss+"/v1/keys"
	} else {
		keyUrl=res.Iss+"/oauth2/v1/keys"
	}

	set, err := jwk.Fetch(keyUrl)

	if err != nil {
		return nil, errors.New("Cannot fetch keys, tried URL"+keyUrl)
	}

	keyID, ok := token.Header["kid"].(string)

	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	if key := set.LookupKeyID(keyID); len(key) == 1 {
		return key[0].Materialize()
	}

	return nil, errors.New("unable to find key")
}

func IsOktaAuthorizationServer ( issuer string) bool {

	if (strings.Contains(issuer,".com/oauth2/")) {
		return true
	} else {
		return false
	}

}

