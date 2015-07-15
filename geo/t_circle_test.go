// Copyright 2015 Dorival Pedroso and Ye Win Tun. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geo

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

func Test_circint01(tst *testing.T) {

	verbose()
	chk.PrintTitle("circint01. horizontal line. 2 intersec")

	var circ Circle
	xc := []float64{10, 10}
	r := 10.0
	circ.Init(xc, r)

	pa, pb := []float64{0, 10}, []float64{20, 10}
	ni := circ.SegmentIntersect(pa, pb)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "circ.Xa", 1e-17, circ.Xa, []float64{0, 10})
	chk.Vector(tst, "circ.Xb", 1e-17, circ.Xb, []float64{20, 10})
}

func Test_circint02(tst *testing.T) {

	verbose()
	chk.PrintTitle("circint02. sloped line (increasing). 2 intersec")

	var circ Circle
	xc := []float64{8, 8}
	r := 9.0
	circ.Init(xc, r)

	pa, pb := []float64{4, 2}, []float64{18, 9}
	ni := circ.SegmentIntersect(pa, pb)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "circ.Xa", 1e-4, circ.Xa, []float64{2.2135, 1.10675})
	chk.Vector(tst, "circ.Xb", 1e-4, circ.Xb, []float64{16.9865, 8.49325})
}

func Test_circint03(tst *testing.T) {

	verbose()
	chk.PrintTitle("circint03. sloped line (decreasing). 2 intersec")

	var circ Circle
	xc := []float64{5, -5}
	r := 3.0
	circ.Init(xc, r)

	pa, pb := []float64{2, 0}, []float64{9, -11.5}
	ni := circ.SegmentIntersect(pa, pb)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("Xa = %v\n", circ.Xa)
	io.Pf("Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "Xa", 1e-2, circ.Xa, []float64{3.4719, -2.4181})
	chk.Vector(tst, "Xb", 1e-2, circ.Xb, []float64{6.5917, -7.5435})
}

/*
//gradient is infinite (vertical line)
func Test_circint04(tst *testing.T) {
	r := 5.0
	pa, pb, pc := []float64{5.5, 3}, []float64{5.5, 7}, []float64{10, 5}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "circ.Xa", 1e-2, circ.Xa, []float64{5.5, 2.82055})
	chk.Vector(tst, "circ.Xb", 1e-2, circ.Xb, []float64{5.5, 7.17944})
}

////gradient is infinite (vertical line) going through the centre of circle
func Test_circint05(tst *testing.T) {
	r := 5.0
	pa, pb, pc := []float64{10, 10}, []float64{10, 0}, []float64{10, 5}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "circ.Xa", 1e-2, circ.Xa, []float64{10, 10})
	chk.Vector(tst, "circ.Xb", 1e-2, circ.Xb, []float64{10, 0})
}

//gradient is less than 0
func Test_circint06(tst *testing.T) {
	r := 2.0
	pa, pb, pc := []float64{-1, 8}, []float64{1, 1}, []float64{0, 4}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "circ.Xa", 1e-4, circ.Xa, []float64{-0.41607, 5.956245})
	chk.Vector(tst, "circ.Xb", 1e-4, circ.Xb, []float64{0.68022, 2.11923})
}

//gradient is less than 0, same as test 6 but the vector a and b were swapped
func Test_circint07(tst *testing.T) {
	r := 2.0
	pa, pb, pc := []float64{1, 1}, []float64{-1, 8}, []float64{0, 4}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 2)
	chk.Vector(tst, "circ.Xa", 1e-4, circ.Xa, []float64{0.68022, 2.11923})
	chk.Vector(tst, "circ.Xb", 1e-4, circ.Xb, []float64{-0.41607, 5.956245})
}

//[one intersection]

//gradient is greater than 0
func Test_circint08(tst *testing.T) {
	r := math.Sqrt(4 * 4 * 2)
	pa, pb, pc := []float64{2, 2}, []float64{14, 14}, []float64{3, 11}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 1)
	chk.Vector(tst, "circ.Xa", 1e-4, circ.Xa, []float64{7, 7})
}

//gradient is 0
func Test_circint09(tst *testing.T) {
	r := 5.0
	pa, pb, pc := []float64{3, 0}, []float64{15, 0}, []float64{10, 5}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 1)
	chk.Vector(tst, "circ.Xa", 1e-4, circ.Xa, []float64{10, 0})
}

//gradient is less than 0
func Test_circint10(tst *testing.T) {
	r := math.Sqrt(8)
	pa, pb, pc := []float64{0, 4}, []float64{4, 0}, []float64{4, 4}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 1)
	chk.Vector(tst, "circ.Xa", 1e-4, circ.Xa, []float64{2, 2})
}

//[no intersection]
func Test_circint11(tst *testing.T) {
	r := 5.0
	pa, pb, pc := []float64{2, 2}, []float64{14, 14}, []float64{3, 11}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 0)
}

func Test_circint12(tst *testing.T) {
	r := 5.0
	pa, pb, pc := []float64{-3, 6}, []float64{10, 6}, []float64{0, 0}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 0)
}

func Test_circint13(tst *testing.T) {
	r := 5.0
	pa, pb, pc := []float64{-6, 10}, []float64{-6, 9}, []float64{0, 0}
	circ.Xa, circ.Xb, ni := circ.SegmentIntersect(pa, pb, pc, r)
	io.Pf("ni = %v\n", ni)
	io.Pf("pa = %v\n", pa)
	io.Pf("pb = %v\n", pb)
	io.Pf("circ.Xa = %v\n", circ.Xa)
	io.Pf("circ.Xb = %v\n", circ.Xb)
	chk.IntAssert(ni, 0)
}

// parallel line
func Test_lineintline01(tst *testing.T) {
	pa1 := []float64{1.0, 1.0}
	pa2 := []float64{2.0, 0.0}
	pa3 := []float64{1.0, 2.0}
	pa4 := []float64{2.0, 1.0}
	x := LineIntLine(pa1, pa2, pa3, pa4)
	io.Pf("pa1 = %v\n", pa1)
	io.Pf("pa2  = %v\n", pa2)
	io.Pf("pa3 = %v\n", pa3)
	io.Pf("pa4 = %v\n", pa4)
	io.Pf("x = %v\n", x)
	// chk.Scalar(tst, "circ.Xa", 1e-3, circ.Xa, 3.714)
	// chk.Scalar(tst, "circ.Xb", 1e-3, circ.Xb, 2.285)
}

//parallel
func Test_lineintline02(tst *testing.T) {
	pa1 := []float64{0.0, 0.0}
	pa2 := []float64{2.0, 2.0}
	pa3 := []float64{0.0, 1.0}
	pa4 := []float64{5.0, 6.0}
	x := LineIntLine(pa1, pa2, pa3, pa4)
	io.Pf("pa1 = %v\n", pa1)
	io.Pf("pa2  = %v\n", pa2)
	io.Pf("pa3 = %v\n", pa3)
	io.Pf("pa4 = %v\n", pa4)
	io.Pf("x = %v\n", x)
	// chk.Scalar(tst, "circ.Xa", 1e-3, circ.Xa, 3.714)
	// chk.Scalar(tst, "circ.Xb", 1e-3, circ.Xb, 2.285)
}

//same line
func Test_lineintline03(tst *testing.T) {
	pa1 := []float64{0.0, 0.0}
	pa2 := []float64{1.0, 1.0}
	pa3 := []float64{1.0, 1.0}
	pa4 := []float64{2.0, 2.0}
	x := LineIntLine(pa1, pa2, pa3, pa4)
	io.Pf("pa1 = %v\n", pa1)
	io.Pf("pa2  = %v\n", pa2)
	io.Pf("pa3 = %v\n", pa3)
	io.Pf("pa4 = %v\n", pa4)
	io.Pf("x = %v\n", x)
	// chk.Scalar(tst, "circ.Xa", 1e-3, circ.Xa, 3.714)
	// chk.Scalar(tst, "circ.Xb", 1e-3, circ.Xb, 2.285)
}

//intersect
func Test_lineintline04(tst *testing.T) {
	pa1 := []float64{2.0, 2.0}
	pa2 := []float64{5.0, 5.0}
	pa3 := []float64{5.0, 2.0}
	pa4 := []float64{3.0, 6.0}
	x := LineIntLine(pa1, pa2, pa3, pa4)
	io.Pf("pa1 = %v\n", pa1)
	io.Pf("pa2  = %v\n", pa2)
	io.Pf("pa3 = %v\n", pa3)
	io.Pf("pa4 = %v\n", pa4)
	io.Pf("x = %v\n", x)
	chk.Vector(tst, "x", 1e-3, x, []float64{4, 4})
}
*/
