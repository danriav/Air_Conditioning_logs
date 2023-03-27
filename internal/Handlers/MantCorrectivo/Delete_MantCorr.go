package mantcorrectivo

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	conn "aires_acondicionados/db"
)


func Borrarmantcorr(w http.ResponseWriter, r *http.Request){
	idMantcorr:= r.URL.Query().Get("id")

	ConexiónEstablecida := conn.DBConnection()
		borrarRegistros, err := ConexiónEstablecida.Prepare("DELETE FROM mantenimiento_correctivo WHERE id=?")

	if err != nil{
		panic(err.Error())
	}
		borrarRegistros.Exec(idMantcorr)

		http.Redirect(w,r,"/mantcorrectivo",http.StatusMovedPermanently)
}