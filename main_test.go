package main

import "testing"

func TestDNF(t *testing.T) {
	for _, tc := range []struct {
		table [][]int
		want  string
	}{
		{
			[][]int{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
				{1, 1, 1},
			},
			"(!a_0 & !a_1) | (a_0 & !a_1) | (a_0 & a_1)",
		},
		{
			[][]int{
				{0, 0, 1},
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			"(!a_0 & !a_1) | (!a_0 & a_1) | (a_0 & !a_1) | (a_0 & a_1)",
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 0, 0},
				{1, 1, 0},
			},
			"",
		},
		{
			[][]int{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
				{1, 1, 0},
			},
			"(!a_0 & !a_1)",
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 0},
			},
			"(!a_0 & a_1) | (a_0 & !a_1)",
		},
	} {
		if got := DNF(tc.table); got != tc.want {
			t.Errorf("DNF(%v), ERROR: got = %v, want = %v", tc.table, got, tc.want)
		}

	}
}
