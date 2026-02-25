package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	// Serve from the directory where the exe is located
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Blad:", err)
		fmt.Println("Nacisnij Enter aby zamknac...")
		fmt.Scanln()
		return
	}
	dir := filepath.Dir(exePath)

	// Check that index.html exists next to the exe
	if _, err := os.Stat(filepath.Join(dir, "index.html")); os.IsNotExist(err) {
		fmt.Println("Nie znaleziono index.html obok pliku .exe")
		fmt.Println("Upewnij sie, ze triage-server.exe znajduje sie w tym samym folderze co index.html")
		fmt.Println()
		fmt.Println("Nacisnij Enter aby zamknac...")
		fmt.Scanln()
		return
	}

	// Find a free port starting from 8080
	port := findFreePort()
	addr := fmt.Sprintf("localhost:%d", port)
	url := fmt.Sprintf("http://%s", addr)

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir(dir)))

	fmt.Println("===========================================")
	fmt.Println("  TRIAGE MCI - Serwer lokalny")
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Printf("  Aplikacja dostepna pod: %s\n", url)
	fmt.Println()
	fmt.Println("  Otwieram przegladarke...")
	fmt.Println()
	fmt.Println("  Aby zainstalowac PWA:")
	fmt.Println("  - Chrome/Edge: kliknij ikone instalacji na pasku adresu")
	fmt.Println("  - Po instalacji mozesz zamknac ten serwer")
	fmt.Println()
	fmt.Println("  Aby zamknac serwer: zamknij to okno lub nacisnij Ctrl+C")
	fmt.Println("===========================================")

	// Open browser
	openBrowser(url)

	// Start server (blocking)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Blad serwera: %v\n", err)
		fmt.Println("Nacisnij Enter aby zamknac...")
		fmt.Scanln()
	}
}

func findFreePort() int {
	for port := 8080; port < 9000; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close()
			return port
		}
	}
	return 8080
}

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	}
}
