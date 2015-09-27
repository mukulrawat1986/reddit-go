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

func Download_images(url string, title string, sema chan bool) {

	fmt.Println("Go routine waiting to start")

	// We will not start work until we get something from the channel
	<-sema

	// Make a http Get request to the url
	// We get back a response and an error
	r, err := http.Get(url)

	if err != nil {
		log.Fatalf("http.Get error %v", err)
		//return err
	}

	// We read all the bytes of the image
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("io.ReadAll error %v", err)
		//return err
	}

	// close the response body
	r.Body.Close()

	// Create a temp file to store image
	file, err := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(file.Name())
	if err != nil {
		log.Fatalf("tempfile creating error %v", err)
	}

	// Store image in temp file
	ioutil.WriteFile(file.Name(), data, 0666)

	// Open temp file
	f, _ := os.Open(file.Name())
	defer f.Close()

	// Find out format of image
	_, format, _ := image.Decode(f)

	if format == "jpeg" || format == "png" || format == "gif" {
		// Now we save the file to disk
		filename := fmt.Sprintf("(%s).%s", title, format)
		log.Println("Saving image ", title)
		ioutil.WriteFile(filename, data, 0666)
	}

	sema <- true

}
