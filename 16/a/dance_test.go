package main

import "testing"

func TestSpin(t *testing.T) {
	alphabet := []string{"a", "b", "c", "d", "e"}
	d := newDance(alphabet)
	want := "eabcd"
	d.spin(1)
	res := d.sequence()

	if res != want {
		t.Errorf(
			"Got %s instead of %s for spin of %d",
			res, want, 1,
		)
	}
}

func TestExchange(t *testing.T) {
	alphabet := []string{"a", "b", "c", "d", "e"}
	d := newDance(alphabet)
	want := "eabdc"

	d.spin(1)
	d.exchange(3, 4)
	res := d.sequence()

	if res != want {
		t.Errorf(
			"Got %s instead of %s for exchange %d,%d",
			res, want, 3, 4,
		)
	}
}

func TestPartner(t *testing.T) {
	alphabet := []string{"a", "b", "c", "d", "e"}
	d := newDance(alphabet)
	want := "baedc"

	d.spin(1)
	d.exchange(3, 4)
	d.partner("e", "b")
	res := d.sequence()

	if res != want {
		t.Errorf(
			"Got %s instead of %s for partner %s,%s",
			res, want, "e", "b",
		)
	}
}
