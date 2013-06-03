//	Copyright (c) 2013, Bogdan S.
//	Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package	interpolation

import	(
	"math"
	"math_tools"
)


/*	HIROSHI AKIMA "A method of smooth curve fitting" 1969

	Summary

	This method is based on a piecewise function composed of a set of polynomials, each of degree three, at most, and applicable to successive intervals of the given points.

	Method assumes that the slope of the curve at each given point is determined locally by the coordinates of five points, with the point in question as a center point, and two points on each side of it ( section 1 ).

	A polynomial of degree three representing a portion of the curve between a pair of given points is determined by the coordinates of and the slopes at the two points ( section 2 ).

	Since the slope of the curve must thus be determined also at the end points of the curve, estimation of two more points is necessary at each end point ( section 3 ).


	Details

	1) We assume that the slope t of the curve at point 3 is determined by :

		t = ( | m4 - m3 | * m2  +  | m2 - m1 | * m3 )  /  ( | m4 - m3 |  +  | m2 - m1 | )

	where m1, m2, m3, and m4 are the slopes of line segments 12, 23, 34, and 45 respectively.

		Note :	m1 [ 12 ] = ( y2 - y1 ) / ( x2 - x1 )


	1.1) Special case if m1 == m2 && m3 == m4	:

		t = ( m2 + m3 ) / 2


	2) Since we have four conditions for determining the polynomial for an interval between two points ( x1, y1 ) and ( x2, y2 ), we assume that the curve between a pair of points can be expressed by a polynomial of, at most, degree three.
	The polynomial, though uniquely determined, can be written in several ways.

	As an example we shall give the following form:

		y = p0 +  p1( x - x1 )  +  p2( x - x1 )^2  +  p3( x - x1 )^3

	where

		p0 = y1
		p1 = t1
		p2 = [ 3( y2 - y1 ) / ( x2 - x1 )  - 2* t1  -  t2 ]  /  ( x2 - x1 )
		p3 = [ t1 + t2 - 2( y2 - y1 ) / ( x2 - x1 ) ]  /  ( x2 - x1 )^2

		t1 and t2 are the slopes at the two points


	3) At each end of the curve, two more points have to be estimated from the given points.
	We assume for this purpose that the end point (x3, y3) and two adjacent given points (x2, y2) and (x1, y1), together with two more points (x4, y4) and (x5, y5) to be estimated, lie on a curve expressed by :

		y = g0 +  g1( x - x3 )  +  g2( x - x3 )^2

	where the g's are constants. Assuming that x5 - x3 == x4 - x2 == x3 - x1, we can determine the ordinates y4 and y5.

		( y5 - y4 ) / ( x5 - x4 )  -  ( y4 - y3 ) / ( x4 - x3 )	=
		( y4 - y3 ) / ( x4 - x3 )  -  ( y3 - y2 ) / ( x3 - x2 )	=
		( y3 - y2 ) / ( x3 - x2 )  -  ( y2 - y1 ) / ( x2 - x1 )
		--------

		x5 - x3 == x4 - x2	=>

			x2	= x3 + x4 - x5
			x5	= x3 + x4 - x2

		x4 - x2 == x3 - x1	=>

			x1	= x2 + x3 - x4
			x4	= x2 + x3 - x1

		Interpolation is better on mirror edge points

		x5 - x3 == x3 - x1	=>

			x1	= 2 * x3 - x5
			x5	= 2 * x3 - x1
		--------

		( y5 - y4 ) / ( x5 - x4 )  -  ( y4 - y3 ) / ( x4 - x3 )  =  ( y4 - y3 ) / ( x4 - x3 )  -  ( y3 - y2 ) / ( x3 - x2 )
			=>
			y2  =  y3 - ( x3 - x2 ) * ( 2 * ( y4 - y3 ) / ( x4 - x3 )  -  ( y5 - y4 ) / ( x5 - x4 ) )
			y5  =  y4 + ( x5 - x4 ) * ( 2 * ( y4 - y3 ) / ( x4 - x3 )  -  ( y3 - y2 ) / ( x3 - x2 ) )


		( y4 - y3 ) / ( x4 - x3 )  -  ( y3 - y2 ) / ( x3 - x2 )  =  ( y3 - y2 ) / ( x3 - x2 )  -  ( y2 - y1 ) / ( x2 - x1 )
			=>
			y1  =  y2 - ( x2 - x1 ) * ( 2 * ( y3 - y2 ) / ( x3 - x2 )  -  ( y4 - y3 ) / ( x4 - x3 ) )
			y4  =  y3 + ( x4 - x3 ) * ( 2 * ( y3 - y2 ) / ( x3 - x2 )  -  ( y2 - y1 ) / ( x2 - x1 ) )

		Interpolation is better on mirror edge points

		( y5 - y4 ) / ( x5 - x4 )  -  ( y4 - y3 ) / ( x4 - x3 )  =  ( y3 - y2 ) / ( x3 - x2 )  -  ( y2 - y1 ) / ( x2 - x1 )
			=>
			y1  =  y2 - ( x2 - x1 ) * ( ( y3 - y2 ) / ( x3 - x2 )  +  ( y4 - y3 ) / ( x4 - x3 )  -  ( y5 - y4 ) / ( x5 - x4 ) )
			y5  =  y4 + ( x5 - x4 ) * ( ( y3 - y2 ) / ( x3 - x2 )  +  ( y4 - y3 ) / ( x4 - x3 )  -  ( y2 - y1 ) / ( x2 - x1 ) )


	Or using Lagrange polynomial

		y2	= y3 * ( x2 - x4 ) * ( x2 - x5 )  /  ( ( x3 - x4 ) * ( x3 - x5 ) )	+
			  y4 * ( x2 - x3 ) * ( x2 - x5 )  /  ( ( x4 - x3 ) * ( x4 - x5 ) )	+
			  y5 * ( x2 - x3 ) * ( x2 - x4 )  /  ( ( x5 - x3 ) * ( x5 - x4 ) )

		y1	=
			y2 * ( x1 - x3 ) * ( x1 - x4 )	/
			   ( ( x2 - x3 ) * ( x2 - x4 ) )	+

			y3 * ( x1 - x2 ) * ( x1 - x4 )	/
			   ( ( x3 - x2 ) * ( x3 - x4 ) )	+

			y4 * ( x1 - x2 ) * ( x1 - x3 )	/
			   ( ( x4 - x2 ) * ( x4 - x3 ) )
*/
//	----------------------------------------

/*	Akima interpolation and smooth curve fitting

	Computes a curve coefficients for the interval where x lies : x1 <= x <= x2

	Method requires at least 5 data points, err might also indicate that x is out of bounds
*/
func Akima_interval_curve  ( data_points  * [][] float64, x  float64 )	( interval_curve  * Akima_curve, err  error )	{

	var points_len	= uint ( len ( * data_points ) )

	if	points_len < 5	||

//		Range Error
		x < ( * data_points ) [ 0 ][ 0 ]	||
		x > ( * data_points ) [ points_len -1 ][ 0 ]	{

		err	= math_tools.Arg_range_error ()
		return	interval_curve, err
	}

	var (
//		Interval points where x lies : x1 <= x <= x2
		y1, y2	float64
		x1, x2	float64

		i	uint

//		Slopes ( 5 point ) of interval points
		t1, t2	float64
	)

//	[ Double side search ] Find the control points where x belong

	for	i_x1, i_x2 := uint ( points_len -2 ), uint ( 1 )
		i_x2 < points_len
		i_x1, i_x2 = i_x1 -1, i_x2 +1	{

		i	= i_x2
		x2	= ( * data_points ) [ i_x2 ][ 0 ]

		if	x <= x2	{	break	}

		x1	= ( * data_points ) [ i_x1 ][ 0 ]

		if	x >= x1	{

			i	= i_x1 +1
			break
		}
	}

	x2	= ( * data_points ) [ i ][ 0 ]
	y2	= ( * data_points ) [ i ][ 1 ]
	t2	= slope_five_point ( data_points, points_len, i )

	i --
	x1	= ( * data_points ) [ i ][ 0 ]
	y1	= ( * data_points ) [ i ][ 1 ]
	t1	= slope_five_point ( data_points, points_len, i )

//	----------------------------------------
//	See section 2 :		y = p0 +  p1( x - x1 )  +  p2( x - x1 )^2  +  p3( x - x1 )^3

	interval_curve	= & Akima_curve {
		X1	: x1,	X2	: x2,
		T1	: t1,	T2	: t2,	Index_x1	: i,
	}
	interval_curve.set_coefficients ( y1, y2 )

	return	interval_curve, err
}

//	----------------------------------------

/*	Smoothing curve on the interval
		x1 <=  x  <= x2
*/
type Akima_curve struct {

	Index_x1	uint
	X1, X2, T1, T2	float64

//	Coefficients for a polynomial y
	p0, p2, p3	float64
}


/*	Calculates a point on a given interval	x1 <=  x  <= x2  ( bounds are not checked )

	By the formula

		y = p0 +  p1( x - x1 )  +  p2( x - x1 )^2  +  p3( x - x1 )^3

	where
		p0 = y1
		p1 = t1
		p2 = [ 3( y2 - y1 ) / ( x2 - x1 )  - 2* t1  -  t2 ]  /  ( x2 - x1 )
		p3 = [ t1 + t2 - 2( y2 - y1 ) / ( x2 - x1 ) ]  /  ( x2 - x1 )^2

		t1 and t2 are 5 point slopes of the two interval points
*/
func ( self  * Akima_curve )	Point  ( x  float64 )		float64	{
	var (
		x_minus_x1      = x  -  self.X1
		x_minus_x1_pow2	= x_minus_x1 * x_minus_x1
	)
	return	self.p0  +  self.p3 * x_minus_x1 * x_minus_x1_pow2  +
			self.T1 * x_minus_x1  +  self.p2 * x_minus_x1_pow2
}

func ( self  * Akima_curve )	Equal  ( other  * Akima_curve )		bool	{

	return	self.X1 == other.X1	&& self.X2 == other.X2	&&
			self.T1 == other.T1	&& self.T2 == other.T2	&&
			self.p2 == other.p2	&& self.p3 == other.p3
}

/*	Computes a curve for the next interval ( Index_x1 +1 )

	Uses property : next.X1, next.T1 = self.X2, self.T2

	Argument tells that the method :
		- is not a getter for a hidden variable ( computations are made on every call )
		- depends on the data points ( so they should not be changed between the calls )

	Returns nil if this interval is the last one
*/
func ( self  * Akima_curve )	Next_curve  ( data_points  * [][] float64 )		( next  * Akima_curve )	{

	var points_len	= uint ( len ( * data_points ) )
	var new_i_x2	= self.Index_x1 +2

	if	new_i_x2 >= points_len	{	return	nil	}

	next	= new ( Akima_curve )
	next.Index_x1	= self.Index_x1 +1

	next.X1	= self.X2
	next.T1	= self.T2
	next.X2	= ( * data_points ) [ new_i_x2 ][ 0 ]
	next.T2	= slope_five_point ( data_points, points_len, new_i_x2 )

	var y1	= ( * data_points ) [ next.Index_x1 ][ 1 ]
	var y2	= ( * data_points ) [ new_i_x2 ][ 1 ]
	next.set_coefficients ( y1, y2 )

	return	next
}

/*	Computes a curve for the prev interval ( Index_x1 -1 )

	Uses property : prev.X2, prev.T2 = self.X1, self.T1

	Returns nil if this interval is the first one
*/
func ( self  * Akima_curve )	Prev_curve  ( data_points  * [][] float64 )		( prev  * Akima_curve )	{

	var points_len	= uint ( len ( * data_points ) )
//	Care : Uint 0 -1 ~ undefined
	if	self.Index_x1 == 0	{	return	nil	}

	prev	= new ( Akima_curve )
	prev.Index_x1	= self.Index_x1 -1

	prev.X1	= ( * data_points ) [ prev.Index_x1 ][ 0 ]
	prev.T1	= slope_five_point ( data_points, points_len, prev.Index_x1 )
	prev.X2	= self.X1
	prev.T2	= self.T1

	var y1	= ( * data_points ) [ prev.Index_x1 ][ 1 ]
	var y2	= ( * data_points ) [ self.Index_x1 ][ 1 ]
	prev.set_coefficients ( y1, y2 )

	return	prev
}

func ( self  * Akima_curve )	set_coefficients ( y1, y2  float64 )	{
	var (
		x2_minus_x1	= self.X2 - self.X1
		y2_minus_y1	= y2 - y1

		m_slope12	= y2_minus_y1 / x2_minus_x1

		t1_plus_t2	= self.T1 + self.T2
	)
	self.p0	= y1
	self.p2	= ( 3 * m_slope12  -  self.T1 - t1_plus_t2 )  /  x2_minus_x1
	self.p3	= ( t1_plus_t2 - 2 * m_slope12 )  /  ( x2_minus_x1 * x2_minus_x1 )
}


func slope_five_point  ( data_points  * [][] float64, points_len, i  uint )		float64	{
	var (
		x1, x2, x3, x4, x5	float64
		y1, y2, y3, y4, y5	float64

//		2 point Slopes
		m12, m23, m34, m45	float64
	)

	x3	= ( * data_points ) [ i ][ 0 ]
	y3	= ( * data_points ) [ i ][ 1 ]

	if	i == 0 || i == 1	{

		x4	= ( * data_points ) [ i +1 ][ 0 ]
		y4	= ( * data_points ) [ i +1 ][ 1 ]

		x5	= ( * data_points ) [ i +2 ][ 0 ]
		y5	= ( * data_points ) [ i +2 ][ 1 ]

		m34	= ( y4 - y3 ) / ( x4 - x3 )
		m45	= ( y5 - y4 ) / ( x5 - x4 )

	} else

	if	i +1 == points_len || i +2 == points_len	{

		x1	= ( * data_points ) [ i -2 ][ 0 ]
		y1	= ( * data_points ) [ i -2 ][ 1 ]

		x2	= ( * data_points ) [ i -1 ][ 0 ]
		y2	= ( * data_points ) [ i -1 ][ 1 ]

		m12	= ( y2 - y1 ) / ( x2 - x1 )
        m23	= ( y3 - y2 ) / ( x3 - x2 )
	}

	switch	i	{

		case	0 :

			x2	= x3 + x4 - x5
			y2	= y3 - ( x3 - x2 ) * ( 2 * m34  - m45 )

			m23	= ( y3 - y2 ) / ( x3 - x2 )

			x1	= 2 * x3 - x5
			y1	= y2 - ( x2 - x1 ) * ( 2 * m23  - m34 )

			m12	= ( y2 - y1 ) / ( x2 - x1 )
			break

		case	1 :

			x2	= ( * data_points ) [ i -1 ][ 0 ]
			y2	= ( * data_points ) [ i -1 ][ 1 ]

			m23	= ( y3 - y2 ) / ( x3 - x2 )

			x1	= 2 * x3 - x5
			y1	= y2 - ( x2 - x1 ) * ( 2 * m23  - m34 )

			m12	= ( y2 - y1 ) / ( x2 - x1 )
			break

		case	points_len -2 :

			x4	= ( * data_points ) [ i +1 ][ 0 ]
			y4	= ( * data_points ) [ i +1 ][ 1 ]

			m34	= ( y4 - y3 ) / ( x4 - x3 )

			x5	= 2 * x3 - x1
			y5	= y4 + ( x5 - x4 ) * ( 2 * m34  - m23 )

			m45	= ( y5 - y4 ) / ( x5 - x4 )
			break

		case	points_len -1 :

			x4	= x3 + x2 - x1
			y4	= y3 + ( x4 - x3 ) * ( 2 * m23 - m12 )

			m34	= ( y4 - y3 ) / ( x4 - x3 )

			x5	= 2 * x3 - x1
			y5	= y4 + ( x5 - x4 ) * ( 2 * m34  - m23 )

			m45	= ( y5 - y4 ) / ( x5 - x4 )
			break

		default	:

			x1, x2	= ( * data_points ) [ i -2 ][ 0 ], ( * data_points ) [ i -1 ][ 0 ]
			y1, y2	= ( * data_points ) [ i -2 ][ 1 ], ( * data_points ) [ i -1 ][ 1 ]

			x5, x4	= ( * data_points ) [ i +2 ][ 0 ], ( * data_points ) [ i +1 ][ 0 ]
			y5, y4	= ( * data_points ) [ i +2 ][ 1 ], ( * data_points ) [ i +1 ][ 1 ]

			m12	= ( y2 - y1 ) / ( x2 - x1 )
			m23	= ( y3 - y2 ) / ( x3 - x2 )
			m34	= ( y4 - y3 ) / ( x4 - x3 )
			m45	= ( y5 - y4 ) / ( x5 - x4 )
			break
	}

	if	m12 == m23	&& m34 == m45	{

		return	( m23 + m34 ) / 2.0
	} else {

		var (
			m45_minus_m34	= math.Abs ( m45 - m34 )
			m23_minus_m12	= math.Abs ( m23 - m12 )
		)
		return	( m45_minus_m34 * m23  +  m23_minus_m12 * m34 )  /  ( m45_minus_m34 + m23_minus_m12 )
	}
}