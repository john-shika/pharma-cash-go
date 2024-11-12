package nokocore

import (
	"fmt"
	"testing"
)

func TestShikaJsonEncodePreview_Ptr(t *testing.T) {
	data := NewMapAny()
	data.SetValueByKey("ptr", uintptr(0))
	message := NewMapAny()
	message.SetValueByKey("text", "Hello, World!")
	data.SetValueByKey("message", message)
	data.SetValueByKey("values", []int{0, 1, 2})
	user1 := NewMapAny()
	user1.SetValueByKey("name", "John")
	user1.SetValueByKey("age", 23)
	user1.SetValueByKey("cities", []string{"Shanghai", "Beijing"})
	user2 := NewMapAny()
	user2.SetValueByKey("name", "Jack")
	user2.SetValueByKey("age", 24)
	user2.SetValueByKey("cities", []string{"Jakarta", "Surabaya"})
	user3 := NewMapAny()
	user3.SetValueByKey("name", "Jill")
	user3.SetValueByKey("age", 25)
	user3.SetValueByKey("cities", []string{"London", "Paris"})
	data.SetValueByKey("users", []MapAnyImpl{user1, user2, user3})

	fmt.Println(ShikaJsonEncode(data))
	fmt.Println(ShikaYamlEncode(data))
}
