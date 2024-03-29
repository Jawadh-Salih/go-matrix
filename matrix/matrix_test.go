package matrix_test

import (
	"testing"

	"github.com/Jawadh-Salih/go-matrix/matrix"
	"github.com/stretchr/testify/assert"
)

func TestMatrixAdd(t *testing.T) {

	t.Run("Matrices of mismatching dimensions should error on addition", func(t *testing.T) {

		m := matrix.Mat{
			Cells: [][]int{
				{1, 2},
				{1, 3},
			},
		}

		n := matrix.Mat{
			Cells: [][]int{
				{1, 2},
			},
		}

		err := m.Add(n)
		assert.NotNil(t, err)
	})

	t.Run("Matrices of mismatching dimensions should error on addition", func(t *testing.T) {

		m := matrix.Mat{
			Cells: [][]int{
				{1, 2},
				{1, 3},
			},
		}

		n := matrix.Mat{
			Cells: [][]int{
				{1, 2},
				{3, 4},
			},
		}

		err := m.Add(n)
		assert.Nil(t, err)
		assert.Equal(t,
			matrix.Mat{
				Cells: [][]int{
					{2, 4},
					{4, 7},
				},
			},
			m,
		)
	})

}
