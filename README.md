# Mangostine-benchmark

# Building the source

Prerequisites

- Go 1.15+ [installed](https://github.com/golang/go)

## Steps for building the project executables

1. Clone the project repository

   `git clone https://gitlab.com/mangostine-benchmark.git`

2. Change to the project directory.

   `cd gitlab.com/mangostine-benchmark`

3. Get all the dependecies and build project the binary

   `make all`

## Running `Mangostine` Full node on the local network

1. Create the account using below code which will create the key.json file

   `mangostine keypair --home ~/.mangostine`

2. Send the transaction
   `mangostine sentTx`
