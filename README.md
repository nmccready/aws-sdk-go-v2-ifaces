# aws-sdk-go-v2-ifaces

To provide a more idiomatic interfaces and mocks to the AWS SDK for Go v2 since AWS refuses to do so.

## Dependencies

- Node.js 20.x (use nvm or [proto](https://moonrepo.dev/proto))
- npm = 8.X
- Go >= 1.20 ([proto](https://moonrepo.dev/proto))
- go mockery `go install github.com/vektra/mockery/v2@v2.43.2`

## Local Development

- npm i (installs node and go dependencies)
- npm run sync (clones repo aws-sdk-go-v2 to ./tmp/aws-sdk-go-v2, then greps interfaces)
