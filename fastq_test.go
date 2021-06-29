package biof_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/mailund/biof"
)

var Reads = `@read0
ctcabbcbtb
+
~~~~~~~~~~
@read1
caatttabaa
+
~~~~~~~~~~
@read2
babaccttca
+
~~~~~~~~~~
@read3
tbbaabacat
+
~~~~~~~~~~
@read4
bbcbtbaabb
+
~~~~~~~~~~`

func TestScanFastq(t *testing.T) {
	expectedNames := []string{
		"read0", "read1", "read2", "read3", "read4",
	}
	expectedReads := []string{
		"ctcabbcbtb", "caatttabaa", "babaccttca", "tbbaabacat", "bbcbtbaabb",
	}
	expectedQuals := []string{
		"~~~~~~~~~~", "~~~~~~~~~~", "~~~~~~~~~~", "~~~~~~~~~~", "~~~~~~~~~~",
	}

	names := []string{}
	reads := []string{}
	quals := []string{}

	collect := func(rec *biof.FastqRecord) {
		names = append(names, rec.Name)
		reads = append(reads, rec.Read)
		quals = append(quals, rec.Qual)
	}

	err := biof.ScanFastq(strings.NewReader(Reads), collect)
	if err != nil {
		t.Errorf("Error scanning reads: %s", err.Error())
	}

	if !reflect.DeepEqual(names, expectedNames) {
		t.Error("Names do not match expected")
	}
	if !reflect.DeepEqual(reads, expectedReads) {
		t.Error("Reads do not match expected")
	}
	if !reflect.DeepEqual(quals, expectedQuals) {
		t.Error("Quals do not match expected")
	}
}
