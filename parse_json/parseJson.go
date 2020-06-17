package parseJson

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Wild struct {
    Name    string `json:"name"`
    Bucket  string `json:"s3_bucket_name"`
    Object  string `json:"s3_object_key"`
    Version string `json:"version"`
    Update  string `json:"update_date"`
}

func parseFile(file string) Wild {

    // Open our jsonFile
    jsonFile, err := os.Open(file)
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully opened " + file)
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var wild Wild

    json.Unmarshal(byteValue, &wild)

    return wild
}
