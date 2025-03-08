// orden.go
package entities

import (
	"time"
)

type Orden struct {
	ID            int       `json:"id"`
	MesaID        int       `json:"mesa_id"`
	Estado        string    `json:"estado"`
	FechaCreacion time.Time `json:"fecha_creacion"`
	Detalles      string    `json:"detalles"` // Añadir el campo Detalles
}

func NewOrden(id, mesaID int, estado string, fechaCreacion time.Time, detalles string) *Orden {
	return &Orden{
		ID:            id,
		MesaID:        mesaID,
		Estado:        estado,
		FechaCreacion: fechaCreacion,
		Detalles:      detalles, // Inicializar el campo Detalles
	}
}

// Getters
func (o *Orden) GetID() int {
	return o.ID
}

func (o *Orden) GetMesaID() int {
	return o.MesaID
}

func (o *Orden) GetEstado() string {
	return o.Estado
}

func (o *Orden) GetFechaCreacion() time.Time {
	return o.FechaCreacion
}

func (o *Orden) GetDetalles() string {
	return o.Detalles
}

// Setters
func (o *Orden) SetID(id int) {
	o.ID = id
}

func (o *Orden) SetMesaID(mesaID int) {
	o.MesaID = mesaID
}

func (o *Orden) SetEstado(estado string) {
	o.Estado = estado
}

func (o *Orden) SetFechaCreacion(fechaCreacion time.Time) {
	o.FechaCreacion = fechaCreacion
}

func (o *Orden) SetDetalles(detalles string) {
	o.Detalles = detalles
}
