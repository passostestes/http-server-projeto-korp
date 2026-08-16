# HTTP Server Projeto Korp

Projeto desenvolvido como desafio técnico utilizando **Go, Docker, Docker Compose, Nginx, Prometheus, Grafana, Ansible, Git e GitHub**.

O objetivo deste projeto é criar um servidor HTTP em Go e evoluí-lo para uma arquitetura containerizada, monitorada e automatizada.

Além de executar a aplicação, o projeto demonstra conceitos importantes utilizados em ambientes reais de desenvolvimento e DevOps, como:

- desenvolvimento de uma aplicação HTTP;
- compilação de uma aplicação Go;
- criação de imagens Docker;
- uso de containers;
- Docker Compose;
- proxy reverso com Nginx;
- redes Docker;
- exposição de métricas;
- monitoramento com Prometheus;
- dashboards com Grafana;
- provisionamento automático;
- automação de deploy com Ansible;
- versionamento com Git;
- publicação do projeto no GitHub.

---

## 1. Visão Geral do Projeto

A aplicação principal foi desenvolvida utilizando a linguagem **Go**.

Ela executa um servidor HTTP que responde às requisições realizadas pelos usuários.

Porém, o projeto não executa apenas o programa Go.

Foram adicionados outros componentes para criar uma infraestrutura mais completa.

A arquitetura utilizada pode ser representada da seguinte forma:

```text
                         USUÁRIO
                            |
                            v
                     +-------------+
                     |    NGINX    |
                     |   Porta 80   |
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
   +-- executa os containers

Docker Compose
   |
   +-- organiza os serviços e suas conexões

Ansible
   |
   +-- automatiza o deploy

Git
   |
   +-- controla as versões

GitHub
   |
   +-- armazena o projeto remotamente
```

Para facilitar o entendimento de quem está começando:

| Tecnologia | Comparação simples |
|---|---|
| Go | Funcionário que executa o trabalho |
| Nginx | Recepção da empresa |
| Prometheus | Funcionário que coleta indicadores |
| Grafana | Painel que apresenta os indicadores |
| Docker | Caixa isolada onde cada serviço funciona |
| Docker Compose | Coordenador das caixas |
| Ansible | Automação que prepara e executa o ambiente |
| Git | Histórico de alterações |
| GitHub | Local onde o projeto fica armazenado na internet |

---

## 2. Tecnologias Utilizadas

### 2.1 Go

**Go**, também conhecido como **Golang**, é a linguagem utilizada para desenvolver a aplicação principal.

O código principal encontra-se em:

```text
main.go
```

De forma simples:

> Go contém a lógica principal da aplicação.

A aplicação Go é responsável por:

- iniciar o servidor HTTP;
- receber requisições;
- processar as requisições;
- devolver respostas;
- disponibilizar a rota da aplicação;
- disponibilizar métricas para monitoramento.

A porta interna utilizada pela aplicação é:

```text
8080
```

---

### 2.2 Docker

**Docker** é uma tecnologia utilizada para executar aplicações dentro de ambientes isolados chamados **containers**.

Um container pode ser entendido como uma pequena caixa contendo determinado serviço e tudo aquilo que ele precisa para funcionar.

Neste projeto existem containers para:

```text
Aplicação Go
Nginx
Prometheus
Grafana
```

Isso facilita a reprodução do ambiente.

Em vez de configurar cada aplicação manualmente em diferentes computadores, o Docker permite criar ambientes padronizados.

Para verificar a versão instalada:

```bash
docker --version
```

Para visualizar containers em execução:

```bash
docker ps
```

Para visualizar também containers parados:

```bash
docker ps -a
```

---

### 2.3 Dockerfile

O arquivo:

```text
Dockerfile
```

contém as instruções necessárias para o Docker construir a imagem da aplicação Go.

O projeto utiliza uma estratégia conhecida como **multi-stage build**.

De maneira simplificada:

```text
ETAPA 1

Imagem com ferramentas de desenvolvimento Go
                |
                v
Código-fonte é compilado
                |
                v
Executável é criado


ETAPA 2

Imagem menor
     |
     v
Recebe somente o executável
     |
     v
Executa a aplicação
```

A principal vantagem é evitar que ferramentas utilizadas apenas durante a compilação fiquem desnecessariamente dentro da imagem final.

Isso ajuda a criar uma imagem menor e mais apropriada para execução.

---

### 2.4 Docker Compose

O arquivo:

```text
compose.yaml
```

centraliza a configuração dos serviços Docker.

O Docker permite executar containers.

O **Docker Compose** facilita a execução de vários containers que fazem parte da mesma aplicação.

Neste projeto o Compose administra:

```text
http-server-projeto-korp
nginx-projeto-korp
prometheus-projeto-korp
grafana-projeto-korp
```

O comando utilizado para iniciar a infraestrutura é:

```bash
docker-compose up -d --build
```

Onde:

```text
up
```

inicia os serviços.

```text
-d
```

executa os containers em segundo plano.

```text
--build
```

solicita a reconstrução da imagem quando necessário.

Para verificar se o arquivo Compose possui configuração válida:

```bash
docker-compose config
```

Também foi utilizado:

```bash
docker-compose config >/dev/null && echo "COMPOSE OK"
```

Quando a configuração está correta, o resultado esperado é:

```text
COMPOSE OK
```

---

## 3. Nginx

O **Nginx** funciona como **proxy reverso**.

Para entender de forma simples, podemos imaginar uma empresa.

Quando uma pessoa chega à empresa, normalmente ela fala primeiro com a recepção.

A recepção identifica para onde aquela pessoa precisa ir.

Neste projeto:

```text
Usuário
   |
   v
Nginx
   |
   v
Aplicação Go
```

O Nginx funciona como essa recepção.

O usuário acessa:

```text
http://localhost/projeto-korp
```

O Nginx recebe a requisição e encaminha internamente para a aplicação Go.

A porta externa utilizada pelo Nginx é:

```text
80
```

O container utilizado é baseado em:

```text
nginx:alpine
```

A configuração encontra-se no diretório:

```text
nginx/
```

A utilização de proxy reverso evita a necessidade de expor diretamente a aplicação Go para o usuário.

### Teste do Nginx

O comando utilizado foi:

```bash
curl -i http://localhost/projeto-korp
```

O resultado esperado inclui:

```text
HTTP/1.1 200 OK
```

e a resposta da aplicação.

---

## 4. Prometheus

O **Prometheus** é responsável pela coleta e armazenamento de métricas.

Uma métrica é um número utilizado para acompanhar o comportamento de uma aplicação.

Por exemplo:

- quantidade de requisições;
- duração das requisições;
- disponibilidade do serviço;
- comportamento da aplicação ao longo do tempo.

A aplicação disponibiliza métricas e o Prometheus consulta essas informações.

O fluxo é:

```text
Aplicação Go
      |
      | /metrics
      v
Prometheus
```

O container utilizado é baseado em:

```text
prom/prometheus:latest
```

O Prometheus utiliza a porta:

```text
9090
```

A configuração encontra-se em:

```text
prometheus/
```

O arquivo principal de configuração é:

```text
prometheus/prometheus.yml
```

### Teste do Prometheus

Foi utilizado:

```bash
curl -s http://localhost:9090/-/ready
```

Resultado esperado:

```text
Prometheus Server is Ready.
```

Isso indica que o servidor Prometheus está funcionando.

---

## 5. Grafana

O **Grafana** é responsável pela visualização das métricas.

Existe uma diferença importante entre Prometheus e Grafana:

```text
PROMETHEUS
     |
     +-- coleta e armazena métricas

GRAFANA
     |
     +-- consulta essas métricas
     |
     +-- apresenta as informações visualmente
```

O fluxo completo é:

```text
Aplicação Go
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

O container utilizado é baseado em:

```text
grafana/grafana:latest
```

O Grafana utiliza a porta:

```text
3000
```

O serviço pode ser verificado através de:

```bash
curl -I http://localhost:3000/login
```

Resultado esperado:

```text
HTTP/1.1 200 OK
```

---

## 6. Provisionamento do Grafana

O projeto utiliza o recurso de **provisioning** do Grafana.

Provisionamento significa deixar configurações preparadas antecipadamente.

Sem provisioning seria necessário:

```text
Iniciar Grafana
      |
      v
Abrir interface
      |
      v
Cadastrar Prometheus
      |
      v
Criar/importar dashboard
```

Com provisioning:

```text
Grafana inicia
      |
      v
Lê os arquivos de configuração
      |
      +-- configura Prometheus
      |
      +-- carrega dashboard
      |
      v
Ambiente preparado
```

A estrutura utilizada é:

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

### Datasource

O arquivo:

```text
grafana/provisioning/datasources/datasource.yml
```

configura o Prometheus automaticamente como fonte de dados.

O Grafana acessa o Prometheus internamente através da rede Docker.

### Dashboard

O arquivo:

```text
grafana/dashboards/http-server-projeto-korp-dashboard.json
```

contém a definição do dashboard.

O arquivo:

```text
grafana/provisioning/dashboards/dashboards.yml
```

informa ao Grafana onde os dashboards estão armazenados.

### Validação do JSON

Durante a configuração foi utilizado Python para verificar se o arquivo JSON era válido:

```bash
python3 -m json.tool grafana/dashboards/http-server-projeto-korp-dashboard.json > /dev/null
```

Depois:

```bash
echo "JSON OK"
```

Resultado:

```text
JSON OK
```

### Teste do datasource pela API do Grafana

Foi utilizado:

```bash
curl -s -u admin:admin http://localhost:3000/api/datasources
```

O resultado confirmou a existência do datasource:

```text
Prometheus
```

com URL interna:

```text
http://prometheus:9090
```

### Teste do dashboard pela API

Também foi realizada uma consulta ao Grafana para verificar o dashboard provisionado:

```bash
curl -s -u admin:admin "http://localhost:3000/api/search?query=HTTP"
```

O dashboard retornado possui o título:

```text
HTTP Server Projeto Korp
```

confirmando o provisionamento.

---

## 7. Ansible

O **Ansible** foi utilizado para automatizar o processo de deploy.

Sem Ansible seria necessário executar manualmente diversas etapas.

Por exemplo:

```text
Verificar Docker
        |
        v
Verificar Docker Compose
        |
        v
Validar configuração
        |
        v
Executar Docker Compose
        |
        v
Verificar containers
```

O Ansible permite colocar essas etapas dentro de um arquivo chamado **playbook**.

A estrutura utilizada é:

```text
ansible/
├── inventory.ini
└── playbook.yml
```

### Verificação da instalação

Foi utilizado:

```bash
ansible --version
```

No ambiente utilizado durante o desenvolvimento foi identificado o Ansible Core instalado e funcionando.

---

## 8. Executando o Ansible

### Inventory

O arquivo:

```text
ansible/inventory.ini
```

define onde o Ansible executará as tarefas.

Neste projeto foi utilizado:

```ini
[projeto_korp]
localhost ansible_connection=local
```

Isso significa:

> Execute as tarefas na própria máquina.

### Playbook

O arquivo:

```text
ansible/playbook.yml
```

contém as tarefas automatizadas.

Entre as verificações realizadas estão:

- coleta de informações do ambiente;
- verificação do Docker;
- exibição da versão do Docker;
- verificação do Docker Compose;
- exibição da versão do Docker Compose;
- validação do Compose;
- execução do deploy;
- verificação dos containers.

### Verificação da sintaxe

Antes de executar o playbook foi utilizado:

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --syntax-check
```

Resultado esperado:

```text
playbook: ansible/playbook.yml
```

### Execução

O comando utilizado foi:

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

Ao final da execução bem-sucedida, o Ansible apresentou um resumo semelhante a:

```text
PLAY RECAP

localhost : ok=10 changed=1 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
```

O ponto mais importante é:

```text
failed=0
```

Isso indica que o playbook terminou sem tarefas com falha.

---

## 9. Arquitetura de Pastas

A estrutura principal do projeto é:

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
│       └── http-server-projeto-korp.conf
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

### `main.go`

É o código-fonte principal da aplicação.

Ele contém o servidor HTTP desenvolvido em Go.

---

### `go.mod`

Define o módulo Go e registra as dependências diretas do projeto.

---

### `go.sum`

Armazena informações de verificação das dependências utilizadas pelo Go.

---

### `Dockerfile`

Define como construir a imagem Docker da aplicação.

---

### `compose.yaml`

É o arquivo central da infraestrutura Docker.

Define:

- aplicação Go;
- Nginx;
- Prometheus;
- Grafana;
- portas;
- volumes;
- dependências;
- rede Docker;
- políticas de reinicialização.

---

### `.gitignore`

Informa ao Git quais arquivos não devem ser adicionados ao repositório.

É útil para evitar arquivos temporários ou desnecessários.

---

### `nginx/`

Contém a configuração do proxy reverso.

---

### `prometheus/`

Contém a configuração do sistema de monitoramento.

---

### `grafana/`

Contém toda a configuração relacionada à visualização das métricas.

---

### `grafana/dashboards/`

Contém o dashboard em formato JSON.

---

### `grafana/provisioning/datasources/`

Configura automaticamente o Prometheus como datasource.

---

### `grafana/provisioning/dashboards/`

Configura o carregamento automático dos dashboards.

---

### `ansible/`

Contém a automação do deploy.

---

### `ansible/inventory.ini`

Define as máquinas gerenciadas pelo Ansible.

---

### `ansible/playbook.yml`

Contém as tarefas executadas automaticamente durante o deploy.

---

## 10. Rede Docker

Os containers precisam conversar entre si.

Para permitir essa comunicação foi utilizada uma rede Docker chamada:

```text
projeto-korp-network
```

A rede pode ser criada através de:

```bash
docker network create --driver bridge projeto-korp-network
```

Para visualizar as redes:

```bash
docker network ls
```

A arquitetura pode ser representada assim:

```text
+------------------------------------------------+
|            projeto-korp-network                |
|                                                |
|   +-------------+                              |
|   | Aplicação Go|                              |
|   +-------------+                              |
|          |                                     |
|   +-------------+                              |
|   |    Nginx    |                              |
|   +-------------+                              |
|                                                |
|   +-------------+                              |
|   | Prometheus  |                              |
|   +-------------+                              |
|          |                                     |
|   +-------------+                              |
|   |   Grafana   |                              |
|   +-------------+                              |
|                                                |
+------------------------------------------------+
```

No `compose.yaml`, a rede foi configurada como externa:

```yaml
networks:
  projeto-korp-network:
    external: true
```

Por isso ela precisa existir antes da inicialização do Compose.

Caso apareça uma mensagem semelhante a:

```text
Network projeto-korp-network declared as external, but could not be found
```

a solução é criar a rede:

```bash
docker network create projeto-korp-network
```

e executar novamente o deploy.

---

## 11. Fluxo Completo de uma Requisição

Quando um usuário acessa:

```text
http://localhost/projeto-korp
```

acontece aproximadamente o seguinte:

```text
1. USUÁRIO
     |
     | HTTP
     v

2. NGINX :80
     |
     | proxy reverso
     v

3. APLICAÇÃO GO :8080
     |
     | processa requisição
     v

4. RESPOSTA
     |
     v

5. NGINX
     |
     v

6. USUÁRIO
```

Paralelamente, existe outro fluxo para monitoramento:

```text
Aplicação Go
     |
     | /metrics
     v
Prometheus
     |
     | consulta métricas
     v
Grafana
     |
     v
Dashboard
```

Portanto, existem dois fluxos principais:

### Fluxo da aplicação

```text
Usuário -> Nginx -> Go
```

### Fluxo de observabilidade

```text
Go -> Prometheus -> Grafana
```

---

## 12. Como Executar o Projeto

### 12.1 Entrar na pasta

```bash
cd /root/http-server-projeto-korp
```

Verifique:

```bash
pwd
```

Resultado esperado:

```text
/root/http-server-projeto-korp
```

---

### 12.2 Verificar arquivos

```bash
ls -la
```

---

### 12.3 Verificar Docker

```bash
docker --version
```

---

### 12.4 Verificar Docker Compose

```bash
docker-compose --version
```

---

### 12.5 Criar a rede

Caso ainda não exista:

```bash
docker network create --driver bridge projeto-korp-network
```

Verifique:

```bash
docker network ls
```

---

### 12.6 Validar o Compose

```bash
docker-compose config >/dev/null && echo "COMPOSE OK"
```

Resultado esperado:

```text
COMPOSE OK
```

---

### 12.7 Construir e iniciar

```bash
docker-compose up -d --build
```

---

### 12.8 Aguardar inicialização

```bash
sleep 10
```

---

### 12.9 Verificar containers

```bash
docker ps
```

Devem aparecer os serviços:

```text
http-server-projeto-korp
nginx-projeto-korp
prometheus-projeto-korp
grafana-projeto-korp
```

---

## 13. Testes Finais

Após iniciar a infraestrutura, os componentes podem ser testados individualmente.

### 13.1 Teste da aplicação através do Nginx

```bash
curl -i http://localhost/projeto-korp
```

Resultado esperado:

```text
HTTP/1.1 200 OK
```

e uma resposta da aplicação semelhante a:

```json
{"nome":"Projeto Korp","horario":"..."}
```

---

### 13.2 Teste do Prometheus

```bash
curl -s http://localhost:9090/-/ready
```

Resultado esperado:

```text
Prometheus Server is Ready.
```

---

### 13.3 Teste do Grafana

```bash
curl -I http://localhost:3000/login
```

Resultado esperado:

```text
HTTP/1.1 200 OK
```

---

### 13.4 Teste dos containers

```bash
docker ps
```

Todos os serviços devem aparecer com status semelhante a:

```text
Up
```

---

### 13.5 Teste do datasource Grafana

```bash
curl -s -u admin:admin http://localhost:3000/api/datasources
```

Deve existir um datasource chamado:

```text
Prometheus
```

---

### 13.6 Teste do dashboard Grafana

```bash
curl -s -u admin:admin "http://localhost:3000/api/search?query=HTTP"
```

Deve aparecer:

```text
HTTP Server Projeto Korp
```

---

### 13.7 Teste da configuração Docker Compose

```bash
docker-compose config >/dev/null && echo "COMPOSE OK"
```

Resultado:

```text
COMPOSE OK
```

---

### 13.8 Teste da sintaxe do Ansible

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --syntax-check
```

---

### 13.9 Teste completo com Ansible

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

O resultado final deve apresentar:

```text
failed=0
```

---

## 14. Deploy Automatizado com Ansible

O deploy manual seria aproximadamente:

```text
Criar rede
    |
    v
Verificar Compose
    |
    v
Executar build
    |
    v
Subir containers
    |
    v
Verificar containers
```

Com Ansible:

```text
                  ANSIBLE
                     |
                     v
          +----------------------+
          | Verifica ambiente    |
          +----------------------+
                     |
                     v
          +----------------------+
          | Verifica Docker      |
          +----------------------+
                     |
                     v
          +----------------------+
          | Verifica Compose     |
          +----------------------+
                     |
                     v
          +----------------------+
          | Valida configuração  |
          +----------------------+
                     |
                     v
          +----------------------+
          | Executa deploy       |
          +----------------------+
                     |
                     v
          +----------------------+
          | Verifica containers  |
          +----------------------+
```

O comando principal é:

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

Isso torna o processo mais padronizado e reduz etapas manuais.

---

## 15. Git e GitHub

O **Git** foi utilizado para controlar as versões do projeto.

O **GitHub** foi utilizado como repositório remoto.

### Verificar status

```bash
git status
```

Esse comando mostra:

- arquivos modificados;
- arquivos novos;
- arquivos preparados para commit;
- situação da branch.

---

### Visualizar histórico

```bash
git log --oneline -5
```

Esse comando mostra os últimos commits de maneira resumida.

---

### Configuração do usuário Git

Durante o desenvolvimento foi necessário configurar a identidade utilizada nos commits.

Exemplo:

```bash
git config --global user.email "seu-email@example.com"
git config --global user.name "Seu Nome"
```

---

### Adicionar alterações

Para adicionar arquivos específicos:

```bash
git add README.md
```

ou:

```bash
git add compose.yaml grafana
```

Para adicionar o Ansible:

```bash
git add ansible
```

---

### Commit do Grafana

Foi utilizado um commit para registrar a implementação do Grafana e do dashboard.

Exemplo:

```bash
git commit -m "feat: adiciona Grafana e dashboard provisionado"
```

---

### Commit do Ansible

Foi utilizado:

```bash
git commit -m "feat: adiciona automacao de deploy com Ansible"
```

---

### Commit da documentação

Para registrar este README:

```bash
git add README.md
git commit -m "docs: adiciona documentacao completa do projeto"
```

---

### Enviar para o GitHub

```bash
git push origin main
```

---

### Verificação final

```bash
git status
```

O resultado ideal é:

```text
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean
```

Isso significa que todas as alterações foram registradas e enviadas ao repositório remoto.

---

## 16. Conclusão

Este projeto demonstra a integração de diferentes tecnologias utilizadas no desenvolvimento de aplicações e em ambientes DevOps.

A aplicação principal foi desenvolvida utilizando:

```text
Go
```

Ela é executada dentro de um:

```text
Container Docker
```

O acesso externo passa pelo:

```text
Nginx
```

As métricas são disponibilizadas pela aplicação e coletadas pelo:

```text
Prometheus
```

Essas métricas são apresentadas visualmente pelo:

```text
Grafana
```

Os serviços são organizados através do:

```text
Docker Compose
```

A comunicação ocorre através da:

```text
projeto-korp-network
```

O processo de deploy é automatizado através do:

```text
Ansible
```

E todas as alterações são controladas através de:

```text
Git + GitHub
```

### Arquitetura final

```text
                         INTERNET / USUÁRIO
                                |
                                |
                                v
                       +----------------+
                       |     NGINX      |
                       |    Porta 80    |
                       | Proxy Reverso  |
                       +----------------+
                                |
                                |
                                v
                       +----------------+
                       |  APLICAÇÃO GO  |
                       |   Porta 8080   |
                       | Servidor HTTP  |
                       +----------------+
                                |
                                |
                         expõe métricas
                                |
                                v
                       +----------------+
                       |   PROMETHEUS   |
                       |   Porta 9090   |
                       | Coleta métricas|
                       +----------------+
                                |
                                |
                                v
                       +----------------+
                       |    GRAFANA     |
                       |   Porta 3000   |
                       |   Dashboard    |
                       +----------------+


                +--------------------------------+
                |          DOCKER                |
                |                                |
                | Executa todos os containers    |
                +--------------------------------+
                               |
                               v
                +--------------------------------+
                |       DOCKER COMPOSE           |
                |                                |
                | Organiza serviços, volumes,    |
                | portas, dependências e rede    |
                +--------------------------------+


                +--------------------------------+
                |            ANSIBLE             |
                |                                |
                | Automatiza o processo de       |
                | verificação e deploy           |
                +--------------------------------+


                +--------------------------------+
                |        GIT + GITHUB            |
                |                                |
                | Versionamento e armazenamento  |
                | do código-fonte                |
                +--------------------------------+
```

### Resumo final das tecnologias

| Tecnologia | Responsabilidade | Explicação para iniciantes |
|---|---|---|
| **Go** | Aplicação | Executa a lógica do servidor |
| **Docker** | Containerização | Cria ambientes isolados |
| **Dockerfile** | Construção da imagem | Define como preparar a aplicação |
| **Docker Compose** | Organização | Gerencia vários containers |
| **Nginx** | Proxy reverso | Funciona como porta de entrada |
| **Prometheus** | Métricas | Coleta informações da aplicação |
| **Grafana** | Dashboard | Exibe as métricas visualmente |
| **Ansible** | Automação | Automatiza o deploy |
| **Git** | Versionamento | Mantém histórico das alterações |
| **GitHub** | Repositório remoto | Armazena e compartilha o projeto |

### Resultado

Ao final do desafio foi obtida uma arquitetura que reúne:

- aplicação HTTP desenvolvida em Go;
- imagem Docker com processo de build;
- múltiplos containers;
- proxy reverso;
- rede Docker;
- Docker Compose;
- coleta de métricas;
- Prometheus;
- dashboard Grafana;
- provisionamento automático;
- automação com Ansible;
- testes dos serviços;
- controle de versão;
- publicação no GitHub.

Dessa forma, o projeto não representa apenas um servidor HTTP simples, mas uma pequena infraestrutura completa contendo conceitos de **desenvolvimento backend, containerização, redes, proxy reverso, observabilidade, monitoramento, infraestrutura como código, automação de deploy e versionamento**.
