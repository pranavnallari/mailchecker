package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read input from the standard input (keyboard).
	scanner := bufio.NewScanner(os.Stdin)
	// Print the header for the output.
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	// Read input line by line until there is no more input.
	for scanner.Scan() {
		// Call the checkDomain function with the current line as the argument.
		checkDomain(scanner.Text())
	}

	// Check if any error occurred while reading the input.
	if err := scanner.Err(); err != nil {
		// Print an error message and exit the program if there was an error.
		log.Fatalf("Error := could not read from input %v\n", err)
	}
}

// Function to check a domain for MX, SPF, and DMARC records.
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool  // Variables to track if the domain has MX, SPF, and DMARC records.
	var spfRecord, dmarcRecord string // Variables to store SPF and DMARC records.

	// Lookup MX records for the domain.
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		// Print an error message if an error occurred during the MX lookup.
		log.Printf("Error occurred: %v\n", err)
	}

	// Check if MX records were found.
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Lookup TXT records for the domain (including SPF records).
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		// Print an error message if an error occurred during the TXT lookup.
		log.Printf("Error occurred: %v\n", err)
	}

	// Check each TXT record for SPF records.
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Lookup TXT records for the _dmarc subdomain of the domain.
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		// Print an error message if an error occurred during the DMARC TXT lookup.
		log.Printf("Error occurred: %v\n", err)
	}

	// Check each TXT record for DMARC records.
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Print the results for the domain, including whether it has MX, SPF, and DMARC records.
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v\n", domain, hasMX, mxRecords, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
