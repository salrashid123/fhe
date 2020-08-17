module main

go 1.14

require (
	github.com/apache/beam v2.23.0+incompatible // indirect
	google.golang.org/api v0.30.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/ldsec/lattigo v1.3.0
	google.golang.org/protobuf v1.25.0 // indirect	
	saltextio v0.0.0
)


replace "saltextio" => "./src"
