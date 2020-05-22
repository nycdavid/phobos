module github.com/nycdavid/phobos

go 1.14

require (
	github.com/gin-contrib/multitemplate v0.0.0-20200514145638-4955c9347179
	github.com/gin-gonic/gin v1.6.3
	github.com/lib/pq v1.5.2
	github.com/nycdavid/phobos/dbconnector v0.0.0-20200515133311-cf4303de984d
	github.com/nycdavid/phobos/migrator v0.0.0-20200516003550-b17bad812e05 // indirect
)

replace github.com/nycdavid/phobos/dbconnector => ./dbconnector/

replace github.com/nycdavid/phobos/migrator => ./migrator/
