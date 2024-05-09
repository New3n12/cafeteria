package main

import (
	"net/http"
	"pro2/rutas"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", rutas.Index)
	mux.HandleFunc("/admin", rutas.Admin)
	mux.HandleFunc("/addProducto", rutas.AddProducto)
	mux.HandleFunc("/orden", rutas.AddPedido)
	mux.HandleFunc("/delete", rutas.DeletePedido)
	mux.HandleFunc("/update", rutas.UpdatePedido)
	mux.HandleFunc("/deleteProducto", rutas.DeleteProducto)
	mux.HandleFunc("/returnProducto", rutas.ReturnProducto)
	mux.HandleFunc("/editarProducto", rutas.ActualizarProducto)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//archivos estáticos hacia mux
	s := http.StripPrefix("/vista/", http.FileServer(http.Dir("./vista/")))
	mux.PathPrefix("/vista/").Handler(s)

	//mux.NotFoundHandler = mux.NewRoute().HandlerFunc(rutas.Pagina404).GetHandler()

	server.ListenAndServe()
}
