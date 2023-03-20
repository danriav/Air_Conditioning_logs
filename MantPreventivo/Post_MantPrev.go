package mantpreventivo

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/Connection"
)

func Insertarmantprev(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
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
		insertarRegistros, err := ConexiónEstablecida.Prepare("insert into mantenimiento_preventivo(Fecha, Edificio, Aula, Equipo, Marca, Modelo, Tipo, Voltaje, Amperaje, Encargado_del_mantenimiento, Observaciones) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	
		if err != nil{
			panic(err.Error())
		}
	
		insertarRegistros.Exec(fecha, edificio, aula, equipo, marca, modelo, tipo, voltaje, amperaje, encargado_del_mantenimiento, observaciones)
	
		http.Redirect(w,r,"/mantpreventivo", http.StatusMovedPermanently)
		}
	}
	