package auth

import (
	"histos-backend/util"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Payload struct {
	Subject    string `json:"sub"`
	IssuedAt   int64  `json:"iat"`
	Expiration int64  `json:"exp"`
}

func GenerateJwt(userId string) (string, error) {
	header := Header{util.EnvironmentVariables.JWT_SIGNING_ALGORITHM, "JWT"}
	jsonHeader, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	now := time.Now()

	payload := Payload{userId, now.Unix(), now.Add(24 * time.Hour).Unix()}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(jsonHeader)
	encodedPayload := base64.RawURLEncoding.EncodeToString(jsonPayload)

	jointString := encodedHeader + "." + encodedPayload
	signature := signJwt(jointString)

	jwt := jointString + "." + signature

	return jwt, nil
}

func signJwt(jwt string) string {
	mac := hmac.New(sha256.New, []byte(util.EnvironmentVariables.JWT_SIGNING_SECRET))
	mac.Write([]byte(jwt))

	signature := mac.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(signature)
}
