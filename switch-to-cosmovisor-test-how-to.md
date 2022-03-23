# Switch to Cosmovisor Test How To

This manual describes how to test switching from the version `0.7.0` which
directly uses `dcld` binary to the version `0.9.0` which uses `cosmovisor` to
control `dcld` binary.

## Docker Pool for Tests

This test uses a customized docker pool originated from the deployment test
(defined in `integration_tests/deploy/docker-compose.yml`). It conatians 2
Ubuntu nodes: each one using systemd service for running DCL node process.

## Test Setup

- Check out
  https://github.com/spivachuk/distributed-compliance-ledger/tree/env-for-switch-to-cosmovisor-test
- Run `./integration_tests/run-all.sh deploy`

In result of these steps a pool with 2 Ubuntu nodes is created. On one of the
nodes (`test_deploy_gvn`) a genesis validator is deployed. Then on another node
(`test_deploy_vn`) a validator is deployed. This validator is added to the pool
of the genesis validator. Each node runs `dcld` as a systemd service. The `dcld`
is built from the local project (however, building is performed within a docker
container). Source code in `env-for-switch-to-cosmovisor-test` branch neary
corresponds to `v0.7.0`. The resulting ledger contains 3 trustees whose keys are
presented in the client wallet on the genesis validator node. They are named
`jack`, `alice` and `bob` there. So the further tx commands are supposed to be
sent from the client on the genesis validator node. Query commands can be sent
from both the client on the genesis validator node and the client on the
validator node.

## Procedure of Node Upgrade from v0.7.0 to v0.9.0/cosmovisor

- Check out
  https://github.com/spivachuk/distributed-compliance-ledger/tree/switch-to-cosmovisor-test

`switch-to-cosmovisor-test` branch contains two scripts
`switch_gvn_to_cosmovisor.sh` and `switch_vn_to_cosmovisor.sh` (in the root
directory). Each of them provisions the corresponding node (the genesis
validator node or the validator node) with `switch_to_cosmovisor` script and
locally gotten/built/installed assets needed by it and then runs
`switch_to_cosmovisor` script within that node. Source code in
`switch-to-cosmovisor-test` branch neary corresponds to `v0.9.0`. To switch
`v0.7.0` node to the local project version and `cosmovisor` run the following
commands:

- `go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0`
- `make build`
- `make install`
- `./switch_gvn_to_cosmovisor.sh` or `./switch_vn_to_cosmovisor.sh`
  (correspondingly to the node to switch)

## Run Tests on Prepared Pool

- Check out
  https://github.com/spivachuk/distributed-compliance-ledger/tree/env-for-switch-to-cosmovisor-test

`env-for-switch-to-cosmovisor-test` branch contains the script
`run_tests_from_gvn.sh` (in the root directory) for running CLI integration
tests on the prepared pool. `run_tests_from_gvn.sh` can be potentially run on
various more or less recent DCL version(s) installed on the pool nodes (the
versions can be potentially different between the nodes). At first,
`run_tests_from_gvn.sh` configures the clients on both the nodes in a proper
way, then copies CLI integration tests from the local project to the genesis
validator node and then run them from the client on the genesis validator node.

`run_tests_from_gvn.sh` runs the following CLI integration tests (in the
specified order):

- `model-demo.sh`
- `modelversion-demo.sh`
- `auth-demo.sh`
- `compliance-demo.sh`
- `pki-demo.sh`
- `vendorinfo-demo.sh`

To run the tests just execute `./run_tests_from_gvn.sh`

Since `run_tests_from_gvn.sh` gets the tests from the local project, so, if you
need some other revision of these tests, you can check out the corresponding
commit but just keep `run_tests_from_gvn.sh` file in the local project to have
ability to run it then.
