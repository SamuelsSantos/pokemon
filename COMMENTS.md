## Análises

Ao usar fmt.Sprintf("%v\n", currentPoint) para mostrar as informações a performance cai drasticamente, como mostram as tabelas a seguir:

    * Com fmt.Sprintf *
  
    goos: darwin
    goarch: amd64
    pkg: pokemon
    Benchmark1000-4            3000     477061 ns/op
    Benchmark50000-4            500    2401046 ns/op
    Benchmark10000-4             20   54163742 ns/op
    Benchmark500000-4             5  286287970 ns/op
    Benchmark1000000-4            2  641935163 ns/op
    Benchmark50000000-4           1 3194655679 ns/op


    * Sem fmt.Sprintf *
  
    goos: darwin
    goarch: amd64
    pkg: pokemon
    Benchmark1000-4           20000      82682 ns/op
    Benchmark50000-4           3000     406000 ns/op
    Benchmark10000-4            100   11639932 ns/op
    Benchmark500000-4            20   78074864 ns/op
    Benchmark1000000-4           10  152001483 ns/op
    Benchmark50000000-4           2  822659591 ns/op