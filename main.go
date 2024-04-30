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
    const n string = "Mani" //here this for don't change the the value
    //we are able to to change the variable
    fmt.Println(n)
    const a = "mani"
    fmt.Println(a)
    //it's not mandatory to declare data type
    fmt.Println(a + n)
    var b = 35 //we must use var or const to declare anything
    //const use for precise
    fmt.Println(b)
    const n = 50
    const d = float64(3 / n)
    fmt.Printf("e = %.2f\n", d)
    i := 0
    for i <= 20 {
        fmt.Println(i) //printing zero to 2
        i = i + 1
    }
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }
    //it's like the same as above it just add j at every loop
    //here j = 0 and range is below 3
    for i := range 3 {
        fmt.Println(i)
    }
    //it's like same here range also like (0, 3)
    */
    a := 0
    for {
        fmt.Println("loop")
        fmt.Println(a)
        a = a + 1
        if a > 50 {
            break
        }
    }
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }
    //|| means or it's like 8%2 ==0 or 7%2 == 0
    

}
