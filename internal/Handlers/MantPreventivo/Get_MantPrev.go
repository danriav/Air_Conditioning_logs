package mantpreventivo

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/db"
)


type Preventivo struct{
	Id int
	Fecha string
	Edificio string
	Aula string
	Equipo string
	Marca string
	Modelo string
	Tipo string
	Voltaje int
	Amperaje int
	Encargado_del_mantenimiento string
	Observaciones string
}

var templates = template.Must(template.ParseGlob("templates/*"))

func MantPreventivo(w http.ResponseWriter, r *http.Request){

	ConexiónEstablecida := conn.DBConnection()

	registros, err := ConexiónEstablecida.Query("SELECT * FROM mantenimiento_preventivo")

	if err != nil{
		panic(err.Error())
	}
	preventivo := Preventivo{}
	arrayPreventivo := []Preventivo{}

	for registros.Next(){
		var id, voltaje, amperaje int
		var fecha, edificio, aula, equipo, marca, modelo, tipo, encargado_del_mantenimiento, observaciones string
		err = registros.Scan(&id, &fecha, &edificio, &aula, &equipo, &marca, &modelo, &tipo, &voltaje, &amperaje, &encargado_del_mantenimiento, &observaciones)
		if err != nil{
			panic(err.Error())
		}

		preventivo.Id = id
		preventivo.Fecha = fecha
		preventivo.Edificio = edificio
		preventivo.Aula = aula
		preventivo.Equipo = equipo
		preventivo.Marca = marca
		preventivo.Modelo = modelo
		preventivo.Tipo = modelo
		preventivo.Voltaje = voltaje
		preventivo.Amperaje = amperaje
		preventivo.Encargado_del_mantenimiento = encargado_del_mantenimiento
		preventivo.Observaciones = observaciones

		arrayPreventivo = append(arrayPreventivo, preventivo)
	}

	templates.ExecuteTemplate(w, "mantpreventivo", arrayPreventivo)
}

func Registrarmantprev(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "registrarmantprev", nil)
}


