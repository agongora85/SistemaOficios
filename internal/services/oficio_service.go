package services

import (
	"sistema-oficios/internal/models"
	"sistema-oficios/internal/repository"
)

func ObtenerOficios() ([]models.Oficio, error) {
	return repository.GetAll()
}
