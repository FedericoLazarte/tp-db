package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"tp-db/database/nosql"
	"tp-db/database/sqlbd"
)

func main() {
	runMenu()
}

func runMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	var opcion string

	for opcion != "0" {
		clearScreen()
		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║        MENÚ PRINCIPAL              ║")
		fmt.Println("╠════════════════════════════════════╣")
		fmt.Println("║ 1. Crear Base de Datos             ║")
		fmt.Println("║ 2. Crear Tablas                    ║")
		fmt.Println("║ 3. Agregar PKs y FKs               ║")
		fmt.Println("║ 4. Eliminar PKs y FKs              ║")
		fmt.Println("║ 5. Cargar Datos                    ║")
		fmt.Println("║ 6. Crear Stored Procedures/Triggers║")
		fmt.Println("║ 7. Iniciar Pruebas                 ║")
		fmt.Println("║ 8. Cargar datos en BoltDB          ║")
		fmt.Println("║ 0. Salir                           ║")
		fmt.Println("╚════════════════════════════════════╝")
		fmt.Print("Seleccione una opción ➤ ")

		scanner.Scan()
		opcion = scanner.Text()

		switch opcion {
		case "1":
			fmt.Println("\n🛠️  Creando Base de Datos...")
			sqlbd.CrearDB()
		case "2":
			fmt.Println("\n🧱 Agregando Tablas...")
			sqlbd.CrearTablas()
		case "3":
			fmt.Println("\n🔐 Agregando PKs y FKs...")
			sqlbd.CrearPksFks()
		case "4":
			fmt.Println("\n🧹 Eliminando PKs y FKs...")
			sqlbd.EliminarPksFks()
		case "5":
			fmt.Println("\n📦 Cargando Datos...")
			sqlbd.CargarDatos()
		case "6":
			fmt.Println("\n⚙️  Creando Stored Procedures y Triggers...")
			sqlbd.CrearSpTriggers()
		case "7":
			fmt.Println("\n🧪 Iniciando Pruebas...")
			sqlbd.IniciarPruebas()
		case "8":
			fmt.Println("\n🗃️  Cargando datos en BoltDB...")
			nosql.Start()
		case "0":
			fmt.Println("\n👋 Saliendo... ¡Hasta luego!")
			return
		default:
			fmt.Println("\n❌ Opción no válida. Intente de nuevo.")
		}

		fmt.Print("\nPresione ENTER para continuar...")
		scanner.Scan()
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Linux, macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
