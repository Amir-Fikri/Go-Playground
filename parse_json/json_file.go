package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "reflect"
)

type Users struct {
    Users []User `json:"users"`
}

type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

type mUser map[string]string

func main() {

    // Open our jsonFile
    jsonFile, err := os.Open("users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    fmt.Println("")
    fmt.Println("Unstructured Data approach")
    fmt.Println("==========================")
    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Println(result)
    fmt.Println(reflect.TypeOf(result))

    fmt.Println(result["users"])
    fmt.Println(reflect.TypeOf(result["users"]))

    fmt.Println("")
    fmt.Println("Structured Data approach")
    fmt.Println("========================")
    var users Users

    json.Unmarshal(byteValue, &users)
    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

    //var result2 map[string]interface{}
    //result2 = result["users"]
    //fmt.Println(result2)
    //fmt.Println(reflect.TypeOf(result2))
}
