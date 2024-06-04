module bankapi

go 1.22.3

require (
   github.com/msft/bank v0.0.1
)

replace (
  github.com/msft/bank => ../bankcore
)
