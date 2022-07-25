package main

import (
	"fmt"
)

type User struct {
	Id   int64
	name string
}

func main() {
	users := []User{
		{1, "Joe"},
		{2, "Mary"},
		{11, "Mary"},
		{22, "Mary"},
		{4, "Mary"},
	}
	fmt.Println(users)
	dm := make(map[int]User, len(users))
	fmt.Println(dm)
	for _, us := range users {
		if _, ok := dm[int(us.Id)]; !ok {
			dm[int(us.Id)] = us
		}
	}
	fmt.Println(dm)
	fmt.Println(fim(2, dm))

}

func fim(id int, dm map[int]User) *User {
	if v, ok := dm[id]; ok {
		return &v
	}
	return nil
}
