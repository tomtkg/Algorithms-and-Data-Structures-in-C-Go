package main

import "mylib"

func Example_mainA() {
	graph := [][3]int{
		{0, 5},
		{0, 1, 3},
		{0, 3, 5},
		{1, 2, 2},
		{1, 3, 1},
		{1, 4, 4},
		{2, 5, 4},
		{3, 6, 2},
		{3, 7, 3},
		{4, 5, 1},
		{5, 7, 2},
		{5, 8, 5},
		{6, 7, 2},
		{7, 8, 4},
	}
	mylib.SetStdin(graph)
	main()
	// Output:
	// node u = [0]
	// dist = [ 0  3 99  5 99 99 99 99 99]
	// flag = [ 1  0  0  0  0  0  0  0  0]
	// path = [-1  0  0  0  0  0  0  0  0]
	// node u = [1]
	// dist = [ 0  3  5  4  7 99 99 99 99]
	// flag = [ 1  1  0  0  0  0  0  0  0]
	// path = [-1  0  1  1  1  0  0  0  0]
	// node u = [3]
	// dist = [ 0  3  5  4  7 99  6  7 99]
	// flag = [ 1  1  0  1  0  0  0  0  0]
	// path = [-1  0  1  1  1  0  3  3  0]
	// node u = [2]
	// dist = [ 0  3  5  4  7  9  6  7 99]
	// flag = [ 1  1  1  1  0  0  0  0  0]
	// path = [-1  0  1  1  1  2  3  3  0]
	// node u = [6]
	// dist = [ 0  3  5  4  7  9  6  7 99]
	// flag = [ 1  1  1  1  0  0  1  0  0]
	// path = [-1  0  1  1  1  2  3  3  0]
	// node u = [4]
	// dist = [ 0  3  5  4  7  8  6  7 99]
	// flag = [ 1  1  1  1  1  0  1  0  0]
	// path = [-1  0  1  1  1  4  3  3  0]
	// node u = [7]
	// dist = [ 0  3  5  4  7  8  6  7 11]
	// flag = [ 1  1  1  1  1  0  1  1  0]
	// path = [-1  0  1  1  1  4  3  3  7]
	// node u = [5]
	// dist = [ 0  3  5  4  7  8  6  7 11]
	// flag = [ 1  1  1  1  1  1  1  1  0]
	// path = [-1  0  1  1  1  4  3  3  7]
	// 0 -> 5 : Distance = 8, Path = 5 4 1 0
}
func Example_mainB() {
	graph := [][3]int{
		{0, 8},
		{0, 1, 4},
		{0, 3, 1},
		{1, 2, 1},
		{1, 4, 1},
		{2, 5, 1},
		{3, 4, 1},
		{3, 6, 1},
		{4, 5, 2},
		{4, 7, 1},
		{5, 8, 2},
		{6, 7, 1},
		{7, 8, 5},
	}
	mylib.SetStdin(graph)
	main()
	// Output:
	// node u = [0]
	// dist = [ 0  4 99  1 99 99 99 99 99]
	// flag = [ 1  0  0  0  0  0  0  0  0]
	// path = [-1  0  0  0  0  0  0  0  0]
	// node u = [3]
	// dist = [ 0  4 99  1  2 99  2 99 99]
	// flag = [ 1  0  0  1  0  0  0  0  0]
	// path = [-1  0  0  0  3  0  3  0  0]
	// node u = [4]
	// dist = [ 0  3 99  1  2  4  2  3 99]
	// flag = [ 1  0  0  1  1  0  0  0  0]
	// path = [-1  4  0  0  3  4  3  4  0]
	// node u = [6]
	// dist = [ 0  3 99  1  2  4  2  3 99]
	// flag = [ 1  0  0  1  1  0  1  0  0]
	// path = [-1  4  0  0  3  4  3  4  0]
	// node u = [1]
	// dist = [ 0  3  4  1  2  4  2  3 99]
	// flag = [ 1  1  0  1  1  0  1  0  0]
	// path = [-1  4  1  0  3  4  3  4  0]
	// node u = [7]
	// dist = [ 0  3  4  1  2  4  2  3  8]
	// flag = [ 1  1  0  1  1  0  1  1  0]
	// path = [-1  4  1  0  3  4  3  4  7]
	// node u = [2]
	// dist = [ 0  3  4  1  2  4  2  3  8]
	// flag = [ 1  1  1  1  1  0  1  1  0]
	// path = [-1  4  1  0  3  4  3  4  7]
	// node u = [5]
	// dist = [ 0  3  4  1  2  4  2  3  6]
	// flag = [ 1  1  1  1  1  1  1  1  0]
	// path = [-1  4  1  0  3  4  3  4  5]
	// node u = [8]
	// dist = [ 0  3  4  1  2  4  2  3  6]
	// flag = [ 1  1  1  1  1  1  1  1  1]
	// path = [-1  4  1  0  3  4  3  4  5]
	// 0 -> 8 : Distance = 6, Path = 8 5 4 3 0
}
