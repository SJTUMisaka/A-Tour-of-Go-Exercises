package main

import "golang.org/x/tour/reader"

type MyReader struct{}

type MyError string

func (e MyError) Error() string{
	return string(e)
}

// Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	if len(b) == 0{
		return 0, MyError("buffer does not have enough length")
	}
	b[0] = 'A'
	return 1, nil
}
func main() {
	reader.Validate(MyReader{})
}
