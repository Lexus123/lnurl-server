# lnurl server

A server doing all kinds of LNURL related things.

## Stuff needed

1. A bitcoin and a lightning node with outbound liquidity.
2. The `lnd.conf` should contain values of the domain or IP the node is running, like so:

```toml
tlsextraip=192.168.1.6
tlsextradomain=some.domain.com
```

3. Generate the TLS files by throwing them away and restarting LND.
4. Bring over your `tls.cert` and all of your macaroons (admin, chainnotifier, invoice, invoices, readonly, router, signer and walletkit) to the machine running the lnurl server.

## Run

Make sure your `.env` has all fields set to their correct values.

You can run lnurl server using this command:

```bash
go run cmd/main.go
```