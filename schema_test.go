package metroninfo

import (
	"bytes"
	"os"
	"testing"
)

func TestRead(t *testing.T) {

	originalFilePath := "./schemas/v1/Sample.xml"
	originalData, err := os.ReadFile(originalFilePath)
	if err != nil {
		t.Fatal(err)
	}

	mi, err := Open("./schemas/v1/Sample.xml")
	if err != nil {
		t.Fatal(err)
	}
	var outputBuffer bytes.Buffer
	if err = mi.Write(&outputBuffer, true); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(originalData, outputBuffer.Bytes()) {
		t.Fatalf("Output does not match the original file")
	}
}
