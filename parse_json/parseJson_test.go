package parseJson

import "testing"

func TestParseFile(t *testing.T) {
    want := "WILD001"
    if got := parseFile("file1.json"); got.Name != want {
        t.Errorf("parseFile(\"file1.json\") = %q, want %q", got, want)
    }

    want2 := ""
    if got2 := parseFile("file5.json"); got2.Name != want2 {
        t.Errorf("parseFile(\"file5.json\") = %q, want %q", got2, want2)
    }
}
