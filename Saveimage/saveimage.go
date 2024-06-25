package saveimage

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SaveImageFromURL(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Image saved successfully ")

	return nil
}
