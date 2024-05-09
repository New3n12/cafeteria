package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func AllMenu(db *sql.DB) []Menu {
	rows, err := db.Query("SELECT * FROM producto")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	menu := []Menu{}

	for rows.Next() {
		m := Menu{}
		err := rows.Scan(&m.Id, &m.Nombre, &m.Descripcion, &m.Img, &m.Precio)
		if err != nil {
			log.Fatal(err)
		}
		menu = append(menu, m)
	}
	return menu
}

func InsertMenu(db *sql.DB, m Menu) int {
	_, err := db.Exec("INSERT INTO producto (nombre, descripcion,img, precio) VALUES (?, ?, ?, ?)", m.Nombre, m.Descripcion, "", m.Precio)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	rows, err := db.Query("SELECT id FROM producto ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer rows.Close()
	var id int

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	msg := strconv.Itoa(id) + m.Img

	_, err2 := db.Exec("UPDATE producto SET img = ? WHERE id = ?", msg, id)

	if err2 != nil {
		fmt.Println(err2.Error())
		return 0
	}
	return id
}

func InsertPedido(db *sql.DB, p Pedido) bool {
	_, err := db.Exec("INSERT INTO `pedidos` ( `id_pedido`, `nombre`, `carrera`, `cantidad`,`status`) VALUES (?, ?, ?, ?, ?)", p.Id, p.Nombre, p.Carrera, p.Cantidad, "0")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func AllPedido(db *sql.DB) []Pedido {
	rows, err := db.Query("SELECT p.id,p.status,p.nombre,p.carrera,p.cantidad,pr.nombre FROM `pedidos` as p INNER JOIN producto as pr ON p.id_pedido = pr.id ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pedido := []Pedido{}

	for rows.Next() {
		p := Pedido{}
		err := rows.Scan(&p.Id, &p.Status, &p.Nombre, &p.Carrera, &p.Cantidad, &p.NombreProducto)
		if err != nil {
			log.Fatal(err)
		}
		pedido = append(pedido, p)
	}
	return pedido
}

func DeletePedido(db *sql.DB, id int) bool {
	_, err := db.Exec("DELETE FROM `pedidos` WHERE `id` = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func UpdatePedido(db *sql.DB, id int) bool {
	_, err := db.Exec("UPDATE `pedidos` SET `status` = 1 WHERE `id` = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func DeleteProducto(db *sql.DB, id int) bool {
	_, err := db.Exec("DELETE FROM `producto` WHERE `id` = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func UpdateProducto(db *sql.DB, m Menu) bool {
	_, err := db.Exec("UPDATE `producto` SET `nombre` = ?, `descripcion` = ?, `precio` = ?,  `img` = ? WHERE `id` = ?", m.Nombre, m.Descripcion, m.Precio, m.Img, m.Id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func ProductoById(db *sql.DB, id int) Menu {
	rows, err := db.Query("SELECT * FROM producto WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	m := Menu{}

	for rows.Next() {
		err := rows.Scan(&m.Id, &m.Nombre, &m.Descripcion, &m.Img, &m.Precio)
		if err != nil {
			log.Fatal(err)
		}
	}
	return m
}
