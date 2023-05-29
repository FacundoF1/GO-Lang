module example.com/electric_vehicle

go 1.20

require (
	example.com/utils v0.0.0-00010101000000-000000000000
	example.com/utils/file_system v0.0.0-00010101000000-000000000000
)

replace example.com/utils => ../../utils

replace example.com/utils/file_system => ../../utils/file_system
