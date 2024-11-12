package nokocore

import (
	"fmt"
	"testing"
)

func TestShikaJsonEncodePreview_Ptr(t *testing.T) {
	data := NewMapAny()
	data.Set("ptr", uintptr(0))
	message := NewMapAny()
	message.Set("text", "Hello, World!")
	data.Set("message", message)
	data.Set("values", []int{0, 1, 2})
	user1 := NewMapAny()
	user1.Set("name", "John")
	user1.Set("age", 23)
	user1.Set("cities", []string{"Shanghai", "Beijing"})
	user2 := NewMapAny()
	user2.Set("name", "Jack")
	user2.Set("age", 24)
	user2.Set("cities", []string{"Jakarta", "Surabaya"})
	user3 := NewMapAny()
	user3.Set("name", "Jill")
	user3.Set("age", 25)
	user3.Set("cities", []string{"London", "Paris"})
	data.Set("users", []MapAnyImpl{user1, user2, user3})

	fmt.Println(ShikaJsonEncode(data))
	fmt.Println(ShikaYamlEncode(data))
}
