package metroninfo

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
)

func Read(r io.Reader) (*MetronInfo, error) {
	var mi MetronInfo
	if err := xml.NewDecoder(r).Decode(&mi); err != nil {
		return nil, err
	}
	mi.SetXMLAttributes()
	return &mi, nil
}

func (mi *MetronInfo) Write(w io.Writer, indented ...bool) error {
	contents, err := func() ([]byte, error) {
		if len(indented) > 0 && indented[0] {
			return xml.MarshalIndent(mi, "", "    ")
		}

		return xml.Marshal(mi)
	}()
	if err != nil {
		return err
	}

	_, err = w.Write(bytes.Join([][]byte{xmlHeader, []byte("\n"), contents}, []byte("")))
	return err
}

func (mi *MetronInfo) Save(path string, indented ...bool) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	return mi.Write(f, indented...)
}

func Open(path string) (*MetronInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return Read(f)
}
