## C++ Bindings for the Lattigo Homomorphic Encryption Library

This library provides partial C++ bindings for the [Lattigo v4.1.0](https://github.com/ldsec/lattigo/releases/tag/v4.1.0) Go-lang homomorphic encryption library. This wrapper does not attempt to provide a bindings for all public Lattigo APIs, but new bindings are quite easy to add and PRs are welcome.

## Building

Status:
* Ubuntu 20.04 is
![Build Status](https://codebuild.us-west-2.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiVkVjR3d0UVVQKys0Rk5sUy9UVEUyNGRhbUJoSWdTZ2pXdXpYVlQ5RVVpYXBraFdkOWpTNHk1QUljUy90a3JIMS84UERuaVBIR2ZOWVJjN2N6QzFFMzZVPSIsIml2UGFyYW1ldGVyU3BlYyI6IjNxdENBZlZNU210VHdndHciLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=main)

## Usage

This library uses the CMake build system. You can build the library by running

```!sh
cmake -Bbuild -GNinja
ninja -Cbuild
```

If you are using clang, you may need to compile dependent libraries with `-Wno-c99-extensions` to suppress warnings in the cgo-generated header files.

This library includes examples which matches some of the corresponding [Lattigo examples](https://github.com/tuneinsight/lattigo/tree/5b707142db0fc16acad96c1e46e7a9d68fb5b014/examples) as closely as possible.
The examples can be run with
```!sh
cmake -Bbuild -GNinja -DLATTICPP_BUILD_EXAMPLES=ON
ninja -Cbuild run_bootstrapexample
ninja -Cbuild run_eulerexample
ninja -Cbuild run_multikeyexample
```

This library's API is in src/latticpp/ckks. This library was tested with Go version 1.15.8. This library makes use of the `unsafe` Go package, so there is a small chance that newer versions of Go might be incompatible with this library.

## API Wrapper Design

This section is intended for people who want to modify this library in some way (e.g., by adding additional Lattigo bindings). This library is organized in three logical levels: the Lattigo library itself, a thin Go wrapper on top of that, and a thin C++ wrapper on top of *that*.

The thin Go wrapper around the Lattigo API closely mirrors the Lattigo API. The reason for this API is that we cannot pass objects between Go and C++. Instead, we pass *references* or *handles* to objects. The Go wrapper maintains a global map from handles to objects. The API accepts handles to objects, then translates these handles into objects using the map, then calls the corresponding Lattigo function, and returns a handle to the result. CGo translates this API into functions which can be called by C++.

The top layer of the API stack is a C++ API that is a one-to-one mapping to the Go wrapper API. Like the Go wrapper, the C++ API only uses object handles and primitive types. This is the user-facing API bindings for Lattigo.

Note that the use of this global map would prevent the Go memory manager from freeing objects. However, the C++ code tracks handles (references) to objects and when there are no more C++ references to an object, C++ tells Go that it may free the corresponding object.

## Example

Consider the Lattigo function `func (encryptor *pkEncryptor) EncryptNew(plaintext *Plaintext) *Ciphertext`. This member function of the `encryptor` interface accepts a `*Plaintext` and returns a `*Ciphertext`. The corresponding Go wrapper is `func lattigo_encryptNew(encryptorHandle Handle, ptHandle Handle) Handle`. This function accepts a *handle* to the `encryptor` object and a *handle* to the plaintext. It returns a handle to the resulting ciphertext. The corresponding C++ function is
`GoHandle<Ciphertext> encryptNew(const GoHandle<Encryptor> &encryptor, const GoHandle<Plaintext> &pt)`.

## Security

Lattigo is an experimental library and should only be used for research purposes. This wrapper directly calls Lattigo, and does not provide any additional security measures. Use it with care, for RESEARCH ONLY.
See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.
