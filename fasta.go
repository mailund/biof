package biof

import (
	"fmt"
	"io"
	"strings"
)

// ReadFasta reads a fasta file into a map from sequence names
// to the sequences
func ReadFasta(r io.Reader) (map[string]string, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	m := map[string]string{}

	records := strings.Split(string(bytes), ">")
	if len(records) == 0 {
		// empty but valid
		return m, nil
	}

	if records[0] != "" {
		return nil, fmt.Errorf("Expected an empty string before first header")
	}

	for i := 1; i < len(records); i++ {
		lines := strings.Split(records[i], "\n")
		header := lines[0]
		seq := strings.Join(lines[1:], "")
		m[header] = seq
	}

	return m, nil
}
