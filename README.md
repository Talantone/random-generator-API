# random-generator-API
generator of the random numbers or strings for my employer

# generate documentation
`swag swag init -g cmd/main.go`

# run project
`go run cmd/main.go`

# endpoints
`127.0.0.1:8000/generator/generate`
METHOD [POST]
{
  "amount": int
}
Generates amount of random numbers and strings

`127.0.0.1:8000/generator/result/:id`
METHOD [GET]
Returns element of the last generated random sequence in the string format
