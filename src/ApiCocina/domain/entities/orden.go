package entities

import (
	"time"
)

type Orden struct {
	ID            int       `json:"id"`
	MesaID        int       `json:"mesa_id"`
	Estado        string    `json:"estado"`
	FechaCreacion time.Time `json:"fecha_creacion"`
	Detalles      string    `json:"detalles"`
}

func NewOrden(id, mesaID int, estado string, fechaCreacion time.Time, detalles string) *Orden {
	return &Orden{
		ID:            id,
		MesaID:        mesaID,
		Estado:        estado,
		FechaCreacion: fechaCreacion,
		Detalles:      detalles,
	}
}

// Getters y Setters
