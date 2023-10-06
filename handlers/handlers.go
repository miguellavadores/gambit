package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/miguellavadores/gambit/auth"
	"github.com/miguellavadores/gambit/routers"
)

func Manejadores(path string, method string, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validoAuthorization(path, method, header)
	if !isOk {
		return statusCode, user
	}

	switch path[1:5] {
	case "user":
		return procesoUser(body, path, method, user, id, request)
	case "prod":
		return procesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return procesoStock(body, path, method, user, idn, request)
	case "addr":
		return procesoAddress(body, path, method, user, idn, request)
	case "cate":
		return procesoCategory(body, path, method, user, idn, request)
	case "orde":
		return procesoOrder(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid"
}

func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token Requerido"
	}

	todoOK, err, msg := auth.ValidoToken(token)
	if !todoOK {
		if err != nil {
			fmt.Println("Error en token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en token " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg

}

func procesoUser(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func procesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func procesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}

	return 400, "Method Invalid"
}

func procesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func procesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func procesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}
