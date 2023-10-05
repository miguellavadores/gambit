package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub        string
	Event_Id   string
	Token_user string
	Scope      string
	Auth_time  int
	Iss        string
	Exp        int
	Iat        int
	Client_id  string
	Username   string
}

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("Token mal formado ")
		return false, nil, "El token no es válido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error al decodificar el token ", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Error al decodificar laestructura JSON ", err.Error())
		return false, err, err.Error()

	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)
	if tm.Before(ahora) {
		fmt.Println("Fecha expiración del token " + tm.String())
		fmt.Println("El token ha expirado !")
		return false, err, "Token Expirado !!"
	}

	return true, nil, string(tkj.Username)
}
