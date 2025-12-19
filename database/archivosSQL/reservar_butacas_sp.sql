create or replace function reservar_butaca(
  _id_funcion int,
  _id_cliente int,
  _nro_fila int,
  _nro_butaca int
) returns boolean as $$
declare 
  v_id_sala int;
  v_nro_filas int;
  v_nro_butacas_por_fila int;
  v_butacas_disponibles int;
  v_estado_actual varchar(20);
begin
  select id_sala, butacas_disponibles into v_id_sala, v_butacas_disponibles
  from funcion
  where id_funcion = _id_funcion;

  if not found then
    insert into error(operacion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('reservar', _nro_fila, _nro_butaca, _id_cliente, now(), '?id de función no válido.');
    return false;
  end if;

  select nro_filas, nro_butacas_por_fila 
  into v_nro_filas, v_nro_butacas_por_fila
  from sala_cine
  where id_sala = v_id_sala;

  if _nro_fila < 1 or _nro_fila > v_nro_filas or
     _nro_butaca < 1 or _nro_butaca > v_nro_butacas_por_fila then
    insert into error(operacion, id_sala, id_funcion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('reservar', v_id_sala, _id_funcion, _nro_fila, _nro_butaca, _id_cliente, now(), '?no existe número de fila o butaca.');
    return false;
  end if;

  if v_butacas_disponibles <= 0 then
    insert into error(operacion, id_funcion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('reservar', _id_funcion, _nro_fila, _nro_butaca, _id_cliente, now(), '?sala completa para la función.');
    return false;
  end if;

  select estado into v_estado_actual
  from butaca_por_funcion
  where id_funcion = _id_funcion and nro_fila = _nro_fila and nro_butaca = _nro_butaca;

  if found and v_estado_actual in ('reservada', 'comprada') then
    insert into error(operacion, id_funcion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('reservar', _id_funcion, _nro_fila, _nro_butaca, _id_cliente, now(), '?butaca no disponible para la función.');
    return false;
  end if;

  if found and v_estado_actual = 'anulada' then
    update butaca_por_funcion
    set id_cliente = _id_cliente, estado = 'reservada'
    where id_funcion = _id_funcion and nro_fila = _nro_fila and nro_butaca = _nro_butaca;
  else
    insert into butaca_por_funcion(id_funcion, nro_fila, nro_butaca, id_cliente, estado)
    values (_id_funcion, _nro_fila, _nro_butaca, _id_cliente, 'reservada');
  end if;

  update funcion
  set butacas_disponibles = v_butacas_disponibles - 1
  where id_funcion = _id_funcion;

  return true;
end;
$$ language plpgsql;

