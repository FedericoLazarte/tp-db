package nosql

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"tp-db/database"
	"tp-db/database/sqlbd"

	bolt "go.etcd.io/bbolt"
)

func Start() {
	sqlDB := sqlbd.ConectarBD()
	defer sqlDB.Close()
	db, err := openConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// migracion sql a nosql
	clientesSlice := sqlbd.GetClientesData(sqlDB)
	peliculasSlice := sqlbd.GetPeliculasData(sqlDB)
	salasCineSlice := sqlbd.GetSalasCineData(sqlDB)
	funcionesSlice := sqlbd.GetFuncionesData(sqlDB)
	butacasCompradasSlice := sqlbd.GetButacasReservadasData(sqlDB)

	loadClients(db, clientesSlice)
	loadPeliculas(db, peliculasSlice)
	loadSalasDeCine(db, salasCineSlice)
	loadFunciones(db, funcionesSlice)
	loadButacasCompradas(db, butacasCompradasSlice)

	readClients(db)
	readPeliculas(db)
	readSalasCine(db)
	readFunciones(db)
	readButacasCompradas(db)
}

func viewData(db *bolt.DB, bucketName string) error {
	// abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		fmt.Printf("\n--- Contenido del bucket %s---\n", bucketName)
		if err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("key: %s, value: %s.\n", k, v)
			return nil
		}); err != nil {
			return err
		}
		return nil
	})

	return err
}

func openConnection() (*bolt.DB, error) {
	// path hardcodeado, se puede pasar como argumento si se quiere
	// crea el archivo .db aunque no exista
	// usen paths relativos para que funcione siempre, independientemente del nombre del directorio del repo
	// tipo: "./database/nosql/cineNosql.db"
	db, err := bolt.Open("./database/nosql/cineNosql.db", 0o600, nil)
	if err != nil {
		return db, err
	}

	fmt.Println("base de datos creada con exito")

	return db, nil
}

func createUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) (*bolt.DB, error) {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return db, err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, val)
	if err != nil {
		return db, err
	}

	// cierra transacción
	if err := tx.Commit(); err != nil {
		return db, err
	}

	return db, nil
}

func loadClients(db *bolt.DB, dataSlice []database.Cliente) {
	for _, el := range dataSlice {
		key := []byte(strconv.Itoa(el.IDCliente))
		value, err := json.Marshal(el)
		if err != nil {
			log.Fatalf("error!")
		}
		_, err = createUpdate(db, "Clientes", key, value)
		if err != nil {
			log.Fatalf("error!")
		}
	}
}

func loadPeliculas(db *bolt.DB, dataSlice []database.Pelicula) {
	for _, el := range dataSlice {
		key := []byte(strconv.Itoa(el.IDPelicula))
		value, err := json.Marshal(el)
		if err != nil {
			log.Fatalf("error!")
		}
		_, err = createUpdate(db, "Peliculas", key, value)
		if err != nil {
			log.Fatalf("error!")
		}
	}
}

func loadSalasDeCine(db *bolt.DB, dataSlice []database.SalaCine) {
	for _, el := range dataSlice {
		key := []byte(strconv.Itoa(el.IDSala))
		value, err := json.Marshal(el)
		if err != nil {
			log.Fatalf("error!")
		}
		_, err = createUpdate(db, "SalasDeCine", key, value)
		if err != nil {
			log.Fatalf("error!")
		}
	}
}

func loadFunciones(db *bolt.DB, dataSlice []database.Funcion) {
	for _, el := range dataSlice {
		key := []byte(strconv.Itoa(el.IDFuncion))
		value, err := json.Marshal(el)
		if err != nil {
			log.Fatalf("error!")
		}
		_, err = createUpdate(db, "Funciones", key, value)
		if err != nil {
			log.Fatalf("error!")
		}
	}
}

func loadButacasCompradas(db *bolt.DB, dataSlice []database.ButacaPorFuncion) {
	for _, el := range dataSlice {
		key := []byte(strconv.Itoa(el.IDFuncion))
		value, err := json.Marshal(el)
		if err != nil {
			log.Fatalf("error!")
		}
		_, err = createUpdate(db, "ButacasCompradas", key, value)
		if err != nil {
			log.Fatalf("error!")
		}
	}
}

func readClients(db *bolt.DB) {
	err := viewData(db, "Clientes")
	if err != nil {
		log.Fatalf("error!")
	}
}

func readPeliculas(db *bolt.DB) {
	err := viewData(db, "Peliculas")
	if err != nil {
		log.Fatalf("error!")
	}
}

func readSalasCine(db *bolt.DB) {
	err := viewData(db, "SalasDeCine")
	if err != nil {
		log.Fatalf("error!")
	}
}

func readButacasCompradas(db *bolt.DB) {
	err := viewData(db, "ButacasCompradas")
	if err != nil {
		log.Fatalf("error!")
	}
}

func readFunciones(db *bolt.DB) {
	err := viewData(db, "Funciones")
	if err != nil {
		log.Fatalf("error!")
	}
}
