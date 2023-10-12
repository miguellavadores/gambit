package routers

import (
	"encoding/json"
	"strconv"

	"github.com/miguellavadores/gambit/bd"
	"github.com/miguellavadores/gambit/models"
)

func InsertProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el Nombre (Title) del producto"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}
