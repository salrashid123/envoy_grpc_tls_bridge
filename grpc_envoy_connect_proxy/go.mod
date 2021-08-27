module main

go 1.15

require (
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	google.golang.org/grpc v1.33.2 // indirect
	echo v0.0.0
)
replace echo => "./src/echo"