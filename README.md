## Lattigo Bindings for C++

![Lattigo](https://github.com/ldsec/lattigo) is a Go (Golang) library for homomorphic encryption. This library provides a C++ interface for the Lattigo HE library.

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

This library includes an example, which matches the corresponding ![Lattigo example](https://github.com/ldsec/lattigo/blob/fa4175901a31d20b5336e7705e6a1243d0477e62/examples/ckks/sigmoid/main.go) as closely as possible.
The example can be run with
```!sh
cmake -Bbuild -GNinja -DLATTICPP_BUILD_EXAMPLES=ON
ninja -Cbuild run_gobindingtest
```

This library's API is in src/latticpp/ckks.

## API Wrapper Design
This section is intended for people who want to modify this library in some way (e.g., by adding additional Lattigo bindings). This library is organized in three logical levels: the Lattigo library itself, a thin Go wrapper on top of that, and a thin C++ wrapper on top of *that*.

The thin Go wrapper around the Lattigo API closely mirrors the Lattigo API. The reason for this API is that we cannot pass objects between Go and C++. Instead, we pass *references* or *handles* to objects. The Go wrapper maintains a global map from handles to objects. The API accepts handles to objects, then translates these handles into objects using the map, then calls the corresponding Lattigo function, and returns a handle to the result. CGo translates this API into functions which can be called by C++.

The top layer of the API stack is a C++ API that is a one-to-one mapping to the Go wrapper API. Like the Go wrapper, the C++ API only uses object handles and primitive types. This is the user-facing API bindings for Lattigo.

Note that the use of this global map would prevent the Go memory manager from freeing objects. However, the C++ code tracks handles (references) to objects and when there are no more C++ references to an object, C++ tells Go that it may free the corresponding object.

## Example
Consider the Lattigo function `func (encryptor *pkEncryptor) EncryptNew(plaintext *Plaintext) *Ciphertext`. This member function of the `encryptor` interface accepts a `*Plaintext` and returns a `*Ciphertext`. The corresponding Go wrapper is `func lattigo_encryptNew(encryptorHandle Handle, ptHandle Handle) Handle`. This function accepts a *handle* to the `encryptor` object and a *handle* to the plaintext. It returns a handle to the resulting ciphertext. The corresponding C++ function is
`GoHandle<Ciphertext> encryptNew(const GoHandle<Encryptor> &encryptor, const GoHandle<Plaintext> &pt)`.

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.
