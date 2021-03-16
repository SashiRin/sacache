# sacache

SaCache is a fast, concurrent in-memory cache service written in pure Go. It uses gRPC for communicating with clients and supports item expiration function. It is still under development. The goal of SaCache is reducing extra costs(GC) as much as possible and making it faster and faster.

## Usage

Please read [sacache_test.go](sacache_test.go) for lib usage purpose and read [server.go](server/server.go) & [client.go](client/client.go) for C-S usage purpose.

## License

SaCache is released under the MIT license (see [LICENSE](LICENSE))
