# nontrivial assumptions

input can be split into these parts:
1. translation of intergalactic units to roman numeral
    - intergalactic units are space separated
    - throws error on:
      - invalid roman numeral
      - duplicate intergalactic unit
    - notes:
      - duplicate roman unit does not throw error since we can still translate intergalactic to roman this way (and we are not required to do it the other way around)
2. declaration of intergalactic items price in _Credits_
    - throws error on:
      - intergalactic units have not been previously defined
      - intergalactic units convert into invalid roman numeral
3. questions on intergalactic unit value or item price

# system design solution

- while a strict parts order validation can be imposed for faster parsing, this repo validates on the go i.e `tests/in/non-ordered.txt`
- in general, error will not be thrown if statements that fall under 1 & 2 are successfully validated, default answer will be used instead
- should one needs to implement
    - a new regex parser with different format or business logic (easily replacing parts mentioned above) -> `Parser`
    - use different storage solution -> `Store`

# compiling and invoking the program

tested on linux, go 1.14, from root directory of this repo:
1. using go file
  - `cd cmd`
  - `go run main.go <path_to_file>` e.g using the current structure it would be `go run main.go ../tests/in/default.txt`
2. using binary
  - `cd cmd`
  - `go build -o prospace`
  - `./prospace <path_to_file>` e.g using the current structure it would be `./prospace ../tests/in/default.txt`
3. using docker
  - why would you want to use docker for this

result will be printed to terminal, to save, simply add `> <path_to_out_file>` after above command (i dont use windows, so have not tested it there) or modify `cmd/main.go`

# you probably do not really care but

- in `internal/unit.go` and `internal/item.go`, both struct and map are mentioned. benchmark using `go test ./... -bench .` reveals fluctuating result on which is more efficient, but usually maps is more efficient, perhaps due to the many conversion needed and inefficienct implementation of `Store` interface (I refrain from implementing other KV database because I would like to do this exercise only using Go stdlib). eventually, struct is used for type safety, with the tradeoff of a bit more complexity (to be fair it is only a few extra lines), hence this experiment is not repeated in `internal/question.go`
- short variable name as per https://talks.golang.org/2014/names.slide
- this repo has successfully ignored `YAGNI` and `KISS`
