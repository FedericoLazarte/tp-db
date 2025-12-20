package sqlbd

import (
	"database/sql"
	"log"

	"tp-db/database"

	_ "github.com/lib/pq"
)

func GetClientesData(db *sql.DB) []database.Cliente {
	var data []database.Cliente
	query := `select * from cliente`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var c database.Cliente
		if err := rows.Scan(&c.IDCliente, &c.Nombre, &c.Apellido, &c.Dni, &c.FechaNacimiento, &c.Telefono,
			&c.Email); err != nil {
			log.Fatal(err)
		}
		data = append(data, c)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func GetSalasCineData(db *sql.DB) []database.SalaCine {
	var data []database.SalaCine
	query := `select * from sala_cine`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s database.SalaCine
		if err := rows.Scan(&s.IDSala, &s.Nombre, &s.Formato, &s.NroFilas, &s.NroButacasPorFila, &s.CapacidadTotal); err != nil {
			log.Fatal(err)
		}
		data = append(data, s)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func GetPeliculasData(db *sql.DB) []database.Pelicula {
	var data []database.Pelicula
	query := `select * from pelicula`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var p database.Pelicula
		if err := rows.Scan(&p.IDPelicula, &p.Titutlo, &p.Duracion, &p.Director, &p.Origen, &p.Formato); err != nil {
			log.Fatal(err)
		}
		data = append(data, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func GetFuncionesData(db *sql.DB) []database.Funcion {
	var data []database.Funcion
	query := `select * from funcion`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var f database.Funcion
		if err := rows.Scan(&f.IDFuncion, &f.IDSala, &f.FechaInicio, &f.HoraInicio, &f.FechaFin,
			&f.HoraFin, &f.IDPelicula, &f.ButacasDisponibles); err != nil {
			log.Fatal(err)
		}
		data = append(data, f)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func GetButacasReservadasData(db *sql.DB) []database.ButacaPorFuncion {
	var data []database.ButacaPorFuncion
	query := `select * from butaca_por_funcion where estado = 'comprada'`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var b database.ButacaPorFuncion
		if err := rows.Scan(&b.IDFuncion, &b.NroFila, &b.NroButaca, &b.IDCliente, &b.Estado); err != nil {
			log.Fatal(err)
		}
		data = append(data, b)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
