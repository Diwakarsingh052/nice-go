package sum

func SumInt(vs []int) int {

	sum := 0
	if vs == nil {
		return 0
	}
	for _, v := range vs {
		sum = v + sum
	}
	return sum

}

// go test -run SumInt/one -v // it is a pattern matching // all tests which have the prefix will run
/*
go test: It's a tool provided by Go which automates testing the packages named by the import paths. It prints a summary of the test results.
-v: This flag stands for 'verbose'. When you use this flag, the go test command will print the names of tests as they are run.
-cover: This flag is used for code coverage analysis. Coverage analysis is a measure used to describe the degree to which the source code of a program is executed when running the test suite. It will output the amount of code, in terms of percentage, covered by the tests.
./...: This to test all packages in your current directory (& subdirectories). The ... is a wildcard referring to all (sub)directories.
*/
