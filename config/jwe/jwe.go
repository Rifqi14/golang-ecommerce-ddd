package jwe

import (
	"encoding/json"
	"fmt"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
)

type Credential struct {
	KeyLocation string
	Passphrase  string
}

func (c *Credential) GenerateJwePayload(payload map[string]interface{}) (res string, err error) {
	privkey, err := rsaConfigSetup(c.KeyLocation, c.Passphrase)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	// Generate payload
	payloadString, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	// Generate JWE
	jweRes, err := jwe.Encrypt([]byte(payloadString), jwa.RSA1_5, &privkey.PublicKey, jwa.A128CBC_HS256, jwa.Deflate)
	res = string(jweRes)

	return res, nil
}

// Rollback
func (c *Credential) Rollback(id string) (res map[string]interface{}, err error) {
	privkey, err := rsaConfigSetup(c.KeyLocation, c.Passphrase)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	// Generate JWE
	jweRes, err := jwe.Decrypt([]byte(id), jwa.RSA1_5, privkey)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	// Unmarshal JWE
	res = map[string]interface{}{}
	err = json.Unmarshal(jweRes, &res)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	return res, nil
}
