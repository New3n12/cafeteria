package controlador

import (
	"net/http"
	"pro2/model"

	"github.com/gorilla/sessions"
)

// inicializa la session
var Store = sessions.NewCookieStore([]byte("session-cafeteria"))

func CargarMenu() []model.Menu {
	menu := []model.Menu{}
	if model.Conectar() {
		menu = model.AllMenu(model.Con)
		model.Close()
		return menu
	}
	return menu
}
func CargarPedido() []model.Pedido {
	pedido := []model.Pedido{}
	if model.Conectar() {
		pedido = model.AllPedido(model.Con)
		model.Close()
		return pedido
	}
	return pedido
}

func AddProducto(m model.Menu) int {
	if model.Conectar() {
		add := model.InsertMenu(model.Con, m)
		if add != 0 {
			return add
		}
		model.Close()
	}
	return 0
}

func AddPedido(p model.Pedido) bool {
	if model.Conectar() {
		add := model.InsertPedido(model.Con, p)
		if add {
			return true
		}
		model.Close()
	}
	return false
}

func DeletePedido(id int) bool {
	if model.Conectar() {
		del := model.DeletePedido(model.Con, id)
		if del {
			return true
		}
		model.Close()
	}
	return false
}

func UpdatePedido(id int) bool {
	if model.Conectar() {
		upd := model.UpdatePedido(model.Con, id)
		if upd {
			return true
		}
		model.Close()
	}
	return false
}

func DeleteProducto(id int) bool {
	if model.Conectar() {
		del := model.DeleteProducto(model.Con, id)
		if del {
			return true
		}
		model.Close()
	}
	return false
}

func UpdateProducto(m model.Menu) bool {
	if model.Conectar() {
		upd := model.UpdateProducto(model.Con, m)
		if upd {
			return true
		}
		model.Close()
	}
	return false
}

func ProductoById(id int) model.Menu {
	menu := model.Menu{}
	if model.Conectar() {
		menu = model.ProductoById(model.Con, id)
		model.Close()
		return menu
	}
	return menu
}

// Se crea un flash para manejo de errores solo cadenas de texto
func CrearMensajesFlash(response http.ResponseWriter, request *http.Request, msg string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(msg, "Error")
	session.Save(request, response)
}

// Se recupera el mensaje flash para manejo de errores
func RetornarMensajesFlash(response http.ResponseWriter, request *http.Request) string {
	session, _ := Store.Get(request, "flash-session")

	fm := session.Flashes("Error")
	session.Save(request, response)
	css_sesion := ""
	if len(fm) == 0 {
		css_sesion = ""
	} else {
		css_sesion = fm[0].(string)
	}

	return css_sesion
}
