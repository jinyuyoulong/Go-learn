package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。这意味着它拥有对底层数据结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map，你可以在这里阅读有关 hash tables 的更多信息。
	// 由于 nil 指针异常，你永远不应该初始化一个空的 map 变量：
	dic := Dictionary{"1": "fan", "2": "jin", "3": "long"}
	t.Run("know", func(t *testing.T) {
		got, _ := dic.Search("1")
		want := "fan"
		assertString(t, got, want)
	})

	t.Run("unknow", func(t *testing.T) {
		_, got := dic.Search("4")
		assertError(t, got, ErrNotFound)

	})
}
func TestAdd(t *testing.T) {
	dic := Dictionary{}
	word := "4"
	definition := "dd"
	err := dic.Add("4", "dd")
	assertDefinition(t, dic, word, definition)
	assertError(t, err, nil)

}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("want '%s' but got %s", want, got)
	}
}
func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want '%s' but got %s", want, got)
	}
}

func assertDefinition(t *testing.T, dic Dictionary, word, definition string) {
	t.Helper()
	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got == definition {
		t.Errorf("have value for key got '%s' want '%s'", got, definition)
	}
}
