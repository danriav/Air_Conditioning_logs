package mantpreventivo

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	conn "aires_acondicionados/db"
)

func Borrarmantprev(w http.ResponseWriter, r *http.Request){
	idMantprev:= r.URL.Query().Get("id")

	ConexiónEstablecida := conn.DBConnection()
		borrarRegistros, err := ConexiónEstablecida.Prepare("DELETE FROM mantenimiento_preventivo WHERE id=?")

	if err != nil{
		panic(err.Error())
	}
		borrarRegistros.Exec(idMantprev)

		http.Redirect(w,r,"/mantpreventivo",http.StatusMovedPermanently)
}