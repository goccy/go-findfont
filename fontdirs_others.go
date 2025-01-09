//go:build wasip1 || js

package findfont

func getFontDirectories() (paths []string) {
	return nil
}

func getUserFontDirs() (paths []string) {
	return nil
}

func getSystemFontDirs() (paths []string) {
	return nil
}
