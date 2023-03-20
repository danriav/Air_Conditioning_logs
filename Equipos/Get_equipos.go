package equipos

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	conn "aires_acondicionados/Connection"
)



var templates = template.Must(template.ParseGlob("templates/*"))

type Equipos struct{
	Id int
	Edificio string
	Aula string
	Equipo string
	Marca string
	Modelo string
	Tipo string
	Capacidad_BTU int
	Refrigerante string
	Voltaje string
	Ubicación string

}


func ObtenerRegistros(w http.ResponseWriter, r *http.Request){

	ConexiónEstablecida := conn.DBConnection()
	registros, err := ConexiónEstablecida.Query("SELECT * FROM registros")

	if err != nil{
		panic(err.Error())
	}

	equipos := Equipos{}
	arrayEquipos := []Equipos{}

	for registros.Next(){
		var id, capacidad_BTU int
		var aula, edificio, equipo, marca, modelo, tipo, refrigerante, voltaje, ubicación string
		err = registros.Scan(&id,&edificio,&aula,&equipo,&marca,&modelo,&tipo,&capacidad_BTU,&refrigerante,&voltaje,&ubicación)
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

		arrayEquipos = append(arrayEquipos, equipos)
	}

	templates.ExecuteTemplate(w, "start", arrayEquipos)
}

func Registrar(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "registrar", nil)
}






