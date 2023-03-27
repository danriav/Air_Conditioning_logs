package mantcorrectivo
import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/db"
)


var templates = template.Must(template.ParseGlob("templates/*"))

type Correctivo struct{
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
	Trabajo_realizado string
}

func MantCorrectivo(w http.ResponseWriter, r *http.Request){

	ConexiónEstablecida := conn.DBConnection()

	registros, err := ConexiónEstablecida.Query("SELECT * FROM mantenimiento_correctivo")

	if err != nil{
		panic(err.Error())
	}
	correctivo := Correctivo{}
	arrayCorrectivo := []Correctivo{}

	for registros.Next(){
		var id, voltaje, amperaje int
		var fecha, edificio, aula, equipo, marca, modelo, tipo, encargado_del_mantenimiento, trabajo_realizado string
		err = registros.Scan(&id, &fecha, &edificio, &aula, &equipo, &marca, &modelo, &tipo, &voltaje, &amperaje, &encargado_del_mantenimiento, &trabajo_realizado)
		if err != nil{
			panic(err.Error())
		}

		correctivo.Id = id
		correctivo.Fecha = fecha
		correctivo.Edificio = edificio
		correctivo.Aula = aula
		correctivo.Equipo = equipo
		correctivo.Marca = marca
		correctivo.Modelo = modelo
		correctivo.Tipo = modelo
		correctivo.Voltaje = voltaje
		correctivo.Amperaje = amperaje
		correctivo.Encargado_del_mantenimiento = encargado_del_mantenimiento
		correctivo.Trabajo_realizado = trabajo_realizado

		arrayCorrectivo = append(arrayCorrectivo, correctivo)
	}

	templates.ExecuteTemplate(w, "mantcorrectivo", arrayCorrectivo)
}






