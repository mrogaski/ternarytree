package ternarytree

import "testing"

func testData() []string {
	return []string{
		"information",
		"expectation",
		"respectable",
		"accumulation",
		"discrimination",
		"complication",
		"intermediate",
		"grandmother",
		"introduction",
		"advertising",
		"acquaintance",
		"interference",
		"preparation",
		"satisfaction",
		"intervention",
		"consideration",
	}
}

func TestNew(t *testing.T) {
	tree := TernaryTree{}
	if tree.head != nil {
		t.Errorf("TernaryTree head incorrectly initialized; expected: %v, observed: %v", nil, tree.head)
	}
	if tree.hasEmpty {
		t.Errorf("TernaryTree data flag incorrectly initialized; expected: %v, observed: %v",
			false, tree.hasEmpty)
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
		"accumulate",
		"intervene",
		"stereophonic",
		"expect",
	}
	for _, v := range targets {
		if tree.Search(v) {
			t.Errorf("Search passed for '%s'", v)
		}
	}
}

func TestEmptyString(t *testing.T) {
	tree := TernaryTree{}
	tree.Insert("foo")
	if tree.Search("") {
		t.Fail()
	}
	tree.Insert("")
	if !tree.Search("") {
		t.Fail()
	}
}
