
<p align="center">
  <img src="https://raw.githubusercontent.com/BrasilAPI/BrasilAPI/master/public/brasilapi-logo-small.png" alt="Sublime's custom image"/>
</p>

<hr>
<br>

SDK desenvolvido com o objetivo de facilitar a integração de sistemas e serviços desenvolvidos em Golang com a [BrasilApi](https://brasilapi.com.br/).



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

#### BANCOS

Buscando informações sobre o sistema bancário brasileiro.

~~~go
  import "github.com/philipelima/brasilapi-go/banks"
~~~

**Consultando todos os bancos:**

~~~go
  banks, err := banks.V1().All()
~~~

**Consultando um banco especifico:**

Para consultar um banco especifico, é necessário saber o código de compensação do banco que você deseja consultar. Veja tabela de códigos aqui [Aqui](https://www.bcb.gov.br/Fis/CODCOMPE/Tabela.pdf). 


~~~go
  code := 237 // Código Banco Bradesco S.A 

  bank, err := banks.V1().Get(code)
~~~

<br>

#### CEP

A api [BrasilApi](https://brasilapi.com.br/) disponibiliza duas versões para consulta de cep, veja a baixo a consulta por meio da versão **V1** e **V2**.


~~~go
  import "github.com/philipelima/brasilapi-go/cep"
~~~

**V1:**

~~~go
  cep, err := cep.V1().Get("47000000")
~~~

**V2:**

~~~go
  cep, err := cep.V2().Get("47000000")
~~~