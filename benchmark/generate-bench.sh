export FILE=./coverage/bench.out
export FNAME=AmountOfCapturedPokemons
go test ./capture  -run=XXX -bench=. > go > ${FILE} 
cat ${FILE}|benchgraph -title="Graph2: Benchmark ${FNAME} ns/op" 