# Tutorial de Como fazer testes unitários em go

## Explicando funções de maneira mais detalhada
> **Se quiser estudar mais funções da lib testing leia a documentação do pacote [testing](https://pkg.go.dev/testing) da stdlib do go**

A função Errorf() == ```  func (c *T) Errorf(format string, args ...any) ``` é quem imprime o erro mostrando o que foi recebido
por parâmetro e o que era esperado que a variável do parâmetro retornasse 
que foi definido na segunda variável passada por parâmetro.

Se quiser entender melhor está função e as demais 
leia a [documentação](https://pkg.go.dev/testing)


Segundo a documentação oficial a função Helper() marca a função de chamada 
como uma função auxiliar de teste.

A função Helper() == ```func (c *T) Helper() ``` ao imprimir informações de arquivo e linha, essa função será ignorada.
Helper pode ser chamado simultaneamente de múltiplas goroutines.
se não sabe o que são goroutines estude pelo [site oficial](https://go.dev/tour/concurrency/1).

a função Run() que vem da variável test 
que é um ponteiro para testing.T
que segundo a documentação oficial:

``` func (t *T) Run(name string, f func(t *T)) bool```

Run executa f como um subteste de t chamado name. 
Ele executa f em uma goroutine separada e bloqueia até que f 
retorne ou chame t.Parallel para se tornar um teste paralelo. 
Run informa se f foi bem-sucedido 
(ou pelo menos não falhou antes de chamar t.Parallel).
Run pode ser chamado simultaneamente de várias goroutines, 
mas todas essas chamadas devem retornar antes que a função de teste externa 
para t retorne.
