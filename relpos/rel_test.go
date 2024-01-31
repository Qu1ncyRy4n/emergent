// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package relpos

import (
	"fmt"
	"testing"

	"cogentcore.org/core/mat32"
)

func TestRels(t *testing.T) {
	rp := Rel{}
	rp.Defaults()
	rp.Rel = RightOf
	rp.YAlign = Center
	rs := rp.Pos(mat32.Vec3{}, mat32.V2(10, 10), mat32.V2(4, 4))
	fmt.Printf("rp: %v rs: %v\n", rp, rs)
	rp.YAlign = Front
	rs = rp.Pos(mat32.Vec3{}, mat32.V2(10, 10), mat32.V2(4, 4))
	fmt.Printf("rp: %v rs: %v\n", rp, rs)
	rp.YAlign = Back
	rs = rp.Pos(mat32.Vec3{}, mat32.V2(10, 10), mat32.V2(4, 4))
	fmt.Printf("rp: %v rs: %v\n", rp, rs)
}
