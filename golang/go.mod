module main

go 1.14

require (
	github.com/google/uuid v1.1.1 // indirect
	github.com/ldsec/lattigo v1.3.0
	github.com/salrashid123/fhe/rideshare v0.0.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/salrashid123/fhe/rideshare => ./src/github.com/salrashid123/fhe/rideshare
