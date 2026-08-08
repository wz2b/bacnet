# Go BACnet Library

This repository is a fork and substantial reworking of [`noahtkeller/go-bacnet`](https://github.com/noahtkeller/go-bacnet), which in turn was heavily influenced by Steve Karg's [`bacnet-stack`](https://github.com/bacnet-stack/bacnet-stack).

That lineage provided a valuable starting point, particularly for BACnet wire formats, protocol constants, and service implementations. This fork retains that foundation while taking the library in a different direction.

## Goals of this Fork

The original library contained useful BACnet functionality, but it was incomplete in a number of areas and retained many implementation patterns inherited from C.

The goals of this fork are:

1. **Fill in missing BACnet functionality**
2. **Make the library feel like Go rather than a transliteration of C**
3. **Provide reusable protocol layers suitable for devices, routers, diagnostic tools, and protocol exploration**

### Expanded BACnet Support

This fork adds and continues to add support for missing or incomplete BACnet functionality, including:

* BACnet/IP BVLC messages
* BBMD and Foreign Device Registration functionality
* BACnet network-layer messages and routing support
* Additional confirmed and unconfirmed application services
* More complete APDU and NPDU encoding and decoding
* BACnet object and property handling
* Generic packet decoding and dispatch
* Handling of unknown or proprietary messages without unnecessarily treating them as malformed packets
* Additional protocol datatypes and encoding/decoding helpers

The intent is not merely to support one application, but to provide useful BACnet protocol primitives that can be reused by higher-level libraries and tools.

## Making It Go

A major goal of this fork is to move away from APIs that closely mirror C implementations.

The original code understandably inherited many patterns from `bacnet-stack`, including preallocated buffers, explicit length fields, mutable decode structures, setter functions, fixed-size storage, and APIs structured around the expectation that the caller would allocate an object and then have a decoder populate it.

Those patterns make sense in C.

They generally do not make sense in Go.

In Go, if decoding a packet produces a `WhoIs`, `IAm`, `ReadPropertyRequest`, `NPDU`, or other protocol value, the decoder can simply create and return that value.

For example, the preferred model is:

```go
msg, err := bacnet.Decode(packet)
```

rather than requiring callers to:

```text
allocate a structure
allocate backing buffers
initialize length fields
pass it into a decoder
copy bytes into it
track how much of the buffer was consumed
```

It is 2026. We can create objects when we decode them.

### Some of the changes made in that direction include

* Decoders return newly constructed Go values instead of mutating caller-allocated structures.
* Encoders return `[]byte` rather than requiring callers to provide oversized scratch buffers.
* Go slices replace C-style fixed-capacity byte arrays where fixed storage is not required by the protocol.
* Explicit length fields are removed where `len(slice)` already represents the information.
* Decode functions return errors for malformed network input rather than panicking.
* Protocol structures contain semantic data rather than encoder scratch state.
* Constructors and setter functions are removed when ordinary Go struct literals are clearer.
* Pointer-to-interface and union-like C patterns are being removed or replaced with more natural Go representations.
* Package boundaries reflect protocol layers rather than the organization of the original C implementation.
* Generic decoding progressively unwraps BVLC, NPDU, APDU, and application-service layers.
* Known messages are returned as concrete Go types that can be handled with ordinary type switches.
* Unknown but structurally valid messages can be preserved as structured values with their raw payload intact, which is useful for logging, diagnostics, and investigation of proprietary BACnet extensions.

The resulting API is intended to support code shaped more like:

```go
msg, err := bacnet.Decode(packet)
if err != nil {
	return err
}

switch value := msg.Value.(type) {
case *services.WhoIs:
	// handle Who-Is

case *services.IAm:
	// handle I-Am

case *services.ReadPropertyRequest:
	// handle ReadProperty

case *npdu.UnknownNetworkMessage:
	// log or inspect an unsupported network-layer message

case *apdu.UnknownAPDUService:
	// log or investigate an unsupported application service
}
```

rather than exposing every detail of BACnet's wire representation to every caller.

## Package Organization

The library is being organized around BACnet protocol responsibilities:

```text
bactypes/   BACnet protocol and domain value types
codec/      Primitive BACnet tags and value encoding/decoding
bvlc/       BACnet/IP BVLC messages
npdu/       BACnet network layer and routing messages
apdu/       BACnet APDU framing
services/   BACnet application services
objects/    BACnet object model
defs/       Shared BACnet protocol constants and enumerations
```

The root package provides higher-level decoding that composes these layers.

A BACnet/IP packet can therefore be decoded progressively:

```text
UDP payload
    ↓
BVLC
    ↓
NPDU
    ↓
network-layer message
    or
APDU
    ↓
application service
```

The higher-level decoder retains the decoded envelopes as well as the final semantic value, allowing applications to inspect protocol context without forcing that context into individual service structures.

## Testing

This fork is also adding protocol-level tests as functionality is reworked.

Of particular importance are **known-wire tests**: tests that begin with or compare against independently known BACnet byte sequences.

Round-trip tests are useful:

```text
Go value → encode → decode → Go value
```

but an encoder and decoder can contain matching bugs.

Known-wire tests additionally verify:

```text
known BACnet bytes ↔ expected Go value
```

and are therefore preferred for important protocol messages.

The test suite currently covers several BVLC, NPDU, APDU, and application-service encoding and decoding paths and will continue to expand as the library grows.

Run the full suite with:

```bash
go test -v ./...
```

## Project Status

This is an actively evolving fork.

The library is being refactored while additional BACnet functionality is implemented, so APIs may change as older C-derived interfaces are replaced with more idiomatic Go ones.

The emphasis is currently on:

* correctness on the wire;
* useful BACnet/IP and routing functionality;
* robust packet decoding;
* diagnostic and discovery tooling;
* clean separation between protocol representation and higher-level device behavior;
* and an API that feels natural to a Go programmer.

## BACnet Trademark and Conformance

BACnet™ is a trademark of the American Society of Heating, Refrigerating and Air-Conditioning Engineers (ASHRAE). The BACnet protocol is defined by ANSI/ASHRAE Standard 135 and related standards, which are copyrighted by ASHRAE.

This project is an independent open-source implementation and is not affiliated with, endorsed by, or approved by ASHRAE.

No claim is made that this library, or any device or application built with it, is BACnet-compliant or has passed BACnet conformance testing. The library implements portions of the BACnet protocol as needed and is under active development. Users requiring formal conformance should evaluate their implementation against the applicable BACnet standards and conformance-testing requirements.


## Attribution

This project is derived from `noahtkeller/go-bacnet`.

That project states that it was heavily influenced by Steve Karg's `bacnet-stack` project.

This fork gratefully acknowledges both projects and their contributors. The intent of this work is not to obscure that lineage, but to build on it while substantially extending the implementation and reshaping the public API for Go.

## License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

This permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
