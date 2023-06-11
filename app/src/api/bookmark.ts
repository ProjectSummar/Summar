import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { z } from "zod";
import { bookmarkSchema } from "@src/types";
import { BASE_URL, serverResponseSchema } from "@src/api/helpers";

const createBookmarkRequestSchema = z.object({
    url: z.string().trim().url(),
    title: z.string().trim(),
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
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: createBookmark,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["bookmarks"] });
        },
    });
};

const getBookmarksResponseSchema = serverResponseSchema.extend({
    bookmarks: z.array(bookmarkSchema).optional(),
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

const getBookmarkRequestSchema = z.object({
    id: z.string().uuid(),
});

const getBookmarkResponseSchema = serverResponseSchema.extend({
    bookmark: bookmarkSchema.optional(),
});

type GetBookmarkRequest = z.infer<typeof getBookmarkRequestSchema>;

const getBookmark = async (req: GetBookmarkRequest) => {
    try {
        const parsedReq = getBookmarkRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/bookmark/${parsedReq.id}`, {
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
        queryFn: () => getBookmark({ id }),
    });
};

const updateBookmarkTitleRequestSchema = z.object({
    id: z.string().uuid(),
    title: z.string().trim(),
});

const updateBookmarkTitleResponseSchema = serverResponseSchema.extend({
    bookmark: bookmarkSchema.optional(),
});

type UpdateBookmarkTitleRequest = z.infer<
    typeof updateBookmarkTitleRequestSchema
>;

const updateBookmarkTitle = async (req: UpdateBookmarkTitleRequest) => {
    try {
        const parsedReq = updateBookmarkTitleRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/bookmark/${parsedReq.id}`, {
            method: "PATCH",
            body: JSON.stringify({ title: parsedReq.title }),
        });

        const parsedRes = updateBookmarkTitleResponseSchema.parse(
            await res.json(),
        );

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useUpdateBookmarkTitle = () => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: updateBookmarkTitle,
        onSuccess: (_, req) => {
            queryClient.invalidateQueries({ queryKey: ["bookmarks"] });
            queryClient.invalidateQueries({ queryKey: ["bookmark", req.id] });
        },
    });
};

const deleteBookmarkRequestSchema = z.object({
    id: z.string().uuid(),
});

const deleteBookmarkResponseSchema = serverResponseSchema.extend({
    bookmark: bookmarkSchema.optional(),
});

type DeleteBookmarkRequest = z.infer<typeof deleteBookmarkRequestSchema>;

const deleteBookmark = async (req: DeleteBookmarkRequest) => {
    try {
        const parsedReq = deleteBookmarkRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/bookmark/${parsedReq.id}`, {
            method: "DELETE",
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
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: deleteBookmark,
        onSuccess: (_, req) => {
            queryClient.invalidateQueries({ queryKey: ["bookmarks"] });
            queryClient.invalidateQueries({ queryKey: ["bookmark", req.id] });
        },
    });
};

const summariseBookmarkRequestSchema = z.object({
    id: z.string().uuid(),
});

type SummariseBookmarkRequest = z.infer<typeof summariseBookmarkRequestSchema>;

const summariseBookmark = async (req: SummariseBookmarkRequest) => {
    try {
        const parsedReq = summariseBookmarkRequestSchema.parse(req);

        const res = await fetch(
            `${BASE_URL}/bookmark/${parsedReq.id}/summarise`,
            { method: "POST" },
        );

        const parsedRes = serverResponseSchema.parse(res);

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useSummariseBookmark = () => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: summariseBookmark,
        onSuccess: (_, req) => {
            queryClient.invalidateQueries({ queryKey: ["bookmarks"] });
            queryClient.invalidateQueries({
                queryKey: ["bookmark", req.id],
            });
        },
    });
};

export {
    useCreateBookmark,
    useDeleteBookmark,
    useGetBookmark,
    useGetBookmarks,
    useSummariseBookmark,
    useUpdateBookmarkTitle,
};
