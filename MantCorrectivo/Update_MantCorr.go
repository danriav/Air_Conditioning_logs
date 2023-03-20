package mantcorrectivo

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/Connection"
)

func Editarmantcorr(w http.ResponseWriter, r *http.Request){
	idMantprev := r.URL.Query().Get("id")

	ConexiónEstablecida := conn.DBConnection()
	editarRegistros, err:= ConexiónEstablecida.Query("SELECT * FROM mantenimiento_correctivo WHERE id=?", idMantprev)

	if err != nil {
		panic(err.Error())
	}

	correctivo := Correctivo{}

	for editarRegistros.Next(){
		var id, voltaje, amperaje int
		var fecha, edificio, aula, equipo, marca, modelo, tipo, encargado_del_mantenimiento, trabajo_realizado string
		err = editarRegistros.Scan(&id, &fecha, &edificio, &aula, &equipo, &marca, &modelo, &tipo, &voltaje, &amperaje, &encargado_del_mantenimiento, &trabajo_realizado)
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
	}

	templates.ExecuteTemplate(w, "editarmantcorr", correctivo)

}

func ActualizarMantcorr(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
		id := r.FormValue("id")
		fecha := r.FormValue("fecha")
		edificio := r.FormValue("edificio")
		aula := r.FormValue("aula")
		equipo := r.FormValue("equipo")
		marca := r.FormValue("marca")
		modelo := r.FormValue("modelo")
		tipo := r.FormValue("tipo")
		voltaje := r.FormValue("voltaje")
		amperaje := r.FormValue("amperaje")
		encargado_del_mantenimiento := r.FormValue("encargado_del_mantenimiento")
		trabajo_realizado := r.FormValue("trabajo_realizado")
	
		ConexiónEstablecida := conn.DBConnection()
		modificarRegistros, err := ConexiónEstablecida.Prepare("UPDATE mantenimiento_correctivo SET Fecha=?, Edificio=?, Aula=?, Equipo=?, Marca=?, Modelo=?, Tipo=?, Voltaje=?, Amperaje=?, Encargado_del_mantenimiento=?, Trabajo_realizado=? WHERE id=?")
	
		if err != nil{
			panic(err.Error())
		}
	
		modificarRegistros.Exec(fecha, edificio, aula, equipo, marca, modelo, tipo, voltaje, amperaje, encargado_del_mantenimiento, trabajo_realizado, id)
	
		http.Redirect(w,r,"/mantcorrectivo", http.StatusMovedPermanently)
		}
}