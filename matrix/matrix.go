package matrix

import "errors"

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
