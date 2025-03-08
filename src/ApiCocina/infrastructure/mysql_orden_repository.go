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
	_, err := mysql.conn.Exec(
		"INSERT INTO Ordenes (mesa_id, estado, fecha_creacion, detalles) VALUES (?, ?, ?, ?)",
		orden.MesaID, orden.Estado, orden.FechaCreacion, orden.Detalles,
	)
	return err
}

func (mysql *MySQLOrdenRepository) Update(id int, orden entities.Orden) error {
	_, err := mysql.conn.Exec(
		"UPDATE Ordenes SET mesa_id = ?, estado = ?, detalles = ? WHERE id = ?",
		orden.MesaID, orden.Estado, orden.Detalles, id,
	)
	return err
}

func (mysql *MySQLOrdenRepository) GetAll() ([]entities.Orden, error) {
	rows, err := mysql.conn.Query("SELECT id, mesa_id, estado, fecha_creacion, detalles FROM Ordenes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordenes []entities.Orden
	for rows.Next() {
		var orden entities.Orden
		var fechaCreacion string
		err := rows.Scan(&orden.ID, &orden.MesaID, &orden.Estado, &fechaCreacion, &orden.Detalles)
		if err != nil {
			return nil, err
		}
		orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)
		ordenes = append(ordenes, orden)
	}
	return ordenes, nil
}

func (mysql *MySQLOrdenRepository) FindByID(id int) (entities.Orden, error) {
	var orden entities.Orden
	var fechaCreacion string
	row := mysql.conn.QueryRow("SELECT id, mesa_id, estado, fecha_creacion, detalles FROM Ordenes WHERE id = ?", id)
	err := row.Scan(&orden.ID, &orden.MesaID, &orden.Estado, &fechaCreacion, &orden.Detalles)
	if err != nil {
		return entities.Orden{}, err
	}
	orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)
	return orden, nil
}
