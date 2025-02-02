#!/usr/bin/env bash

set -e

git_root=$(git rev-parse --show-toplevel)

# shellcheck disable=SC1091
source "$git_root"/docs/tools/aperturectl/validate_common.sh

generate_from_values \
	load-ramping/base \
	values.yaml \
	tmp

# copy the generated policy and graph to this (assets) directory so that they can be used in the docs
cp tmp/policies/feature-rollout-cr.yaml policy.yaml
cp tmp/graphs/feature-rollout-cr.mmd graph.mmd

# git add the generated policy and graph
"$git_root"/scripts/git_add_safely.sh policy.yaml graph.mmd

# remove the tmp directory
rm -rf tmp
