# Go E-commerce

Uma aplicação web de e-commerce simples e moderna construída em Go, com autenticação de usuário, carrinho de compras e um painel administrativo para gerenciamento de produtos. O projeto segue uma arquitetura limpa, separando responsabilidades em *handlers*, *services* e *repositories*.

## Recursos

* **Gerenciamento de Usuários:** Registro, login e logout seguros.
* **Catálogo de Produtos:** Navegação pelos produtos na página inicial e visualização de detalhes de cada item.
* **Carrinho de Compras:** Adicionar, remover e atualizar quantidades de produtos no carrinho.
* **Processo de Checkout:** Fluxo de checkout em múltiplas etapas com sistema de pagamento simulado.
* **Rotas Protegidas:** Middleware para proteger páginas específicas do usuário, como dashboard e carrinho.
* **Painel Administrativo:** Área dedicada para administradores realizarem operações CRUD (Criar, Ler, Atualizar e Excluir) em produtos.
* **Arquitetura em Camadas:** Organizado em camadas distintas (*handlers*, *services*, *repositories*) para maior manutenibilidade e escalabilidade.

## Stack Tecnológico

* **Backend:** Go
* **Roteador:** [chi/v5](https://github.com/go-chi/chi)
* **Banco de Dados:** MongoDB
* **Variáveis de Ambiente:** [joho/godotenv](https://github.com/joho/godotenv)

## Início Rápido

Siga estas instruções para obter uma cópia do projeto e rodá-lo localmente para desenvolvimento e testes.

### Pré-requisitos

* Go (versão 1.18 ou superior)
* Uma instância do MongoDB em execução (local ou remota, como MongoDB Atlas)

### Instalação

1. **Clone o repositório:**

   ```sh
   git clone https://github.com/MarcosAndradeV/go-ecommerce.git
   cd go-ecommerce
   ```

2. **Configure as variáveis de ambiente:**
   Crie um arquivo `.env` na raiz do projeto e adicione as seguintes variáveis. Um arquivo `.env.example` é fornecido como modelo.

   ```env
   # Porta em que o servidor web irá escutar
   PORT=8080

   # Sua string de conexão com o MongoDB
   MONGO_URI="mongodb://localhost:27017"

   # Nome do banco de dados a ser usado
   DB_NAME="ecommerce"
   ```

3. **Instale as dependências:**
   O projeto usa Go Modules. As dependências serão baixadas automaticamente na primeira execução ou compilação do projeto.

4. **Execute a aplicação:**

   ```sh
   go run ./cmd/web/main.go
   ```

   O servidor estará rodando em `http://localhost:8080` (ou na porta especificada por você).

## Estrutura do Projeto

O projeto segue um layout padrão de projetos Go para separar responsabilidades:

```
.
├── cmd/web/          # Ponto de entrada principal da aplicação
├── internal/         # Código interno e privado da aplicação
│   ├── database/     # Conexão com o banco e lógica de persistência
│   ├── handlers/     # Handlers HTTP (controladores)
│   ├── repository/   # Camada de acesso a dados (interage com o banco)
│   ├── routes/       # Definição de rotas (usando chi)
│   └── service/      # Lógica de negócios
├── static/           # Arquivos estáticos (CSS, JS, imagens)
├── templates/        # Templates HTML para renderização dinâmica
└── go.mod            # Dependências do módulo Go
```

## Rotas da Aplicação

A aplicação expõe as seguintes rotas, gerenciadas pelo roteador `chi`:

### Rotas Públicas

* `GET /`: Página inicial com listagem de produtos.
* `GET /product/{id}`: Detalhes de um produto específico.
* `GET /register`, `POST /register`: Registro de usuário.
* `GET /login`, `POST /do-login`: Login de usuário.
* `GET /logout`: Logout do usuário.

### Rotas Protegidas (Requer Autenticação)

* `GET /dashboard`: Dashboard do usuário.
* `GET /cart`: Visualizar o carrinho de compras.
* `GET /add-to-cart`, `GET /remove-from-cart`, `POST /update-cart`: Gerenciamento do carrinho.
* `GET /checkout`, `POST /checkout`: Processo de checkout.
* `POST /payment`, `POST /purchase`, `POST /purchase/simulate/{id}`: Fluxo de pagamento e compra.

### Rotas de Administração

* `GET /admin/dashboard`: Dashboard do administrador.
* `POST /admin/create`: Criar um novo produto.
* `GET /admin/edit/product/{product_id}`, `POST /admin/edit/product`: Editar um produto.
* `POST /admin/delete/product/{id}`: Excluir um produto.
