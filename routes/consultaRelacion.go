package routes

import (
	"api/bd"
	"api/models"
	"encoding/json"
	"net/http"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioId = IDUsuario
	relacion.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion
	status, err := bd.ConsultaRelacion(relacion)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
