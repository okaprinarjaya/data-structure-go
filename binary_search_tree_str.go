package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Key   string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(newKey string) {
	if newKey > n.Key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: newKey}
			return
		} else {
			n.Right.Insert(newKey)
		}
	} else if newKey < n.Key {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: newKey}
			return
		} else {
			n.Left.Insert(newKey)
		}
	}
}

func (n *Node) Search(key string) string {
	if n == nil {
		return ""
	}

	if strings.Compare(key, n.Key) == +1 {
		// move right
		return n.Right.Search(key)
	} else if strings.Compare(key, n.Key) == -1 {
		// move left
		return n.Left.Search(key)
	} else {
		return n.Key
	}
}

func main() {
	tree := &Node{Key: "oka"}
	tree.Insert("redi")
	tree.Insert("najih")
	tree.Insert("blessy")
	tree.Insert("tasya")
	tree.Insert("sani")
	tree.Insert("nico")
	tree.Insert("sasa")
	tree.Insert("heri")
	tree.Insert("jagoenk")
	tree.Insert("solut")

	tree.Insert("dul")
	tree.Insert("pepenk")
	tree.Insert("suarjani")
	tree.Insert("gungis")
	tree.Insert("ilma")
	tree.Insert("leo")
	tree.Insert("deni")
	tree.Insert("rata")
	tree.Insert("ucil")
	tree.Insert("roby")
	tree.Insert("galuh")
	tree.Insert("adit")
	tree.Insert("sari")
	tree.Insert("srinadi")
	tree.Insert("doyok")
	tree.Insert("kolok")
	tree.Insert("michael")
	tree.Insert("ariper")
	tree.Insert("dewacenk")
	tree.Insert("dewablonk")

	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Search name: ")
		name, _ := r.ReadString('\n')
		name = strings.Replace(name, "\n", "", -1)

		found := tree.Search(name)

		if found == "" {
			fmt.Println("not found")
		} else {
			fmt.Printf("found: %s\n", name)
		}
	}
}
