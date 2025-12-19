create or replace function comprar_butaca (
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
  v_cliente_existente int;
  v_estado char(15);

begin
  select id_sala, v_butacas_disponibles
  into v_id_sala, v_butacas_disponibles
  from funcion
  where id_funcion = _id_funcion;

  if not found then
    insert into error(operacion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('comprar butaca', _nro_fila, _nro_butaca, _id_cliente, now(), '?id de función no válido.');
    return false;
  end if;

  select nro_filas, nro_butacas_por_fila
  into v_nro_filas, v_nro_butacas_por_fila
  from sala_cine
  where id_sala = v_id_sala;

  if _nro_fila < 1 or _nro_fila > v_nro_filas
    or _nro_butaca < 1 or _nro_butaca > v_nro_butacas_por_fila then
    insert into error(operacion, id_sala, id_funcion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('compra butaca', v_id_sala, _id_funcion, _nro_fila, _nro_butaca, _id_cliente, now(), '?no existe número de fila o butaca.');
    return false;
  end if;

  select id_cliente, estado
  into v_cliente_existente, v_estado
  from butaca_por_funcion
  where id_funcion = _id_funcion
  and nro_fila = _nro_fila
  and nro_butaca = _nro_butaca;

  if found and v_estado in ('reserva', 'comprada') and v_cliente_existente != _id_cliente then
    insert into error(operacion, id_funcion, nro_fila, nro_butaca, id_cliente, f_error, motivo)
    values ('compra butaca', _id_funcion, _nro_fila, _nro_butaca, _id_cliente, now(), '?butaca ocupada por otro cliente.');
    return false;
  end if;
  
  if found then
    if v_estado = 'reservada' and v_cliente_existente = _id_cliente then
      update butaca_por_funcion
      set estado = 'comprada'
      where id_funcion = _id_funcion and nro_fila = _nro_fila and nro_butaca = _nro_butaca;

    elsif v_estado = 'anulada' then
      update butaca_por_funcion
      set id_cliente = _id_cliente, estado = 'comprada'
      where id_funcion = _id_funcion and nro_fila = _nro_fila and nro_butaca = _nro_butaca;

      update funcion
      set butacas_disponibles = v_butacas_disponibles - 1
      where id_funcion = _id_funcion
    end if;
  else 
    insert into butaca_por_funcion(id_funcion, nro_fila, nro_butaca, id_cliente, estado)
    values (_id_funcion, _nro_fila, _nro_butaca, _id_cliente, 'comprada')

    update funcion
    set butacas_disponibles = v_butacas_disponibles - 1
    where id_funcion = _id_funcion;
  end if;
  return true;
end;
$$ language plpgsql;
