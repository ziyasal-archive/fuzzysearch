
package main

import "testing"

func TestFuzzySearch(t *testing.T) {
  var v bool
  v = FuzzySearch("twl", "cartwheel")
  if v != true {
    t.Error("Expected true, got ", v)
  }
}
