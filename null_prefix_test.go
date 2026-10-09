package jsoniter

import "testing"

func TestUnmarshalInvalidNullPrefixDoesNotClearMap(t *testing.T) {
	data := map[string]interface{}{"keep": "yes"}
	err := Unmarshal([]byte("none"), &data)
	if err == nil {
		t.Fatal("expected error")
	}
	if data == nil {
		t.Fatal("map was cleared")
	}
	if data["keep"] != "yes" {
		t.Fatalf("map changed: %#v", data)
	}
}

func TestUnmarshalNullClearsMap(t *testing.T) {
	data := map[string]interface{}{"keep": "yes"}
	err := Unmarshal([]byte("null"), &data)
	if err != nil {
		t.Fatal(err)
	}
	if data != nil {
		t.Fatalf("null should clear the map, got %#v", data)
	}
}

func TestUnmarshalInvalidNullPrefixDoesNotClearSlice(t *testing.T) {
	data := []int{1, 2}
	err := Unmarshal([]byte("none"), &data)
	if err == nil {
		t.Fatal("expected error")
	}
	if len(data) != 2 || data[0] != 1 || data[1] != 2 {
		t.Fatalf("slice changed: %#v", data)
	}
}

func TestUnmarshalNullClearsSlice(t *testing.T) {
	data := []int{1, 2}
	err := Unmarshal([]byte("null"), &data)
	if err != nil {
		t.Fatal(err)
	}
	if data != nil {
		t.Fatalf("null should clear the slice, got %#v", data)
	}
}
