package mantcorrectivo

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	
	conn "aires_acondicionados/db"
)

func Insertarmantcorr(w http.ResponseWriter, r *http.Request){
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
		trabajo_realizado := r.FormValue("trabajo_realizado")
	
		ConexiónEstablecida := conn.DBConnection()
		insertarRegistros, err := ConexiónEstablecida.Prepare("insert into mantenimiento_correctivo(Fecha, Edificio, Aula, Equipo, Marca, Modelo, Tipo, Voltaje, Amperaje, Encargado_del_mantenimiento, Trabajo_realizado) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	
		if err != nil{
			panic(err.Error())
		}
	
		insertarRegistros.Exec(fecha, edificio, aula, equipo, marca, modelo, tipo, voltaje, amperaje, encargado_del_mantenimiento, trabajo_realizado)
	
		http.Redirect(w,r,"/mantcorrectivo", http.StatusMovedPermanently)
		}
	}

	func Registrarmantcorr(w http.ResponseWriter, r *http.Request){
		templates.ExecuteTemplate(w, "registrarmantcorr", nil)
	}
	