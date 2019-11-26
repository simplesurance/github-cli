package command

import (
	"fmt"
	"os"

	"github.com/simplesurance/baur/format"
	"github.com/simplesurance/baur/format/csv"
	"github.com/simplesurance/baur/format/table"
)

const ErrorPrefix = "ERROR: "

func printErrf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, ErrorPrefix+format, a...)
}

func printErrln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, append([]interface{}{ErrorPrefix}, a...)...)
}

// mustWriteRows creates a CSV or table formatter depending on the
// rootFlags.csv value.
// It writes all rows to the formatter and flushes it.
// If an error happens, the program is terminated.
func mustWriteRows(rows [][]interface{}) {
	var formatter format.Formatter

	if rootFlags.csv {
		formatter = csv.New(nil, os.Stdout)
	} else {
		formatter = table.New(nil, os.Stdout)
	}

	for _, row := range rows {
		err := formatter.WriteRow(row)
		if err != nil {
			fmt.Fprintf(os.Stderr, "writing to stdout failed")
			os.Exit(1)
		}
	}

	formatter.Flush()
}
