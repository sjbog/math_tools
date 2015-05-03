//	Copyright (c) 2013, Bogdan S.
//	Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

/*	Package provides advanced mathematical functions and tools.

	Functions that start with "Bit" use bit operations ( hacks ) OR, AND, XOR
*/
package	math_tools

import	(
	"math/big"
	"strconv"
)


//	---------------------------
type  arg_range_error  struct {

	Msg	string
}


func ( self  arg_range_error )	Error ()		string	{
	return	self.Msg
}

/*	Returns an error "Error : passed argument(s) out of range"
*/
func Arg_range_error ()		error	{

	return	arg_range_error { "Error : passed argument(s) out of range" }
}
//	---------------------------


/*	Property :
		( x ^ y ) ^ y == x
		( x ^ y ) ^ x == y
*/
func Bit_swap ( x, y int )		( int, int )	{


	x = x ^ y
	y = x ^ y
	x = x ^ y

	return	x, y
}


/*	Doesn't work : Ideal bool error

		return y ^ ( ( x ^ y ) & -( x < y ) )
*/
func Bit_min ( x, y int )		int	{

	if	x <= y	{	return x	}

	return y
}


/*	Doesn't work : Ideal bool error

		return y ^ ( (x ^ y) & -(x > y) )
		or     x ^ ( (x ^ y) & -(x < y) )
		or ( x >= y ) * x + ( x < y ) * y
*/
func Bit_max ( x, y int )		int {


	if	x >= y	{	return x	}

	return y
}

/*	Average of Integers

	Computes average of 2 integers ( negative values are ok ot use ) without int overflow, but floors the result to int.

		return	math.Floor ( ( x + y ) / 2.0 )

	Formula

		( x & y ) +  // common bits remain
		(( x ^ y )   // xor : different bits = 1, same = 0
		     >> 1 )  // difference should be divided by 2

	Example

		x, y = 5, 7 = 101, 111
		    101		    101
		AND 111		XOR 111
		-------		-------
		    101		+   010 >> 1	= 101 + 001  = 110 = 6


	Source :	http://aggregate.org/MAGIC/#Average%20of%20Integers

	This is actually an extension of the "well known" fact that for binary integer values x and y, (x+y) equals ((x&y)+(x|y)) equals ((x^y)+2*(x&y)).

	Given two integer values x and y, the (floor of the) average normally would be computed by (x+y)/2; unfortunately, this can yield incorrect results due to overflow. A very sneaky alternative is to use (x&y)+((x^y)/2). If we are aware of the potential non-portability due to the fact that C does not specify if shifts are signed, this can be simplified to (x&y)+((x^y)>>1). In either case, the benefit is that this code sequence cannot overflow.

*/
func BitAvgFloor( x, y int )	int	{
	return ( x & y ) + ( ( x ^ y ) >> 1 )
}

func BitAvgCeil( x, y int )	int	{
	return ( x | y ) - (( x ^ y ) >> 1 )
}

func BitAbs( x int )	int {
	var y int = x >> ( strconv.IntSize -1 )
	return ( x + y ) ^ y
}

func BitAbs32( x int32 )	int32 {
	var y int32 = x >> 31
	return ( x + y ) ^ y
}

func BitAbs64( x int64 )	int64 {
	var y int64 = x >> 63
	return ( x + y ) ^ y
}


/*	Uses Pascal's triangle row for calculating "n choose k",
	in math formula - it is a cancellation of numerator and denominator

	Example : 5 choose 2	= 5 !  /  ( 2 ! * ( 5 - 2 ) ! )

	= 2 * 3 * 4 * 5  /  ( 2  *  3 * 4 * 5 )
	= 4 * 5  /  2
*/
func Binomial_coefficient ( n, k uint )		uint	{

	if	n <= 0	|| k <= 0	|| k >= n	{	return 1	}
	if	k == 1	|| k +1 == n			{	return	n	}

//	k	= Bit_min ( k, n - k )
	if	k > n - k	{
		k = n - k
	}

	switch	k	{

		case 2 :	return	n * ( n - 1 ) / 2
		case 3 :	return	n * ( n - 1 ) * ( n - 2 ) / 6
		case 4 :	return	n * ( n - 1 ) * ( n - 2 ) * ( n - 3 ) / 24
		case 5 :
			if	n < 7134	{
					return	n * ( n - 1 ) * ( n - 2 ) * ( n - 3 ) * ( n - 4 ) / 120
			}
			break
	}

	var result	= uint ( 1 )

/*	# Classic
	for	i := result ; i <= k ; i ++	{
		result	= result * ( n - i + 1 ) / i
	}
*/
	for	i := result ; i <= k ; i, n = i +1, n -1	{

		result	= result * n / i
	}

	return	result
}


/*	Same as Binomial_coefficient but internally uses big.Int
	from math/big package ( which doesn't overflow during int multiplication )
*/
func Binomial_coeff_big ( n, k uint64 )		uint64	{

	if	n <= 0	|| k <= 0	|| k >= n	{	return 1	}
	if	k == 1	|| k +1 == n			{	return n	}

//	k	= Bit_min ( k, n - k )
	if	k > n - k	{
		k = n - k
	}


	var (
		numerator	= big.NewInt ( 1 )
		denominator	= big.NewInt ( 1 )
	)

	numerator	.MulRange ( int64 ( n - k +1 ), int64 ( n ))
	denominator	.MulRange ( 1, int64 ( k ) )

	numerator.Quo ( numerator, denominator )

	return	numerator.Uint64 ()
}