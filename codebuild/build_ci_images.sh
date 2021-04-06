#!/bin/bash -ex
# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

if [ -z ${1+x} ]; then
  ECS_REPO="465824340922.dkr.ecr.us-west-2.amazonaws.com/latticpp-ci"
else
  ECS_REPO=$1
fi

echo "Uploading docker images to ${ECS_REPO}."

$(aws ecr get-login --no-include-email --region us-west-2)

docker build -f ci_ubuntu-20.04_clang-10.docker -t latticpp_ubuntu-20.04:base .
docker tag latticpp_ubuntu-20.04:base ${ECS_REPO}:latticpp_ubuntu-20.04_base
docker push ${ECS_REPO}:latticpp_ubuntu-20.04_base
