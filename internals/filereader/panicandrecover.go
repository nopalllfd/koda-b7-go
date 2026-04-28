package filereader

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Continue :", r)
		}
	}()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Gagal membaca isi file: ", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic("Gagal membuka file")
	}
	fmt.Println(string(data))
	return

}
