# Projeto: Upload de Repositório Git como ZIP no S3 LocalStack

Este projeto clona um repositório Git, compacta o conteúdo em um arquivo ZIP e faz upload para um bucket S3 simulado pelo LocalStack.

## Pré-requisitos

- [Go](https://golang.org/) instalado
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) instalados
- [AWS CLI](https://aws.amazon.com/cli/) instalada (opcional, para testes)
- [Git](https://git-scm.com/) instalado

## Passo 1: Clone este repositório
````bash
git clone https://github.com/joaobcandido/upload-repositorio-git-no-s3-localStack.git
cd upload-repositorio-git-no-s3-localStack
````

## Passo 2: Suba o LocalStack com Docker Compose
````bash
docker-compose up -d
````
## Passo 3: Crie o bucket S3 no LocalStack

````bash
aws --endpoint-url=http://localhost:4566 s3 mb s3://meu-bucket
````
## Passo 4: Instale as dependências Go
````bash
go mod tidy
````
Se não existir, inicialize o módulo:
````bash
go mod init nome-do-projeto
go mod tidy
````
## Passo 5: Execute o código
````bash
go run main.go
````
O programa irá:
1. Clonar o repositório definido no código.
2. Compactar o conteúdo clonado em um arquivo ZIP.
3. Fazer upload do arquivo ZIP para o bucket S3 do LocalStack.
   
## Passo 6: Verifique o arquivo no S3
````bash
aws --endpoint-url=http://localhost:4566 s3 ls s3://meu-bucket
````
Você deverá ver o arquivo ZIP enviado.
## Observações
O endpoint do S3 no LocalStack é sempre http://localhost:4566.
O bucket utilizado é meu-bucket (pode ser alterado no código).
O repositório Git clonado é definido na variável repoURL no código Go.

## Dicas
- Para reiniciar o LocalStack, use:
 ````bash
  docker-compose down 
  docker-compose up -d
````
- Para limpar o bucket, use:
````bash
aws --endpoint-url=http://localhost:4566 s3 rm s3://meu-bucket --recursive
````







  
