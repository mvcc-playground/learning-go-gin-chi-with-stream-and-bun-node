const http = require("http");

// Frase a ser transmitida em partes
const streamPhrase = "Streaming responses are fun and useful";

// Criação do servidor HTTP
const server = http.createServer((req, res) => {
  if (req.url === "/stream" && req.method === "GET") {
    // Configura os cabeçalhos para streaming
    res.setHeader("Content-Type", "application/json");
    res.setHeader("Transfer-Encoding", "chunked");

    const words = streamPhrase.split(" ");

    const sendWord = (index) => {
      if (index < words.length) {
        const chunk = JSON.stringify({ word: words[index] }) + "\n";
        res.write(chunk);
        setTimeout(() => sendWord(index + 1), 500);
      } else {
        res.end();
      }
    };

    sendWord(0);
  } else {
    res.statusCode = 404;
    res.end("Not Found");
  }
});

// Inicia o servidor na porta 8000
server.listen(8000, () => {
  console.log("Node.js server running at http://localhost:8000/stream");
});
