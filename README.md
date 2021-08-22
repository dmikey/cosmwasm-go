## Overview

This is a demo for cosmwasm-go build, it is in alpha state, but it is compatible
with CosmWasm 0.16.

This is currently in a **private pre-release**. This repo should only be shared
with select developers who will work on improving it. When we are happy with
the stability, we will publish this.

## Features

The following areas have been tested, both in unit tests and in integration
tests using `go-cosmwasm` to ensure compatibility. Please look at the 
[`hackatom` contract](./example/hackatom) for a good example with test coverage.

* Export `instantiate`, `execute`, `query`, `migrate` and return success response or errors
* Get and Set from persistent storage
* Canonicalize and Humanize addresses with the Api
* Query into the bank module (use querier)
* Efficiently parse json structs
* Simple API so you can focus on the business logic
* Easy setup and similar format for unit test and integration tests, to allow
  easy porting them.
  
The following areas have some code but have not been tested and *likely buggy*:

* Remove from storage
* Storage iterators
* Querying staking/wasm modules

These can be called, but will likely not function 100% and will require
some debugging to get them working.

## Caveats

Beyond the "likely buggy" Apis above, the following "features" are 
definitely not implemented:

* Support for Uint128 (parsing strings, dealing with large integers)
* Helper functions for creating messages or queries

## JSON

We have a major fork of [easyjson](https://github.com/mailru/easyjson)
that removes all calls to reflect and net and fmt and encoding/json. It also removes
all floats and parses numbers directly into ints.

We call it [tinyjson](https://github.com/CosmWasm/tinyjson).
v0.8.x is the first version after forking (easyjson is at 0.7.x)
and already working quite well.

It works like `encoding/json`, but uses code generation to create
`MarshalJSON` and `UnmarshalJSON` methods on all structs when you
run a generation stage. After that, you can simply call those
methods on the structs to do all parsing.

It handles slices and pointers and all kinds of other things just
like `encoding/json`. It automatically snake_cases's the variable
names if you compile with `snake_case`, which saves some work in the structs.
It turns `[]byte` into base64 strings, and it properly escapes other `string`s.

A special feature is the custom `emptyslice` option, which you can add
just like `omitempty`. With that, it means any slice of length 0 will
serialize as `[]` rather than `null` (the default behavior).
This avoids the need for most of the custom JSON marshallers that we just doing this.  
### Usage

We require `docker` installed and executable by the current user. This
is used to compile the contracts to Wasm. You can code contracts and
write unit test with only a normal Golang installation (requires 1.14+)

You can try the following top-level commands:

```
# Run unit tests on the standard library as well as `erc20` and `hackatom` contracts.
# Also compiles the contracts to wasm and runs integration tests
make test

# This will compile wasm binaries for all contracts and leave them in this directory
make examples
```

To try out a contract, either check out [hackatom](./example/hackatom),
running the tests and editing the code. Or start your own contract
by going to the [template](./example/template) directory and follow the
instructions on how to get started.

Both of these support the following commands: `make unit-test`, `make build`, 
and `make test`.

## Build system

We use docker tooling to get consistent builds acorss dev machines.
In my minimal experience, these seem to also produce deterministic
builds, but I would like others to try this out on other machines.
The following produces the same sha256 hash everytime I run it:

```
cd example/hackatom
make build && sha256sum hackatom.wasm

# this will test the wasm code
go test ./integration
```

However, the docker image for our custom TinyGo is not yet published.
In order to build locally you can do the following:

```
git clone https://github.com/confio/tinygo.git
git checkout cw-0.19.0
docker build -t cosmwasm/tinygo:0.19.0 -f Dockerfile.wasm .
```

Once it is finished, you should be able to successfully run `make build` on hackatom

## Performance

Many people ask how these compare to the rust contracts. I have yet to
do a detailed comparion, but now that we have two versions of the same
hackatom contract, we can do a rough side-by-side analysis.

**The good:**

The contract size is significantly lower than the Rust version, that is
97kB for the TinyGo version compared to 179kB for the Rust version.
(There is a sha256 algorithm in the Rust version missing from the Go version,
but that is only about 20Kb of the size)

**The bad:**

Before I switched to tinyjson, this used significantly more gas.
Now, it seems to use about 1.5-2x more gas than the Rust contracts.
And the compiled size is about 50% of the Rust size.
