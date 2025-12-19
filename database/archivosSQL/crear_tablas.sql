create table cliente (
  id_cliente int,
  nombre text,
  apellido text,
  dni int,
  fecha_nacimiento date,
  telefono char(12),
  email text
);

create table sala_cine (
  id_sala int,
  nombre text,
  formato char(10),
  nro_filas int,
  nro_butacas_por_fila int,
  capacidad_total int
);

create table pelicula (
  id_pelicula int,
  titulo text,
  duracion interval,
  director varchar(40),
  origen varchar(60),
  formato char(10)
);

create table funcion (
  id_funcion int,
  id_sala int,
  fecha_inicio date,
  hora_inicio time,
  fecha_fin date,
  hora_fin time,
  id_pelicula int,
  butacas_disponibles int
);

create table butaca_por_funcion (
  id_funcion int,
  nro_fila int,
  nro_butaca int,
  id_cliente int,
  estado char(15)
);

create table error (
  id_error int,
  operacion char(20),
  id_sala int,
  f_inicio_funcion timestamp,
  id_pelicula int,
  id_funcion int,
  nro_fila int,
  nro_butaca int,
  id_cliente int,
  f_error timestamp,
  motivo varchar(80)
);

create table envio_email (
  id_email int,
  f_generacion timestamp,
  email_cliente text,
  asunto text,
  cuerpo text,
  f_envio timestamp,
  estado char(10)
);

-- Tabla para probar el sistema
create table datos_de_prueba (
  id_orden int,
  operacion char(20),
  id_sala int,
  f_inicio_funcion timestamp,
  id_pelicula int,
  id_funcion int,
  nro_fila int,
  nro_butaca int,
  id_cliente int
);
