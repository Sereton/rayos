package camera

import (
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
)

func View_transformation(from, to, up primitives.Tuple) matrix.Matrix {
	forward := to.Sub(from).Normalize()
	up = up.Normalize()
	left := forward.Cross(up)
	trueUp := left.Cross(forward)
	M := matrix.Identity(4)
	M.Set(0, 0, left.X)
	M.Set(0, 1, left.Y)
	M.Set(0, 2, left.Z)
	M.Set(1, 0, trueUp.X)
	M.Set(1, 1, trueUp.Y)
	M.Set(1, 2, trueUp.Z)
	M.Set(2, 0, -forward.X)
	M.Set(2, 1, -forward.Y)
	M.Set(2, 2, -forward.Z)

	T := matrix.Translation(-from.X, -from.Y, -from.Z)
	M = M.MMulti(T)
	return M
}
