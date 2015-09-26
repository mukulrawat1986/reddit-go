// The part which deals with download images and storing them 
// on a local drive
package redditgo

import (
    "fmt"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func init() {

      image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

      image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

      image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
 }


func Download_images(url string, title string) error{

    // Make a http Get request to the url
    // We get back a response and an error
    r, err := http.Get(url)

    if err != nil{
        log.Fatalf("http.Get error %v", err)
        return err
    }

    // We read all the bytes of the image
    data, err := ioutil.ReadAll(r.Body)

    if err != nil{
        log.Fatalf("io.ReadAll error %v", err)
        return err
    }

    // close the response body
    r.Body.Close()

    // Now we save the file to disk
    filename := fmt.Sprintf("%s.jpg", title)
    log.Println("Saving image")
    ioutil.WriteFile(filename, data, 0666)

    // Open the file and check its format
    f, _ := os.Open(filename)
    defer f.Close()

    _, format, _ := image.Decode(f)

    fmt.Printf("Format is %s\n", format)


    return err
}
