package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Input domain to enumerate subdomains
	domain := "example.com"

	// Run subfinder to enumerate subdomains
	subfinderCmd := exec.Command("subfinder", "-d", domain, "-o", "subdomains.txt")
	err := subfinderCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute subfinder:", err)
	}

	// Run amass to further enumerate subdomains
	amassCmd := exec.Command("amass", "enum", "-d", domain, "-o", "subdomains-amass.txt")
	err = amassCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute amass:", err)
	}

	// Merge subdomains from subfinder and amass
	mergeCmd := exec.Command("cat", "subdomains.txt", "subdomains-amass.txt", "|", "sort", "|", "uniq", ">", "merged-subdomains.txt")
	err = mergeCmd.Run()
	if err != nil {
		log.Fatal("Failed to merge subdomains:", err)
	}

	// Run httpx to find live web apps
	httpxCmd := exec.Command("httpx", "-l", "merged-subdomains.txt", "-o", "webapps.txt")
	err = httpxCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute httpx:", err)
	}

	// Run nuclei to perform threat detection
	nucleiCmd := exec.Command("nuclei", "-l", "webapps.txt", "-t", "nuclei-templates/", "-o", "output.txt")
	err = nucleiCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute nuclei:", err)
	}

	// Run aquatone to capture screenshots and perform HTTP and DNS analysis
	aquatoneCmd := exec.Command("aquatone", "-out", "aquatone", "-d", "merged-subdomains.txt")
	err = aquatoneCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute aquatone:", err)
	}

	// Run dirsearch to perform directory brute-forcing
	dirsearchCmd := exec.Command("python3", "dirsearch.py", "-l", "merged-subdomains.txt", "-e", "php,asp,aspx,jsp", "-x", "400,403,404", "-t", "50", "-r", "-w", "common.txt")
	dirsearchCmd.Dir = "dirsearch"
	err = dirsearchCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute dirsearch:", err)
	}

	// Run wfuzz to perform web app fuzzing
	wfuzzCmd := exec.Command("wfuzz", "-c", "-z", "file,common.txt", "-Z", "--hc", "404,403", "-t", "50", "-f", "wfuzz-output.html", "--script", "xss-scan", "-H", "Host: FUZZ.example.com", "-u", "http://FUZZ.example.com/FUZZ")
	err = wfuzzCmd.Run()
	if err != nil {
		log.Fatal("Failed to execute wfuzz:", err)
	}

	// Convert the nuclei output to PDF format
	pdfConversionCmd := exec.Command("pandoc", "-s", "output.txt", "-o", "output.pdf")
	err = pdfConversionCmd.Run()
	if err != nil {
		log.Fatal("Failed to convert output to PDF:", err)
	}

	fmt.Println("Threat detection completed. Output saved as output.pdf")
}
