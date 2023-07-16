
<p align="center">
  <img src="https://raw.githubusercontent.com/BrasilAPI/BrasilAPI/master/public/brasilapi-logo-small.png" alt="Sublime's custom image"/>
</p>

<hr>
<br>

SDK desenvolvido com o objetivo de facilitar a integração de sistemas e API's desenvolvidos em Golang com a [BrasilApi](https://brasilapi.com.br/).



### Importando o SDK para seu projeto

<br>

Para que você possa utilizar este SDK em seu projeto, primeiramente é necessário rodar o seguinte comando para que o go possa instala-lo:

<br>

~~~bash
    go get -u https://github.com/philipelima/brasilapi-go
~~~

<br>

Após a instalação, você poderá importar os packages da seguinte forma:

~~~go
    import "https://github.com/philipelima/brasilapi-go/{package_name}"
~~~


<br>

### Utilizando o SDK

Veja a baixo, exemplos de utilizações do SDK.

<br>

#### CEP

<br>

A api BrasilApi disponibiliza duas versões para consulta de cep, veja a baixo a consulta por meio da versão **V1** e **V2**.


~~~go

  package main

  import "github.com/philipelima/brasilapi-go/cep"


  func main() {

    // Consulta de Cep na versão 1
    cepV1, errV1 := cep.V1().Get("47000000")

    // Consulta de Cep na versão 2
    cepV2, errV2 := cep.V2.Get("47000000")

  }
  
~~~