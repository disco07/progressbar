package main

import (
	"github.com/disco07/progressbar"
	"io"
	"net/http"
	"os"
)

func main() {
	req, _ := http.NewRequest("GET", "https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile("Docker%20Desktop%20Installer.exe", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(resp.ContentLength)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}
