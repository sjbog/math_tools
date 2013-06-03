//	Copyright (c) 2013, Bogdan S.
//	Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package	interpolation_test

import	(
	"fmt"
	"testing"
	"math"
	. "math_tools/interpolation"
)


/*	Akima spline fitting of SIN(X):
*                                                      *
*   X       SIN(X)       AKIMA INTERPOLATION    ERROR  *
* ---------------------------------------------------- *
* 0.00    0.0000000          0.0000000       0.0000000 *
* 0.05    0.0499792          0.0500402      -0.0000610 *
* 0.10    0.0998334          0.0998435      -0.0000101 *
* 0.15    0.1494381          0.1494310       0.0000072 *
* 0.20    0.1986693          0.1986458       0.0000235 *
* 0.25    0.2474040          0.2474157      -0.0000118 *
* 0.30    0.2955202          0.2955218      -0.0000016 *
* 0.35    0.3428978          0.3428916       0.0000062 *
* 0.40    0.3894183          0.3894265      -0.0000081 *
* 0.45    0.4349655          0.4349655       0.0000000 *
* 0.50    0.4794255          0.4794204       0.0000051 *
* 0.55    0.5226872          0.5226893      -0.0000021 *
* 0.60    0.5646425          0.5646493      -0.0000068 *
* 0.65    0.6051864          0.6051821       0.0000043 *
* 0.70    0.6442177          0.6442141       0.0000035 *
* 0.75    0.6816388          0.6816405      -0.0000017 *
* 0.80    0.7173561          0.7173609      -0.0000048 *
* 0.85    0.7512804          0.7512811      -0.0000007 *
* 0.90    0.7833269          0.7833267       0.0000002 *
* 0.95    0.8134155          0.8134114       0.0000041 *
* ---------------------------------------------------- *
*/
func Test_Akima_spline ( t  * testing.T )	{

	t.Parallel ()

	var (
		control_points_sin	= [][] float64 {
			[] float64 { 0.000, 0.000 },
			[] float64 { 0.125, 0.12467473 },
			[] float64 { 0.217, 0.21530095 },
			[] float64 { 0.299, 0.29456472 },
			[] float64 { 0.376, 0.36720285 },
			[] float64 { 0.450, 0.43496553 },
			[] float64 { 0.520, 0.49688014 },
			[] float64 { 0.589, 0.55552980 },
			[] float64 { 0.656, 0.60995199 },
			[] float64 { 0.721, 0.66013615 },
			[] float64 { 0.7853981634, 0.7071067812 },
			[] float64 { 0.849, 0.75062005 },
			[] float64 { 0.911, 0.79011709 },
			[] float64 { 0.972, 0.82601466 },
		}

		results_sin	= [] float64 {
			0.00,	0.0000000,
            0.05,	0.0500402,
            0.10,	0.0998435,
            0.15,	0.1494310,
            0.20,	0.1986458,
            0.25,	0.2474157,
            0.30,	0.2955218,
            0.35,	0.3428916,
            0.40,	0.3894265,
            0.45,	0.4349655,
            0.50,	0.4794204,
            0.55,	0.5226893,
            0.60,	0.5646493,
            0.65,	0.6051821,
            0.70,	0.6442141,
            0.75,	0.6816405,
            0.80,	0.7173609,
            0.85,	0.7512811,
            0.90,	0.7833267,
            0.95,	0.8134114,
		}

		control_points	= [][] float64 {}

		ctrl_points_step	= math.Pi / 35.0
		ctrl_points_start	= 0.0
		ctrl_points_stop	= 2 * math.Pi

		ctrl_points_func	= math.Sin

		test_points_step	= 0.05
		test_points_abs_err	= 0.00005
		test_points_stop	float64

		x, result, expected, measured_err	float64

		curve	* Akima_curve
		err		error
	)


//	Hand test of sin

	for	i, size := 0, len ( results_sin ) ; i < size ; i += 2	{

		x, expected	= results_sin [ i ], results_sin [ i +1 ]

		curve, err	:= Akima_interval_curve ( & control_points_sin, x )

		if	err != nil || curve == nil	{
			t.Error ( err )
			continue
		}

		result	= curve.Point ( x )

		if	fmt.Sprintf ( "%9.7f", result )	!=
			fmt.Sprintf ( "%9.7f", expected )	{

			t.Errorf ( "Error : Sin ( %2.2f ) = %9.7f, expected %9.7f, got %9.7f\n",
				x, ctrl_points_func ( x ), expected, result,
			)
			t.FailNow ()
		}
	}


//	Fill control_points

	for	x = ctrl_points_start ; x <= ctrl_points_stop ; x += ctrl_points_step	{

		control_points	= append ( control_points, [] float64 {
			x, ctrl_points_func ( x ),
		})
	}

	test_points_stop	= control_points [ len ( control_points ) -1 ][ 0 ]

	curve, err	= Akima_interval_curve ( & control_points, ctrl_points_start )

	if	err != nil || curve == nil	{

		t.Error ( err, curve )
		return
	}

//	Test of Sin ( x ) with absolute error

	for	x = ctrl_points_start ; x <= test_points_stop ; x += test_points_step	{

		expected	= ctrl_points_func ( x )

//		step might be bigger than the interval
		for	curve.X2 < x	{

			curve	= curve.Next_curve ( & control_points )

			if	curve == nil	{
				t.Error ( "Out of range" )
				t.FailNow ()
				return
			}
		}

		result	= curve.Point ( x )
		measured_err	= math.Abs ( expected - result )

		if	test_points_abs_err < measured_err	{

			t.Errorf (
				"Sin ( %2.2f ) = %9.7f, with step %v, abs_error %9.7f\n" +
				"Got %9.7f, measured_err = %9.7f\n",
				x, expected,
				test_points_step, test_points_abs_err,
				result, measured_err,
			)
			t.FailNow ()
		}
	}
}


func Test_Akima_struct ( t  * testing.T )	{

	t.Parallel ()

	var (	// y = x
		control_points	= [][] float64 {
			[] float64 { 1.0, 1.0 },
			[] float64 { 2.0, 2.0 },
			[] float64 { 3.0, 3.0 },
			[] float64 { 4.0, 4.0 },
			[] float64 { 5.0, 5.0 },
		}

		result, expected	* Akima_curve
		err		error
	)

//	Test error
	result, err	= Akima_interval_curve ( & control_points, -1.0 )

	if	err == nil	{
		t.Error ( "Argument is out of range, but there is no error ; result : ", result )
	}

//	Test Next_curve
	result, err	= Akima_interval_curve ( & control_points, 1.5 )

	if	err != nil	|| result == nil	{
		t.Error ( "Curve ", err, result )
		t.FailNow ()
	}

	result	= result.Next_curve ( & control_points )

	if	result == nil	{
		t.Error ( "Curve should not be nil" )
		t.FailNow ()
	}

	expected, err	= Akima_interval_curve ( & control_points, 2.5 )

	if	err != nil	|| result == nil	{
		t.Error ( "Expected curve ", err, result )
		t.FailNow ()
	}

	if	! result.Equal ( expected )	{

		t.Errorf ( "Expected curve : %v [ %v, %v ] [ %v, %v ],\tgot %v [ %v, %v ] [ %v, %v ]\n%v\n%v",
			expected.Index_x1, expected.X1, expected.X2, expected.T1, expected.T2,
			result.Index_x1, result.X1, result.X2, result.T1, result.T2,
			expected, result,
		)
	}

//	Next == nil
	result, err	= Akima_interval_curve ( & control_points, control_points [ len ( control_points ) -1 ][ 0 ] )

	if	err != nil	|| result == nil	{
		t.Error ( "Curve ", err, result )
		t.FailNow ()
	}

	result	= result.Next_curve ( & control_points )

	if	result != nil	{
		t.Error ( "After last interval, curve should be nil", result )
	}

//	Prev == nil
	result, err	= Akima_interval_curve ( & control_points, control_points [ 0 ][ 0 ] )

	if	err != nil	|| result == nil	{
		t.Error ( "Curve ", err, result )
		t.FailNow ()
	}

	result	= result.Prev_curve ( & control_points )

	if	result != nil	{
		t.Error ( "Before first interval, curve should be nil", result )
	}
}


func ExampleAkima_interval_curve_1  ()	{
	var (
//		y = x , method requires at least 5 points

		control_points	= [][] float64 {
			[] float64 { 10.0, 10.0 },
			[] float64 { 20.0, 20.0 },
			[] float64 { 30.0, 30.0 },
			[] float64 { 40.0, 40.0 },
			[] float64 { 50.0, 50.0 },
		}

		size	= len ( control_points )
		step	= 16.0

		curve	* Akima_curve
		err		error
	)


	fmt.Println ( "\nForward" )

//	Get the initial smooth interval
	curve, err	= Akima_interval_curve ( & control_points, control_points [ 0 ][ 0 ] )

	if	err != nil || curve == nil	{

		fmt.Println ( "Error : couldn't get the initial interval", err, curve )
		return
	}

	for	x, x_end := control_points [ 0 ][ 0 ], control_points [ size -1 ][ 0 ]
		x <= x_end
		x += step	{

//		step might be bigger than the interval
		for	curve.X2 < x	{

			curve	= curve.Next_curve ( & control_points )

			if	curve == nil	{
				break
//				fmt.Println ( "Error : Out of range" )
//				return
			}
		}

		fmt.Printf ( "y ( %v ) = %v, got %v\n", x, x, curve.Point ( x ) )
	}


	fmt.Println ( "\nBackward" )
	step	= 15

//	Get the last smooth interval
	curve, err	= Akima_interval_curve ( & control_points, control_points [ size -1 ][ 0 ] )

	if	err != nil || curve == nil	{

		fmt.Println ( "Error : couldn't get the last interval", err, curve )
		return
	}

	for	x_start, x := control_points [ 0 ][ 0 ], control_points [ size -1 ][ 0 ]
		x_start <= x
		x -= step	{

//		step might be bigger than the interval
		for	x < curve.X1	{

			curve	= curve.Prev_curve ( & control_points )

			if	curve == nil	{	break	}
		}

		fmt.Printf ( "y ( %v ) = %v, got %v\n", x, x, curve.Point ( x ) )
	}

//	Output:
//
//Forward
//y ( 10 ) = 10, got 10
//y ( 26 ) = 26, got 26
//y ( 42 ) = 42, got 42
//
//Backward
//y ( 50 ) = 50, got 50
//y ( 35 ) = 35, got 35
//y ( 20 ) = 20, got 20
}