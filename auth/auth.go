package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type TokenJSON struct {
	Sub      string
	EventID  string `json:"event_id"`
	TokenUse string `json:"token_use"`
	Scope    string
	AuthTime int `json:"auth_time"`
	Iss      string
	Exp      int
	Iat      int
	ClientID string `json:"client_id"`
	Username string
}

func ValidoToken(token string) (bool, error, string) {
	userInfo, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		fmt.Println("No se puede decodificar la parte del token: ", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura JSON ", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha expiración token = " + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expirado !!"
	}

	return true, nil, tkj.Username
}
