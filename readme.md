# Grinnix Judge

O **Grinnix Judge** é um serviço backend responsável por executar código enviado por usuários em um ambiente isolado, inspirado em plataformas como LeetCode.

Fluxo principal:

Recebe código → Executa em container isolado → Retorna resultado

Objetivo:

Avaliar código com segurança, isolamento e controle de recursos

## Modelo lógico da execução dos códigos

[Client]
↓
[API Go]
↓
[Executor]
↓
[Docker Container (Language)]
↓
[stdout / stderr]
↓
[Response JSON]

## Estrutura do Projeto

.
├── cmd # Entry points da aplicação
│ └── api
│ └── main.go # Inicialização da API HTTP
├── docker # Ambientes de execução isolados
│ └── ts-runner
│ ├── dockerfile # Imagem para executar TypeScript
│ └── run.sh # Script de execução dentro do container
├── go.mod # Dependências do projeto Go
└── internal # Código interno da aplicação
└── executor
└── docker.go # Executor que interage com Docker

| Arquivo                       | Responsabilidades                                                                                                               |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `cmd/api/main.go`             | - Sobe servidor HTTP <br> - Define rotas <br> - Recebe código do usuário <br> - Retorna resultado da execução                   |
| `internal/executor/docker.go` | - Executa containers Docker <br> - Aplica limites de recursos <br> - Captura stdout/stderr <br> - Retorna resultado estruturado |
| `docker/ts-runner/dockerfile` | - Define ambiente Node.js <br> - Instala tsx <br> - Cria usuário não-root <br> - Configura ambiente seguro                      |
| `docker/ts-runner/run.sh`     | - Recebe código via variável de ambiente <br> - Cria arquivo main.ts <br> - Executa com timeout                                 |

## Segurança aplicada

| Medida de Segurança            | Descrição                                                           | Motivo / Benefício                                                                                |
| ------------------------------ | ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Container efêmero (`--rm`)     | Remove automaticamente o container após a execução                  | Evita acúmulo de containers, reduz superfície de ataque e impede persistência de código malicioso |
| Sem rede (`--network none`)    | Desabilita completamente o acesso à rede dentro do container        | Impede exfiltração de dados, ataques externos e uso indevido da infraestrutura                    |
| Limite de memória (`--memory`) | Restringe a quantidade de RAM disponível para o container           | Evita consumo excessivo de memória (DoS) e travamento do host                                     |
| Limite de CPU (`--cpus`)       | Controla a quantidade de CPU que o container pode utilizar          | Previne abuso de processamento e garante estabilidade do sistema                                  |
| Timeout de execução            | Interrompe o processo após um tempo limite definido                 | Evita loops infinitos e execuções prolongadas maliciosas                                          |
| Usuário não-root               | Executa o código dentro do container com um usuário sem privilégios | Reduz impacto de possíveis exploits e impede acesso a recursos sensíveis do sistema               |

---

## Como rodar

### 1. Dar permissão ao script

```
chmod +x ./docker/ts-runner/run.sh
```

### 2. Build da imagem Docker

```
docker build -t ts-runner ./docker/ts-runner
```

### 3. Rodar a API

```
go run ./cmd/api
```

Servidor disponível em:
http://localhost:8080

[Link com os testes rápidos de execução](./docs/test.md)
