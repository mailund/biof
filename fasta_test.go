package biof_test

import (
	"os"
	"testing"

	"github.com/mailund/biof"
)

func TestReadFasta(t *testing.T) {
	fname := "/tmp/biof.fasta"

	f, err := os.Create(fname)
	if err != nil {
		t.Fatalf("Error creating file: %s", err.Error())
	}

	_, err = f.WriteString(">rec1\nabbcca\nbbabac\n>rec2\nadcat\nafacc\n")
	if err != nil {
		t.Fatalf("Error writing test fasta file: %s", err.Error())
	}

	f.Close()

	f, err = os.Open(fname)
	if err != nil {
		t.Fatalf("Error opening file: %s", err.Error())
	}

	m, err := biof.ReadFasta(f)
	if err != nil {
		t.Fatalf("error reading fasta records: %s", err.Error())
	}

	if seq, ok := m["rec1"]; !ok {
		t.Error("rec1 should be in the map")
	} else if seq != "abbccabbabac" {
		t.Errorf("We got the wrong sequence from rec1: %s", seq)
	}

	if seq, ok := m["rec2"]; !ok {
		t.Error("rec2 should be in the map")
	} else if seq != "adcatafacc" {
		t.Errorf("We got the wrong sequence from rec2: %s", seq)
	}
}
