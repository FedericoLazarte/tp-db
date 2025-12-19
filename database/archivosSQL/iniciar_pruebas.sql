create or replace function iniciar_pruebas()
				returns void as $$
				declare
					p record;
				begin
						for p in select * from datos_de_prueba order by id_orden loop
						 --raise notice e'\n--- prueba %---', prueba.id_orden
								if p.operacion = 'nueva funcion' then
								perform apertura_funcion(p.id_sala, p.f_inicio_funcion, p.id_pelicula);
								elsif p.operacion = 'reserva butaca' then
                                perform reservar_butaca(p.id_funcion, p.id_cliente, p.nro_fila, p.nro_butaca);
								elsif p.operacion = 'compra butaca' then
                                perform comprar_butaca(p.id_funcion, p.id_cliente, p.nro_fila, p.nro_butaca);
								end if;
                                end loop;
				end;
				$$ language plpgsql;
            
select iniciar_pruebas();
insert into funcion (id_sala, fecha_inicio, hora_inicio, fecha_fin, hora_fin, id_pelicula, butacas_disponibles) values (5, '2025-08-23', '09:00:00', '2025-08-23', '10:30:00', 3, 125);
insert into funcion (id_sala, fecha_inicio, hora_inicio, fecha_fin, hora_fin, id_pelicula, butacas_disponibles) values (5, '2025-08-23', '12:00:00', '2025-08-23', '13:30:00', 3, 125);
insert into funcion (id_sala, fecha_inicio, hora_inicio, fecha_fin, hora_fin, id_pelicula, butacas_disponibles) values (5, '2025-08-23', '18:00:00', '2025-08-23', '19:30:00', 3, 125);

insert into butaca_por_funcion values (3, 10, 10, 6, 'comprada');
insert into butaca_por_funcion values (4, 2, 5, 7, 'comprada');
insert into butaca_por_funcion values (5, 3, 5, 8, 'comprada');

