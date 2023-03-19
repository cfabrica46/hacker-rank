package drawingbook

const (
	lastPageOffset = 1
)

func PageCount(totalPages int32, desiredPage int32) int32 {
	if desiredPage == totalPages {
		// Caso especial: la pagina deseada es la ultima pagina
		return 0
	}

	// Calcula el numero de paginas a partir de la derecha y de la izquierda
	fromRight := (totalPages - desiredPage) / 2
	fromLeft := desiredPage / 2

	// Ajusta fromRight si la pagina deseada esta en la utlima pagina del libro
	if fromRight == 0 && totalPages%2 == 0 {
		fromRight = lastPageOffset
	}

	// Devuelve el numero minimo de paginas para llegar a la pagina deseada
	if fromRight < fromLeft {
		return fromRight
	}

	return fromLeft
}

// [|1] -> [2|3] -> [4|5] -> [6|7] -> [8|9] -> [10]
