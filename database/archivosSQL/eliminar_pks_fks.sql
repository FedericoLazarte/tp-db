alter table error drop constraint if exists fk_error_sala;
alter table error drop constraint if exists fk_error_pelicula;
alter table error drop constraint if exists fk_error_funcion;
alter table error drop constraint if exists fk_error_butaca;
alter table error drop constraint if exists fk_error_cliente;

alter table butaca_por_funcion drop constraint if exists fk_butaca_funcion;
alter table butaca_por_funcion drop constraint if exists fk_butaca_cliente;

alter table funcion drop constraint if exists fk_funcion_sala;
alter table funcion drop constraint if exists fk_funcion_pelicula;


alter table cliente drop constraint if exists pk_cliente;
alter table sala_cine drop constraint if exists pk_sala_cine;
alter table pelicula drop constraint if exists pk_pelicula;
alter table funcion drop constraint if exists pk_funcion;
alter table butaca_por_funcion drop constraint if exists pk_butaca_por_funcion;
alter table error drop constraint if exists pk_error;
alter table envio_email drop constraint if exists pk_envio_email;
alter table datos_de_prueba drop constraint if exists pk_datos_de_prueba;

alter table funcion alter column id_funcion drop identity if exists;
alter table error alter column id_error drop identity if exists;
alter table envio_email alter column id_email drop identity if exists;

alter table funcion alter column id_funcion drop not null;
alter table error alter column id_error drop not null;
alter table envio_email alter column id_email drop not null;

alter table funcion alter column id_funcion type int;
alter table error alter column id_error type int;
alter table envio_email alter column id_email type int;

