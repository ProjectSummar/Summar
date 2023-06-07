import { useMutation, useQuery } from "@tanstack/react-query";
import { z } from "zod";
import { bookmarkSchema } from "@src/types";
import { BASE_URL, serverResponseSchema } from "@src/api/helpers";

const createBookmarkRequestSchema = z.object({
    url: z.string().url(),
});

const createBookmarkResponseSchema = serverResponseSchema.extend({
    bookmark: bookmarkSchema.optional(),
});

type CreateBookmarkRequest = z.infer<typeof createBookmarkRequestSchema>;

const createBookmark = async (req: CreateBookmarkRequest) => {
    try {
        const parsedReq = createBookmarkRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/bookmark`, {
            method: "POST",
            body: JSON.stringify(parsedReq),
        });

        const parsedRes = createBookmarkResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useCreateBookmark = () => {
    return useMutation({
        mutationFn: createBookmark,
    });
};

const getBookmarksResponseSchema = serverResponseSchema.extend({
    bookmarks: z.array(bookmarkSchema.optional()),
});

const getBookmarks = async () => {
    try {
        const res = await fetch(`${BASE_URL}/bookmark`, {
            method: "GET",
        });

        const parsedRes = getBookmarksResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes.bookmarks;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useGetBookmarks = () => {
    return useQuery({
        queryKey: ["bookmarks"],
        queryFn: getBookmarks,
    });
};

const getBookmarkResponseSchema = serverResponseSchema.extend({
    bookmark: bookmarkSchema.optional(),
});

const getBookmark = async (id: string) => {
    try {
        const res = await fetch(`${BASE_URL}/bookmark/${id}`, {
            method: "GET",
        });

        const parsedRes = getBookmarkResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes.bookmark;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useGetBookmark = (id: string) => {
    return useQuery({
        queryKey: ["bookmark", id],
        queryFn: () => getBookmark(id),
    });
};

const deleteBookmarkResponseSchema = serverResponseSchema.extend({
    bookmark: bookmarkSchema.optional(),
});

const deleteBookmark = async (id: string) => {
    try {
        const res = await fetch(`${BASE_URL}/bookmark/${id}`, {
            method: "POST",
        });

        const parsedRes = deleteBookmarkResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useDeleteBookmark = () => {
    return useMutation({
        mutationFn: deleteBookmark,
    });
};

export {
    useCreateBookmark,
    useDeleteBookmark,
    useGetBookmark,
    useGetBookmarks,
};
