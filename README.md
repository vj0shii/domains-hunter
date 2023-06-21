

# Installation

Make sure to give it execution permissions:

```bash
chmod +x setup.sh
```

To run the script, execute the following command:

```bash
./setup.sh
```

The script will update the system, install Go and the required tools (subfinder, amass, httpx, nuclei, and pandoc). It will also install additional tools (dirsearch, aquatone, and wfuzz) and clone the necessary repositories (dirsearch and nuclei-templates). Finally, it will compile the Go script.

After running the script, you can execute the Go script with integrated tools using:

```bash
go run script.go example.com
```

Please note that the script assumes a clean Ubuntu system and may require adjustments for different versions or distributions. Additionally, it's always a good practice to review and validate any script or commands before executing them on a production system.
