<br />
<p align="center">
    <a href="#">
        <img src="https://cdn.labstack.com/images/echo-logo.svg" alt="Logo" width="240" height="80">
    </a>
    <h3 align="center">goecho-projectbase-api</h3>

   <p align="center">
        Projeto construído em Go com Echo framework para ser utilizado como base de uma API.
   </p>
</p>

<!-- ABOUT THE PROJECT -->
## About the Project

Este projeto tem como objetivo fornecer uma base para iniciar novas APIs em GO utilizando o Echo Framework.



### Environments

* [https://echo.labstack.com/](https://echo.labstack.com/)



<!-- GETTING STARTED -->
## Getting Started

Após efetuar um git clone da API, todos os comandos, tanto de build, quanto de debug, são feitos a partir da pasta raiz do projeto.
Existe um Makefile para auxiliar no build do projeto.
Outros arquivos como o ".devcontainer/devcontainer.json" e o docker-compose.yml vão auxiliar a efetuar o remote debug do seu container.

### Prerequisites

* Docker
  ```sh
  docker --version
  ```

* Docker Compose
  ```sh
  docker-compose --version
  ```  

* Make


### Installation / Usage

1. Clone the repo
   ```sh
   git clone https://github.com/felipeolivers/goecho-projectbase-api.git
   ```

2. Execute go command to download dependencies
   ```sh
   go mod tidy
   ```

3. Execute make command to build
   ```sh
   make build
   ```

### Debug
Para efetuar o debug localmente na máquina do desenvolvedor necessita copiar o .env para dev.env e alterar os valores respectivos para conexão com o Elastic, banco de dados e demais variáveis.
Também é necessário alterar a instrução no arquivo main.go: 

Mudar de:
   ```sh
  err := godotenv.Load()
   ```
Para:
   ```sh
  err := godotenv.Load("dev.env")
   ``` 

Ao subir o container, para rodar a aplicação de forma debugável e com live reload.
Usar o seguinte comando no bash:
   ```sh
  sh ./scripts/run.sh 
   ```
Este script permitirá que a cada save no projeto, um reload na API seja feito.

### k8s Set Image Deployment (DEV)
Para efetuar o set image da imagem no k8s, necessita criar no diretório raiz no projeto uma pasta com o nome de "k8s" e dentro da mesma criar o arquivo "c3.conf".
* Solicitar o modelo do arquivo para uma pessoa do time.   