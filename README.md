# Proof-of-Space PoC

This repository contains a proof-of-concept implementation of a Proof-of-Space inspired blockchain system. It demonstrates basic plotting (disk allocation) and farming (challenge validation) mechanics in Go.

## Overview
Unlike Proof-of-Work, which consumes CPU/GPU resources, Proof-of-Space leverages disk space as the primary resource. The PoC here does not aim to replicate Chia or any production-ready system, but to provide an educational example of how such concepts can be modeled.

Main components:
- **Plotting**: Creates files on disk with pseudo-random data to simulate allocated storage.
- **Farming**: Validates random challenges against existing plots to prove storage commitment.
- **Difficulty**: Adjustable challenge parameters to control farming probability.

## Getting Started

### Requirements
- Go 1.20+
- Linux/MacOS/Windows

### Clone the repository
```
git clone https://github.com/<your-username>/proof-of-space-poc.git
cd proof-of-space-poc
```

### Run Plotting

This will create a plot file of approximately 50 MB in the plots/ directory.
```
go run ./cmd/plotter -size 50
```


### Run Farming

The farmer will scan the plots and attempt to provide a valid proof for the given challenge.
```
go run ./cmd/farmer -challenge random-seed
```

### Educational Purpose

This project is intended for learning and experimentation only. It is not a production-ready cryptocurrency implementation.

















