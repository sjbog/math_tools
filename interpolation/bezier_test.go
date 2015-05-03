//	Copyright (c) 2013, Bogdan S.
//	Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package	interpolation

import	(
	"fmt"
	"testing"
)


func Test_Bernstein_basis ( t * testing.T )	{

	t.Parallel ()

	var (
		cases	= [...] uint {
    		 0, 0, 50, 10000,
    		 1, 0, 50, 5000,		1, 1, 50, 5000,
    		 2, 0, 50, 2500,		2, 1, 50, 5000,		2, 2, 50, 2500,
    		 3, 0, 50, 1250,		3, 1, 50, 3750,		3, 2, 50, 3750,		3, 3, 50, 1250,
    		 3, 0, 0, 10000,		3, 1, 00, 0000,		3, 2, 00, 0000,		3, 3, 00, 0000,
    		 3, 0, 100, 000,		3, 1, 100, 000,		3, 2, 100, 000,		3, 3, 100, 10000,
    		 4, 3, 50, 2500,		4, 2, 00, 0000,		4, 2, 50, 3750,		4, 2, 100, 000,
    		 4, 2, 10,  486,		4, 2, 20, 1536,		4, 2, 30, 2646,		4, 3, 40, 1536,		5, 4, 40, 768,
    	}

		divisor_offset, divisor_result	= 100.0, 10000.0

		result	float64
	)

	for	i, cases_num := 0, len ( cases ) ; i < cases_num ; i += 4	{

		result	= Bernstein_basis ( cases [ i ], cases [ i +1 ], float64 ( cases [ i +2 ] ) / divisor_offset )

		if	fmt.Sprintf ( "%.4f", result ) !=
			fmt.Sprintf ( "%.4f", float64 ( cases [ i +3 ] ) / divisor_result )	{

			t.Errorf (
				"Bernstein_basis ( %d, %d, %.4f ) expected = %.4f, got : %.4f",
				cases [ i ], cases [ i +1 ],

				float64 ( cases [ i +2 ] ) / divisor_offset,
				float64 ( cases [ i +3 ] ) / divisor_result,
				result,
			)
			t.FailNow ()
		}
	}
}

func Test_Bezier_point ( t * testing.T )	{

	t.Parallel ()

	type Bezier_test struct	{
		points	[][] float64
		result	map [ float64 ] [] float64
	}

	var (
		cases	= [...] Bezier_test {

//			3 Points - upper half of a cirle ( [ 1.0, 0.0 ], 1 )
			Bezier_test {

				points	: [][] float64 {
					[] float64 { 0.0, 0.0 },
					[] float64 { 1.0, 2.0 },
					[] float64 { 2.0, 0.0 },
				},

				result	: map [ float64 ] [] float64 {
					0.0	: [] float64 { 0.0, 0.0 },
					0.5	: [] float64 { 1.0, 1.0 },
					1.0	: [] float64 { 2.0, 0.0 },
				},
			},

//			Single point
			Bezier_test {

				points	: [][] float64 {
					[] float64 { 0.0, 0.0 },
				},

				result	: map [ float64 ] [] float64 {
					0.0	: [] float64 { 0.0, 0.0 },
					0.5	: [] float64 { 0.0, 0.0 },
					1.0	: [] float64 { 0.0, 0.0 },
				},
			},

//			2 Points
			Bezier_test {

				points	: [][] float64 {
					[] float64 { 0.0, 0.0 },
					[] float64 { 1.0, 0.0 },
				},

				result	: map [ float64 ] [] float64 {
					0.0	: [] float64 { 0.0, 0.0 },
					0.5	: [] float64 { 0.5, 0.0 },
					1.0	: [] float64 { 1.0, 0.0 },
				},
			},


//			3 Points - upper half of a cirle ( [ 0.0, 0.0 ], 1 )
			Bezier_test {

				points	: [][] float64 {
					[] float64 { -1.0, 0.0 },
					[] float64 { +0.0, 2.0 },
					[] float64 { +1.0, 0.0 },
				},

				result	: map [ float64 ] [] float64 {
					0.00	: [] float64 { -1.0, 0.0 },
					0.50	: [] float64 { +0.0, 1.0 },
					1.00	: [] float64 { +1.0, 0.0 },
				},
			},


			Bezier_test {

				points	: [][] float64 {
					[] float64 { 0.0, 0.0 },
					[] float64 { 0.0, 8.0 },
					[] float64 { 16.0, 8.0 },
					[] float64 { 16.0, 0.0 },
				},

				result	: map [ float64 ] [] float64 {
					0.00	: [] float64 { 0.0, 0.0 },
					0.25	: [] float64 { 2.5, 4.5 },
					0.50	: [] float64 { 8.0, 6.0 },
					0.75	: [] float64 { 13.5, 4.5 },
					1.00	: [] float64 { 16.0, 0.0 },
				},
			},


			Bezier_test {

				points	: [][] float64 {
					[] float64 { -2.0,  0.0 },
					[] float64 { -1.0,  2.0 },
					[] float64 {  0.0,  0.0 },
					[] float64 {  1.0, -2.0 },
					[] float64 {  2.0,  0.0 },
				},

				result	: map [ float64 ] [] float64 {
					0.00	: [] float64 { -2.0, 0.0 },
					0.25	: [] float64 { -1.0, 0.75 },
					0.50	: [] float64 { 0.0, 0.0 },
					0.75	: [] float64 { 1.0, -0.75 },
					1.00	: [] float64 { 2.0, 0.0 },
				},
			},
    	}

    	ok	= true
    	result	[] float64
	)

	for	ci, cases_num := 0, len ( cases ) ; ci < cases_num ; ci ++	{

		for	step, expected := range cases [ ci ].result	{

			result	= Bezier_point ( & cases [ ci ].points, step )


			for	degree, dsize := 0, len ( expected ) ; degree < dsize ; degree ++	{

				ok	= ok &&
					fmt.Sprintf ( "%.4f", expected [ degree ] ) ==
					fmt.Sprintf ( "%.4f", result [ degree ] )
			}

			if	! ok	{

				t.Errorf (
					"Step %.4f, expected = %v, got : %v.\tPoints %v",
					step,
					expected,
					result,
					cases [ ci ].points,
				)
				t.FailNow ()
			}
		}
	}

	result	= Bezier_point ( new ( [][] float64 ), 0.0 )

	if	result != nil	{
		t.Error ( "Arguments are wrong but there is no error, result : ", result )
	}
}