# gobook

# GoBook API - README

A GoBook API é uma aplicação baseada em Go que utiliza a framework Gin que oferece funcionalidades CRUD com integração ao MongoDB.

## Funcionalidades Principais
A GoBook API oferece as seguintes funcionalidades principais:

1. **POST**: Permite a criação de novos registros no banco de dados, permitindo adicionar novos livros à biblioteca.

2. **GET**: Permite a recuperação de registros existentes no banco de dados, fornecendo acesso aos detalhes dos livros cadastrados.

3. **DELETE**: Permite a exclusão de registros do banco de dados, possibilitando a remoção de livros da biblioteca.

## Banco de Dados
A GoBook API está conectada a um banco de dados, fornecendo um armazenamento persistente para os dados da biblioteca. O banco de dados utilizado pode ser configurado de acordo com as necessidades do ambiente de execução.

## Configuração
Antes de executar a aplicação, é necessário configurar as informações de conexão ao banco de dados no arquivo de configuração correspondente. Certifique-se de fornecer as credenciais corretas e outras configurações relevantes, conforme documentado.

## Uso
Após configurar a conexão com o banco de dados, você pode iniciar ao utilizar o comando `make run` no linux ou `mingw32-make run` no windows. Utilize as rotas adequadas para realizar as operações CRUD, enviando as requisições HTTP correspondentes aos endpoints disponíveis.

## Configuração de Testes
Antes de executar os testes, é necessário configurar a conexão com o banco de dados no arquivo de configuração correspondente, fornecendo as informações corretas para o ambiente de testes.

## Teste: PostBookHandler
Este teste verifica a funcionalidade de criação de um novo livro na biblioteca. Um livro de teste é criado e enviado como uma requisição HTTP POST para o endpoint "/books". O teste verifica se a resposta HTTP retornada é igual a "StatusCreated" (201), indicando que o livro foi criado com sucesso.

## Teste: GetBooksHandler
Este teste verifica a funcionalidade de recuperação de todos os livros da biblioteca. Uma requisição HTTP GET é enviada para o endpoint "/books", e o teste verifica se a resposta HTTP retornada é igual a "StatusOK" (200), indicando que a recuperação foi bem-sucedida.

## Teste: DeleteBooksHandler
Este teste verifica a funcionalidade de exclusão de um livro da biblioteca. Primeiro, um livro de teste é criado e enviado como uma requisição HTTP POST para o endpoint "/books" para adicioná-lo à biblioteca. Em seguida, uma requisição HTTP DELETE é enviada para o endpoint "/books/{uuid}", onde {uuid} é o identificador único do livro criado. O teste verifica se a resposta HTTP retornada é igual a "StatusOK" (200), indicando que a exclusão do livro foi bem-sucedida.
