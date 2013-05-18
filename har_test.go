package har

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	inFile  *os.File
	outFile *os.File
	log     *Log
	err     error
)

func init() {
	inFile, err = os.Open("session.har")
	if err != nil {
		panic("session.har file is missing or not accessible")
	}
}

func TestParse(t *testing.T) {
	defer inFile.Close()
	log, err = NewLog(inFile)
	if err != nil {
		t.Errorf("Failed to parse: %v", err)
	}
	t.Logf("Parsed. HAR Version on session.har is %s", log.Version)
}

func TestDump(t *testing.T) {
	outFile, err = ioutil.TempFile(".", "test_")
	if err != nil {
		t.Errorf("Failed to create temp file for writing: %v", err)
	}
	defer outFile.Close()
	err = Dump(outFile, log)
	if err != nil {
		t.Errorf("Failed to dump: %v", err)
	}
}
