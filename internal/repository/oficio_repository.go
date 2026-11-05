package repository

import (
	"sistema-oficios/internal/database"
	"sistema-oficios/internal/models"
)

func GetAll() ([]models.Oficio, error) {
	rows, err := database.DB.Query(`
		SELECT 
			id, 
			folio, 
			asunto, 
			remitente, 
			fecha_recepcion, 
			estatus
		FROM Oficios
		`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var oficios []models.Oficio
	for rows.Next() {
		var oficio models.Oficio
		err := rows.Scan(&oficio.ID, &oficio.Folio, &oficio.Asunto, &oficio.Remitente, &oficio.FechaRecepcion, &oficio.Estatus)
		if err != nil {
			return nil, err
		}
		oficios = append(oficios, oficio)
	}

	return oficios, nil
}
