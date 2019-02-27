package ternarytree

import "testing"

func testData() []string {
	return []string{
		"firewater",
		"stereolab",
		"tirewater",
		"tidewader",
		"tidewader",
		"firewater",
		"stereo",
	}
}

func TestNew(t *testing.T) {
	tree := TernaryTree{}
	if tree.head != nil {
		t.Errorf("TernaryTree head incorrectly initialized; expected: %v, observed: %v", nil, tree.head)
	}
	if tree.terminal != false {
		t.Errorf("TernaryTree terminal flag incorrectly initialized; expected: %v, observed: %v",
			false, tree.terminal)
	}
}

func TestInsert(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
}

func TestSearchSuccess(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	for _, v := range data {
		if !tree.Search(v) {
			t.Errorf("Search failed for '%s'", v)
		}
	}
}

func TestSearchFail(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	targets := []string{
		"fire",
		"tide",
		"stereophonic",
		"clutch",
	}
	for _, v := range targets {
		if tree.Search(v) {
			t.Errorf("Search passed for '%s'", v)
		}
	}
}
