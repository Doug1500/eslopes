// Copyright 2015 Dorival Pedroso and Ye Win Tun. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geo

import (
	"math"

	"github.com/cpmech/gosl/chk"
)

// Circle holds data for performing operations with circles
//  Note: works in 2D only
type Circle struct {

	// input
	Xc []float64 // centre
	R  float64   // radius

	// sandbox variables
	Xa []float64 // intersection computed by SegmentIntersect
	Xb []float64 // intersection computed by SegmentIntersect

	// auxiliary
	e  []float64 // vector aligned with input segment
	e0 []float64 // unit vector aligned with input segment
	v  []float64 // vector from centre to pa
	p  []float64 // normal projection
	q  []float64 // orthogona projection
}

// Init initialises the circle
func (o *Circle) Init(xc []float64, r float64) {

	// check: circle works in 2D only
	ndim := 2
	chk.IntAssert(len(xc), ndim)

	// save input data
	o.Xc = make([]float64, ndim)
	copy(o.Xc, xc)
	o.R = r

	// allocate auxiliary variables
	o.Xa = make([]float64, ndim)
	o.Xb = make([]float64, ndim)
	o.e = make([]float64, ndim)
	o.e0 = make([]float64, ndim)
	o.v = make([]float64, ndim)
	o.p = make([]float64, ndim)
	o.q = make([]float64, ndim)
}

// SegmentIntersect computes circle-segment intersection
//
//  Circle centre @ pc           pb o
//  Circle radius = r              /
//                                /
//                      _-----_  /
//                    .'       `o xb
//                   /         / \
//                   |     +  /  |
//                   \   xc  /   /
//                    '.    /  .'
//                      `--o--'
//                        / xa
//                       /
//                   pa o
//
//
//  Input:
//   pa, pb -- two points on line
//  Output:
//   o.Xa, o.Xb -- intersections (stored in the sanbox variables)
//   nint       -- number of intersections: 0, 1 or 2
func (o *Circle) SegmentIntersect(pa, pb []float64) (nint int) {

	// e and v and the norm of e
	ndim := 2
	var e_norm float64
	for i := 0; i < ndim; i++ {
		o.e[i] = pb[i] - pa[i]
		o.v[i] = o.Xc[i] - pa[i]
		e_norm += o.e[i] * o.e[i]
	}
	e_norm = math.Sqrt(e_norm)

	// unit vector of e and its inner product with v
	var e0_dot_v float64
	for i := 0; i < ndim; i++ {
		o.e0[i] = o.e[i] / e_norm
		e0_dot_v += o.e0[i] * o.v[i]
	}

	// parallel and orthogonal projections of v along e (the line)
	var q_norm float64
	for i := 0; i < ndim; i++ {
		o.p[i] = e0_dot_v * o.e0[i]
		o.q[i] = o.v[i] - o.p[i]
		q_norm += o.q[i] * o.q[i]
	}
	q_norm = math.Sqrt(q_norm)

	// intersections
	switch {
	case q_norm > o.R: // no intersections
		nint = 0
	case q_norm < o.R: // 2 intersections
		nint = 2
		m := math.Sqrt(o.R*o.R - q_norm*q_norm)
		for i := 0; i < ndim; i++ {
			o.Xa[i] = pa[i] + o.p[i] - m*o.e0[i]
			o.Xb[i] = pa[i] + o.p[i] + m*o.e0[i]
		}
	default: // 1 intersection => tangent
		nint = 1
		for i := 0; i < ndim; i++ {
			o.Xa[i] = pa[i] + o.p[i]
		}
	}
	return
}

/*
//dot product
func LineIntLine(pa1, pa2, pa3, pa4 []float64) (xa []float64) {
	v1 := []float64{pa1[0] - pa2[0], pa1[1] - pa2[1]}
	v2 := []float64{pa3[0] - pa4[0], pa3[1] - pa4[1]}
	leftdot := ((pa2[0] - pa1[0]) * (pa3[1] - pa2[1])) - ((pa2[1] - pa1[1]) * (pa3[0] - pa2[0]))
	rightdot := ((pa2[0] - pa1[0]) * (pa4[1] - pa2[1])) - ((pa2[1] - pa1[1]) * (pa4[0] - pa2[0]))
	if leftdot == rightdot {
		// panic("Error: two lines are parallel")
		// fmt.Println("Error: two lines are parallel")
	}
	t := (v2[1]*(pa2[0]-pa4[0]) - v2[0]*(pa2[1]-pa4[1])) / (v2[0]*v1[1] - v2[1]*v1[0])
	xa1, xb1 := pa2[0]+t*v1[0], pa2[1]+t*v1[1]
	xa = []float64{xa1, xb1}
	return
}

//Linear Algebra
func LineIntLineLA(pa1, pa2, pa3, pa4 []float64) (xa []float64) {
	v1 := []float64{pa1[0] - pa2[0], pa1[1] - pa2[1]}
	v2 := []float64{pa3[0] - pa4[0], pa3[1] - pa4[1]}
	t := (v2[1]*(pa2[0]-pa4[0]) - v2[0]*(pa2[1]-pa4[1])) / (v2[0]*v1[1] - v2[1]*v1[0])
	s := (pa4[0] - pa2[0] + (v1[0])*t) / v2[0]
	if s == t {
		// panic("Error: two lines are parallel")
		// fmt.Println("Error: two lines are parallel")

	}
	xa1, xb1 := pa2[0]+t*v1[0], pa2[1]+t*v1[1]
	xa = []float64{xa1, xb1}
	return
}

//Unit Vector
//still need to find the easier method than this
func LineIntLineUV(pa1, pa2, pa3, pa4 []float64) (xa []float64) {
	v1 := []float64{pa1[0] - pa2[0], pa1[1] - pa2[1]}
	v2 := []float64{pa3[0] - pa4[0], pa3[1] - pa4[1]}
	v1_norm := math.Sqrt((v1[0])*(v1[0]) + (v1[1])*(v1[1]))
	v2_norm := math.Sqrt((v2[0])*(v2[0]) + (v2[1])*(v2[1]))
	unit_v1 := []float64{math.Abs(v1[0] / v1_norm), math.Abs(v1[1] / v1_norm)}
	unit_v2 := []float64{math.Abs(v2[0] / v2_norm), math.Abs(v2[1] / v2_norm)}
	if unit_v1[0] == unit_v2[0] && unit_v1[1] == unit_v2[1] {
		// panic("Error: two lines are parallel")
	}
	t := (v2[1]*(pa2[0]-pa4[0]) - v2[0]*(pa2[1]-pa4[1])) / (v2[0]*v1[1] - v2[1]*v1[0])
	xa1, xb1 := pa2[0]+t*v1[0], pa2[1]+t*v1[1]
	xa = []float64{xa1, xb1}
	return
}
*/
