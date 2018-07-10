package main

import (
    "fmt"
)

type Llist struct {
    next *Llist
    data string
}

func main() {
    array := make([]Llist, 10)

    // Setup
    insert(array, "Prestige")
    insert(array, "Go")
    insert(array, "Dont")
    insert(array, "Wont")
    insert(array, "Hash")
    insert(array, "Table")
    insert(array, "rm -rf *")
    insert(array, "Oh")

    // Prints the Linked List key value pairs
    print(array)
}

//func (l *Llist) find(data string) (*Llist, error) {
    //TODO
        //Traverse List until desired value is found
        //Return error if value does not exist
//}

// Prints a Linked List
func print(array []Llist) {
    for i := 0; i < len(array); i++ {
        fmt.Printf("%d -> ", i)
        PrintList(array[i])
        fmt.Println("")
    }
}

// Helper for print
func PrintList(llist Llist) {
    if len(llist.data) > 0 {
        fmt.Printf("%s, ", llist.data)
    }

    if llist.next != nil {
        PrintList(*llist.next)
    }
}

// Insert into an array of Linked Lists
func insert(array []Llist, data string) *Llist {
    // Simple hash index calc
    index := len(data)
    return array[index].insert(data)
}

// Recursive Insertion
func (l *Llist) insert(data string) *Llist {
    if len(l.data) < 1 {
        l.data = data
        l.next = &Llist{}
    } else {
        l = l.next.insert(data)
    }

    return l
}
