package main

import (
	"fmt"
)

// user
type user struct {
	name string
	age  int
}

// iIterator interface
type iIterator interface {
	hasNext() bool
	next()
	currentItem() *user
}

// userCollection interface
type iUserCollection interface {
	iterator() iIterator
	add(u *user)
	size() int
	get(index int) *user
}

// userSlice concrete userCollection
type userSlice struct {
	users []*user
}

func (u *userSlice) iterator() iIterator {
	return newUserIterator(u)
}

func (u *userSlice) add(user *user) {
	u.users = append(u.users, user)
}

func (u *userSlice) size() int {
	return len(u.users)
}

func (u *userSlice) get(index int) *user {
	return u.users[index]
}

// userIterator concrete iterator
type userIterator struct {
	cursor int
	data   iUserCollection
}

func newUserIterator(data iUserCollection) iIterator {
	return &userIterator{
		0,
		data,
	}
}

func (i *userIterator) hasNext() bool {
	return i.cursor < i.data.size()
}

func (i *userIterator) next() {
	i.cursor++
}

func (i *userIterator) currentItem() *user {
	if !i.hasNext() {
		return nil
	}
	return i.data.get(i.cursor)
}

// main
func main() {
	u1 := &user{
		"John",
		13,
	}
	u2 := &user{
		"Marry",
		11,
	}
	users := []*user{u1}

	dataCollection := &userSlice{
		users,
	}

	dataCollection.add(u2)

	iterator1 := dataCollection.iterator()

	for iterator1.hasNext() {
		fmt.Println(*iterator1.currentItem())
		iterator1.next()
	}
}
