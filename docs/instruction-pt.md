# Manual de utilização TRIAGE MCI

## 1. O que é o TRIAGE MCI

O TRIAGE MCI é uma aplicação web (PWA) para triagem e gestão de incidentes com múltiplas vítimas (MCI), baseada no algoritmo START. Permite:
- Classificação de pacientes em 4 categorias (T1–T4)
- Gestão de transporte para hospitais
- Monitorização da capacidade hospitalar
- Geração e envio de relatórios
- Importação de dados do operador de central

A aplicação funciona totalmente offline. Todos os dados são armazenados exclusivamente no dispositivo do utilizador — nenhuma informação é enviada para servidores externos.

Suporta 7 idiomas: polaco, inglês, italiano, francês, alemão, checo, português.

---

## 2. Início do incidente

### Ecrã inicial

Ao abrir a aplicação, surge o ecrã de configuração do incidente:

1. **Seleção de idioma** — seletor no topo do ecrã. A alteração do idioma traduz imediatamente toda a interface.
2. **Nome do incidente** — campo obrigatório. Introduza a localização e o tipo de incidente (ex.: "Acidente A1 km 312"). Aparecerá nos relatórios.
3. **KAM** — campo opcional. Responsável pela Ação Médica — nome da pessoa que coordena as operações no local.
4. **Mostrar dicas** — interruptor que ativa balões com sugestões contextuais em toda a aplicação. Útil na primeira utilização.
5. **INICIAR TRIAGE** — inicia o incidente e arranca o cronómetro.

### Retomar incidente

Se existir um incidente guardado na memória do dispositivo, ao iniciar surgirá a opção:
- **CONTINUAR** — retoma o incidente guardado com todos os dados
- **NOVO INCIDENTE** — apaga os dados anteriores e começa de novo

---

## 3. Triagem (algoritmo START)

### Navegação

A aplicação tem 3 separadores na parte superior do ecrã:
- **TRIAGEM** — classificação de pacientes
- **DESPACHO** — transporte e gestão
- **RELATÓRIO** — resumo e exportação

A barra de estatísticas abaixo da navegação mostra o número de pacientes em cada categoria: 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4 e o total.

### Assistente de triagem START

O assistente guia-o através do algoritmo START em 6 passos. Em cada passo responde SIM ou NÃO:

**Passo 1 — O paciente caminha?**
- SIM → T3 Verde (diferido)
- NÃO → avançar para o passo 2

**Passo 2 — Respira espontaneamente?**
- SIM → avançar para o passo 4 (avaliação da frequência respiratória)
- NÃO → avançar para o passo 3

**Passo 3 — Permeabilize a via aérea. Respira após permeabilização?**
- SIM → T1 Vermelho (imediato)
- NÃO → T4 Preto (óbito)

**Passo 4 — A frequência respiratória é normal (≤ 30/min)?**
- SIM → avançar para o passo 5
- NÃO → T1 Vermelho (imediato)

**Passo 5 — O tempo de preenchimento capilar é normal (≤ 2 seg.) e o pulso radial está presente?**
- SIM → avançar para o passo 6
- NÃO → T1 Vermelho (imediato)

**Passo 6 — Obedece a comandos simples?**
- SIM → T2 Amarelo (urgente)
- NÃO → T1 Vermelho (imediato)

Cada pergunta tem uma dica (hint) visível abaixo — descreve como avaliar o respetivo parâmetro.

### Cartão de resultado

Após a conclusão da triagem surge o cartão de resultado com:

- **Marcador colorido** — quadrado grande com a categoria (T1/T2/T3/T4)
- **Percurso decisório** — mostra o caminho seguido no algoritmo
- **Sexo** — botões M (masculino), F (feminino), ? (desconhecido)
- **Idade** — campo numérico com botões -5, -1, +1, +5 (intervalo 0–150)
- **Lesões corporais** — botão que abre o diagrama de lesões (ver secção 4)
- **Notas** — campo para informações adicionais (lesões, mecanismo de trauma, etc.)

Ações disponíveis:
- **✓ CONFIRMAR PACIENTE** — adiciona o paciente à lista
- **↺ RECOMEÇAR** — reinicia o assistente a partir do passo 1
- **ALTERAR COR MANUALMENTE** — abre uma fila de 4 botões (T1–T4) para alteração manual da categoria

### Lista de pacientes no local

Abaixo do assistente é apresentada a lista "Pacientes no local", ordenada por prioridade (T1 → T2 → T3 → T4).

Cada cartão de paciente contém:
- Marcador colorido (ex.: P-001)
- Categoria, sexo/idade, hora de registo
- Campo de notas (editável)
- Botão de lesões (🩻) com resumo
- 4 pontos de retriagem (T1–T4) — clique para alterar a categoria
- Botão de eliminação (✕)
- Marcador de retriagem (↺n) caso a categoria tenha sido alterada

### Lista em transporte

A secção "Em transporte" aparece quando há pacientes a caminho do hospital. Cada cartão mostra:
- Marcador do paciente com a cor da categoria
- Rota: Equipa → Hospital → hora de saída
- Botão "alterar" — alteração do hospital de destino
- Botão "✓ Entregue" — marca como entregue e liberta a equipa de emergência

---

## 4. Diagrama de lesões corporais

O botão "🩻 Lesões corporais" abre um diagrama corporal interativo.

### Vista do diagrama

São apresentados dois contornos do corpo lado a lado:
- **FRENTE** (front) — cabeça, tórax, abdómen, braços, pernas
- **COSTAS** (back) — cabeça, parte superior das costas, parte inferior das costas, braços, pernas

Nos lados são visíveis as indicações de lateralidade: **E** (esquerda) e **D** (direita), invertidas na vista posterior.

Clique numa zona do corpo para a selecionar.

### Painel de lesões da zona

Ao selecionar uma zona, abre-se um painel com 7 tipos de lesão para assinalar:
- Fratura
- Hemorragia
- Queimadura
- Ferida
- Edema
- Dor
- Amputação

Os tipos assinalados são marcados com ✓. Cada um pode ser ativado/desativado com um clique.

Na parte inferior do painel:
- **✕ NÃO** — descarta as alterações nesta zona
- **✓ SIM** — confirma as lesões para a zona selecionada

### Confirmação do diagrama

Após assinalar as lesões em todas as zonas necessárias:
- **✓ SIM** — confirma todo o diagrama de lesões
- **✕ NÃO** — descarta todas as alterações e regressa ao estado anterior à abertura

As lesões assinaladas são visíveis:
- No cartão do paciente como resumo abreviado (ex.: "🩻 Cabeça: Fratura, Hemorragia")
- No relatório junto ao respetivo paciente

---

## 5. Despacho

### Janelas modais de configuração

Ao entrar pela primeira vez no separador Despacho, surgem sequencialmente 3 janelas modais:

**1. Diretor de Gestão Médica (GDM)**
- Campo para nome do GDM
- Botão de importação de dados do operador de central (ver secção 6)
- SALTAR ou GUARDAR

**2. Equipas de emergência**
- Pergunta: "Conhece as equipas de emergência enviadas para o local?"
- SIM → formulário para adicionar equipas (nome + botão adicionar)
- NÃO → utilizar equipas predefinidas (S-01, S-02, P-01, P-02, P-03, LPR)
- CONFIRMAR ou SALTAR (PREDEFINIDAS)

**3. Hospitais de destino**
- Formulário para adicionar hospitais com capacidade:
  - Nome do hospital
  - 🔴 Capacidade vermelha (T1) — número de camas
  - 🟡 Capacidade amarela (T2) — número de camas
- CONFIRMAR ou SALTAR (PREDEFINIDOS)

Cada janela modal aparece apenas uma vez por incidente.

### Formulário de despacho

A secção "Despachar transporte" contém:

1. **Paciente** — lista pendente com pacientes no local, ordenada por prioridade
2. **Equipa de emergência** — lista de equipas disponíveis (equipas ocupadas ficam ocultas)
3. **Hospital** — lista de hospitais com informação de ocupação:
   - Formato: "Nome do hospital [🔴 ocupados/total 🟡 ocupados/total]"
   - Hospitais sem vagas na categoria do paciente assinalados com aviso
4. **DESPACHAR TRANSPORTE** — botão ativo quando todos os campos estão preenchidos

### Painéis de gestão

Abaixo do formulário encontram-se dois painéis (lado a lado em ecrãs largos, sobrepostos em ecrãs estreitos):

**Hospitais** (camas TOTAIS):
- Lista de hospitais com campos de capacidade editáveis (🔴 e 🟡)
- Adição de novos hospitais
- Eliminação de hospitais (✕)
- Expandir/recolher a lista

**Equipas de emergência:**
- Lista de equipas com botões de eliminação
- Adição de novas equipas
- Expandir/recolher a lista

### Histórico de transportes

Na parte inferior da secção de despacho (recolhido por defeito). Expanda para ver a lista cronológica de todos os transportes com:
- Marcador do paciente com cor
- Equipa → Hospital → hora de saída
- Estado (em transporte / entregue)

---

## 6. Importação de dados do operador de central

Esta função permite ao operador de central enviar dados (nome do incidente, GDM, equipas de emergência, hospitais com capacidade) para o socorrista no terreno.

### Método 1 — Link por SMS (um clique)

O operador de central envia um SMS com um link no formato:
```
https://0xjaqbek.github.io/triage/?i=DADOS_BASE64
```

Ao clicar no link:
1. A aplicação abre a partir da cache (não necessita de internet se estiver instalada)
2. Surge uma pré-visualização dos dados: nome do incidente, GDM, equipas, hospitais
3. **IMPORTAR** — importa os dados e abre a triagem
4. **REJEITAR** — ignora os dados

### Método 2 — Colar dados

Se o link não abrir a aplicação (dispositivos bloqueados):
1. Abra a aplicação, vá ao separador Despacho
2. Na janela modal do GDM clique em **📥 IMPORTAR DADOS DO OPERADOR**
3. Cole o conteúdo do SMS (o link completo ou apenas o código Base64)
4. A pré-visualização dos dados surgirá automaticamente
5. **IMPORTAR** — confirma os dados

### Importação durante o incidente

Se o incidente já estiver em curso e tiver pacientes triados:
- **Pacientes** — mantidos (nunca são eliminados)
- **KAM** — mantido
- **GDM** — substituído pelos dados do operador
- **Equipas de emergência** — substituídas pelos dados do operador
- **Hospitais** — substituídos pelos dados do operador
- **Nome do incidente** — combinado: "O seu nome | Nome do operador (operador)"

Um aviso amarelo na pré-visualização informa sobre a manutenção dos pacientes e do KAM.

---

## 7. Modo operador de central

Acessível através do link "Modo operador de central →" no ecrã inicial ou no endereço `/dyspozytor/`.

Esta página permite ao operador de central preparar e enviar dados para a equipa no terreno.

### Formulário

1. **Nome do incidente** — ex.: "MCI Lisboa-Expo"
2. **Diretor de Gestão Médica (GDM)** — nome completo
3. **Equipas de emergência** — adicionar uma a uma (nome + botão "+")
4. **Hospitais de destino** — nome + capacidade 🔴 e 🟡 (+ botão adicionar)

### Envio

- **📋 Pré-visualização dos dados** — mostra o resumo com o tamanho em caracteres e número de SMS
- **📱 Enviar SMS** — abre a aplicação de SMS com o link pronto
- **📋 Copiar dados** — copia o link para a área de transferência

Os dados são codificados em Base64 (compatibilidade com SMS GSM-7 — até 1530 caracteres em 10 segmentos de SMS). Comporta aproximadamente 20 hospitais e 20 equipas.

---

## 8. Relatório

### Relatório visual

A secção "Relatório para o operador" é gerada automaticamente e contém:

**Cabeçalho:**
- Nome do incidente, hora de início, hora do relatório
- KAM e GDM (se indicados)
- Hora de encerramento (se o incidente estiver encerrado)

**Resumo:**
- Marcadores coloridos com o número de pacientes em cada categoria
- Número total de pacientes

**Hospitais** (bloco separado para cada hospital com transportes):
- Nome do hospital, número de pacientes, informação de capacidade
- Lista de pacientes com: marcador, idade/sexo, equipa, hora, estado (EM TRANSPORTE / ENTREGUE)
- Histórico de retriagem (ex.: ↺T1→T2→T1)
- Histórico de alterações de hospital (ex.: ⇄Hospital1→Hospital2)
- Resumo de lesões

**Pacientes no local** (se alguém permanecer):
- Lista com: marcador, idade/sexo, categoria, notas, lesões

### Relatório textual

O botão "📋 Relatório textual" abre uma janela modal com o relatório completo em formato de texto. Contém a lista completa de pacientes com notas, sexo, idade e percurso de triagem.

### Envio do relatório

- **📋 Copiar para a área de transferência** — copia o texto completo do relatório
- **📤 Enviar** — abre opções:
  - **SMS** — abre a aplicação nativa de SMS com o conteúdo do relatório
  - **Email** — abre o cliente de email com o assunto e conteúdo do relatório

### Encerramento do incidente

Botão **ENCERRAR INCIDENTE** no cabeçalho do relatório:
- Regista a hora de encerramento
- Transforma-se em **REVERTER ENCERRAMENTO** (reversível)
- A hora de encerramento aparece no relatório

---

## 9. Retriagem e alteração de hospital

### Retriagem

No cartão do paciente na secção "Pacientes no local" encontram-se 4 pontos coloridos (T1–T4). Clicar numa categoria diferente da atual:
1. Abre uma janela de confirmação: "Alterar a categoria do paciente P-001 de T1 para T2?"
2. **CONFIRMAR** — altera a categoria, regista no histórico
3. **CANCELAR** — sem alterações

Histórico de retriagem:
- Marcador ↺n no cartão do paciente (n = número de alterações)
- Histórico completo no relatório (ex.: T1→T2→T1 com horas)

### Alteração de hospital durante o transporte

No cartão do paciente "Em transporte", botão "alterar":
1. Abre uma janela modal com o hospital atual
2. Lista pendente com hospitais disponíveis (com informação de capacidade)
3. **CONFIRMAR** — altera o hospital, regista no histórico
4. **CANCELAR** — sem alterações

O histórico de alterações de hospital é visível no relatório (⇄Hospital1→Hospital2).

---

## 10. Instalação e modo offline

### Instalação como PWA

**Android (Chrome):**
1. Abra a aplicação no Chrome
2. Toque em ⋮ (menu) no canto superior direito
3. Selecione "Adicionar ao ecrã inicial"
4. Toque em "Instalar"

**iOS (Safari):**
1. Abra a aplicação no Safari
2. Toque em ⎙ (Partilhar)
3. Deslize para baixo e selecione "Adicionar ao ecrã inicial"

Após a instalação, a aplicação funciona como uma aplicação independente com modo offline completo.

### Modo offline

Após o primeiro carregamento, a aplicação é guardada na cache do dispositivo. Todas as funcionalidades operam sem ligação à internet:
- Triagem de pacientes
- Despacho de transporte
- Geração de relatórios
- Importação de dados (colagem)

A internet é necessária apenas para: o primeiro download, atualizações e envio de SMS/email (através das aplicações nativas do dispositivo).

### Download como ficheiro HTML

O botão **⬇ Descarregar** no ecrã inicial descarrega a aplicação como um ficheiro HTML autónomo. Pode ser transferido para uma pen USB e executado em qualquer computador com navegador.

---

## 11. Privacidade

- Todos os dados são armazenados **exclusivamente no dispositivo** (localStorage)
- Nenhum dado **é enviado** para servidores externos
- Sem analítica, rastreamento ou telemetria
- Sem cookies de rastreamento
- As funções de SMS e email utilizam as aplicações nativas do dispositivo — a aplicação não envia mensagens diretamente
- Eliminação de dados: utilize "Encerrar incidente" ou limpe os dados do navegador

A política de privacidade completa está disponível em `/privacy/`.
