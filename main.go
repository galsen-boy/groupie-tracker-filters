package main

import (
	"fmt"
	handlers "groupie-tracker/handlers"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	clearScreen()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists", handlers.HandleArtists)
	http.HandleFunc("/artists/", handlers.ViewArtist)
	http.HandleFunc("/search", handlers.Search)
	http.HandleFunc("/filters", handlers.Filters)
	fmt.Println("Server start on port :8080")
	http.ListenAndServe(":8080", nil)

}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
