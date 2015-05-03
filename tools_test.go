//	Copyright (c) 2013, Bogdan S.
//	Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package	math_tools

import	(
	"math/big"
	"testing"
)


var Pascals_triangle	= [][] uint {
	[] uint { 1 },
	[] uint { 1, 1 },
	[] uint { 1, 2, 1 },
	[] uint { 1, 3, 3, 1 },
	[] uint { 1, 4, 6, 4, 1 },
	[] uint { 1, 5, 10, 10, 5, 1 },
	[] uint { 1, 6, 15, 20, 15, 6, 1 },
	[] uint { 1, 7, 21, 35, 35, 21, 7, 1 },
	[] uint { 1, 8, 28, 56, 70, 56, 28, 8, 1 },
	[] uint { 1, 9, 36, 84, 126, 126, 84, 36, 9, 1 },
	[] uint { 1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1 },
	[] uint { 1, 11, 55, 165, 330, 462, 462, 330, 165, 55, 11, 1 },
	[] uint { 1, 12, 66, 220, 495, 792, 924, 792, 495, 220, 66, 12, 1 },
	[] uint { 1, 13, 78, 286, 715, 1287, 1716, 1716, 1287, 715, 286, 78, 13, 1 },
}

func Test_Binomial_coefficient ( t * testing.T )	{

	t.Parallel ()

	var cases	= [...] uint {
		 7, 2, 21,
		 7, 5, 21,
		20, 3, 1140,
		15, 1, 15,
		15, 2, 105,
		 19, 4, 3876,
		 20, 4, 4845,
		 20, 5, 15504,
		 24, 14, 1961256,
		 53, 13, 841392966470,

		 1024, 2, 523776,
		 1999, 2, 1997001,
		 2048, 2, 2096128,
		 4001, 2, 8002000,
		 5003, 2, 12512503,
		 5333, 2, 14217778,
		 6007, 2, 18039021,
		 7877, 2, 31019626,

		 1024, 3, 178433024,
		 1999, 3, 1329336999,
		 2048, 3, 1429559296,
		 4001, 3, 10666666000,
		 5003, 3, 20858342501,
		 5333, 3, 25264991506,
		 6007, 3, 36108107035,
		 7877, 3, 81426518250,

		 1024, 4, 45545029376,
		 1999, 4, 663339162501,
		 2048, 4, 730862190080,
		 4001, 4, 10661332667000,
		 5003, 4, 26072928126250,
		 5333, 4, 33665601181745,
		 6007, 4, 54198268659535,
		 7877, 4, 160288101175125,

		 14, 5, 2002,
		 15, 5, 3003,
		 16, 5, 4368,
		 17, 5, 6188,
		 18, 5, 8568,
		 19, 5, 11628,

		 1024, 5, 9291185992704,
		 1999, 5, 264672325837899,
		 2048, 5, 298776463304704,
		 4001, 5, 8522669333999800,
		 5003, 5, 26067713540624750,
		 5333, 5, 35880797739503821,
		 6007, 5, 65070441352637721,
		 7500, 5, 197490357398439000,
		 7877, 5, 252389644110351825,
	}

	var (
		result		uint
		result64	uint64
		result_big	= big.NewInt ( 0 )
	)

	for	i := uint ( 7000 ) ; i < 7250 ; i ++	{

		if	uint64 ( Binomial_coefficient ( i, 5 ) ) !=
			Binomial_coeff_big ( uint64 ( i ), 5 )	{

			t.Errorf (
				"Uint overflow : C ( %d, 5 ) expected = %d, got : %d",
				i,
				Binomial_coeff_big ( uint64 ( i ), 5 ),
				Binomial_coefficient ( i, 5 ),
			)
			t.FailNow ()
		}
	}


	for	i, cases_num := 0, len ( cases ) ; i < cases_num ; i += 3	{

		result	= Binomial_coefficient ( cases [ i ], cases [ i +1 ] )

		if	result != cases [ i +2 ]	{
			t.Errorf (
				"C ( %d, %d ) expected = %d, got : %d",
				cases [ i ], cases [ i +1 ],
				cases [ i +2 ], result,
			)
			t.FailNow ()
		}

		result64	= Binomial_coeff_big ( uint64 ( cases [ i ] ), uint64 ( cases [ i +1 ] ) )

		if	result64 != uint64 ( cases [ i +2 ] )	{
			t.Errorf (
				"C ( %d, %d ) expected = %d, got : %d",
				cases [ i ], cases [ i +1 ],
				cases [ i +2 ], result64,
			)
			t.FailNow ()
		}


		result_big.Binomial ( int64 ( cases [ i ] ), int64 ( cases [ i +1 ] ) )


		if	result_big.Cmp (
				big.NewInt ( int64 ( cases [ i +2 ] ) ),
			) != 0	{

			t.Errorf (
				"C ( %d, %d ) expected = %d, got : %d",
				cases [ i ], cases [ i +1 ],
				cases [ i +2 ], result_big,
			)
			t.FailNow ()
		}
	}


	for	row, rows_num := uint ( 0 ), uint ( len ( Pascals_triangle ) ) ; row < rows_num ; row ++	{

		for	col := uint ( 0 ) ; col <= row ; col ++	{

			result	= Binomial_coefficient ( row, col )

			if	result != Pascals_triangle [ row ][ col ]	{

				t.Errorf (
					"C ( %d, %d ) expected = %d, got : %d",
					row, col,
					Pascals_triangle [ row ][ col ],
					result,
				)
				t.FailNow ()
			}

			result64	= Binomial_coeff_big ( uint64 ( row ), uint64 ( col ) )

			if	result64 != uint64 ( Pascals_triangle [ row ][ col ] )	{

				t.Errorf (
					"C ( %d, %d ) expected = %d, got : %d",
					row, col,
					Pascals_triangle [ row ][ col ],
					result64,
				)
				t.FailNow ()
			}


			result_big.Binomial ( int64 ( row ), int64 ( col ) )

			if	result_big.Cmp (
					big.NewInt ( int64 ( Pascals_triangle [ row ][ col ] ) ),
				) != 0	{

				t.Errorf (
					"C ( %d, %d ) expected = %d, got : %d",
					row, col,
					Pascals_triangle [ row ][ col ],
					result_big,
				)
				t.FailNow ()
			}
		}
	}
}

func Benchmark_Binomial_coefficient ( b * testing.B )	{

	for	row, size := uint ( 0 ), uint ( b.N ) ; row < size ; row ++	{

			Binomial_coefficient ( row, 2 )
			Binomial_coefficient ( row, 3 )
			Binomial_coefficient ( row, 4 )
			Binomial_coefficient ( row, 5 )
			Binomial_coefficient ( row, 6 )
			Binomial_coefficient ( row, 7 )
			Binomial_coefficient ( row, 8 )
			Binomial_coefficient ( row, 9 )
			Binomial_coefficient ( row, 10 )
			Binomial_coefficient ( row, row / 2 )
	}
}

func Benchmark_Binomial_coeff_big ( b * testing.B )	{

	for	row, size := uint64 ( 0 ), uint64 ( b.N ) ; row < size ; row ++	{

			Binomial_coeff_big ( row, 2 )
			Binomial_coeff_big ( row, 3 )
			Binomial_coeff_big ( row, 4 )
			Binomial_coeff_big ( row, 5 )
			Binomial_coeff_big ( row, 6 )
			Binomial_coeff_big ( row, 7 )
			Binomial_coeff_big ( row, 8 )
			Binomial_coeff_big ( row, 9 )
			Binomial_coeff_big ( row, 10 )
			Binomial_coeff_big ( row, row / 2 )
	}
}

func Benchmark_Binomial_internal ( b * testing.B )	{

	var result	= big.NewInt ( 0 )

	for	row, size := int64 ( 0 ), int64 ( b.N ) ; row < size ; row ++	{

			result.Binomial ( row, 2 )
			result.Binomial ( row, 3 )
			result.Binomial ( row, 4 )
			result.Binomial ( row, 5 )
			result.Binomial ( row, 6 )
			result.Binomial ( row, 7 )
			result.Binomial ( row, 8 )
			result.Binomial ( row, 9 )
			result.Binomial ( row, 10 )
			result.Binomial ( row, row / 2 )
	}
}

func Test_Bit_swap ( t * testing.T )	{

	t.Parallel ()

	var (
		cases	= [] int {
			21, 113,
			-21, 11,
			-21, -11,
		}

		a, b	int
	)

	for	i, size := 0, len ( cases ) ; i < size ; i += 2	{

		a, b	= Bit_swap ( cases [ i ], cases [ i +1 ] )

		if	a != cases [ i +1 ]	|| b != cases [ i ]	{

			t.Error ( "Bit swap error" )
		}
	}
}

func Test_Bit_min ( t * testing.T )	{

	t.Parallel ()

	if	Bit_min ( 21, 113 ) != 21	{
		t.Error ( "Bit minimum error" )
		t.FailNow ()
	}

	if	Bit_min ( 21, 21 ) != 21	{
		t.Error ( "Bit minimum error" )
		t.FailNow ()
	}

	if	Bit_min ( 21, 11 ) != 11	{
		t.Error ( "Bit minimum error" )
		t.FailNow ()
	}
}

func Test_Bit_max ( t * testing.T )	{

	t.Parallel ()

	if	Bit_max ( 21, 113 ) != 113	{
		t.Error ( "Bit max error" )
		t.FailNow ()
	}

	if	Bit_max ( 21, 21 ) != 21	{
		t.Error ( "Bit max error" )
		t.FailNow ()
	}

	if	Bit_max ( 11, 21 ) != 21	{
		t.Error ( "Bit max error" )
		t.FailNow ()
	}
}

func Test_BitAvgFloor ( t * testing.T )	{

	t.Parallel ()

	var (
		cases	= [] int {
			 4,  2,  3,
			-4, -2, -3,
			-1,  1,  0,

//			FLoor cases
			 4, 5,  4,
			-1, 0, -1,
			-1,-2, -2,
		}

		result	int
	)

	for	i, size := 0, len ( cases ) ; i < size ; i += 3	{
		result	= BitAvgFloor( cases [ i ], cases [ i +1 ] )

		if	result != cases [ i +2 ]	{
			t.Errorf ( "Expected %v, got %v", cases [ i +2 ], result )
		}
	}
}

func Test_BitAvgCeil ( t * testing.T )	{

	t.Parallel ()

	var (
		cases	= [] int {
			 4,  2,  3,
			-4, -2, -3,
			-1,  1,  0,

//			FLoor cases
			 4, 5,  5,
			-1, 0, 0,
			-1,-2, -1,
		}

		result	int
	)

	for	i, size := 0, len ( cases ) ; i < size ; i += 3	{
		result	= BitAvgCeil( cases [ i ], cases [ i +1 ] )

		if	result != cases [ i +2 ]	{
			t.Errorf ( "Expected %v, got %v", cases [ i +2 ], result )
		}
	}
}

func Test_BitAbs ( t * testing.T )	{

	t.Parallel ()

	if	BitAbs ( -1 ) != 1 || BitAbs ( 5 ) != 5 || BitAbs ( -0 ) != 0	{
		t.Error( "Bit absolute int error" )
		t.FailNow()
	}
	if	BitAbs32 ( -1 ) != 1 || BitAbs32 ( 5 ) != 5 || BitAbs32 ( -0 ) != 0	{
		t.Error( "Bit absolute int32 error" )
		t.FailNow()
	}
	if	BitAbs64 ( -1 ) != 1 || BitAbs64 ( 5 ) != 5 || BitAbs64 ( -0 ) != 0	{
		t.Error( "Bit absolute int64 error" )
		t.FailNow()
	}
}