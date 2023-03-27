package equipos

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/db"
)

func Borrar(w http.ResponseWriter, r *http.Request){
	idEquipo:= r.URL.Query().Get("id")

	ConexiónEstablecida := conn.DBConnection()
		borrarRegistros, err := ConexiónEstablecida.Prepare("DELETE FROM registros WHERE id=?")

	if err != nil{
		panic(err.Error())
	}
		borrarRegistros.Exec(idEquipo)

		http.Redirect(w,r,"/",http.StatusMovedPermanently)

}