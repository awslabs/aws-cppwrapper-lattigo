module lattigo-cpp

go 1.15

require github.com/ldsec/lattigo/v2 v2.1.1

replace (
	github.com/ldsec/lattigo/v2 => ../../external/lattigo/
	lattigo-cpp => ./
)
