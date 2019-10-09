package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	r, err := git.PlainOpen("./")
	if err != nil {
		fmt.Printf("[ERROR] %#v", err)
	}
	w, err := r.Worktree()
	if err != nil {
		fmt.Printf("[ERROR] %#v", err)
	}
	s, err := w.Status()
	if err != nil {
		fmt.Printf("[ERROR] %#v", err)
	}
	for k, v := range s {
		fmt.Printf("%s, %#v\n", k, v)
		fmt.Println(s.IsClean())
		fmt.Println(s.IsUntracked(k))
	}
}
