module github.com/playsistemico/curso-go

go 1.20

require (
	example.com/electric_vehicle v0.0.0-00010101000000-000000000000
	example.com/examples v0.0.0-00010101000000-000000000000
	example.com/tuiter v0.0.0-00010101000000-000000000000
)

replace example.com/electric_vehicle => ./internal/electric_vehicle

replace example.com/tuiter => ./internal/tuiter

replace example.com/examples => ./internal/examples
