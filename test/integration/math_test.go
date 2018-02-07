//+build !windows
//+build integration

package integration

import (
	. "gopkg.in/check.v1"
)

type MathSuite struct{}

var _ = Suite(&MathSuite{})

func (s *MathSuite) TestMath(c *C) {
	inOutTest(c, `{{ math.Add 1 2 3 4 }} {{ add -5 5 }}`, "10 0")
	inOutTest(c, `{{ math.Sub 10 5 }} {{ sub -5 5 }}`, "5 -10")
	inOutTest(c, `{{ math.Mul 1 2 3 4 }} {{ mul -5 5 }}`, "24 -25")
	inOutTest(c, `{{ math.Div 5 3 }} {{ div -5 5 }}`, "1 -1")
	inOutTest(c, `{{ math.Rem 5 3 }} {{ rem 2 2 }}`, "2 0")
	inOutTest(c, `{{ math.Pow 8 4 }} {{ pow 2 2 }}`, "4096 4")
	inOutTest(c, `{{ math.Seq 0 }}, {{ seq 0 3 }}, {{ seq -5 -10 2 }}`, "[1 0], [0 1 2 3], [-5 -7 -9]")
}
