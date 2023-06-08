import { z } from "zod";

const userSchema = z.object({
    id: z.string().uuid(),
    email: z.string().email(),
    createdAt: z.string(),
});

type User = z.infer<typeof userSchema>;

const bookmarkSchema = z.object({
    id: z.string().uuid(),
    userId: z.string().uuid(),
    url: z.string().url(),
    title: z.string(),
    summary: z.string(),
});

type Bookmark = z.infer<typeof bookmarkSchema>;

export { Bookmark, bookmarkSchema, User, userSchema };
