create or replace function apertura_funcion(_id_sala int, _fecha_inicio timestamp, _id_pelicula int) return int as $$
declare
  _id_funcion int := -1;
  _fecha_fin date;
  _hora_fin time;
  _pelicula record;
  _sala record;

begin
  select * into _pelicula from pelicula where id_pelicula = _id_pelicula;
  if not found then
    insert into error (
      operacion, id_sala, f_inicio_funcion,
      id_pelicula, id_funcion, nro_fila, nro_butaca, id_cliente,
      f_error, motivo
    ) values (
    'apertura_funcion', _id_sala, _fecha_inicio,
    _id_pelicula, null, null, null, null,
    now(), 'id de película no válido.'
  );
  return _id_funcion;
  end if;

  _fecha_fin := (_fecha_inicio + _pelicula.duracion)::date;
  _hora_fin := (_fecha_inicio + _pelicula.duracion)::time;

  select * into _sala from sala_cine where id_sala = _id_sala;

  if not found then
    insert into error (
      operacion, id_sala, f_inicio_funcion,
      id_pelicula, id_funcion, nro_fila, nro_butaca, id_cliente,
      f_error, motivo
    ) values (
    'apertura_funcion', _id_sala, _fecha_inicio,
    _id_pelicula, null, null, null, null,
    now(), 'id de sala no válido.'
  );
  return _id_funcion;
  end if;

  if not (_fecha_inicio > now()) then
    insert into error (
      operacion, id_sala, f_inicio_funcion,
      id_pelicula, id_funcion, nro_fila, nro_butaca, id_cliente,
      f_error, motivo
    ) values (
    'apertura_funcion', _id_sala, _fecha_inicio,
    _id_pelicula, null, null, null, null,
    now(), 'no se permite abrir una nuev función con retroactividad'
  );
  retur _id_funcion;
  end if;

  if exists (
    select 1 from funccion f
    where f.id_sala = _id_sala
    and not (_fecha_inicio > f.fecha_fin or _hora_fin < f.hora_inicio)
    ) then
      insert into error (
        operacion, id_sala, f_inicio_funcion,
        id_pelicula, id_funcion, nro_fila, nro_butaca, id_cliente,
        f_error, motivo 
      ) values (
      'apertura_funcion', _id_sala, _fecha_inicio,
      _id_pelicula, null, null, null, null,
      now(), 'no se permite solapar funciones en una sala'
    );
    return _id_funcion;
  end if;

  if lower(_pelicula.formato) <> lower(_sala.formato) then
    insert into error (
      operacion, id_sala, f_inicio_funcion,
      id_pelicula, id_funcion, nro_fila, nro_butaca, id_cliente,
      f_error, motivo
    ) values (
    'apertura_funcion', _id_sala, _fecha_inicio,
    _id_pelicula, null, null, null, null,
    now(), 'sala no habilitada para formato de película'
  );
  return _id_funcion;
  end if;

  insert into funcion (
    id_sala, fecha_inicio, hora_inicio,
    fecha_fin, hora_fin, id_pelicula, butacas_disponibles
  ) values (
  _id_sala,
  _fecha_inicio::date,
  _fecha_inicio::time,
  _fecha_fin,
  _hora_fin,
  _id_pelicula,
  _sala.capacidad_total
)
returning id_funcion into _id_funcion;
return _id_funcion;
end;
$$ language plpgsql;
