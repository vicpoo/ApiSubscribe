// mysql_orden_repository.go
package infrastructure

import (
	"database/sql"
	"time"

	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"
	"github.com/vicpoo/ApiSubscribe/src/core"
)

type MySQLOrdenRepository struct {
	conn *sql.DB
}

func NewMySQLOrdenRepository() domain.OrdenRepository {
	conn := core.GetDB()
	return &MySQLOrdenRepository{conn: conn}
}

func (mysql *MySQLOrdenRepository) Save(orden entities.Orden) error {
	// Iniciar una transacción
	tx, err := mysql.conn.Begin()
	if err != nil {
		return err
	}

	// Validar la fecha de creación
	fechaCreacion := orden.FechaCreacion
	if fechaCreacion.IsZero() || fechaCreacion.Year() < 1000 {
		// Si la fecha no es válida, usar la fecha actual
		fechaCreacion = time.Now()
	}

	// Insertar la orden
	result, err := tx.Exec(
		"INSERT INTO Ordenes (mesa_id, estado, fecha_creacion) VALUES (?, ?, ?)",
		orden.MesaID, orden.Estado, fechaCreacion,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Obtener el ID de la orden insertada
	ordenID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insertar los detalles de la orden
	for _, detalle := range orden.Detalles {
		_, err := tx.Exec(
			"INSERT INTO DetallesOrden (orden_id, platillo_id, cantidad) VALUES (?, ?, ?)",
			ordenID, detalle.PlatilloID, detalle.Cantidad,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Confirmar la transacción
	return tx.Commit()
}

func (mysql *MySQLOrdenRepository) Update(id int, orden entities.Orden) error {
	// Iniciar una transacción
	tx, err := mysql.conn.Begin()
	if err != nil {
		return err
	}

	// Validar la fecha de creación
	fechaCreacion := orden.FechaCreacion
	if fechaCreacion.IsZero() || fechaCreacion.Year() < 1000 {
		// Si la fecha no es válida, usar la fecha actual
		fechaCreacion = time.Now()
	}

	// Actualizar la orden
	_, err = tx.Exec(
		"UPDATE Ordenes SET mesa_id = ?, estado = ?, fecha_creacion = ? WHERE id = ?",
		orden.MesaID, orden.Estado, fechaCreacion, id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Eliminar los detalles antiguos de la orden
	_, err = tx.Exec("DELETE FROM DetallesOrden WHERE orden_id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insertar los nuevos detalles de la orden
	for _, detalle := range orden.Detalles {
		_, err := tx.Exec(
			"INSERT INTO DetallesOrden (orden_id, platillo_id, cantidad) VALUES (?, ?, ?)",
			id, detalle.PlatilloID, detalle.Cantidad,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Confirmar la transacción
	return tx.Commit()
}

func (mysql *MySQLOrdenRepository) GetAll() ([]entities.Orden, error) {
	rows, err := mysql.conn.Query("SELECT id, mesa_id, estado, fecha_creacion FROM Ordenes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordenes []entities.Orden
	for rows.Next() {
		var orden entities.Orden
		var fechaCreacion string
		err := rows.Scan(&orden.ID, &orden.MesaID, &orden.Estado, &fechaCreacion)
		if err != nil {
			return nil, err
		}
		orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)

		// Obtener los detalles de la orden
		detallesRows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE orden_id = ?", orden.ID)
		if err != nil {
			return nil, err
		}
		defer detallesRows.Close()

		for detallesRows.Next() {
			var detalle entities.DetalleOrden
			err := detallesRows.Scan(&detalle.ID, &detalle.OrdenID, &detalle.PlatilloID, &detalle.Cantidad)
			if err != nil {
				return nil, err
			}
			orden.Detalles = append(orden.Detalles, detalle)
		}

		ordenes = append(ordenes, orden)
	}
	return ordenes, nil
}

func (mysql *MySQLOrdenRepository) FindByID(id int) (entities.Orden, error) {
	var orden entities.Orden
	var fechaCreacion string
	row := mysql.conn.QueryRow("SELECT id, mesa_id, estado, fecha_creacion FROM Ordenes WHERE id = ?", id)
	err := row.Scan(&orden.ID, &orden.MesaID, &orden.Estado, &fechaCreacion)
	if err != nil {
		return entities.Orden{}, err
	}
	orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)

	// Obtener los detalles de la orden
	detallesRows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE orden_id = ?", id)
	if err != nil {
		return entities.Orden{}, err
	}
	defer detallesRows.Close()

	for detallesRows.Next() {
		var detalle entities.DetalleOrden
		err := detallesRows.Scan(&detalle.ID, &detalle.OrdenID, &detalle.PlatilloID, &detalle.Cantidad)
		if err != nil {
			return entities.Orden{}, err
		}
		orden.Detalles = append(orden.Detalles, detalle)
	}

	return orden, nil
}
