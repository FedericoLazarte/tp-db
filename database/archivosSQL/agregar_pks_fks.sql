alter table funcion alter column id_funcion set not null;
alter table error alter column id_error set not null;
alter table envio_email alter column id_email set not null;

-- auto incremento
alter table funcion alter column id_funcion add generated always as identity;
alter table error alter column id_error add generated always as identity;
alter table envio_email alter column id_email add generated always as identity;

-- pks
alter table cliente add constraint pk_cliente primary key (id_cliente);
alter table sala_cine add constraint pk_sala_cine primary key (id_sala);
alter table pelicula add constraint pk_pelicula primary key (id_pelicula);
alter table funcion add constraint pk_funcion primary key (id_funcion);
alter table butaca_por_funcion add constraint pk_butaca_por_funcion primary key (id_funcion, nro_fila, nro_butaca);
alter table error add constraint pk_error primary key (id_error);
alter table envio_email add constraint pk_envio_email primary key (id_email);
alter table datos_de_prueba add constraint pk_datos_de_prueba primary key (id_orden);

-- fks
alter table funcion add constraint fk_funcion_sala foreing key (id_sala) references sala_cine(id_sala);
alter table funcion add constraint fk_funcion_pelicula foreing key (id_pelicula) references pelicula(id_pelicula);

alter table butaca_por_funcion add constraint fk_butaca_funcion foreing key (id_funcion) references funcion(id_funcion);
alter table butaca_por_funcion add constraint fk_butaca_cliente foreing key (id_cliente) references cliente(id_cliente);

alter table error add constraint fk_error_sala foreing key (id_sala) references sala_cine(id_sala);
alter table error add constraint fk_error_pelicula foreing key (id_pelicula) references pelicula(id_pelicula);
alter table error add constraint fk_error_funcion foreing key (id_funcion) references funcion(id_funcion);
alter table error add constraint fk_error_cliente foreing key (id_cliente) references cliente(id_cliente);
