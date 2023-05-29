module github.com/playsistemico/curso-go

go 1.20

require (
	example.com/electric_vehicle v0.0.0-00010101000000-000000000000
	example.com/examples v0.0.0-00010101000000-000000000000
	example.com/tuiter v0.0.0-00010101000000-000000000000
	example.com/utils v0.0.0-00010101000000-000000000000
	example.com/utils/benchmark v0.0.0-00010101000000-000000000000
)

require example.com/utils/file_system v0.0.0-00010101000000-000000000000 // indirect

replace example.com/electric_vehicle => ./internal/electric_vehicle

replace example.com/tuiter => ./internal/tuiter

replace example.com/examples => ./internal/examples

replace example.com/utils => ./utils

replace example.com/utils/benchmark => ./utils/benchmark

replace example.com/utils/file_system => ./utils/file_system
