// orden.go
package entities

import (
	"time"
)

type DetalleOrden struct {
	ID         int `json:"id"`
	OrdenID    int `json:"orden_id"`
	PlatilloID int `json:"platillo_id"`
	Cantidad   int `json:"cantidad"`
}

type Orden struct {
	ID            int            `json:"id"`
	MesaID        int            `json:"mesa_id"`
	Estado        string         `json:"estado"`
	FechaCreacion time.Time      `json:"fecha_creacion"`
	Detalles      []DetalleOrden `json:"detalles"` // Cambia a un slice de DetalleOrden
}

func NewOrden(id, mesaID int, estado string, fechaCreacion time.Time, detalles []DetalleOrden) *Orden {
	// Validar la fecha de creación
	if fechaCreacion.IsZero() || fechaCreacion.Year() < 1000 {
		// Si la fecha no es válida, usar la fecha actual
		fechaCreacion = time.Now()
	}

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

func (o *Orden) GetDetalles() []DetalleOrden {
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
	// Validar la fecha de creación
	if fechaCreacion.IsZero() || fechaCreacion.Year() < 1000 {
		// Si la fecha no es válida, usar la fecha actual
		fechaCreacion = time.Now()
	}
	o.FechaCreacion = fechaCreacion
}

func (o *Orden) SetDetalles(detalles []DetalleOrden) {
	o.Detalles = detalles
}
