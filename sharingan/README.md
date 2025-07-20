# Sharingan

**Sharingan** is a polyglot DevSecOps CLI toolkit that:

- 🔍 Scans your code for **secrets** (Ruby)  
- 🔐 Tests **security headers** (Ruby)  
- 🕵️‍♂️ Performs a **deep API vulnerability scan** (Python)  
- 🏅 Generates a **security badge** (Ruby)  
- 🛠️ Generates a **secure Express.js API template** (Node.js)  
- ⚙️ Orchestrates everything via a **Go/Cobra** CLI  

---

## Quick Start

1. **Clone the repo**  
   ```bash
   git clone https://github.com/RochdiFERjaoui1234/sharingan.git
   cd sharingan
Install dependencies

Go (with Cobra):

bash
Copy
Edit
go mod tidy
Python:

bash
Copy
Edit
pip3 install requests
Ruby: none (uses stdlib + ERB)

Node.js (for API generator):

bash
Copy
Edit
npm install express helmet express-rate-limit
Make scripts executable (if not already)

bash
Copy
Edit
chmod +x scripts/*.rb scripts/*.py scripts/*.js
Run the CLI

bash
Copy
Edit
# View help
go run main.go --help

# Scan a target for secrets, headers, and deep vulnerabilities
go run main.go scan https://example.com

# Generate a badge (outputs SVG to stdout)
go run main.go badge https://example.com > badge.svg

# Generate a secure Express.js API stub
go run main.go gen-api

# Run the full pipeline (scan + badge + API gen)