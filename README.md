# Aprendendo-Go

🛠️ Desafio de Aquecimento (Nível: Fácil)
Para fixar e já avançar um pouco na lógica, quero que você modifique ou crie um novo programa que resolva o seguinte:

O Problema: Crie um programa que classifique a temperatura de um banho de acordo com o valor em Celsius:

Se a temperatura for menor que 15°C, exiba: "Água muito fria!".

Se estiver entre 15°C e 30°C, exiba: "Água morna/agradável".

Se for maior que 30°C, exiba: "Água quente!".

Requisitos:

Use uma variável para a temperatura.

Use estruturas if / else if / else.

Tente usar o operador fmt.Scan() para que o usuário digite a temperatura no terminal.

🔄 Desafio 2 (Nível: Médio) - Entrando nos Loops
Agora que a leitura de dados funciona, imagine que é chato ter que rodar o programa toda vez que quisermos testar uma nova temperatura. Seria legal se o programa ficasse pedindo temperaturas continuamente.

Em Go, só existe um tipo de laço de repetição: o for. Não temos while ou do while. O for do Go é poderoso e faz o trabalho de todos eles.

Seu novo desafio:
Coloque a lógica do seu código (o Print, o Scan e os if/else) dentro de um laço for infinito.
Mas atenção: o programa precisa ter uma forma de parar! Se o usuário digitar um valor impossível para a água, digamos -100, o programa deve exibir "Saindo..." e encerrar.

Dica: Para criar um loop infinito em Go, você pode apenas escrever for { ... }. Para "quebrar" o loop e sair dele, usamos a palavra-chave break.

🛠️ Desafio 3 (Nível: Fácil/Médio) - Refatorando para Funções
Vamos melhorar o nosso código do termômetro organizando ele com funções!

Sua missão:

Crie uma função nova (fora da main) chamada avaliarTemperatura.

Essa função deve receber a temperatura (float64) como parâmetro.

Coloque toda aquela lógica de if / else if / else dentro dessa função.

A função deve retornar uma string (o texto "A temperatura está fria", etc.).

Dentro da sua main, você vai continuar lendo a temperatura no seu loop for, mas vai chamar a sua nova função e imprimir o resultado que ela devolver.