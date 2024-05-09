package model

type Menu struct {
	Id          int
	Nombre      string
	Descripcion string
	Img         string
	Precio      int
}
type Pedido struct {
	Nombre         string
	Carrera        string
	Status         string
	Cantidad       int
	Id             int
	NombreProducto string
}
