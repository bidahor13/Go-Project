
/** 

Name:Babatunde 

**/
// Copyright © 2018 Babatunde
//
// Golang stuff
// Version:1.0

package main

import "fmt"

const lang_1 = "Ghana"
const lang_2 = "Bini"
const lang_3 = "French"

func main(){
    
    fmt.Println("Welcome to my Gopher domain.")
    var ghMessage = "Akwaaba"
    var biMessage = "koyoo"
    var frMessage = "Bonjour"
    fmt.Println("The word",ghMessage, "originated from" ,lang_1,".")
    fmt.Println("The phrase",biMessage,"originated from",lang_2,".")
    fmt.Printf("Language of origin : %s, %s, %q " , lang_1, lang_2, lang_3)
    fmt.Printf("%s\n", frMessage)
    
    // printVariables()
    // showData()
    // Sensei()
}

