package main

import "fmt"

func main() {
	//Slice
	fmt.Println("Slice")
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}
	links = append(links, "http://amazon.com")
	for index, link := range links {
		fmt.Println(index, link)
	}

	//Struct use
	fmt.Println("Struct")
	li := student{name: "LJH", neptun: "abcdef"}
	fmt.Println(li)
	li.neptun = "nnnnn"
	fmt.Printf("%+v", li)
	fmt.Println()

	//use receiver function and not receiver function
	fmt.Println("Receiver")
	fmt.Println(li.getName())
	fmt.Println(getName(li))

	//Pointer use
	fmt.Println("Pointer")
	fmt.Println(li)
	modifyStudentName(&li, "HZX")
	fmt.Println(li)
	fmt.Println(links[0])
	modifyLinksFirstItem(links, "baidu.com")
	fmt.Println(links[0])

} //main function

//Struct
type student struct {
	name   string
	neptun string
}

//Receiver function
func (s student) getName() string {
	return s.name
}

//Not Receiver function
func getName(s student) string {
	return s.name
}

//modify original value type variable
func modifyStudentName(pointerToStudent *student, newName string) {
	(*pointerToStudent).name = newName
}

//modify orignal reference type variable
func modifyLinksFirstItem(links []string, newFirstItem string) {
	links[0] = newFirstItem
}
