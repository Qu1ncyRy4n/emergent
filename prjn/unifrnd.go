// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prjn

import (
	"math"
	"math/rand"
	"sort"

	"github.com/emer/emergent/erand"
	"github.com/emer/etable/etensor"
)

// UnifRnd implements uniform random pattern of connectivity between two layers
// uses a permuted (shuffled) list for without-replacement randomness
// and maintains its own local random seed for fully replicable results
// (if seed is not set when run, then random number generator is used to create seed)
// should reset seed after calling to resume sequence appropriately
type UnifRnd struct {
	PCon    float32 `min:"0" max:"1" desc:"probability of connection (0-1)"`
	RndSeed int64   `view:"-" desc:"the current random seed"`
	SelfCon bool    `desc:"if true, and connecting layer to itself (self projection), then make a self-connection from unit to itself"`
	Recip   bool    `desc:"reciprocal connectivity: if true, switch the sending and receiving layers to create a symmetric top-down projection -- ESSENTIAL to use same RndSeed between two prjns to ensure symmetry"`
}

func NewUnifRnd() *UnifRnd {
	return &UnifRnd{PCon: 0.5}
}

func (ur *UnifRnd) Name() string {
	return "UnifRnd"
}

func (ur *UnifRnd) Connect(send, recv *etensor.Shape, same bool) (sendn, recvn *etensor.Int32, cons *etensor.Bits) {
	if ur.PCon >= 1 {
		return ur.ConnectFull(send, recv, same)
	}
	if ur.Recip {
		return ur.ConnectRecip(send, recv, same)
	}
	sendn, recvn, cons = NewTensors(send, recv)
	slen := send.Len()
	rlen := recv.Len()

	noself := same && !ur.SelfCon
	var nsend int
	if noself {
		nsend = int(math.Round(float64(ur.PCon) * float64(slen-1)))
	} else {
		nsend = int(math.Round(float64(ur.PCon) * float64(slen)))
	}

	// NOTE: this is reasonably accurate: mean + 3 * SEM, but we can just use
	// empirical values more easily and safely.

	// recv number is even distribution across recvs plus some imbalance factor
	// nrMean := float32(rlen*nsend) / float32(slen)
	// // add 3 * SEM as corrective factor
	// nrSEM := nrMean / mat32.Sqrt(nrMean)
	// nrecv := int(nrMean + 3*nrSEM)
	// if nrecv > rlen {
	// 	nrecv = rlen
	// }

	rnv := recvn.Values
	for i := range rnv {
		rnv[i] = int32(nsend)
	}

	if ur.RndSeed == 0 {
		ur.RndSeed = int64(rand.Uint64())
	}
	rand.Seed(ur.RndSeed)

	sordlen := slen
	if noself {
		sordlen--
	}

	sorder := rand.Perm(sordlen)
	slist := make([]int, nsend)
	for ri := 0; ri < rlen; ri++ {
		if noself { // need to exclude ri
			ix := 0
			for j := 0; j < slen; j++ {
				if j != ri {
					sorder[ix] = j
					ix++
				}
			}
			erand.PermuteInts(sorder)
		}
		copy(slist, sorder)
		sort.Ints(slist) // keep list sorted for more efficient memory traversal etc
		for si := 0; si < nsend; si++ {
			off := ri*slen + slist[si]
			cons.Values.Set(off, true)
		}
		erand.PermuteInts(sorder)
	}

	// 	set send n's empirically
	snv := sendn.Values
	for si := range snv {
		nr := 0
		for ri := 0; ri < rlen; ri++ {
			off := ri*slen + si
			if cons.Values.Index(off) {
				nr++
			}
		}
		snv[si] = int32(nr)
	}
	return
}

// ConnectRecip does reciprocal connectvity
func (ur *UnifRnd) ConnectRecip(send, recv *etensor.Shape, same bool) (sendn, recvn *etensor.Int32, cons *etensor.Bits) {
	sendn, recvn, cons = NewTensors(send, recv)
	slen := recv.Len() // swapped
	rlen := send.Len()

	slenR := send.Len() // NOT swapped

	noself := same && !ur.SelfCon
	var nsend int
	if noself {
		nsend = int(math.Round(float64(ur.PCon) * float64(slen-1)))
	} else {
		nsend = int(math.Round(float64(ur.PCon) * float64(slen)))
	}

	rnv := sendn.Values // swapped
	for i := range rnv {
		rnv[i] = int32(nsend)
	}

	if ur.RndSeed == 0 {
		ur.RndSeed = int64(rand.Uint64())
	}
	rand.Seed(ur.RndSeed)

	sordlen := slen
	if noself {
		sordlen--
	}

	sorder := rand.Perm(sordlen)
	slist := make([]int, nsend)
	for ri := 0; ri < rlen; ri++ {
		if noself { // need to exclude ri
			ix := 0
			for j := 0; j < slen; j++ {
				if j != ri {
					sorder[ix] = j
					ix++
				}
			}
			erand.PermuteInts(sorder)
		}
		copy(slist, sorder)
		sort.Ints(slist) // keep list sorted for more efficient memory traversal etc
		for si := 0; si < nsend; si++ {
			off := slist[si]*slenR + ri
			cons.Values.Set(off, true)
		}
		erand.PermuteInts(sorder)
	}

	// set send n's empirically
	snv := recvn.Values // swapped
	for si := range snv {
		nr := 0
		for ri := 0; ri < rlen; ri++ { // actually si
			off := si*slenR + ri
			if cons.Values.Index(off) {
				nr++
			}
		}
		snv[si] = int32(nr)
	}
	return
}

func (ur *UnifRnd) ConnectFull(send, recv *etensor.Shape, same bool) (sendn, recvn *etensor.Int32, cons *etensor.Bits) {
	sendn, recvn, cons = NewTensors(send, recv)
	cons.Values.SetAll(true)
	nsend := send.Len()
	nrecv := recv.Len()
	if same && !ur.SelfCon {
		for i := 0; i < nsend; i++ { // nsend = nrecv
			off := i*nsend + i
			cons.Values.Set(off, false)
		}
		nsend--
		nrecv--
	}
	rnv := recvn.Values
	for i := range rnv {
		rnv[i] = int32(nsend)
	}
	snv := sendn.Values
	for i := range snv {
		snv[i] = int32(nrecv)
	}
	return
}
