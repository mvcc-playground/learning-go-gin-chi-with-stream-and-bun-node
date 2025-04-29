import { serve } from "bun";

// Frase a ser transmitida em partes
const streamPhrase = "Streaming responses are fun and useful";

// Servidor de streaming
serve({
  port: 8000,
  fetch(req) {
    if (req.url === "/stream") {
      const { readable, writable } = new TransformStream();
      const writer = writable.getWriter();

      const words = streamPhrase.split(" ");

      const sendWord = async (index: number) => {
        if (index < words.length) {
          const chunk = JSON.stringify({ word: words[index] }) + "\n";
          await writer.write(chunk);
          setTimeout(() => sendWord(index + 1), 500);
        } else {
          writer.close();
        }
      };

      sendWord(0);
      return new Response(readable, {
        headers: {
          "Content-Type": "application/json",
          "Transfer-Encoding": "chunked",
        },
      });
    }

    return new Response("Not Found", { status: 404 });
  },
});
