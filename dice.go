package main

import (
	"math/rand"
)

// World Creation SRD https://www.traveller-srd.com/core-rules/world-creation/
// says roll 2D6 - 2, which is simpler in code as 2 dice roles with -1 on each:
// 2x 0-5 D6 rolls. Two dice roll, each with 0-5, so 5 is most likely at
// 1 in 6, like 7 normally (see the 0-5 rather than 1-6) 0 & 10 are least
//
//	likely at 1 in 36 (like 2 and 12 normally)
func zero_to_ten() int8 {
	return int8(6*rand.Float32()) + int8(6*rand.Float32())
}

// single die, with range 0 to 5, all equally likely
func zero_to_five() int8 {
	return int8(6 * rand.Float32())
}

// Three possible values, all equally likely
func zero_to_two() int8 {
	return int8(3 * rand.Float32())
}

// Two values, 50/50 (all equally likely)
func zero_to_one() int8 {
	return int8(2 * rand.Float32())
}

// Nine possible values, all equally likely, 1-9
func one_to_nine() int8 {
	return int8(9*rand.Float32()) + 1
}

// Normal 2D6 roll, 7 most common, 2 & 12 least likely
func two_to_twelve() int8 {
	return zero_to_ten() + 2
}

// World Creation SRD https://www.traveller-srd.com/core-rules/world-creation/
// says roll 3D6 - 3, which is simpler in code as 3 dice roles with -1 on each:
// 3x 0-5 D6 rolls. 7 and 8 are most likely, 0 and 15 are least likely at 1 in 216
// each (like 3 and 18 normally on 3D6)
func zero_to_fifteen() int8 {
	return zero_to_ten() + zero_to_five()
}
