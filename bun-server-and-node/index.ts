import { z } from "@zod/mini";

const responseSchema = z.object({
  word: z.string(),
});
