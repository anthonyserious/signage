# `signage`
Work in progress.  Proof of concept service for maintaining a rolling hash of
signed events.

# Design
`signage` provides an HTTP interface to a mechanism for signing and maintaining
a sequential blockchain of data.  It's envisioned that this can help with event auditing (i.e. tamper detection) and verification.

As an example, consider a plane ticket booking service that uses a microservice architecture with a bus.  You can add tamper detection by storing the continuous hash for each event along with any transactions, then periodically validating that the chain is consistent.  Additionally, you could ask `signage` to verify that a signed message is indeed from the service it's supposed to be from.

### Config
There will be a config.yml file containing, at a minimum:
- Path to server's private key
- Path to directory containing namespace public keys of the form `{namespace}.pub`
- Path to hash store directory where hash sequences will be stored per namespace.  In the future, this structure should support partitioning, with hashing performed in `goroutines` or even distributed.

### Stores
Data stores are required for a) server's private key, b) namespaces' publis keys, and c) hash logs.

### Routes
- /sign (POST) - sign and hash even.  Store and return the new hash.
  - `{ "namespace": "ticketBooking", "message": "Martha gets 1 round trip ticket to Aruba", "previousHash": "base64-encoded sha256 hash" }`
