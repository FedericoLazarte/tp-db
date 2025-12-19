create or replace function generar_email_butaca()
returns trigger as $$
declare
    v_cliente_record record;
    v_funcion_record record;
    v_asunto text;
    v_cuerpo text;
begin
    select nombre, apellido, email into v_cliente_record
    from cliente where id_cliente = new.id_cliente;

    select f.*, p.titulo, p.formato as formato_pelicula, s.nombre as nombre_sala
    into v_funcion_record
    from funcion f
    join pelicula p on f.id_pelicula = p.id_pelicula
    join sala_cine s on f.id_sala = s.id_sala
    where f.id_funcion = new.id_funcion;

    if new.estado = 'reservada' then
        v_asunto := 'Jardines de Noviembre - Reserva de butaca';
        v_cuerpo := 'Estimade ' || v_cliente_record.nombre || ' ' || v_cliente_record.apellido || 
                   ', se registró su reserva para la película "' || v_funcion_record.titulo || 
                   '" (' || v_funcion_record.formato_pelicula || ') en la sala ' || v_funcion_record.nombre_sala || 
                   ' el día ' || v_funcion_record.fecha_inicio || 
                   ' a las ' || v_funcion_record.hora_inicio || 
                   '. Butaca: Fila ' || new.nro_fila || ', Asiento ' || new.nro_butaca || '.';
    
    elsif new.estado = 'comprada' then
        v_asunto := 'Jardines de Noviembre - Compra de butaca';
        v_cuerpo := 'Estimade ' || v_cliente_record.nombre || ' ' || v_cliente_record.apellido || 
                   ', se registró su compra para la película "' || v_funcion_record.titulo || 
                   '" (' || v_funcion_record.formato_pelicula || ') en la sala ' || v_funcion_record.nombre_sala || 
                   ' el día ' || v_funcion_record.fecha_inicio || 
                   ' a las ' || v_funcion_record.hora_inicio || 
                   '. Butaca: Fila ' || new.nro_fila || ', Asiento ' || new.nro_butaca || '.';
	end if;

    insert into envio_email (
        f_generacion, 
        email_cliente, 
        asunto, 
        cuerpo, 
        f_envio, 
        estado
    ) values (
        now(),
        v_cliente_record.email,
        v_asunto,
        v_cuerpo,
        null,
        'pendiente'
    );
    
    return new;
end;
$$ language plpgsql;

create trigger tr_generar_email_butaca
after insert or update of estado on butaca_por_funcion
for each row
when (new.estado in ('reservada', 'comprada'))
execute function generar_email_butaca();

