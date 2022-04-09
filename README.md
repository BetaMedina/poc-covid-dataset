# POC povoamento e consumo de dataset coid

### O que é

 - Alguns dias para tás eu queria validar alguns conhecimentos téoricos sobre Channels no golang, durante minhas pesquisas acabei encontrando o covid dataset cujo o qual tive grande vontade de trabalhar em cima, então basicamente o projeto que irão encontrar nesse repositório é uma POC.


### Escolhas de desenvolvimento:

- Golang
- Mongodb
- Gorilla mux
- Makefile

### Funções:

- Importação de csv com dados sobre o covid
- Disponibilização de api para consumo de dados

### Como utilizar esse projeto:

- Run `make import` para povoar o banco
- Run `make run` para iniciar em desenvolvimento.


#### Dicas
- Será nescessario baixar o csv: [Covid dataset](https://brasil.io/dataset/covid19/caso_full/), após o download mova o dataset para a pasta /public

- O mongo pode ter alguns problemas de performance que irão impactar  em sua utilização devido aos mecanismos de busca, sugestão pessoal é criar indexes para trabalhar em cima dele
```
db.covid_data.createIndex({"state":1})
db.covid_data.createIndex({"city":1})
```
Com esses indexes o tempo de resposta irã diminuir colossalmente 
