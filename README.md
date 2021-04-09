## Lattigo Bindings for C++

![Lattigo](https://github.com/ldsec/lattigo) is a Go (Golang) library for homomorphic encryption. This library provides a C++ interface for the Lattigo HE library.

## Building
Status:
* Ubuntu 20.04 is
![Build Status](https://codebuild.us-west-2.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiVkVjR3d0UVVQKys0Rk5sUy9UVEUyNGRhbUJoSWdTZ2pXdXpYVlQ5RVVpYXBraFdkOWpTNHk1QUljUy90a3JIMS84UERuaVBIR2ZOWVJjN2N6QzFFMzZVPSIsIml2UGFyYW1ldGVyU3BlYyI6IjNxdENBZlZNU210VHdndHciLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=main)

## Usage
This library includes an example, which matches the corresponding Lattigo example as closely as possible. This library uses the CMake build system. You can build the library by running

```!sh
cmake -Bbuild -GNinja
ninja -Cbuild
```

The example can be run with
```!sh
cmake -Bbuild -GNinja -DLATTICPP_BUILD_EXAMPLES=ON
ninja -Cbuild run_gobindingtest
```

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.
