package main

import "fmt"

type IntSet interface {
	isEmpty() bool
	contains(int) bool
	insert(int) IntSet
	union(IntSet) IntSet
}

type Empty struct {
}

func (Empty) isEmpty() bool {
	return true
}

func (Empty) contains(int) bool {
	return false
}

func (e Empty) insert(n int) IntSet {
	return Insert{e, n}
}

func (Empty) union(s IntSet) IntSet {
	return s
}

type Insert struct {
	s IntSet
	n int
}

func (Insert) isEmpty() bool {
	return false
}

func (i Insert) contains(n int) bool {
	return n == i.n || i.contains(n)
}

func (i Insert) insert(n int) IntSet {
	if i.s.contains(n) {
		return i.s
	}
	return Insert{i, n}
}

func (i Insert) union(s IntSet) IntSet {
	return Union{i, s}
}

type Union struct {
	s1 IntSet
	s2 IntSet
}

func (u Union) isEmpty() bool {
	return u.s1.isEmpty() && u.s2.isEmpty()
}

func (u Union) contains(n int) bool {
	return u.s1.contains(n) || u.s2.contains(n)
}

func (u Union) insert(n int) IntSet {
	return Insert{u, n}
}

func (u Union) union(s IntSet) IntSet {
	return Union{u, s}
}

func main() {
	var empty = Empty{}
	var one = Insert{empty, 1}
	fmt.Println(one.contains(1))
}
