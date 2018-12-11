
/** 

Name:Babatunde 

**/

package main

import "fmt"

const lang_1 = "Ghana"
const lang_2 = "Bini"

func main(){
    fmt.Println("Welcome to my Gopher domain.")
    var ghMessage = "Akwaaba"
    var biMessage = "koyoo"
    fmt.Println("The word",ghMessage, "originated from" ,lang_1,".")
    fmt.Println("The phrase",biMessage,"originated from", lang_2,".")
    
    printVariables()
    showData()
    Sensei()
}

