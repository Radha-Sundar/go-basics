package calc

import (
	"fmt"
	"testing"
)

func TestRandStrN(t *testing.T) {
	ct := RandStrN(15)
	pr := fmt.Sprintf("RandStrN=%v", ct)
	fmt.Println(pr)
}

func TestRandIntN(t *testing.T) {
	ct := RandIntN(15)
	pr := fmt.Sprintf("RandIntN=%v", ct)
	fmt.Println(pr)
}

func TestStringPt(t *testing.T) {
	ct := StringPt("abcde")
	pr := fmt.Sprintf("StringPt=%v", ct)
	fmt.Println(pr)
}

func TestGetBasicAuthHeader(t *testing.T) {
	ct := GetBasicAuthHeader("abcde", "gfdse")
	pr := fmt.Sprintf("GetBasicAuthHeader=%v", ct)
	fmt.Println(pr)
}

func TestJsonForPrint(t *testing.T) {
	ct := JsonForPrint(28192)
	pr := fmt.Sprintf("JSON Marshal Int=%v", ct)
	fmt.Println(pr)

	ct = JsonForPrint(true)
	pr = fmt.Sprintf("JSON Marshal Bool=%v", ct)
	fmt.Println(pr)

	ct = JsonForPrint([]string{"foo", "bar", "baz"})
	pr = fmt.Sprintf("JSON Marshal Slice=%v", ct)
	fmt.Println(pr)

	ct = JsonForPrint(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	pr = fmt.Sprintf("JSON Marshal Map=%v", ct)
	fmt.Println(pr)

	type User struct {
		Name        string
		Age         int
		Active      bool
		lastLoginAt string
	}

	ct = JsonForPrint(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	pr = fmt.Sprintf("JSON Marshal Struct=%v", ct)
	fmt.Println(pr)
}

func TestStructToMap(t *testing.T) {

	type User struct {
		Name   string
		Age    int
		Active bool
	}
	req := &User{
		Name:   "Bob",
		Age:    10,
		Active: true,
	}
	ct := StructToMap(req)
	pr := fmt.Sprintf("StructToMap=%v", ct)
	fmt.Println(pr)
}

func TestSanitize(t *testing.T) {
	ct := Sanitize("this #is #for #testing", "A-Za-z ", 30)
	pr := fmt.Sprintf("Sanitize=%v", ct)
	fmt.Println(pr)
}
