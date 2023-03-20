package mantpreventivo

import (

	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/Connection"
)

func Editarmantprev(w http.ResponseWriter, r *http.Request){
	idMantprev := r.URL.Query().Get("id")

	ConexiónEstablecida := conn.DBConnection()
	editarRegistros, err:= ConexiónEstablecida.Query("SELECT * FROM mantenimiento_preventivo WHERE id=?", idMantprev)

	if err != nil {
		panic(err.Error())
	}

	preventivo := Preventivo{}

	for editarRegistros.Next(){
		var id, voltaje, amperaje int
		var fecha, edificio, aula, equipo, marca, modelo, tipo, encargado_del_mantenimiento, observaciones string
		err = editarRegistros.Scan(&id, &fecha, &edificio, &aula, &equipo, &marca, &modelo, &tipo, &voltaje, &amperaje, &encargado_del_mantenimiento, &observaciones)
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
	}

	templates.ExecuteTemplate(w, "editarmantprev", preventivo)

}

func ActualizarMantprev(w http.ResponseWriter, r *http.Request){
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
		observaciones := r.FormValue("observaciones")
	
		ConexiónEstablecida := conn.DBConnection()
		modificarRegistros, err := ConexiónEstablecida.Prepare("UPDATE mantenimiento_preventivo SET Fecha=?, Edificio=?, Aula=?, Equipo=?, Marca=?, Modelo=?, Tipo=?, Voltaje=?, Amperaje=?, Encargado_del_mantenimiento=?, Observaciones=? WHERE id=?")
	
		if err != nil{
			panic(err.Error())
		}
	
		modificarRegistros.Exec(fecha, edificio, aula, equipo, marca, modelo, tipo, voltaje, amperaje, encargado_del_mantenimiento, observaciones, id)
	
		http.Redirect(w,r,"/mantpreventivo", http.StatusMovedPermanently)
		}
}