
# IP-MONITORING

O projeto IP Monitoring é uma aplicação desenvolvida em Go que tem como principal objetivo monitorar o status de diversos endereços IP e URLs. Esta aplicação verifica periodicamente o status HTTP, a latência e outras informações relacionadas a uma lista de URLs fornecidas, oferecendo insights valiosos para administradores de sistemas e desenvolvedores.


## Melhorias

- **Melhorias de performance:** Otimizei o código para melhorar a performance. Por exemplo, utilizamos o pacote "sync" do Go para gerenciar o acesso simultâneo à estrutura de dados que contém os IPs a serem monitorados, reduzindo a chance de condições de corrida.

- **Integração com webhook do Discord:** Uma funcionalidade para enviar notificações automáticas para um canal Discord através de um webhook. Isso permite uma forma fácil e automatizada de manter-se atualizado sobre o status dos IPs sendo monitorados.

- **Manuseio de erro robusto:** Melhorei o manuseio de erros em todo o código. Agora, o programa faz uma verificação de erros após cada operação que pode potencialmente falhar, e lida com esses erros de uma maneira que não causa uma falha completa do programa.

- **Acessibilidade e usabilidade:** Adicionei mensagens de log claras e informativas em todo o código para facilitar a depuração e entender o que o programa está fazendo em qualquer momento.


## Roadmap

- **Testes Automatizados:** Planejo implementar mais testes automatizados para melhorar a qualidade do código e garantir que a aplicação esteja sempre funcionando como esperado.

- **Configuração Externa:** Atualmente, a URL do webhook do Discord e a lista de URLs monitorados estão codificadas diretamente no código-fonte. Planejo tornar isso configurável externamente para facilitar o uso da aplicação.

- **Notificações de Alerta:** Além de enviar atualizações de status, também planejo adicionar recursos de notificações de alerta. Por exemplo, se um endereço IP passar a retornar um status HTTP de erro, ou se a latência aumentar significativamente, a aplicação poderá enviar uma notificação de alerta.

- **Suporte para Mais Serviços de Webhook:** Atualmente, a aplicação suporta webhooks do Discord. No futuro, planejamos adicionar suporte para mais serviços de webhook, como Slack e Microsoft Teams.

## FAQ

#### Como a aplicação monitora os endereços IP?

A aplicação lê uma lista de URLs a partir de um arquivo chamado "urls.txt". Para cada URL, a aplicação faz uma solicitação para obter o status do IP, incluindo o endereço IP, o status HTTP, a latência, o registrador do Whois, a data de criação do Whois e a data de expiração do Whois. Este processo é repetido a cada hora.

#### Como posso configurar o webhook do Discord?

Você precisa fornecer a URL do webhook do Discord diretamente no código-fonte, no arquivo "main.go". Uma vez que você tenha fornecido a URL do webhook, a aplicação enviará automaticamente as atualizações de status para esse webhook.

#### Como posso modificar os endereços URL que estão sendo monitorados?

Os URLs monitorados são lidos a partir de um arquivo chamado "urls.txt". Para monitorar URLs diferentes, basta modificar este arquivo para incluir os URLs que você deseja monitorar, um por linha. A aplicação irá ler este arquivo e começar a monitorar os novos URLs na próxima vez que for iniciada.


## Autores

- [@kevinfinalboss](https://www.github.com/kevinfinalboss)


## Rodando localmente

Clone o projeto

```bash
  git clone git@github.com:kevinfinalboss/ip-monitoring.git
```

Entre no diretório do projeto

```bash
  cd ip-monitoring
```

Instale as dependências

```bash
  go mod tidy
```

Inicie o servidor

```bash
  go run .
```


## Rodando os testes

Antes certifique que está no diretorio de testes e depois execute o seguinte comando.

```bash
  go test -v ./...
```

