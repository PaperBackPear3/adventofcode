package utils

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sorted array",
			args: args{array: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "reverse sorted array",
			args: args{array: []int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "unsorted array",
			args: args{array: []int{2, 3, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "array with duplicates",
			args: args{array: []int{3, 1, 2, 1}},
			want: []int{1, 1, 2, 3},
		},
		{
			name: "empty array",
			args: args{array: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayAtoi(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "array of numbers",
			args: args{array: []string{"1", "2", "3"}},
			want: []int{1, 2, 3},
		},
		{
			name: "array with negative numbers",
			args: args{array: []string{"-1", "-2", "-3"}},
			want: []int{-1, -2, -3},
		},
		{
			name: "array with mixed numbers",
			args: args{array: []string{"1", "-2", "3"}},
			want: []int{1, -2, 3},
		},
		{
			name: "array with non-numeric strings",
			args: args{array: []string{"a", "2", "3"}},
			want: []int{0, 2, 3},
		},
		{
			name: "empty array",
			args: args{array: []string{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayAtoi(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayValuesDiffAndMerge(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "arrays of same length",
			args: args{array1: []int{1, 2, 3}, array2: []int{4, 5, 6}},
			want: []int{3, 3, 3},
		},
		{
			name: "first array longer",
			args: args{array1: []int{1, 2, 3, 7}, array2: []int{4, 5, 6}},
			want: []int{3, 3, 3, 7},
		},
		{
			name: "second array longer",
			args: args{array1: []int{1, 2, 3}, array2: []int{4, 5, 6, 8}},
			want: []int{3, 3, 3, 8},
		},
		{
			name: "arrays with negative numbers",
			args: args{array1: []int{-1, -2, -3}, array2: []int{4, -5, 6}},
			want: []int{5, 3, 9},
		},
		{
			name: "arrays with mixed numbers",
			args: args{array1: []int{1, -2, 3}, array2: []int{-4, 5, -6}},
			want: []int{5, 7, 9},
		},
		{
			name: "empty arrays",
			args: args{array1: []int{}, array2: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayValuesDiffAndMerge(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayValuesDiffAndMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraySum(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArraySum(tt.args.array); got != tt.want {
				t.Errorf("ArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayFindOccasions(t *testing.T) {
	type args struct {
		array  []int
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayFindOccasions(tt.args.array, tt.args.number); got != tt.want {
				t.Errorf("ArrayFindOccasions() = %v, want %v", got, tt.want)
			}
		})
	}
}
