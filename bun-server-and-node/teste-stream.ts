import process from "node:process"; // Importa o módulo process para usar stdout

// URL da API de streaming
const url = "http://localhost:8000/stream";

// Função para consumir a stream
async function consumeStream() {
  const response = await fetch(url);

  if (!response.body) {
    console.error("A resposta não contém um corpo de stream.");
    return;
  }

  const reader = response.body.getReader();
  const decoder = new TextDecoder("utf-8");

  console.log("Iniciando o consumo da stream:");

  let fullSentence = ""; // Variável para armazenar a frase completa

  while (true) {
    const { done, value } = await reader.read();

    if (done) {
      console.log("Stream finalizada.");
      console.log(fullSentence.trim()); // Exibe a frase completa
      break;
    }

    // Decodifica o chunk recebido e processa imediatamente
    const chunk = decoder.decode(value, { stream: true });

    // Extrai a palavra do JSON recebido e exibe imediatamente
    try {
      const parsed = JSON.parse(chunk);
      if (parsed.word) {
        process.stdout.write(parsed.word + " "); // Exibe a palavra sem quebra de linha
      }
    } catch (err) {
      console.error("Erro ao processar o chunk:", err);
    }
  }
}

// Inicia o consumo da stream
consumeStream().catch((err) =>
  console.error("Erro ao consumir a stream:", err)
);
