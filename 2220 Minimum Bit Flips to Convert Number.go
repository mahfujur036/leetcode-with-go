// The intuition is pretty simple and straight-forward. The XOR between the given
// two numbers will denote the changed bits. Now we just need to find out the number
// of set bits in the result of the XOR calculation to get to the result

import "math/bits"

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal)) 
}
