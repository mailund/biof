package biof

import (
	"fmt"
	"io"
)

func PrintSam(w io.Writer, qname, rname string, pos int, cigar, read, qual string) error {
	_, err := fmt.Fprintf(w, "%s\t0\t%s\t%d\t0\t%s\t*\t0\t0\t%s\t%s", qname, rname, pos+1, cigar, read, qual)
	return err
}
