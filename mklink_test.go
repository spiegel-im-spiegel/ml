package mklink

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var typesTests2 = []typesTestCase{
	{"[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)\n", StyleMarkdown},
	{"[https://github.com/spiegel-im-spiegel/mklink GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format]\n", StyleWiki},
	{"<a href=\"https://github.com/spiegel-im-spiegel/mklink\">GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format</a>\n", StyleHTML},
	{"\"https://git.io/vFR5M\",\"https://github.com/spiegel-im-spiegel/mklink\",\"GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format\",\"mklink - Make Link with Markdown Format\"\n", StyleCSV},
	{"", StyleUnknown},
}

func TestEncode(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/mklink", Title: "GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format", Description: "mklink - Make Link with Markdown Format"}
	for _, tst := range typesTests2 {
		r := lnk.Encode(tst.t)
		buf := new(bytes.Buffer)
		io.Copy(buf, r)
		str := buf.String()
		if str != tst.name {
			t.Errorf("Encode(%v)  = \"%v\", want \"%v\".", tst.t, str, tst.name)
		}
	}
}

func TestString(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/mklink", Title: "GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format", Description: "mklink - Make Link with Markdown Format"}
	str := lnk.String()
	res := `{
  "url": "https://git.io/vFR5M",
  "location": "https://github.com/spiegel-im-spiegel/mklink",
  "title": "GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format",
  "description": "mklink - Make Link with Markdown Format"
}
`
	if str != res {
		t.Errorf("New()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestNewErr(t *testing.T) {
	_, err := New("https://foo.bar")
	if err == nil {
		t.Error("New()  = nil error, not want nil error.")
	} else {
		fmt.Fprintf(os.Stderr, "info: %v\n", err)
	}
}

func ExampleNew() {
	link, err := New("https://git.io/vFR5M")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(link.Encode(StyleMarkdown))
	// Output:
	// [GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
}
