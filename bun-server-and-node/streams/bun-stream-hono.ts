import { Hono } from "hono";

const port = 8000;

// Frase a ser transmitida em partes
const streamPhrase = "Streaming responses are fun and useful";

// Rota de streaming
const app = new Hono().get("/stream", (c) => {
  c.header("Content-Type", "application/json");
  c.header("Transfer-Encoding", "chunked");

  const words = streamPhrase.split(" ");

  const readable = new ReadableStream({
    start(controller) {
      let index = 0;

      const push = () => {
        if (index < words.length) {
          const chunk = JSON.stringify({ word: words[index] }) + "\n";
          controller.enqueue(new TextEncoder().encode(chunk));
          index++;
          setTimeout(push, 500);
        } else {
          controller.close();
        }
      };

      push();
    },
  });

  return new Response(readable);
});

Bun.serve({
  fetch: app.fetch,
});
