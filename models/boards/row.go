package boards

// Una fila contiene un slice de cajas
type Row struct {
	boxs []Box
}

/**
 * Devuelve la lista de casillas que tiene una fila.
 */
func (row *Row) GetBoxs() []Box{
	return	row.boxs
}