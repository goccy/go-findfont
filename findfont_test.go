package findfont

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	// List must find something
	fonts := List()
	if len(fonts) == 0 {
		t.Errorf("No font files found in system folders")
	}

	// ListWithSuffix using bad suffix
	bad_suffixes := make([]string, 0, 1)
	bad_suffixes = append(bad_suffixes, ".bad-suffix")
	bad_suffix_fonts := ListWithSuffixes(bad_suffixes)
	if len(bad_suffix_fonts) != 0 {
		t.Errorf("Unexpectedly found font files with bad suffix, e.g. %s", bad_suffix_fonts[0])
	}

	// ListWithSuffixes using good suffix
	good_suffixes := make([]string, 0, 1)
	good_suffixes = append(good_suffixes, filepath.Ext(fonts[0]))
	good_suffix_fonts := ListWithSuffixes(good_suffixes)
	if len(good_suffix_fonts) == 0 {
		t.Errorf("No font files with suffix %s", good_suffixes[0])
	}
}

func TestFind(t *testing.T) {
	// Try to find a non-existing font
	font, err := Find("this-font-does-not-exist.ttf")
	if err == nil {
		t.Errorf("Expected match when searching for non-existant font: %s", font)
	}

	fonts := List()
	if len(fonts) > 0 {
		// Direct search for existing font file
		font, err = Find(fonts[0])
		if err != nil {
			t.Errorf("Direct search failed: %v", err)
		}
		if font != fonts[0] {
			t.Errorf("Unexpected match for direct search: %s, expected: %s", font, fonts[0])
		}

		// Search only for basename
		needle := filepath.Base(fonts[0])
		needle = strings.TrimSuffix(needle, filepath.Ext(needle))
		_, err = Find(needle)
		if err != nil {
			t.Errorf("Basename search failed: %v", err)
		}
	}
}
