# Projeto: Upload de Repositório Git como ZIP no S3 LocalStack

Este projeto clona um repositório Git, compacta o conteúdo em um arquivo ZIP e faz upload para um bucket S3 simulado pelo LocalStack.

## Pré-requisitos

- [Go instalado](https://go.dev/doc/install)
- [Docker e Docker Compose instalados](https://docs.docker.com/get-docker/)
- [AWS CLI instalada (opcional, para testes)](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- Git instalado

## Passo 1: Clone este repositório

```bash
git clone https://github.com/seu-usuario/seu-repo.git
cd seu-repo