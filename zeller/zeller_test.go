package zeller

import (
	"testing"
)

func TestZeller(t *testing.T) {
	w, err := Zeller("2022-03-26 000")
	if err == nil {
		t.Fatalf(`Zeller("2022-03-26 000")=%v, %v, want error, error\n`, w, err)
	}

	w, err = Zeller("2022-03-26")
	if w != 6 {
		t.Fatalf(`Zeller("2022-03-26")=%v, %v, want "6", error\n`, w, err)
	}

	w, err = Zeller("1582-10-04")
	if w != 4 {
		t.Fatalf(`Zeller("1582-10-04")=%v, %v, want "0", error\n`, w, err)
	}

	w, err = Zeller("1582-10-05")
	if err == nil {
		t.Fatalf(`Zeller("1582-10-05")=%v, %v, want error, error\n`, w, err)
	}

	w, err = Zeller("1582-10-15")
	if w != 5 {
		t.Fatalf(`Zeller("1582-10-15")=%v, %v, want "5", error\n`, w, err)
	}

}
