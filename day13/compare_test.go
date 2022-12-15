package day13

import (
	"log"
	"testing"
)

func TestGetPartsFromString(t *testing.T) {
	tests := []struct {
		name   string
		expect []string
		input  string
	}{
		{"empty parsing", []string{}, "[]"},
		{"basic parsing", []string{"1", "1", "3", "1", "1"}, "[1,1,3,1,1]"},
		{"nested parsing", []string{"[1]", "[2,3,4]"}, "[[1],[2,3,4]]"},
		{"deeply nested parsing", []string{"1", "[2,[3,[4,[5,6,7]]]]", "8", "9"}, "[1,[2,[3,[4,[5,6,7]]]],8,9]"},
	}

	for _, test := range tests {
		expect := test.expect
		got := getPartsFromString(test.input)
		if len(got) != len(expect) {
			log.Println("test: ", test.name)
			log.Println("Wrong length")
			log.Println("got: ", len(got))
			log.Println("expect: ", len(expect))
			t.FailNow()
		}
		for i := range got {
			if got[i] != expect[i] {
				log.Println("test: ", test.name)
				log.Println("i: ", i)
				log.Println("got: ", got)
				log.Println("expect: ", expect)
				t.Fail()
			}
		}
	}
}

// Test cases from example
func TestIsPacketsInRightOrder(t *testing.T) {
	// Pair 1
	expect := true
	got := isPacketsInRightOrder("[1,1,3,1,1]", "[1,1,5,1,1]")
	if got != expect {
		log.Println("Pair 1")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// Pair 2
	expect = true
	got = isPacketsInRightOrder("[[1],[2,3,4]]", "[[1],4]")
	if got != expect {
		log.Println("Pair 2")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// Pair 3
	expect = false
	got = isPacketsInRightOrder("[9]", "[[8,7,6]]")
	if got != expect {
		log.Println("Pair 3")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// Pair 4
	expect = true
	got = isPacketsInRightOrder("[[4,4],4,4]", "[[4,4],4,4,4]")
	if got != expect {
		log.Println("Pair 4")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// Pair 5
	expect = false
	got = isPacketsInRightOrder("[7,7,7,7]", "[7,7,7]")
	if got != expect {
		log.Println("Pair 5")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// Pair 6
	expect = true
	got = isPacketsInRightOrder("[]", "[3]")
	if got != expect {
		log.Println("Pair 6")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// pAIR 7
	expect = false
	got = isPacketsInRightOrder("[[[]]]", "[[]]")
	if got != expect {
		log.Println("Pair 7")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// Pair 8
	expect = false
	got = isPacketsInRightOrder("[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]")
	if got != expect {
		log.Println("Pair 8")
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}
