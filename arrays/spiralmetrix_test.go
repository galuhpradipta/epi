// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestClockwise(t *testing.T) {
	for _, test := range []struct {
		in   [][]int
		want []int
	}{
		{ // 1x1
			[][]int{{1}},
			[]int{1}},

		{ // 2x2
			[][]int{
				{1, 2},
				{3, 4}},
			[]int{1, 2, 4, 3}},

		{ // 3x3
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},

		{ // 4x4
			[][]int{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15}},
			[]int{0, 1, 2, 3, 7, 11, 15, 14, 13, 12, 8, 4, 5, 6, 10, 9}},
		{ // 5x5
			[][]int{
				{0, 1, 2, 3, 4},
				{15, 16, 17, 18, 5},
				{14, 23, 24, 19, 6},
				{13, 22, 21, 20, 7},
				{12, 11, 10, 9, 8}},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}},
	} {
		if got := Clockwise(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Clockwise(%d)", test.in)
			t.Errorf(" got %d", got)
			t.Errorf("want %d", test.want)
		}
	}
}

func benchClockwise(b *testing.B, size int) {
	ints := rand.New(rand.NewSource(int64(size))).Perm(size)
	data := make([][]int, len(ints))
	for j := range data {
		data[j] = append([]int(nil), ints...)
	}
	for i := 0; i < b.N; i++ {
		Clockwise(data)
	}
}

func BenchmarkClockwise1e1(b *testing.B) { benchClockwise(b, 1e1) }
func BenchmarkClockwise1e2(b *testing.B) { benchClockwise(b, 1e2) }
func BenchmarkClockwise1e3(b *testing.B) { benchClockwise(b, 1e3) }
