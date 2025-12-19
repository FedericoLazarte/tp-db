package database

type Cliente struct {
	IDCliente       int    `json:"id_cliente"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Dni             int    `json:"dni"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Telefono        string `json:"telefono"`
	Email           string `json:"email"`
}

type SalaCine struct {
	IDSala            int    `json:"id_sala"`
	Nombre            string `json:"nombre"`
	Formato           string `json:"formato"`
	NroFilas          int    `json:"nro_filas"`
	NroButacasPorFila int    `json:"nro_butacas_por_fila"`
	CapacidadTotal    int    `json:"capacidad_total"`
}

type Pelicula struct {
	IDPelicula int    `json:"id_pelicula"`
	Titutlo    string `json:"titulo"`
	Duracion   string `json:"duracion"`
	Director   string `json:"director"`
	Origen     string `json:"origen"`
	Formato    string `json:"formato"`
}

type Funcion struct {
	IDFuncion          int    `json:"id_funcion"`
	IDSala             int    `json:"id_sala"`
	FechaInicio        string `json:"fecha_inicio"`
	HoraInicio         string `json:"hora_inicio"`
	FechaFin           string `json:"fecha_fin"`
	HoraFin            string `json:"hora_fin"`
	IDPelicula         int    `json:"id_pelicula"`
	ButacasDisponibles int    `json:"butacas_disponibles"`
}

type ButacaPorFuncion struct {
	IDFuncion int    `json:"id_funcion"`
	NroFila   int    `json:"nro_fila"`
	NroButaca int    `json:"nro_butaca"`
	IDCliente int    `json:"id_cliente"`
	Estado    string `json:"estado"`
}

type Error struct {
	IDError        int    `json:"id_error"`
	Operacion      string `json:"operacion"`
	IDSala         int    `json:"id_sala"`
	FInicioFuncion string `json:"f_inicio_funcion"`
	IDPelicula     int    `json:"id_pelicula"`
	IDFuncion      int    `json:"id_funcion"`
	NroFila        int    `json:"nro_fila"`
	NroButaca      int    `json:"nro_butaca"`
	IDCliente      int    `json:"id_cliente"`
	FError         string `json:"f_error"`
	Motivo         string `json:"motivo"`
}

type EnvioEmail struct {
	IDEmail      int    `json:"id_email"`
	FGeneracion  string `json:"f_generacion"`
	EmailCliente string `json:"email_cliente"`
	Asunto       string `json:"asunto"`
	Cuerpo       string `json:"cuerpo"`
	FEnvio       string `json:"f_envio"`
	Estado       string `json:"estado"`
}

// Tabla que no forma parte del modelo de datos
// Prueba la funcionalidad del sistema

type DatosDePrueba struct {
	IDOrden       int     `json:"id_orden"`
	Operacion     string  `json:"operacion"`
	IDSala        int     `json:"id_sala"`
	FInicioSesion *string `json:"f_inicio"`
	IDPelicula    int     `json:"id_pelicula"`
	IDFuncion     int     `json:"id_funcion"`
	NroFila       int     `json:"nro_fila"`
	NroButaca     int     `json:"nro_butaca"`
	IDCliente     int     `json:"id_cliente"`
}
