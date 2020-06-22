# Pokemon

## Problema:

O Ash está a apanhar pokémons num mundo que consiste numa grelha bidimensional infinita de casas.
Em cada casa há exatamente um pokémon.
O Ash começa por apanhar o pokémon que está na casa onde começa. A seguir, move-se para a casa
imediatamente a norte, sul, este ou oeste de onde se encontra e apanha o pokémon que aí se encontrar,
e assim sucessivamente. Atenção: se ele passar numa casa onde já passou (e, portanto, onde já apanhou
um pokémon), já lá não está um pokémon para ele apanhar!
O que queremos saber é: começando com um mundo cheio de pokémons (um em cada casa!), quantos
pokémons o Ash apanha para uma dada sequência de movimentos?

### Formato do input
O programa deve ler uma linha do stdin, que contém uma sequência de movimentos. Cada movimento é
descrito por uma letra N, S, Eou O(respetivamente: norte, sul, este, oeste).

### Formato do output
O programa deve escrever uma linha para o stdout, apenas com um número: quantos pokémons o Ash
apanhou?

    Input: E -> Output: 2
    Dado uma tabela bidimensional *E*   
    Onde cada casa possui exatamente um pokemon
    Então quando o sistema contar 
    então o sistema de deve retornar 2

    Input: NESO -> Output: 4

    Dado uma tabela bidimensional *NESO*   
    Onde cada casa possui exatamente um pokemon
    Então quando o sistema contar 
    então o sistema de deve retornar 4

    Input: NSNSNSNSNS -> Output: 2

    Dado uma tabela bidimensional *NSNSNSNSNS*   
    Onde cada casa possui exatamente um pokemon
    Então quando o sistema contar 
    então o sistema de deve retornar 4


## Solução:

Para contar a quantidade de pokemons capturados durante o percurso do Ash, devemos levar em consideração que ele pode passar pelo mesmo ponto durante o trajeto. Para identificarmos isso, podemos afirmar que cada ponto (X,Y), ou seja, linha e coluna é único.

Pensando no plano cartesiano, o ponto de partida é o encontro dos eixos x e y, ou seja X: 0, Y: 0. Para avaliar seu deslocamento devemos seguir os cenários descritos abaixo:

    Ao movimentar-se sentido *NORTE* então a coordenada *X* deve se manter com o último valor atribuido e a coordenada *Y* deve ser incrementada. (X, Y++)

    Ao movimentar-se sentido *SUL* então a coordenada *X* deve-se manter com o último valor atribuido e a coordenada *Y* deve ser decrementada. (X, Y--)

    Ao movimentar-se sentido *LESTE* então a coordenada *X* deve ser incrementada e a coordenada *Y* deve-se manter com o último valor atribuido. (X++, Y)

    Ao movimentar-se sentido *OESTE* então a coordenada *X* deve ser decrementada e a coordenada *Y* deve-se manter com o último valor atribuido. (X--, Y)

Ao movimentar-se os pontos serão adicionados ao mapa para garantir que eles serão únicos.


## Setup

### Build

```bash
   # Build Unix
   go build

   # Build Windows
   env GOOS=windows GOARCH=amd64 go build
```

### run
```bash
   # Run Unix/Linux
   ./pokemon

   # Run Windows
   pokemon.exe
```

### tests
```bash
   go test ./...

   # Ouput coverage std out
   go test ./... -v -cover  

   # Generate coverage file out
   go test ./...  -coverprofile=./coverage/tests.out

   # Generate coverage HTML file out
   go tool cover -html=./coverage/tests.out -o ./coverage/coverage-report.html

```

### Benchnmark
```bash
   go test ./...  -run=XXX -bench=.
```

### 