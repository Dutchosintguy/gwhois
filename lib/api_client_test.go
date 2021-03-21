package lib

import "testing"

func TestValidIdentifierShouldReturnNormally(t *testing.T) {
	info, _, err := GetFileInformation("1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms")
	if err != nil {
		t.Errorf("GetFileInformation threw error for valid document: %s", err)
	}
	if info.ID != "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms" {
		t.Errorf("ID does not match")
	}
}

func TestInvalidIdentifierShouldThrowErrors(t *testing.T) {
	_, _, err := GetFileInformation("invalid")
	if err == nil {
		t.Errorf("GetFileInformation threw no error for invalid document: %s", err)
	}
}
