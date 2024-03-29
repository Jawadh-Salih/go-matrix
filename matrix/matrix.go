package matrix

import (
	"errors"
	"fmt"
)

type Mat struct {
	// matrix is originally a 2D array
	// it has dimensions
	Cells [][]int
}

func (m Mat) Dims() (int, int) {
	return len(m.Cells), len(m.Cells[0])
}

func (m Mat) Add(n Mat) error {
	rm, cm := m.Dims()
	rn, cn := n.Dims()
	if rm != rn || cm != cn {
		return errors.New("mismatching dimensions for matrix addition")
	}

	for i := range m.Cells {
		for j := range m.Cells[i] {
			m.Cells[i][j] = m.Cells[i][j] + n.Cells[i][j]
		}
	}

	return nil

}

func (m Mat) Multiply(n Mat) (Mat, error) {
	rm, cm := len(m.Cells), len(m.Cells[0])
	rn, cn := len(n.Cells), len(n.Cells[0])

	if cm != rn {
		return Mat{}, errors.New(fmt.Sprintf("mismatching dimensions %d and %d", cm, rn))
	}

	o := Mat{
		Cells: [][]int{},
	}

	// loop through the rows of m (But we actuallt loop throught the rows of C)
	for r := 0; r < rm; r++ {
		// loop through the cols of n (But we actuallr loop through the cols of C)
		arr := make([]int, 0)
		for c := 0; c < cn; c++ {
			val := 0
			// loop the column of m and multiplr row to col combination of m and n
			for j := 0; j < cm; j++ {
				// cm should be equal to rn
				// Thus
				val += m.Cells[r][j] * n.Cells[j][c] // this gives the C[0][0] -> SUM(m.Cells[0][j]*n.Cells[j][0])
				// how to get C[0][1] -> SUM(m.Cells[0][j] * n.Cells[j][1])
				// therefore C[r][c] -> SUM of m.Cells[r][j] * n.Cells[j][c]
			}
			arr = append(arr, val)
		}
		o.Cells = append(o.Cells, arr)
	}

	return o, nil
}
