//	Copyright (c) 2013, Bogdan S.
//	Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

/*	Package provides interpolation and curves smoothing mathematical functions and tools.
*/
package	interpolation	;	import	( "math" ; "math_tools" )

/*	Calculates the point ( by offset percent ) from a Bézier curve

	Uses formulas for special cases : single point, linear, quadratic and cubic curves. See http://en.wikipedia.org/wiki/B%C3%A9zier_curve#Examination_of_cases

	Arguments

		0.0 <= offset <= 1.0	: offset is a percentage of a full Bezier curve, where 0.0 == P0 ( control point 0 ) and 1.0 == Pn

	Return

		result	: a point ( dimensions are the same as P0 ) from a Bezier curve on a specified offset ( percentage ) position

	One should use Goroutines for calculating curves
*/
func Bezier_point ( control_points  * [][] float64, offset  float64 )		( result  [] float64 )	{

	var points_len	= uint ( len ( * control_points ) )

	if	points_len == 0		{	return	result	}

//	Care : Inclusive N
	points_len --

	var (
		berstein_basis	float64
		degree			= len ( ( * control_points ) [ 0 ] )
		offset_complementary	= 1.0 - offset
	)

//	Fill resulting point from P0 :	P0 * ( 1 - t )

	berstein_basis	= math.Pow ( offset_complementary, float64 ( points_len ) )

	for	di := 0 ; di < degree ; di ++	{
		result	= append ( result,	( * control_points ) [ 0 ][ di ] * berstein_basis )
	}

//	Possible optimization : vector / matrix computation of point's dimensions


/*	# Classic solution :
	# We can improve it by going through Pascal's triangle row ( binomial ) sequentially

	for	point_i := uint ( 0 ) ; point_i <= points_len ; point_i ++	{

		berstein_basis	= Bernstein_basis ( points_len, point_i, offset )

		for	di := 0 ; di < degree ; di ++	{

			result [ di ]	+= berstein_basis * ( * control_points ) [ point_i ][ di ]
		}

	}*/

	switch	points_len	{

//		P0 only	: result = P0
		case 0 :	return	( * control_points ) [ 0 ]

//		Linear	: result = ( 1 - t ) * P0  +  t * P1
		case 1 :

			var P0, P1	float64

			for	di := 0 ; di < degree ; di ++	{

				P0, P1	= ( * control_points ) [ 0 ][ di ] ,	( * control_points ) [ 1 ][ di ]
				result [ di ]	= P0 * offset_complementary + P1 * offset
			}
			return

//		Quadratic	: result = P0 * ( 1 − t )^2  +  2 * P1 * ( 1 − t ) * t + P2 * t^2
		case 2 :

			var (
				P0, P1, P2		float64

				offset_Mul_complementary	= offset_complementary * offset
				offset_Pow2					= offset * offset
				offset_complementary_Pow2	= offset_complementary * offset_complementary
			)

			for	di := 0 ; di < degree ; di ++	{

				P0	= ( * control_points ) [ 0 ][ di ]
				P1	= ( * control_points ) [ 1 ][ di ]
				P2	= ( * control_points ) [ 2 ][ di ]

				result [ di ]	= P0 * offset_complementary_Pow2 +
					2 * P1 * offset_Mul_complementary + P2 * offset_Pow2
			}
			return

//		Cubic	: result = P0 * ( 1 − t )^3  +  3 * P1 * t * ( 1 − t )^2  +  3 * P2 * ( 1 − t ) * t^2  +  P3 * t^3
		case 3 :

			var (
				P0, P1, P2, P3		float64

				offset_Mul_complementary_Pow2	= offset * offset_complementary * offset_complementary
				offset_Pow2_Mul_complementary	= offset * offset * offset_complementary
				offset_Pow3					= offset * offset * offset
				offset_complementary_Pow3	= offset_complementary * offset_complementary * offset_complementary
			)

			for	di := 0 ; di < degree ; di ++	{

				P0	= ( * control_points ) [ 0 ][ di ]
				P1	= ( * control_points ) [ 1 ][ di ]
				P2	= ( * control_points ) [ 2 ][ di ]
				P3	= ( * control_points ) [ 3 ][ di ]

				result [ di ]	= P0 * offset_complementary_Pow3  +  P3 * offset_Pow3  +  3 *
					( P1 * offset_Mul_complementary_Pow2  +  P2 * offset_Pow2_Mul_complementary )
			}
			return
	}


	for	point_i, binomial_coeff := uint ( 1 ), uint ( 1 )
		point_i <= points_len
		point_i ++	{


		binomial_coeff	= binomial_coeff * ( points_len - point_i + 1 ) / point_i

		berstein_basis	= float64 ( binomial_coeff ) *
			math.Pow ( offset, float64 ( point_i ) ) *
			math.Pow ( offset_complementary, float64 ( points_len - point_i ) )


		for	di := 0 ; di < degree ; di ++	{

			result [ di ]	+= ( * control_points ) [ point_i ][ di ]  *  berstein_basis
		}
	}
	return
}


/*	Bernstein basis polynomials of degree n

	Polynomials on http://mathworld.wolfram.com/BernsteinPolynomial.html

	Arguments:
		control_points_num	: total number of control points
		control_point_index	: ( 0 based ) index of the control point of interest [ 0 <= control_point_index < control_points_num ]
		offset	: 0.0 <= offset <= 1.0 , see func Bezier_point
*/
func Bernstein_basis ( control_points_num, control_point_index uint, offset float64 )		float64	{

	if	control_points_num == 0	{	return 1.0	}

	if	control_points_num == control_point_index	{

		var t	= offset

		for	; control_point_index > 1 ; control_point_index --	{
			offset	= t * offset
		}

		return offset
	}

	switch	control_point_index	{

		case 0 :
			var t	= 1.0 - offset
			offset	= t

			for	; control_points_num > 1 ; control_points_num --	{
				offset	= t * offset
			}

			return	offset

		case 1 :
			var t	= 1.0 - offset
			offset	= float64 ( control_points_num ) * offset

			for	; control_points_num > 1 ; control_points_num --	{
				offset	= t * offset
			}

			return	offset

		case 2 :
			if	control_points_num == 3	{
				return 3.0 * ( 1.0 - offset ) * offset * offset
			}
			if	control_points_num == 4	{
				return 6.0 * ( 1.0 - offset ) * ( 1.0 - offset ) * offset * offset
			}
	}


	if	control_points_num == control_point_index +1	{

		var t	= float64 ( control_points_num ) * ( 1.0 - offset ) * offset

		for	; control_point_index > 1 ; control_point_index --	{
			t	= t * offset
		}

		return t
	}


	return	float64 ( math_tools.Binomial_coefficient ( control_points_num, control_point_index ) ) *
		math.Pow ( offset, float64 ( control_point_index ) ) *
		math.Pow ( ( 1.0 - offset ), float64 ( control_points_num - control_point_index ) )
}