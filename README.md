# HTTP Server Projeto Korp

Projeto desenvolvido como desafio técnico utilizando **Go, Docker, Docker Compose, Nginx, Prometheus, Grafana, Ansible, Git e GitHub**.

O objetivo é demonstrar a construção de uma aplicação HTTP simples e sua evolução para uma arquitetura containerizada, monitorada, automatizada e reproduzível.

O projeto reúne conceitos de:

- desenvolvimento backend;
- API/servidor HTTP;
- containerização;
- Docker multi-stage build;
- proxy reverso;
- redes Docker;
- orquestração de containers;
- métricas;
- observabilidade;
- monitoramento;
- dashboards;
- provisionamento automático;
- automação de deploy;
- infraestrutura como código;
- Git e GitHub.

---

## Sumário

1. [Visão Geral do Projeto](#1-visão-geral-do-projeto)
2. [Roteiro de Avaliação Técnica para Recrutadores](#2-roteiro-de-avaliação-técnica-para-recrutadores)
3. [Tecnologias Utilizadas](#3-tecnologias-utilizadas)
4. [Aplicação Go](#4-aplicação-go)
5. [Docker e Docker Compose](#5-docker-e-docker-compose)
6. [Nginx](#6-nginx)
7. [Prometheus](#7-prometheus)
8. [Grafana e Provisionamento](#8-grafana-e-provisionamento)
9. [Ansible](#9-ansible)
10. [Arquitetura de Pastas](#10-arquitetura-de-pastas)
11. [Rede Docker](#11-rede-docker)
12. [Fluxo Completo da Aplicação](#12-fluxo-completo-da-aplicação)
13. [Execução Manual](#13-execução-manual)
14. [Testes e Diagnóstico](#14-testes-e-diagnóstico)
15. [Git e GitHub](#15-git-e-github)
16. [Conclusão](#16-conclusão)

---

# 1. Visão Geral do Projeto

A aplicação principal foi desenvolvida utilizando **Go (Golang)**.

Ela disponibiliza um servidor HTTP que recebe requisições e retorna uma resposta em JSON.

Porém, o projeto não consiste apenas na aplicação Go.

Ao redor da aplicação foi construída uma pequena infraestrutura utilizando tecnologias comuns em ambientes DevOps.

A arquitetura principal é:

```text
                         USUÁRIO
                            |
                            | HTTP
                            v
                     +-------------+
                     |    NGINX    |
                     |   Porta 80   |
                     |Reverse Proxy|
                     +-------------+
                            |
                            v
                     +-------------+
                     | APLICAÇÃO   |
                     |     GO      |
                     | Porta 8080  |
                     +-------------+
                            |
                            | /metrics
                            v
                     +-------------+
                     | PROMETHEUS  |
                     | Porta 9090  |
                     +-------------+
                            |
                            v
                     +-------------+
                     |   GRAFANA   |
                     | Porta 3000  |
                     +-------------+
```

Além desses componentes:

```text
Docker
   |
   +-- executa os serviços em containers

Docker Compose
   |
   +-- organiza e conecta os containers

Ansible
   |
   +-- automatiza o processo de deploy

Git
   |
   +-- registra o histórico do projeto

GitHub
   |
   +-- mantém o repositório remoto
```

### Explicação simples

Para alguém que não conhece essas tecnologias, podemos imaginar uma empresa:

| Tecnologia | Comparação simples |
|---|---|
| **Go** | Funcionário que realiza o trabalho |
| **Nginx** | Recepção que recebe o usuário |
| **Prometheus** | Funcionário que coleta indicadores |
| **Grafana** | Painel onde os indicadores são apresentados |
| **Docker** | Caixa isolada onde cada serviço funciona |
| **Docker Compose** | Coordenador das caixas |
| **Ansible** | Automação que executa tarefas de infraestrutura |
| **Git** | Histórico de todas as alterações |
| **GitHub** | Local remoto onde o projeto fica armazenado |

---

# 2. Roteiro de Avaliação Técnica para Recrutadores

Esta seção permite que um avaliador que esteja acessando o projeto pela primeira vez consiga executar e validar a solução.

O roteiro considera preferencialmente um ambiente **Linux** com Docker, Docker Compose, Git e Ansible disponíveis.

---

## 2.1 Clonar o projeto

Clone o repositório:

```bash
git clone https://github.com/passostestes/http-server-projeto-korp.git
```

Entre na pasta:

```bash
cd http-server-projeto-korp
```

Confirme:

```bash
pwd
```

Liste os arquivos:

```bash
ls -la
```

A estrutura principal deverá ser semelhante a:

```text
http-server-projeto-korp/
├── ansible/
├── grafana/
├── nginx/
├── prometheus/
├── .gitignore
├── Dockerfile
├── compose.yaml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## 2.2 Verificar os pré-requisitos

### Git

```bash
git --version
```

### Docker

```bash
docker --version
```

### Docker Compose

```bash
docker-compose --version
```

### Ansible

```bash
ansible --version
```

### Go

Opcionalmente:

```bash
go version
```

> O Go instalado diretamente no host não é obrigatório para executar a aplicação via Docker, pois a compilação é realizada durante a construção da imagem.

---

## 2.3 Criar a rede Docker

O projeto utiliza uma rede externa chamada:

```text
projeto-korp-network
```

Primeiro verifique:

```bash
docker network ls
```

Caso a rede não exista:

```bash
docker network create --driver bridge projeto-korp-network
```

Confirme novamente:

```bash
docker network ls
```

Deve aparecer:

```text
projeto-korp-network
```

> Se o Docker informar que a rede já existe, basta continuar.

---

## 2.4 Validar o Docker Compose

Antes de iniciar os serviços:

```bash
docker-compose config
```

Para uma verificação mais simples:

```bash
docker-compose config >/dev/null && echo "COMPOSE OK"
```

Resultado esperado:

```text
COMPOSE OK
```

Isso confirma que o arquivo `compose.yaml` possui sintaxe válida.

---

## 2.5 Construir e iniciar o ambiente

Execute:

```bash
docker-compose up -d --build
```

Aguarde alguns segundos:

```bash
sleep 10
```

Verifique:

```bash
docker ps
```

Os quatro serviços esperados são:

```text
http-server-projeto-korp
nginx-projeto-korp
prometheus-projeto-korp
grafana-projeto-korp
```

Todos devem apresentar status semelhante a:

```text
Up
```

---

## 2.6 Verificar as portas

Execute:

```bash
docker ps
```

O comportamento esperado é:

| Serviço | Porta | Exposição |
|---|---:|---|
| Nginx | 80 | Host |
| Go | 8080 | Somente rede Docker |
| Prometheus | 9090 | Host |
| Grafana | 3000 | Host |

A aplicação Go utiliza a porta `8080`, mas ela não precisa ser publicada diretamente no host.

O acesso externo ocorre através do Nginx.

Isso significa que não é esperado encontrar:

```text
0.0.0.0:8080->8080/tcp
```

para a aplicação Go.

---

## 2.7 Testar a aplicação

Execute:

```bash
curl -i http://localhost/projeto-korp
```

Resultado esperado:

```text
HTTP/1.1 200 OK
```

Além do status HTTP, será retornado um JSON semelhante a:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-08-16T14:32:51Z"
}
```

O horário é gerado dinamicamente pela aplicação.

Execute novamente:

```bash
curl http://localhost/projeto-korp
```

Uma nova requisição deve retornar um novo horário.

---

## 2.8 Confirmar o funcionamento do proxy reverso

O fluxo testado anteriormente é:

```text
curl
  |
  v
localhost:80
  |
  v
Nginx
  |
  v
Go:8080
```

O retorno `HTTP 200` através da porta 80 confirma que o Nginx consegue encaminhar a requisição para a aplicação Go.

---

## 2.9 Confirmar que a porta 8080 não está publicada

Execute:

```bash
curl --max-time 3 http://localhost:8080/projeto-korp
```

Como a porta não está publicada no host, é esperado que a conexão direta falhe.

Por exemplo:

```text
curl: (7) Failed to connect to localhost port 8080
```

Isso não representa falha da aplicação.

É consequência da arquitetura adotada:

```text
Usuário -> Nginx -> Go
```

em vez de:

```text
Usuário -> Go diretamente
```

---

## 2.10 Gerar tráfego para as métricas

Para gerar várias requisições:

```bash
for i in $(seq 1 30); do
  curl -s http://localhost/projeto-korp > /dev/null
done
```

Isso realiza 30 requisições contra a aplicação.

Essas requisições ajudam a gerar dados para o Prometheus e Grafana.

---

## 2.11 Testar o Prometheus

Verifique se o Prometheus está pronto:

```bash
curl -s http://localhost:9090/-/ready
```

Resultado esperado:

```text
Prometheus Server is Ready.
```

---

## 2.12 Consultar o Prometheus

Uma consulta básica pode ser realizada utilizando:

```bash
curl -s 'http://localhost:9090/api/v1/query?query=up'
```

A métrica `up` é utilizada pelo Prometheus para indicar se um target está disponível.

Em termos simples:

```text
up = 1
```

significa que o target está sendo coletado corretamente.

```text
up = 0
```

indica problema na coleta daquele target.

---

## 2.13 Testar o Grafana

Execute:

```bash
curl -I http://localhost:3000/login
```

Resultado esperado:

```text
HTTP/1.1 200 OK
```

O Grafana está disponível em:

```text
http://localhost:3000
```

No ambiente de demonstração foram configuradas credenciais administrativas para facilitar os testes.

> Em um ambiente real de produção, credenciais não devem permanecer fixas diretamente em arquivos versionados. O recomendado é utilizar secrets ou variáveis protegidas.

---

## 2.14 Verificar o datasource do Grafana

O datasource Prometheus é provisionado automaticamente.

Durante o desenvolvimento foi utilizado:

```bash
curl -s -u admin:admin http://localhost:3000/api/datasources
```

O resultado deve conter informações semelhantes a:

```text
"name":"Prometheus"
```

e:

```text
"url":"http://prometheus:9090"
```

Isso confirma que o Grafana consegue localizar o Prometheus pela rede Docker.

---

## 2.15 Verificar o dashboard provisionado

Execute:

```bash
curl -s -u admin:admin "http://localhost:3000/api/search?query=HTTP"
```

Deve aparecer um dashboard com título semelhante a:

```text
HTTP Server Projeto Korp
```

Isso confirma que o dashboard foi carregado automaticamente.

---

## 2.16 Validar o Ansible

Antes da execução:

```bash
ansible-playbook \
  -i ansible/inventory.ini \
  ansible/playbook.yml \
  --syntax-check
```

Resultado esperado:

```text
playbook: ansible/playbook.yml
```

---

## 2.17 Executar o Ansible

Execute:

```bash
ansible-playbook \
  -i ansible/inventory.ini \
  ansible/playbook.yml
```

Ao final deverá aparecer um resumo semelhante a:

```text
PLAY RECAP

localhost : ok=10 changed=1 unreachable=0 failed=0
```

O ponto principal é:

```text
failed=0
```

Isso indica que nenhuma tarefa do playbook falhou.

---

## 2.18 Verificar novamente os containers

Depois da execução do Ansible:

```bash
docker ps
```

Os quatro serviços devem continuar funcionando:

```text
http-server-projeto-korp
nginx-projeto-korp
prometheus-projeto-korp
grafana-projeto-korp
```

---

## 2.19 Teste consolidado

Para uma verificação rápida:

```bash
echo "=== APLICACAO VIA NGINX ==="
curl -i http://localhost/projeto-korp

echo
echo "=== PROMETHEUS ==="
curl -s http://localhost:9090/-/ready

echo
echo "=== GRAFANA ==="
curl -I http://localhost:3000/login

echo
echo "=== CONTAINERS ==="
docker ps
```

Resultado esperado:

```text
Aplicação  -> HTTP 200
Prometheus -> Ready
Grafana    -> HTTP 200
Containers -> todos Up
```

---

## 2.20 Encerrar o ambiente

Depois dos testes:

```bash
docker-compose down
```

Verifique:

```bash
docker ps
```

Como `projeto-korp-network` é uma rede externa, ela não é necessariamente removida pelo Compose.

Caso seja necessário removê-la:

```bash
docker network rm projeto-korp-network
```

---

## 2.21 Resultado esperado da avaliação

Ao final deste roteiro, o avaliador deverá conseguir verificar:

- aplicação Go funcionando;
- endpoint HTTP respondendo;
- resposta JSON;
- horário gerado dinamicamente;
- Nginx funcionando como proxy reverso;
- porta 8080 não publicada diretamente;
- containers funcionando;
- rede Docker funcionando;
- Prometheus disponível;
- Prometheus coletando o target;
- Grafana disponível;
- datasource Prometheus provisionado;
- dashboard Grafana provisionado;
- Docker Compose válido;
- Ansible com sintaxe válida;
- playbook Ansible executado sem falhas.

---

# 3. Tecnologias Utilizadas

## 3.1 Go

**Go**, ou Golang, é a linguagem utilizada para desenvolver o servidor HTTP.

De maneira simples:

> Go é o componente que executa a lógica principal da aplicação.

O código encontra-se principalmente em:

```text
main.go
```

Responsabilidades:

- iniciar o servidor;
- receber requisições;
- processar requisições;
- gerar respostas;
- expor a aplicação;
- fornecer métricas.

---

## 3.2 Docker

Docker permite executar aplicações em ambientes isolados chamados **containers**.

Neste projeto existem containers independentes para:

- Go;
- Nginx;
- Prometheus;
- Grafana.

A vantagem é permitir que cada serviço possua seu próprio ambiente.

---

## 3.3 Dockerfile

O:

```text
Dockerfile
```

define como a imagem da aplicação Go é construída.

Foi utilizado **multi-stage build**.

```text
Primeira etapa
      |
      v
Compilação da aplicação
      |
      v
Binário Go
      |
      v
Segunda etapa
      |
      v
Imagem final de execução
```

Essa técnica permite separar o ambiente utilizado para compilar daquele utilizado para executar.

---

## 3.4 Docker Compose

O:

```text
compose.yaml
```

define os serviços que compõem a solução.

Ele centraliza configurações relacionadas a:

- imagens;
- containers;
- portas;
- volumes;
- redes;
- dependências;
- reinicialização.

---

## 3.5 Nginx

Nginx funciona como proxy reverso.

É a porta de entrada da aplicação.

---

## 3.6 Prometheus

Prometheus realiza coleta e armazenamento de métricas.

---

## 3.7 Grafana

Grafana consulta as métricas e apresenta dashboards.

---

## 3.8 Ansible

Ansible automatiza tarefas relacionadas ao deploy e verificação da infraestrutura.

---

## 3.9 Git

Git registra o histórico de desenvolvimento.

---

## 3.10 GitHub

GitHub mantém uma cópia remota e versionada do projeto.

---

# 4. Aplicação Go

A aplicação principal está localizada em:

```text
main.go
```

Ela disponibiliza o endpoint:

```text
/projeto-korp
```

A aplicação retorna informações em JSON.

Exemplo:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-08-16T14:32:51Z"
}
```

A porta utilizada internamente é:

```text
8080
```

Por decisão de arquitetura, essa porta não é publicada diretamente para o host.

O acesso externo ocorre através do Nginx.

```text
Usuário
   |
   v
Nginx :80
   |
   v
Go :8080
```

Isso mantém a aplicação atrás do proxy reverso.

---

# 5. Docker e Docker Compose

## 5.1 Dockerfile

O Dockerfile é responsável por transformar a aplicação em uma imagem Docker.

A estratégia multi-stage permite:

```text
Código Go
   |
   v
Compilação
   |
   v
Binário
   |
   v
Imagem final
```

---

## 5.2 Docker Compose

O arquivo:

```text
compose.yaml
```

coordena os serviços.

Os principais containers são:

```text
http-server-projeto-korp
nginx-projeto-korp
prometheus-projeto-korp
grafana-projeto-korp
```

### Validar

```bash
docker-compose config
```

### Construir e iniciar

```bash
docker-compose up -d --build
```

### Verificar

```bash
docker ps
```

### Parar

```bash
docker-compose down
```

---

# 6. Nginx

O Nginx funciona como **reverse proxy**.

Para uma explicação simples:

```text
Cliente chega
    |
    v
Recepção (Nginx)
    |
    v
Aplicação correta (Go)
```

A configuração encontra-se dentro de:

```text
nginx/
```

O Nginx recebe requisições pela porta:

```text
80
```

e encaminha para o servidor Go pela rede Docker.

Teste:

```bash
curl -i http://localhost/projeto-korp
```

Resultado esperado:

```text
HTTP/1.1 200 OK
```

---

# 7. Prometheus

Prometheus é utilizado para monitoramento.

A aplicação disponibiliza métricas e o Prometheus realiza coletas periódicas.

```text
Aplicação
   |
   | métricas
   v
Prometheus
```

Sua configuração encontra-se em:

```text
prometheus/prometheus.yml
```

A porta publicada é:

```text
9090
```

Teste:

```bash
curl -s http://localhost:9090/-/ready
```

Resultado:

```text
Prometheus Server is Ready.
```

Uma consulta pela API pode ser realizada com:

```bash
curl -s 'http://localhost:9090/api/v1/query?query=up'
```

---

# 8. Grafana e Provisionamento

Grafana é responsável pela visualização das métricas coletadas pelo Prometheus.

```text
Go
 |
 v
Prometheus
 |
 v
Grafana
 |
 v
Dashboard
```

A porta utilizada é:

```text
3000
```

## 8.1 Estrutura

```text
grafana/
├── dashboards/
│   └── http-server-projeto-korp-dashboard.json
│
└── provisioning/
    ├── dashboards/
    │   └── dashboards.yml
    │
    └── datasources/
        └── datasource.yml
```

## 8.2 Datasource

O datasource Prometheus é configurado automaticamente.

A comunicação interna utiliza:

```text
http://prometheus:9090
```

## 8.3 Dashboard

O dashboard também é carregado automaticamente através do provisioning.

## 8.4 Validação do JSON

Durante o desenvolvimento foi utilizado:

```bash
python3 -m json.tool grafana/dashboards/http-server-projeto-korp-dashboard.json > /dev/null && echo "DASHBOARD JSON OK"
```

Resultado esperado:

```text
DASHBOARD JSON OK
```

## 8.5 Teste do Grafana

```bash
curl -I http://localhost:3000/login
```

Resultado:

```text
HTTP/1.1 200 OK
```

## 8.6 Teste do datasource

```bash
curl -s -u admin:admin http://localhost:3000/api/datasources
```

## 8.7 Teste do dashboard

```bash
curl -s -u admin:admin "http://localhost:3000/api/search?query=HTTP"
```

---

# 9. Ansible

Ansible automatiza tarefas de infraestrutura.

Os arquivos estão em:

```text
ansible/
├── inventory.ini
└── playbook.yml
```

## 9.1 Inventory

O inventário utilizado é baseado em:

```ini
[projeto_korp]
localhost ansible_connection=local
```

Isso informa ao Ansible que as tarefas devem ser executadas na própria máquina.

## 9.2 Playbook

O playbook realiza tarefas como:

```text
Verificar Docker
       |
       v
Verificar Docker Compose
       |
       v
Validar Compose
       |
       v
Executar build/deploy
       |
       v
Verificar containers
```

## 9.3 Verificar versão

```bash
ansible --version
```

## 9.4 Validar sintaxe

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --syntax-check
```

## 9.5 Executar

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

No teste realizado, o playbook terminou sem falhas:

```text
unreachable=0
failed=0
```

---

# 10. Arquitetura de Pastas

A estrutura principal é:

```text
http-server-projeto-korp/
│
├── .git/
├── .gitignore
├── README.md
├── Dockerfile
├── compose.yaml
├── go.mod
├── go.sum
├── main.go
│
├── nginx/
│   └── conf.d/
│
├── prometheus/
│   └── prometheus.yml
│
├── grafana/
│   ├── dashboards/
│   │   └── http-server-projeto-korp-dashboard.json
│   │
│   └── provisioning/
│       ├── dashboards/
│       │   └── dashboards.yml
│       │
│       └── datasources/
│           └── datasource.yml
│
└── ansible/
    ├── inventory.ini
    └── playbook.yml
```

## `main.go`

Código principal da aplicação.

## `go.mod`

Define o módulo Go e suas dependências.

## `go.sum`

Mantém informações de integridade das dependências.

## `Dockerfile`

Define a construção da imagem da aplicação.

## `compose.yaml`

Define a infraestrutura Docker.

## `.gitignore`

Define arquivos que não devem ser versionados.

## `nginx/`

Configuração do proxy reverso.

## `prometheus/`

Configuração da coleta de métricas.

## `grafana/dashboards/`

Dashboard do projeto.

## `grafana/provisioning/datasources/`

Configuração automática do Prometheus no Grafana.

## `grafana/provisioning/dashboards/`

Configuração automática dos dashboards.

## `ansible/inventory.ini`

Define os hosts administrados pelo Ansible.

## `ansible/playbook.yml`

Define as tarefas automatizadas.

---

# 11. Rede Docker

Os containers precisam conseguir conversar.

Foi utilizada:

```text
projeto-korp-network
```

A rede é do tipo:

```text
bridge
```

Criação:

```bash
docker network create --driver bridge projeto-korp-network
```

Verificação:

```bash
docker network ls
```

Inspeção:

```bash
docker network inspect projeto-korp-network
```

A arquitetura lógica é:

```text
+-----------------------------------------+
|       projeto-korp-network              |
|                                         |
|   +------+       +----------+           |
|   |Nginx | ----> |    Go    |           |
|   +------+       +----------+           |
|                       |                 |
|                       v                 |
|                 +------------+          |
|                 | Prometheus |          |
|                 +------------+          |
|                       |                 |
|                       v                 |
|                   +---------+           |
|                   | Grafana |           |
|                   +---------+           |
|                                         |
+-----------------------------------------+
```

No Compose a rede é externa.

Isso significa que ela precisa existir antes do Compose utilizá-la.

Caso apareça:

```text
Network projeto-korp-network declared as external, but could not be found
```

execute:

```bash
docker network create --driver bridge projeto-korp-network
```

---

# 12. Fluxo Completo da Aplicação

## 12.1 Requisição do usuário

```text
1. Usuário
      |
      | GET /projeto-korp
      v

2. Nginx
      |
      | proxy_pass
      v

3. Aplicação Go
      |
      | processa
      v

4. Resposta JSON
      |
      v

5. Nginx
      |
      v

6. Usuário
```

## 12.2 Monitoramento

Paralelamente:

```text
Aplicação Go
      |
      | métricas
      v
Prometheus
      |
      | datasource
      v
Grafana
      |
      v
Dashboard
```

## 12.3 Automação

```text
Ansible
   |
   v
Docker Compose
   |
   +-- Go
   +-- Nginx
   +-- Prometheus
   +-- Grafana
```

---

# 13. Execução Manual

## 13.1 Entrar no projeto

```bash
cd /root/http-server-projeto-korp
```

ou, após um clone normal:

```bash
cd http-server-projeto-korp
```

## 13.2 Verificar diretório

```bash
pwd
```

## 13.3 Ver arquivos

```bash
ls -la
```

## 13.4 Criar rede

```bash
docker network create --driver bridge projeto-korp-network
```

Caso já exista, continue normalmente.

## 13.5 Validar Compose

```bash
docker-compose config >/dev/null && echo "COMPOSE OK"
```

## 13.6 Iniciar

```bash
docker-compose up -d --build
```

## 13.7 Verificar

```bash
docker ps
```

## 13.8 Parar

```bash
docker-compose down
```

---

# 14. Testes e Diagnóstico

## 14.1 Aplicação

```bash
curl -i http://localhost/projeto-korp
```

Esperado:

```text
HTTP/1.1 200 OK
```

---

## 14.2 Prometheus

```bash
curl -s http://localhost:9090/-/ready
```

Esperado:

```text
Prometheus Server is Ready.
```

---

## 14.3 Grafana

```bash
curl -I http://localhost:3000/login
```

Esperado:

```text
HTTP/1.1 200 OK
```

---

## 14.4 Containers

```bash
docker ps
```

---

## 14.5 Todos os containers, inclusive parados

```bash
docker ps -a
```

---

## 14.6 Logs da aplicação

```bash
docker logs http-server-projeto-korp
```

---

## 14.7 Logs do Nginx

```bash
docker logs nginx-projeto-korp
```

---

## 14.8 Logs do Prometheus

```bash
docker logs prometheus-projeto-korp
```

---

## 14.9 Logs do Grafana

```bash
docker logs grafana-projeto-korp
```

---

## 14.10 Logs em tempo real

Exemplo:

```bash
docker logs -f http-server-projeto-korp
```

Para sair:

```text
Ctrl + C
```

---

## 14.11 Rede

```bash
docker network inspect projeto-korp-network
```

---

## 14.12 Validar dashboard JSON

```bash
python3 -m json.tool grafana/dashboards/http-server-projeto-korp-dashboard.json > /dev/null && echo "DASHBOARD JSON OK"
```

---

## 14.13 Validar Compose

```bash
docker-compose config >/dev/null && echo "COMPOSE OK"
```

---

## 14.14 Validar Ansible

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --syntax-check
```

---

## 14.15 Executar Ansible

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

Esperado:

```text
failed=0
```

---

# 15. Git e GitHub

Git foi utilizado para versionar o desenvolvimento.

GitHub foi utilizado como repositório remoto.

## 15.1 Verificar status

```bash
git status
```

---

## 15.2 Visualizar últimos commits

```bash
git log --oneline -5
```

---

## 15.3 Adicionar arquivos

Exemplo:

```bash
git add compose.yaml
```

Grafana:

```bash
git add compose.yaml grafana
```

Ansible:

```bash
git add ansible
```

README:

```bash
git add README.md
```

---

## 15.4 Criar commits

Exemplos utilizados durante a evolução do projeto:

```bash
git commit -m "feat: adiciona Grafana e dashboard provisionado"
```

```bash
git commit -m "feat: adiciona automacao de deploy com Ansible"
```

Para a documentação:

```bash
git commit -m "docs: adiciona README completo com roteiro de avaliacao tecnica"
```

---

## 15.5 Enviar para o GitHub

```bash
git push origin main
```

---

## 15.6 Conferência final

```bash
git status
```

Resultado ideal:

```text
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean
```

Isso significa que não existem alterações locais pendentes e que a branch está sincronizada com o repositório remoto.

---

# 16. Conclusão

O projeto começou com uma aplicação HTTP desenvolvida em Go e foi evoluído para uma arquitetura contendo diferentes componentes de infraestrutura.

A solução final integra:

```text
Go
 |
 v
Docker
 |
 v
Docker Compose
 |
 +---------> Nginx
 |
 +---------> Prometheus
 |
 +---------> Grafana
 |
 v
Ansible

Versionamento:
Git -> GitHub
```

### Arquitetura final

```text
                         USUÁRIO
                            |
                            v
                  +-------------------+
                  |       NGINX       |
                  |      Porta 80     |
                  |   Proxy Reverso   |
                  +-------------------+
                            |
                            v
                  +-------------------+
                  |    APLICAÇÃO GO   |
                  |     Porta 8080    |
                  |    Servidor HTTP  |
                  +-------------------+
                            |
                            | métricas
                            v
                  +-------------------+
                  |    PROMETHEUS     |
                  |     Porta 9090    |
                  | Coleta de métricas|
                  +-------------------+
                            |
                            v
                  +-------------------+
                  |      GRAFANA      |
                  |     Porta 3000    |
                  |     Dashboard     |
                  +-------------------+


             +--------------------------------+
             |             DOCKER             |
             |                                |
             | Executa serviços em containers |
             +--------------------------------+
                            |
                            v
             +--------------------------------+
             |        DOCKER COMPOSE          |
             |                                |
             | Organiza containers, rede,     |
             | portas, volumes e dependências |
             +--------------------------------+
                            ^
                            |
             +--------------------------------+
             |            ANSIBLE             |
             |                                |
             | Automatiza tarefas de deploy   |
             +--------------------------------+


             +--------------------------------+
             |         GIT + GITHUB           |
             |                                |
             | Versionamento e repositório    |
             +--------------------------------+
```

### Responsabilidade de cada tecnologia

| Tecnologia | Responsabilidade | Explicação simples |
|---|---|---|
| **Go** | Backend | Executa a aplicação HTTP |
| **Docker** | Containerização | Isola os serviços |
| **Dockerfile** | Build | Define como a imagem Go é construída |
| **Docker Compose** | Orquestração local | Gerencia os containers |
| **Nginx** | Proxy reverso | Recebe o acesso externo |
| **Prometheus** | Monitoramento | Coleta métricas |
| **Grafana** | Observabilidade | Exibe métricas em dashboards |
| **Ansible** | Automação | Automatiza tarefas de deploy |
| **Git** | Versionamento | Registra alterações |
| **GitHub** | Repositório remoto | Armazena e compartilha o projeto |

### Resultado final

Ao final do desenvolvimento foram implementados:

- servidor HTTP em Go;
- resposta JSON;
- Dockerfile;
- multi-stage build;
- containerização da aplicação;
- Nginx como proxy reverso;
- Docker Compose;
- rede Docker;
- Prometheus;
- métricas da aplicação;
- Grafana;
- datasource provisionado;
- dashboard provisionado;
- Ansible;
- automação do deploy;
- validação do playbook;
- testes dos serviços;
- Git;
- GitHub;
- documentação;
- roteiro de avaliação técnica.

Dessa forma, o projeto demonstra de maneira prática conceitos relacionados a:

**Backend + Containers + Redes + Proxy Reverso + Monitoramento + Observabilidade + Automação + Infraestrutura como Código + Versionamento.**

---

## Autor

Projeto desenvolvido como parte de um desafio técnico para demonstração prática de conhecimentos em desenvolvimento, containers, infraestrutura, observabilidade e automação.
