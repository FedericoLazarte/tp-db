package sqlbd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"tp-db/database"

	_ "github.com/lib/pq"
)

func CrearDB() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`drop database if exists cinemark with (force)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`create database cinemmark`)
	if err != nil {
		log.Fatal(err)
	}
}

func ConectarBD() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=cinemark sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CrearTablas() {
	db := ConectarBD()
	defer db.Close()
	ejecutarSql(db, "database/archivosSQL/crear_tablas.sql")
}

func CrearPksFks() {
	db := ConectarBD()
	defer db.Close()
	ejecutarSql(db, "database/archivosSQL/agregar_pks_fks.sql")
}

func EliminarPksFks() {
	db := ConectarBD()
	defer db.Close()
	ejecutarSql(db, "database/archivosSQL/elimiar_pks_fks.sql")
}

func CargarDatos() {
	db := ConectarBD()
	defer db.Close()

	clientes := leerJson[database.Cliente]("database/json/clientes.json")
	peliculas := leerJson[database.Pelicula]("database/json/peliculas.json")
	salas := leerJson[database.SalaCine]("database/json/salas_de_cine.json")
	datosPrueba := leerJson[database.DatosDePrueba]("database/json/datos_de_prueba.json")

	for _, c := range clientes {
		_, err := db.Exec(`insert into cliente values ($1, $2, $3, $4, $5, $6, $7)`,
			c.IDCliente, c.Nombre, c.Apellido, c.Dni, c.FechaNacimiento, c.Telefono, c.Email)
		if err != nil {
			fmt.Println("Error insertando cliente:", err)
		}
	}

	for _, p := range peliculas {
		_, err := db.Exec(`insert into pelicula values ($1, $2, $3, $4, $5, $6)`,
			p.IDPelicula, p.Titutlo, p.Duracion, p.Director, p.Origen, p.Formato)
		if err != nil {
			fmt.Println("Error insertando película:", err)
		}
	}

	for _, s := range salas {
		capacidad := s.NroFilas * s.NroButacasPorFila
		_, err := db.Exec(`insert into sala_cine values ($1, $2, $3, $4, $5, $6)`,
			s.IDSala, s.Nombre, s.Formato, s.NroFilas, s.NroButacasPorFila, capacidad)
		if err != nil {
			fmt.Println("Error insertando sala:", err)
		}
	}

	for _, d := range datosPrueba {
		_, err := db.Exec(`insert into datos_de_prueba values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			d.IDOrden, d.Operacion, d.IDSala, d.FInicioSesion, d.IDPelicula, d.IDFuncion, d.NroFila, d.NroButaca, d.IDCliente)
		if err != nil {
			fmt.Println("Error insertando datos de prueba:", err)
		}
	}
}

func CrearSpTriggers() {
	db := ConectarBD()
	defer db.Close()
	ejecutarSql(db, "database/archivosSQL/reservar_butacas_sp.sql")
	ejecutarSql(db, "database/archivosSQL/compra_de_butaca_sp.sql")
	ejecutarSql(db, "database/archivosSQL/apertura_funcion_sp.sql")
	ejecutarSql(db, "database/archivosSql/enviar_email.sql")
}

func IniciarPruebas() {
	db := ConectarBD()
	defer db.Close()
	ejecutarSql(db, "database/archivosSQL/iniciar_pruebas.sql")
}

func leerJson[T any](archivo string) []T {
	jsonFile, err := os.Open(archivo)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo %s: %v", archivo, err)
	}
	defer jsonFile.Close()

	datosJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("No se pudo leer el archivo %s: %v", archivo, err)
	}

	var datos []T
	err = json.Unmarshal(datosJSON, &datos)
	if err != nil {
		log.Fatalf("No se pudo convertir el JSON %s: %v", archivo, err)
	}
	return datos
}

func ejecutarSql(db *sql.DB, rutaArchivo string) {
	archivo, err := ioutil.ReadFile(rutaArchivo)
	if err != nil {
		log.Fatal(err)
	}
	sentenciaSql := string(archivo)
	_, err = db.Exec(sentenciaSql)
	if err != nil {
		log.Fatal(err)
	}
}
