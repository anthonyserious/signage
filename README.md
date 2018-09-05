# `signage`
Work in progress.  Proof of concept service for maintaining a rolling hash of
signed events.

## Design
`signage` provides an HTTP interface to a mechanism for signing and maintaining
a sequential blockchain of data.  It's envisioned that this can help with event auditing (i.e. tamper detection) and verification.

Routes:
- /sign (POST) - sign and store event.
  - `{ "namespace": "ticketBooking", "message": "Martha gets 1 ticket for the show", "previousHash": "base64-encoded sha256 hash" }`
- /verify - verify that a message was signed by the stored namespace's key
