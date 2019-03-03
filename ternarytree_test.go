package ternarytree

import (
	"reflect"
	"testing"
)

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
		"grandfather",
		"introduction",
		"advertising",
		"acquaintance",
		"interference",
		"preparation",
		"satisfaction",
		"intervention",
		"consideration",
		"tire",
		"fire",
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

func TestPartialMatchSearch(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	result := tree.PartialMatchSearch("in..........", '.')
	expected := []string{
		"interference",
		"intermediate",
		"intervention",
		"introduction",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	result = tree.PartialMatchSearch("XXtXXXXXtion", 'X')
	expected = []string{
		"intervention",
		"introduction",
		"satisfaction",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
}

func TestPartialMatchSearchEmptyString(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	var expected []string
	result := tree.PartialMatchSearch("", '.')
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	tree.Insert("")
	result = tree.PartialMatchSearch("", '.')
	expected = []string{""}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
}

func TestNearNeighborSearch(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	result := tree.NearNeighborSearch("grandmother", 9)
	expected := []string{
		"grandfather",
		"grandmother",
		"information",
		"preparation",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	result = tree.NearNeighborSearch("grandmother", 8)
	expected = []string{
		"grandfather",
		"grandmother",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	result = tree.NearNeighborSearch("grandmother", 2)
	expected = []string{
		"grandfather",
		"grandmother",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
}

func TestNearNeighborSearchEmptyString(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	var expected []string
	result := tree.NearNeighborSearch("", 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	result = tree.NearNeighborSearch("", 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	tree.Insert("")
	expected = []string{""}
	result = tree.NearNeighborSearch("", 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
	result = tree.NearNeighborSearch("", 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
}

func collector() (func(string), func() []string) {
	var store []string

	gather := func(s string) {
		store = append(store, s)
	}

	report := func() []string {
		return store
	}

	return gather, report
}

func TestTraverse(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	expected := []string{
		"accumulation",
		"acquaintance",
		"advertising",
		"complication",
		"consideration",
		"discrimination",
		"expectation",
		"fire",
		"grandfather",
		"grandmother",
		"information",
		"interference",
		"intermediate",
		"intervention",
		"introduction",
		"preparation",
		"respectable",
		"satisfaction",
		"tire",
	}

	gather, report := collector()
	tree.Traverse(gather)
	result := report()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
}

func TestSort(t *testing.T) {
	tree := TernaryTree{}
	data := testData()
	for _, v := range data {
		tree.Insert(v)
	}
	expected := []string{
		"accumulation",
		"acquaintance",
		"advertising",
		"complication",
		"consideration",
		"discrimination",
		"expectation",
		"fire",
		"grandfather",
		"grandmother",
		"information",
		"interference",
		"intermediate",
		"intervention",
		"introduction",
		"preparation",
		"respectable",
		"satisfaction",
		"tire",
	}

	result := tree.Sort()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v, observed: %v", expected, result)
	}
}
