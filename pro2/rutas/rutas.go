package rutas

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	ctr "pro2/controlador"
	"pro2/model"
	"strconv"
	"strings"
)

type data struct {
	Menu    []model.Menu
	Pedido  []model.Pedido
	Mensaje string
}

func Index(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/index.html"))
	data := data{Menu: ctr.CargarMenu(), Pedido: ctr.CargarPedido()}
	template.Execute(w, data)
}
func Admin(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/admin.html"))
	data := data{Pedido: ctr.CargarPedido(), Menu: ctr.CargarMenu(), Mensaje: ctr.RetornarMensajesFlash(w, r)}
	template.Execute(w, data)
}
func AddProducto(w http.ResponseWriter, r *http.Request) {

	// Obtiene los campos del formulario
	nombreProducto := r.FormValue("nombre")
	descripcionProducto := r.FormValue("descri")
	precioProducto, _ := strconv.Atoi(r.FormValue("precio"))
	fmt.Println(nombreProducto)
	// Obtiene la imagen del formulario
	file, path, err := r.FormFile("imagen")
	if err != nil {
		fmt.Println("Error al obtener el archivo:", err)
		http.Error(w, "Error al obtener el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := strings.Split(path.Filename, ".")

	producto := model.Menu{
		Nombre:      nombreProducto,
		Descripcion: descripcionProducto,
		Img:         "." + ext[1],
		Precio:      precioProducto,
	}

	add := ctr.AddProducto(producto)

	if add != 0 {
		imageBytes, _ := ioutil.ReadAll(file)
		// Guarda la imagen en el servidor
		err = ioutil.WriteFile("vista/img-producto/"+strconv.Itoa(add)+"."+ext[1], imageBytes, 0666)
		if err != nil {
			fmt.Println("Error al guardar el archivo en el servidor:", err)
			http.Error(w, "Error al guardar el archivo en el servidor", http.StatusInternalServerError)
			return
		}
	}
	ctr.CrearMensajesFlash(w, r, "agregado")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id_producto"))

	// Obtiene los campos del formulario
	nombreProducto := r.FormValue("nombre")
	descripcionProducto := r.FormValue("descri")
	precioProducto, _ := strconv.Atoi(r.FormValue("precio"))
	fmt.Println(nombreProducto)
	// Obtiene la imagen del formulario
	file, path, err := r.FormFile("imagen")
	if err != nil {
		fmt.Println("Error al obtener el archivo:", err)
		http.Error(w, "Error al obtener el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := strings.Split(path.Filename, ".")

	producto := model.Menu{
		Id:          id,
		Nombre:      nombreProducto,
		Descripcion: descripcionProducto,
		Img:         strconv.Itoa(id) + "." + ext[1],
		Precio:      precioProducto,
	}

	add := ctr.UpdateProducto(producto)

	if add {
		imageBytes, _ := ioutil.ReadAll(file)
		// Guarda la imagen en el servidor
		err = ioutil.WriteFile("vista/img-producto/"+strconv.Itoa(id)+"."+ext[1], imageBytes, 0666)
		if err != nil {
			fmt.Println("Error al guardar el archivo en el servidor:", err)
			http.Error(w, "Error al guardar el archivo en el servidor", http.StatusInternalServerError)
			return
		}
	}
	ctr.CrearMensajesFlash(w, r, "actualizado")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func AddPedido(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	carrera := r.FormValue("carrera")
	//semestre := r.FormValue("semestre")
	cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))
	id, _ := strconv.Atoi(r.FormValue("id"))

	pedido := model.Pedido{
		Nombre:  nombre,
		Carrera: carrera,
		//Status: semestre,
		Cantidad: cantidad,
		Id:       id,
	}

	res := ctr.AddPedido(pedido)

	if res {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func DeletePedido(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	res := ctr.DeletePedido(id)

	if res {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func UpdatePedido(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	res := ctr.UpdatePedido(id)

	if res {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	res := ctr.DeleteProducto(id)

	if res {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func ReturnProducto(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	producto := ctr.ProductoById(id)

	// Convertir el producto a JSON
	jsonData, err := json.Marshal(producto)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Escribir el JSON en el cuerpo de la respuesta
	w.Write(jsonData)

}
