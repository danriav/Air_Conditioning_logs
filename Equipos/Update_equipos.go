package equipos

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	conn "aires_acondicionados/Connection"
)


func Editar(w http.ResponseWriter, r *http.Request){
	idEquipo:= r.URL.Query().Get("id")

	ConexiónEstablecida := conn.DBConnection()
	registro, err := ConexiónEstablecida.Query("SELECT * FROM registros WHERE id=?", idEquipo)


	if err != nil{
		panic(err.Error())
	}

	equipos := Equipos{}

	for registro.Next(){
		var id, capacidad_BTU int
		var aula, edificio, equipo, marca, modelo, tipo, refrigerante, voltaje, ubicación string
		err = registro.Scan(&id,&edificio,&aula,&equipo,&marca,&modelo,&tipo,&capacidad_BTU,&refrigerante,&voltaje,&ubicación)
		if err != nil{
			panic(err.Error())
		}
		equipos.Id = id
		equipos.Edificio = edificio
		equipos.Aula = aula
		equipos.Equipo = equipo
		equipos.Marca = marca
		equipos.Modelo = modelo
		equipos.Tipo = tipo
		equipos.Capacidad_BTU = capacidad_BTU
		equipos.Refrigerante = refrigerante
		equipos.Voltaje = voltaje
		equipos.Ubicación = ubicación
}

	//Colocar el mismo nombre que en templates
	templates.ExecuteTemplate(w, "edit",equipos)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		id := r.FormValue("id")
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
		modificarRegistros, err := ConexiónEstablecida.Prepare("UPDATE registros SET Edificio=?, Aula=?, Equipo=?, Marca=?, Modelo=?, Tipo=?, Capacidad_BTU=?, Refrigerante=?, Voltaje=?, Ubicación=? WHERE id=?")

	if err != nil{
		panic(err.Error())
	}
		modificarRegistros.Exec(edificio, aula, equipo, marca, modelo, tipo, capacidad_BTU, refrigerante, voltaje, ubicación, id)

		http.Redirect(w,r,"/",http.StatusMovedPermanently)
	}
}