# Email Status Checker

This Go program checks the status of email-related records (MX, SPF, and DMARC) for a given domain.

## Description

The program reads input from the standard input (keyboard) line by line, expecting each line to contain a domain name. For each domain name provided, it performs the following checks:

- **MX Record Lookup**: It uses the `net.LookupMX` function to retrieve the mail exchange (MX) records associated with the domain. MX records specify the mail server responsible for receiving email messages on behalf of the domain.

- **SPF Record Lookup**: It uses the `net.LookupTXT` function to retrieve text (TXT) records associated with the domain. It searches for a record starting with "v=spf1", indicating the presence of a Sender Policy Framework (SPF) record. SPF records specify which IP addresses are allowed to send emails on behalf of the domain.

- **DMARC Record Lookup**: It performs a TXT record lookup for the `_dmarc` subdomain of the given domain. DMARC (Domain-based Message Authentication, Reporting, and Conformance) records specify email authentication and reporting policies for the domain. It searches for a record starting with "v=DMARC1".

For each domain, the program prints a comma-separated line containing the following information:

- Domain name
- Whether it has MX records
- MX records (if present)
- Whether it has SPF records
- SPF record (if present)
- Whether it has DMARC records
- DMARC record (if present)

## Theory

### MX Records

MX (Mail Exchange) records are DNS records that specify the mail server responsible for receiving email messages on behalf of a domain. When an email is sent to a domain, the sender's email server queries the DNS system for the MX records of the recipient domain to determine where to deliver the email.

### SPF Records

SPF (Sender Policy Framework) records are DNS records that specify which IP addresses are allowed to send emails on behalf of a domain. SPF records help prevent email spoofing and phishing by allowing domain owners to specify authorized senders.

### DMARC Records

DMARC (Domain-based Message Authentication, Reporting, and Conformance) records are DNS records that specify email authentication and reporting policies for a domain. DMARC records allow domain owners to specify how email providers should handle emails that fail SPF or DKIM (DomainKeys Identified Mail) authentication, and to receive reports on email authentication failures.
