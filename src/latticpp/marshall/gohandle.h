// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include <cstdint>
#include "cgo/storage.h"
#include <iostream>

namespace latticpp {

    enum class GoType {
        Parameters,
        Encoder,
        KeyGenerator,
        KeyPair,
        RelinKey,
        Encryptor,
        Decryptor,
        Evaluator,
        SecretKey,
        PublicKey,
        Plaintext,
        Ciphertext,
        RotationKeys
    };

    template<GoType t>
    struct GoHandle {
    public:
        // Valid Go handles are > 0, so 0 can never refer to a real Go object.
        GoHandle() : handle(0) { }

        // constructor: This should only be called by Go code, which automatically sets the reference count to 1
        GoHandle(uint64_t handle) : handle(handle) { }

        // destructor: decrement the references to this handle
        ~GoHandle() {
            decref(handle);
        }

        // copy constructor: increment the references to this handle
        GoHandle(const GoHandle& other) {
            handle = other.handle;
            incref(handle);
        }

        // copy assignment operator: we are overwriting the contents of this handle,
        // so decrement the references to the current handle, then increment the
        // references to the new handle
        GoHandle& operator= (const GoHandle& other) {
            if (handle == other.handle) {
                return *this;
            }
            decref(handle);
            handle = other.handle;
            incref(handle);
            return *this;
        }

        // move contructor: the moved-from object *will still be destructed*, so increment the ref counter
        // https://stackoverflow.com/a/20589077/925978
        GoHandle (const GoHandle&& other) {
            handle = other.handle;
            incref(handle);
        }

        // move assignment operator: we are overwriting the contents of this handle,
        // so decrement references to the current handle, and increment references to the moved-from handle
        // (see move constructor for details)
        GoHandle& operator= (const GoHandle&& other) {
            if (handle == other.handle) {
                return *this;
            }
            decref(handle);
            handle = other.handle;
            incref(handle);
            return *this;
        }

        uint64_t getRawHandle() const {
            return handle;
        }

    private:
        uint64_t handle;
    };

    using Parameters = GoHandle<GoType::Parameters>;
    using Encoder = GoHandle<GoType::Encoder>;
    using KeyGenerator = GoHandle<GoType::KeyGenerator>;
    using KeyPair = GoHandle<GoType::KeyPair>;
    using RelinKeys = GoHandle<GoType::RelinKey>;
    using Encryptor = GoHandle<GoType::Encryptor>;
    using Decryptor = GoHandle<GoType::Decryptor>;
    using Evaluator = GoHandle<GoType::Evaluator>;
    using SecretKey = GoHandle<GoType::SecretKey>;
    using PublicKey = GoHandle<GoType::PublicKey>;
    using Plaintext = GoHandle<GoType::Plaintext>;
    using Ciphertext = GoHandle<GoType::Ciphertext>;
    using GaloisKeys = GoHandle<GoType::RotationKeys>;

}  // namespace latticpp
