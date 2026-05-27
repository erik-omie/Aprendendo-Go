# Aprendendo-Go

Documentação do Projeto: Calculadora
Simples em Go
Versão: 1.0
Linguagem: Go (Golang)
Autor: Desenvolvedor Go

1. Visão Geral do Sistema
Este projeto consiste em uma calculadora baseada em console (CLI) desenvolvida em Go. O
sistema roda em um loop contínuo, permitindo que o usuário selecione operações matemáticas
básicas (Adição, Subtração, Multiplicação e Divisão), insira dois números decimais, veja o
resultado formatado e continue utilizando o programa até que decida sair explicitamente.
O projeto está estruturado de forma modular, separando a lógica de exibição de menus/captura
de dados da lógica de controle principal e das operações matemáticas.

2. Fluxo de Execução
O ciclo de vida do programa segue os seguintes passos dentro de um laço de repetição:
1. O menu principal é renderizado na tela.
2. O sistema aguarda a entrada do usuário para definir a operação desejada e os dois
valores numéricos.
3. A opção digitada é avaliada por uma estrutura de decisão.
4. A função matemática correspondente é invocada e o resultado é exibido de forma
formatada.
5. O ciclo recomeça, a menos que a opção de saída (5) tenha sido acionada.

3. Explicação das Funções
Função / Assinatura Descrição Técnico-Funcional
func Menu() (int, float64, float64) Responsável por desenhar a interface gráfica
textual no console. Ela utiliza ponteiros
através do pacote fmt.Scan para capturar três
valores digitados pelo usuário (opção,
número A, número B) e os retorna
simultaneamente usando o recurso de

Função / Assinatura Descrição Técnico-Funcional
múltiplos retornos do Go.

func main() Ponto de entrada (entrypoint) do executável.
Contém o laço de repetição infinito que
gerencia o estado da calculadora e a
estrutura de decisão switch para direcionar o
fluxo baseado na escolha do usuário.
func Adicao(a, b float64) float64 Recebe dois parâmetros decimais e retorna a

soma aritmética entre eles.

func Subtracao(a, b float64) float64 Recebe dois parâmetros decimais e retorna a

diferença aritmética (a - b).

func Multiplicacao(a, b float64) float64 Recebe dois parâmetros decimais e retorna o

produto da multiplicação (a * b).

func Divisao(a, b float64) float64 Recebe dois parâmetros decimais e retorna o
quociente da divisão (a / b). Nota de boa
prática: deve conter uma validação para
evitar divisões por zero.

4. Dicionário de Palavras-Chave e Conceitos Go
Abaixo estão explicados os termos nativos e pacotes utilizados na arquitetura desta
calculadora:
● package main: Define que este arquivo específico gera um arquivo executável binário
após a compilação, e não apenas uma biblioteca compartilhada.
● import "fmt": Importa o pacote nativo de formatação (Format). Ele provê funções de
entrada e saída de dados, essenciais para interagir com o terminal.
● func: Palavra-chave utilizada para declarar uma nova função ou método no Go.
● int e float64: Tipos de dados nativos. int representa números inteiros (ex: 1, 5, -10).
float64 representa números reais/decimais com precisão de 64 bits (ex: 3.14, 2.0, -0.5).
● := (Operador de Curta Declaração): Cria e inicializa variáveis em uma única instrução,
permitindo que o Go infira o tipo do dado automaticamente sem a necessidade de usar
explicitamente a palavra var.
● for { ... }: Cria um laço de repetição (loop) infinito. Em Go, não existe a palavra-chave
while; o for sem condições cumpre esse papel de rodar o bloco continuamente.
● switch / case: Estrutura de controle condicional limpa. Avalia uma variável contra
múltiplos cenários (cases). No Go, o switch possui um "break implícito", o que significa
que ele executa apenas o caso correspondente e sai da estrutura automaticamente.
● default: A cláusula de escape do switch. É executada caso o valor avaliado não coincida
com nenhum dos case listados (no projeto, serve para tratar opções inválidas).
● return: Finaliza a execução da função atual e devolve os valores especificados para

quem a chamou. Quando usado dentro da função main(), encerra o programa por
completo.
● fmt.Scan(&variavel): Lê o que o usuário digitou no console. O caractere comercial (&)
indica um ponteiro, informando ao Go o endereço de memória exato onde o valor
capturado deve ser salvo.
● fmt.Printf(): Imprime texto formatado na tela. Permite o uso de marcadores como %.2f
(exibir número decimal com apenas duas casas após o ponto) e \n (pular linha no
console).

