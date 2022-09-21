package boards 
/**
 * Representa el tablero
 * Un tablaro está formado por un array de casillas y cada una 
 * contiene un array de cajas
 * [row_0] [Box_00, Box_01,Box_02,.......Box_07]
 * [row_1] [Box_10, Box_11, Box_12,......Box_17]
 * ......
 * [row_7] [Box_70, Box_71, Box_72,......Box_77]
 */

//Rows
type Board struct {
	//TODO: Ver qué hacemos con el tamaño, lo asignamos o controlamos que no
	//sea mayor que siete
	rows []Row
}

/**
 * Devulve las filas que tiene un tablero
 */
func (board *Board) GetRows() []Row{
	return board.rows
}

/**
 * Añade una fila al tablero.
 */
func(board *Board) AddRow(row Row) {
	if(len(board.rows)<7){
		board.rows = append(board.rows, row)
	}
}

/** 
 * Devuelve las caja que tiene asociadas una fila
 */
func (board *Board) GetBoxFromRow(row int) []Box {
	return board.rows[1].GetBoxs()
}
