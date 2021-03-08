package main

import (
	"fmt"
)

    type name struct {
        fname [20]string
        lname [20]string
    }

    type listOfNames struct {
        Items []name
    }

    func (names *listOfNames) AddItem(item name) {
        listOfNames.Items = append(listOfNames.Items, item)
    }


    func main() {

       
        // Get name of text file
        // Open file

        // Function one - fill data from file to slice
        // Read file in a loop till EOF
        // fill the fname and lname from file into the struct of name
        // Add the name type to the listOfNames slice using the AddItem() method
        
        //item1 := MyBoxItem{Name: "Test Item 1"}
        //item2 := MyBoxItem{Name: "Test Item 2"}

        //box := MyBox{}

        //box.AddItem(item1)
        //box.AddItem(item2)

        // Close file

        // Function two - iterate through the slice and print each name
        // checking the output
        fmt.Println(len(listOfNames.Items))
        fmt.Println(listOfNames.Items)
    }