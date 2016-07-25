package cerrors

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	InitLog(LOG_Verbose)
	if IsVerbose() {
		log.Printf("Test string: %v", "Dag")
	}
	InitLog(LOG_Debug)
	if IsVerbose() {
		log.Printf("Test string: %v", "Dag")
	}

}

func TestErr(t *testing.T) {
	err := errors.New("File not found.")
	cerr := Wrap(err, 27, "While testing")
	fmt.Println("Short:")
	fmt.Println(cerr)
	fmt.Println("Long:")
	fmt.Printf("%+v\n", cerr)
	fmt.Println("Recursive:")
	ccerr := Wrap(cerr, 28, "Recursive")
	fmt.Println("Short:")
	fmt.Println(ccerr)
	fmt.Println("Long:")
	fmt.Printf("%+v\n", ccerr)
}
