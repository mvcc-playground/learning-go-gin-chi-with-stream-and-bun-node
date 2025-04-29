import express from "express";

const app = express();
const port = 8000;

// Frase a ser transmitida em partes
const streamPhrase = "Streaming responses are fun and useful";

// Rota de streaming
app.get("/stream", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  res.setHeader("Transfer-Encoding", "chunked");

  const words = streamPhrase.split(" ");

  const sendWord = (index: number) => {
    if (index < words.length) {
      const chunk = JSON.stringify({ word: words[index] }) + "\n";
      res.write(chunk);
      setTimeout(() => sendWord(index + 1), 500);
    } else {
      res.end();
    }
  };

  sendWord(0);
});

app.listen(port, () => {
  console.log(`Node.js server running at http://localhost:${port}/stream`);
});
