package main

import (
	"strings"
	"testing"
)

func TestReplace1(t *testing.T) {
	after := replace_url("http:")
	after = strings.TrimRight(after, "\n")
	if after != "- http:" {
		t.Errorf("after = %s", after)
	}
}

func TestReplace2(t *testing.T) {
	after := replace_url("あい～う?　,えおほ!げ::()   end")
	after = strings.TrimRight(after, "\n")
	if after != "## あいう_えおほげ_end.mp4" {
		t.Errorf("after = %s", after)
	}
}

func TestReplace3(t *testing.T) {
	after := replace_url("あい&う&&")
	after = strings.TrimRight(after, "\n")
	if after != "## あいう.mp4" {
		t.Errorf("after = %s", after)
	}
}

func TestReplace4(t *testing.T) {
	after := replace_url("aa~bb")
	after = strings.TrimRight(after, "\n")
	if after != "## aa_bb.mp4" {
		t.Errorf("after = %s", after)
	}
}
