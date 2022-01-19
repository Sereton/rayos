package matrix

import (
	p "bolijollo/rayos/primitives"
	"math"
)

type Matrix struct {
	Dimension int
	Elements  [][]float64
}

func CreateMatrix(dimension int) Matrix {
	elem := make([][]float64, dimension)
	for i := 0; i < dimension; i++ {
		elem[i] = make([]float64, dimension)
	}
	return Matrix{Dimension: dimension, Elements: elem}
}

func (m *Matrix) Get(i, j int) float64 {
	return m.Elements[i][j]
}

func (m *Matrix) Set(i, j int, value float64) {
	m.Elements[i][j] = value
}

func Identity(dimension int) Matrix {
	m := CreateMatrix(dimension)
	for i := 0; i < dimension; i++ {
		m.Elements[i][i] = 1
	}
	return m
}

func Translation(x, y, z float64) Matrix {
	m := Identity(4)
	m.Elements[0][3] = x
	m.Elements[1][3] = y
	m.Elements[2][3] = z
	return m
}

func Scaling(x, y, z float64) Matrix {
	m := Identity(4)
	m.Elements[0][0] = x
	m.Elements[1][1] = y
	m.Elements[2][2] = z
	return m
}

func Rotation_Z(theta float64) Matrix {
	m := Identity(4)
	m.Elements[0][0] = math.Cos(theta)
	m.Elements[0][1] = -math.Sin(theta)
	m.Elements[1][0] = math.Sin(theta)
	m.Elements[1][1] = math.Cos(theta)
	return m
}

func Rotation_X(theta float64) Matrix {
	m := Identity(4)
	m.Elements[1][1] = math.Cos(theta)
	m.Elements[1][2] = -math.Sin(theta)
	m.Elements[2][1] = math.Sin(theta)
	m.Elements[2][2] = math.Cos(theta)
	return m
}

func Rotation_Y(theta float64) Matrix {
	m := Identity(4)
	m.Elements[0][0] = math.Cos(theta)
	m.Elements[0][2] = math.Sin(theta)
	m.Elements[2][0] = -math.Sin(theta)
	m.Elements[2][2] = math.Cos(theta)
	return m
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	m := Identity(4)
	m.Elements[0][1] = xy
	m.Elements[0][2] = xz
	m.Elements[1][0] = yx
	m.Elements[1][2] = yz
	m.Elements[2][0] = zx
	m.Elements[2][1] = zy
	return m
}

func (m *Matrix) Equal(other *Matrix) bool {
	if m.Dimension != other.Dimension {
		return false
	}
	for i := 0; i < m.Dimension; i++ {
		for j := 0; j < m.Dimension; j++ {
			if m.Elements[i][j] != other.Elements[i][j] {
				return false
			}
		}
	}
	return true

}

func (m Matrix) MMulti(other Matrix) Matrix {
	if m.Dimension != other.Dimension {
		panic("Matrix dimensions do not match")
	}
	result := CreateMatrix(m.Dimension)
	for i := 0; i < m.Dimension; i++ {
		for j := 0; j < m.Dimension; j++ {

			for k := 0; k < m.Dimension; k++ {
				result.Elements[i][j] += m.Elements[i][k] * other.Elements[k][j]
			}
		}
	}
	return result
}

func (m Matrix) VMulti(t *p.Tuple) p.Tuple {
	if m.Dimension != 4 {
		panic("Matrix is not a 4x4 matrix")
	}
	result := make([]float64, 4)
	result[0] = m.Elements[0][0]*t.X + m.Elements[0][1]*t.Y + m.Elements[0][2]*t.Z + m.Elements[0][3]*t.W
	result[1] = m.Elements[1][0]*t.X + m.Elements[1][1]*t.Y + m.Elements[1][2]*t.Z + m.Elements[1][3]*t.W
	result[2] = m.Elements[2][0]*t.X + m.Elements[2][1]*t.Y + m.Elements[2][2]*t.Z + m.Elements[2][3]*t.W
	result[3] = m.Elements[3][0]*t.X + m.Elements[3][1]*t.Y + m.Elements[3][2]*t.Z + m.Elements[3][3]*t.W
	tup := p.Tuple{X: result[0], Y: result[1], Z: result[2], W: result[3]}

	return tup
}

func (m Matrix) Transpose() Matrix {
	result := CreateMatrix(m.Dimension)
	for i := 0; i < m.Dimension; i++ {
		for j := 0; j < m.Dimension; j++ {
			result.Elements[i][j] = m.Elements[j][i]
		}
	}
	return result
}

func (m *Matrix) Cofactor(i, j int) float64 {
	sign := math.Pow(-1, float64(i+j))
	minor := CreateMatrix(m.Dimension - 1)
	l := 0
	mm := 0
	for x := 0; x < m.Dimension; x++ {
		if x == i {
			continue
		}
		for y := 0; y < m.Dimension; y++ {
			if y == j {
				continue
			}
			minor.Elements[l][mm] = m.Elements[x][y]
			mm++
		}
		l++
		mm = 0
	}
	return sign * minor.Determinant()
}

func (m *Matrix) Determinant() float64 {

	if m.Dimension == 1 {
		return m.Elements[0][0]
	} else if m.Dimension == 2 {
		return m.Elements[0][0]*m.Elements[1][1] - m.Elements[0][1]*m.Elements[1][0]
	} else {
		sum := 0.0
		for i := 0; i < m.Dimension; i++ {
			sum += m.Elements[0][i] * m.Cofactor(0, i)
		}
		return sum
	}

}

func (m *Matrix) Inverse() Matrix {
	if m.Determinant() == 0 {
		panic("Matrix is not invertible")
	}
	result := CreateMatrix(m.Dimension)
	for i := 0; i < m.Dimension; i++ {
		for j := 0; j < m.Dimension; j++ {
			result.Elements[i][j] = m.Cofactor(j, i) / m.Determinant()
		}
	}
	return result
}
