package routers

import (
	"encoding/json"
	"strconv"
	
	//"github.com/aws/aws-lambda-go/events"
	"github.com/miguellavadores/gambit/bd"
	"github.com/miguellavadores/gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category
	
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos "+err.Error()
	}
	
	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (title) de la Categoría"
	}
	
	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la Categoría"
	}
	
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}
	
	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la Categoría"+ t.CategName + err2.Error()
	}
	
	return 200, "{ CategID: "+ strconv.Itoa(int(result)) + " }"
}