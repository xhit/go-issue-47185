package main

import (
	"sdi/drsqlite3"
	"sdi/p1"
)

func main() {
	// declare "j" variable because inline it in "c" make build success
	var j map[string]longStruct1
	c := j[""]
	eD(&c)

	// reading this also produce the build error
	_ = drsqlite3.DBimport.A
}

func eD(c *longStruct1) {
	p1.Fp1()
}

// total variables here can cause the build error, so less variables of defined is fine
type longStruct1 struct {
	a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z                           string
	aa, ab, ac, ad, ae, af, ag, ah, ai, aj, ak, al, am, an, ao, ap, aq, ar, as, at, au, av, aw, ax, ay, az string
	ba, bb, bc, bd, be, bf, bg, bh, bi, bj, bk, bl, bm, bn                                                 string
}
