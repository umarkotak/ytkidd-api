# Jodoh Express Backend Service

The code is written in "clean architecture" style. It will have 3 layers:
```
  Handler (http handler, web socket handler)
  |
  ---- Service (business logic)
      |
      ---- Repos (DB call, external API call)
```

## Code convention

- The upper layer must call the lower layer, it is allowed for handler to call directly to repository (not encouraged)
- All response contract must be stored in `models/resp_contract`
- Http request body contract and function contract can be stored in `models/contract`
- Core model must be exactly representing table in DB
- Other model can be used as needed to define core object (somewhat hard to distinguish between model and contract) - compromised to an extent

## Requirements

- golang 1.22
- postgresdb
- redis

## Running the server

```
make run
```
other command can be checked on Makefile


## Notes

if you run your golang app on wsl, you might need to get the wsl ip
`ip addr show eth0 | grep -oP '(?<=inet\s)\d+(\.\d+){3}'`
source: https://github.com/postmanlabs/postman-app-support/issues/11204#issuecomment-1605502449
