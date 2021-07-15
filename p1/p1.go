package p1

import "sync"

type longStruct2 struct {
	a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z                           string
	aa, ab, ac, ad, ae, af, ag, ah, ai, aj, ak, al, am, an, ao, ap, aq, ar, as, at, au, av, aw, ax, ay, az string
	ba, bb, bc, bd, be, bf, bg, bh, bi, bj, bk, bl, bm, bn, bo                                             string
}

var m = &sync.RWMutex{}

var v map[string]longStruct2

func Fp1() {
	m.RLock()
	_ = v[""]
	m.RUnlock()
}
