# Chatbot (GPT-3.5-Turbo) em Golang

Este projeto é uma aplicação em Golang que faz integração com a API do OpenAI GPT-3.5-Turbo para criar um chatbot capaz de responder perguntas e realizar conversas com os usuários.

## Requisitos
Para executar a aplicação, você precisará do seguinte:

Go 1.16 ou superior instalado em sua máquina
Uma chave de API válida do OpenAI para acessar a API do GPT-3.5-Turbo
Como usar
Antes de começar, certifique-se de que possui uma chave de API válida do OpenAI. Você pode se inscrever para obter uma chave aqui.

Clone este repositório em sua máquina:

```bash
git clone https://github.com/seu-usuario/gpt-3.5-turbo-chatbot.git
```
Na pasta do projeto, crie um arquivo .env e adicione sua chave de API do OpenAI como uma variável de ambiente, como mostrado abaixo:
```makefile
OPENAI_API_KEY=SUA_CHAVE_DE_API
```
Execute o seguinte comando para baixar as dependências do projeto:
```golang
go mod tidy
```
Agora, execute o seguinte comando para iniciar o chatbot:
```go
go run ./cmd/main.go
```
O chatbot está pronto para uso! Você pode digitar suas perguntas ou conversar com ele.
Detalhes técnicos
O chatbot utiliza a api https://api.openai.com/v1/chat/completions, da OpenAI, para receber respostas e exibi-las ao usuário.

O bot usa o modelo de geração de texto do GPT-3.5-Turbo para gerar respostas para as perguntas do usuário.

Para obter mais informações sobre como usar a API do OpenAI, consulte a documentação oficial em https://beta.openai.com/docs/api-reference.

## License

[MIT](https://choosealicense.com/licenses/mit/)