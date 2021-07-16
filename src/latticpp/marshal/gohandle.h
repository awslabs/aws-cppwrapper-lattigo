// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#ifndef DEFINE_GOHANDLE_H
#define DEFINE_GOHANDLE_H

#include <cstdint>
#include "cgo/storage.h"
#include <iostream>

namespace latticpp {

    enum class GoType {
        Bootstrapper,
        BootstrappingParameters,
        BootstrappingKey,
        EvaluationKey,
        Parameters,
        Encoder,
        KeyGenerator,
        KeyPair,
        RelinearizationKey,
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
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                decref(handle);
            }
        }

        // copy constructor: increment the references to this handle
        GoHandle(const GoHandle& other) {
            handle = other.handle;
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                incref(handle);
            }
        }

        // copy assignment operator: we are overwriting the contents of this handle,
        // so decrement the references to the current handle, then increment the
        // references to the new handle
        GoHandle& operator= (const GoHandle& other) {
            if (handle == other.handle) {
                return *this;
            }
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                decref(handle);
            }
            handle = other.handle;
            if (handle != 0) {
                incref(handle);
            }
            return *this;
        }

        // move contructor: the moved-from object *will still be destructed*, so increment the ref counter
        // https://stackoverflow.com/a/20589077/925978
        GoHandle (const GoHandle&& other) noexcept {
            handle = other.handle;
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                incref(handle);
            }
        }

        // move assignment operator: we are overwriting the contents of this handle,
        // so decrement references to the current handle, and increment references to the moved-from handle
        // (see move constructor for details)
        GoHandle& operator= (const GoHandle&& other) noexcept {
            if (handle == other.handle) {
                return *this;
            }
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                decref(handle);
            }
            handle = other.handle;
            if (handle != 0) {
                incref(handle);
            }
            return *this;
        }

        // move contructor: the moved-from object *will still be destructed*, so increment the ref counter
        // https://stackoverflow.com/a/20589077/925978
        GoHandle (GoHandle&& other) noexcept {
            handle = other.handle;
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                incref(handle);
            }
        }

        // move assignment operator: we are overwriting the contents of this handle,
        // so decrement references to the current handle, and increment references to the moved-from handle
        // (see move constructor for details)
        GoHandle& operator= (GoHandle&& other) noexcept {
            if (handle == other.handle) {
                return *this;
            }
            // a handle of 0 is an invalid Go reference (my equivalent of a nil/null pointer)
            if (handle != 0) {
                decref(handle);
            }
            handle = other.handle;
            if (handle != 0) {
                incref(handle);
            }
            return *this;
        }

        bool operator == (const GoHandle& other) const {
           return handle == other.handle;
        }

        bool operator != (const GoHandle& other) const {
           return handle != other.handle;
        }

        uint64_t getRawHandle() const {
            return handle;
        }

    private:
        uint64_t handle;
    };

    using Bootstrapper = GoHandle<GoType::Bootstrapper>;
    using BootstrappingKey = GoHandle<GoType::BootstrappingKey>;
    using BootstrappingParameters = GoHandle<GoType::BootstrappingParameters>;
    using Parameters = GoHandle<GoType::Parameters>;
    using Encoder = GoHandle<GoType::Encoder>;
    using KeyGenerator = GoHandle<GoType::KeyGenerator>;
    using KeyPair = GoHandle<GoType::KeyPair>;
    using RelinearizationKey = GoHandle<GoType::RelinearizationKey>;
    using EvaluationKey = GoHandle<GoType::EvaluationKey>;
    using Encryptor = GoHandle<GoType::Encryptor>;
    using Decryptor = GoHandle<GoType::Decryptor>;
    using Evaluator = GoHandle<GoType::Evaluator>;
    using SecretKey = GoHandle<GoType::SecretKey>;
    using PublicKey = GoHandle<GoType::PublicKey>;
    using Plaintext = GoHandle<GoType::Plaintext>;
    using Ciphertext = GoHandle<GoType::Ciphertext>;
    using RotationKeys = GoHandle<GoType::RotationKeys>;

}  // namespace latticpp
#endif
