/*package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
//this is the way to comment single line
/*this is the way to comment multiple lines
*this how to*/
package main

import "fmt"

func main() {
    /*
    fmt.Println("go" + "lang")//This just like adding two string and give golang

    fmt.Println("1+1 =", 1+1)//This is just add two number which are not in doble inverted comma give 1 + 1 = 2
    //1 + 1 is in string
    fmt.Println("7.0/3.0 =", 7.0/3.0) // here 7.0/ 3.0 = is in string and 7.0/3.0 in integer so int operations applicable
 

    fmt.Println(true && false)//here && means and so true and false is false
    fmt.Println(true || false)//here || is or so true or false is true
    fmt.Println(!true)// here ! is use for not
    //variable declaration
    var a = "mani"
    var b = "gupta"
    var c, d = 15, 25 //it's in string not in integer
    var e, f int = 245, 265 //here it's in integer
    fmt.Println(e + f)
    fmt.Println(e, f)
    var g = true //here true is in boolean
    var h int //If we don't determine the value then it take 0
    i := "apple"//this is also a way to define the variable
    fmt.Println(f)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c + d)
    fmt.Println(g)
    fmt.Println(h)
    fmt.Println(i)
    */
    const n string = "Mani"
    fmt.Println(n)
    const a = "mani"
    fmt.Println(a)
    //it's not mandatory to declare data type
    fmt.Println(a + n)
    var b = 35 //we must use var or const to declare anything
    //const use for precise
    fmt.Println(b)
}
