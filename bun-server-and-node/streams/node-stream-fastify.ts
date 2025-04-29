import Fastify from "fastify";

const fastify = Fastify();
const port = 8000;

// Frase a ser transmitida em partes
const streamPhrase = "Streaming responses are fun and useful";

// Rota de streaming
fastify.get("/stream", async (request, reply) => {
  reply.raw.setHeader("Content-Type", "application/json");
  reply.raw.setHeader("Transfer-Encoding", "chunked");

  const words = streamPhrase.split(" ");

  const sendWord = async (index: number) => {
    if (index < words.length) {
      const chunk = JSON.stringify({ word: words[index] }) + "\n";
      reply.raw.write(chunk);
      setTimeout(() => sendWord(index + 1), 500);
    } else {
      reply.raw.end();
    }
  };

  sendWord(0);
});

fastify.listen({ port }, (err, address) => {
  if (err) {
    console.error(err);
    process.exit(1);
  }
  console.log(`Fastify server running at ${address}/stream`);
});
