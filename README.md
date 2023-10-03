
# Simple Automatization test cases in Golang 

This project contains a test suite for verifying the correctness of a square root function. The tests are driven by data from a CSV file (`sqrt_cases.csv`), making it easy to extend with additional test cases.

## Overview

The core of the project is the `TestSqrt` function, which reads test cases from the `sqrt_cases.csv` file and then runs a series of subtests. Each subtest checks if the square root function (`Sqrt`) correctly calculates the square root of a given value, within a specified tolerance.

## File Structure

- `main.go`: This file should contain the main square root function `Sqrt`.
- `sqrt_test.go`: Contains the testing logic and the utility function to load test cases.
- `sqrt_cases.csv`: Contains the simple test cases for Sqrt function
## How the Test Works

1. The `loadSqrtCases` function reads test cases from the `sqrt_cases.csv` file. Each line in the CSV represents a test case with two values: the input number and the expected square root value.
2. For each test case, `TestSqrt` creates a subtest using the input value as its name.
3. The subtest then calls the `Sqrt` function, checks if there are no errors, and asserts that the computed square root is close to the expected value (within a tolerance of 0.001).

## Usage

1. Ensure you have the `github.com/stretchr/testify/require` package installed. If not, you can install it using:
   ```
   go get github.com/stretchr/testify/require
   ```

2. To run the tests, navigate to the directory containing the code and execute:
   ```
   go test
   ```

3. Add or modify test cases by editing the `sqrt_cases.csv` file. Follow the format of input value, and expected result.

