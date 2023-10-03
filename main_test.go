package awesomeProject3

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strconv"
	"testing"
)

type testCase struct {
	value    float64
	expected float64
}

func loadSqrtCases(t *testing.T) []testCase {
	file, err := os.Open("sqrt_cases.csv")
	require.NoError(t, err)
	defer file.Close()
	var cases []testCase
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)

		val, err := strconv.ParseFloat(record[0], 64)
		require.NoError(t, err)
		expected, err := strconv.ParseFloat(record[1], 64)
		require.NoError(t, err)
		tc := testCase{val, expected}
		cases = append(cases, tc)
	}
	return cases
}

func TestSqrt(t *testing.T) {
	for _, tc := range loadSqrtCases(t) {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			require.NoError(t, err)
			require.InDelta(t, tc.expected, out, 0.001)
		})
	}
}
