package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	Equipos "aires_acondicionados/Equipos"
	MantCorr "aires_acondicionados/MantCorrectivo"
	MantPrev "aires_acondicionados/MantPreventivo"
)





func main() {
	http.HandleFunc("/", Equipos.ObtenerRegistros)
	http.HandleFunc("/registrar", Equipos.Registrar)
	http.HandleFunc("/insertar", Equipos.Insertar)
	http.HandleFunc("/borrar", Equipos.Borrar)
	http.HandleFunc("/editar", Equipos.Editar)
	http.HandleFunc("/actualizar", Equipos.Actualizar)

	//mantenimiento preventivo
	http.HandleFunc("/mantpreventivo", MantPrev.MantPreventivo)
	http.HandleFunc("/registrarmantprev", MantPrev.Registrarmantprev)
	http.HandleFunc("/insertarmantprev", MantPrev.Insertarmantprev)
	http.HandleFunc("/borrarmantprev", MantPrev.Borrarmantprev)
	http.HandleFunc("/editarmantprev", MantPrev.Editarmantprev)
	http.HandleFunc("/actualizarmantprev", MantPrev.ActualizarMantprev)

	//mantenimiento correctivo
	http.HandleFunc("/mantcorrectivo", MantCorr.MantCorrectivo)
	http.HandleFunc("/registrarmantcorr", MantCorr.Registrarmantcorr)
	http.HandleFunc("/insertarmantcorr", MantCorr.Insertarmantcorr)
	http.HandleFunc("/borrarmantcorr", MantCorr.Borrarmantcorr)
	http.HandleFunc("/editarmantcorr", MantCorr.Editarmantcorr)
	http.HandleFunc("/actualizarmantcorr", MantCorr.ActualizarMantcorr)

	log.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}