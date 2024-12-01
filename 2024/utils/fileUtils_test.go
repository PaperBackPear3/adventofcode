package utils

import (
	"os"
	"reflect"
	"testing"
)

func TestReadAndSplitColumns(t *testing.T) {
	type args struct {
		filePath          string
		fakeData          string
		removeWhiteSpaces bool
		delimiter         string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "Test with comma delimiter",
			args: args{
				filePath: "testfile1.txt",
				fakeData: `68684,75847
98153,38453
67237,72168
25854,59466
`,
				removeWhiteSpaces: false,
				delimiter:         ",",
			},
			want: [][]string{
				{"68684", "98153", "67237", "25854"},
				{"75847", "38453", "72168", "59466"},
			},
		},
		{
			name: "Test with space delimiter and white spaces removed",
			args: args{
				filePath: "testfile2.txt",
				fakeData: `68684   75847
98153   38453
67237   72168
25854   59466
`,
				removeWhiteSpaces: true,
				delimiter:         " ",
			},
			want: [][]string{
				{"68684", "98153", "67237", "25854"},
				{"75847", "38453", "72168", "59466"},
			},
		},
		{
			name: "Test with space delimiter and white spaces removed",
			args: args{
				filePath: "testfile2.txt",
				fakeData: `68684   75847
98153   38453
67237   72168
25854   59466
`,
				removeWhiteSpaces: false,
				delimiter:         " ",
			},
			want: [][]string{
				{"68684", "98153", "67237", "25854"},
				{"", "", "", ""},
				{"", "", "", ""},
				{"75847", "38453", "72168", "59466"},
			},
		},
	}
	for _, tt := range tests {
		err := os.WriteFile(tt.args.filePath, []byte(tt.args.fakeData), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(tt.args.filePath)
		t.Run(tt.name, func(t *testing.T) {

			if got := ReadAndSplitColumns(tt.args.filePath, tt.args.removeWhiteSpaces, tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAndSplitColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadAndSplitRows(t *testing.T) {
	type args struct {
		filePath          string
		removeWhiteSpaces bool
		fakeData          string
		delimiter         string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "Test with comma delimiter",
			args: args{
				filePath: "testfile1.txt",
				fakeData: `68684,75847
98153,38453
67237,72168
25854,59466
`,
				removeWhiteSpaces: false,
				delimiter:         ",",
			},
			want: [][]string{
				{"68684", "75847"},
				{"98153", "38453"},
				{"67237", "72168"},
				{"25854", "59466"},
			},
		},
		{
			name: "Test with space delimiter and white spaces removed",
			args: args{
				filePath: "testfile2.txt",
				fakeData: `68684   75847
98153   38453
67237   72168
25854   59466
`,
				removeWhiteSpaces: true,
				delimiter:         " ",
			},
			want: [][]string{
				{"68684", "75847"},
				{"98153", "38453"},
				{"67237", "72168"},
				{"25854", "59466"},
			},
		},
		{
			name: "Test with space delimiter and white spaces removed",
			args: args{
				filePath: "testfile2.txt",
				fakeData: `68684   75847
98153   38453
67237   72168
25854   59466
`,
				removeWhiteSpaces: false,
				delimiter:         " ",
			},
			want: [][]string{
				{"68684", "", "", "75847"},
				{"98153", "", "", "38453"},
				{"67237", "", "", "72168"},
				{"25854", "", "", "59466"},
			},
		},
	}
	for _, tt := range tests {
		err := os.WriteFile(tt.args.filePath, []byte(tt.args.fakeData), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(tt.args.filePath)
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadAndSplitRows(tt.args.filePath, tt.args.removeWhiteSpaces, tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAndSplitRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
