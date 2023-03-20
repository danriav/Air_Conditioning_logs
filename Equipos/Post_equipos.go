package equipos

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/Connection"
)

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		edificio := r.FormValue("edificio")
		aula := r.FormValue("aula")
		equipo := r.FormValue("equipo")
		marca := r.FormValue("marca")
		modelo := r.FormValue("modelo")
		tipo := r.FormValue("tipo")
		capacidad_BTU := r.FormValue("capacidad_BTU")
		refrigerante := r.FormValue("refrigerante")
		voltaje := r.FormValue("voltaje")
		ubicación := r.FormValue("ubicación")

		ConexiónEstablecida := conn.DBConnection()
		insertarRegistros, err := ConexiónEstablecida.Prepare("insert into registros(Edificio, Aula, Equipo, Marca, Modelo, Tipo, Capacidad_BTU, Refrigerante, Voltaje, Ubicación) VALUES(?,?,?,?,?,?,?,?,?,?)")

	if err != nil{
		panic(err.Error())
	}
		insertarRegistros.Exec(edificio, aula, equipo, marca, modelo, tipo, capacidad_BTU, refrigerante, voltaje, ubicación)

		http.Redirect(w,r,"/",http.StatusMovedPermanently)
	}
}