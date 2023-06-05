# random-generator-API
generator of the random numbers or strings for my employer

# generate documentation
`swag swag init -d cmd`

# run project
`go run cmd/main.go`

# endpoints
`127.0.0.1:8000/generator/generate`
METHOD [POST]
{
  "amount": int
}
Generates amount of random numbers and strings

`127.0.0.1:8000/generator/result`
METHOD [GET]
Returns last generated random sequence in the string format
