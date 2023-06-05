import { useMutation, useQuery } from "@tanstack/react-query";
import { z } from "zod";
import { bookmarkSchema, userSchema } from "@src/types";
import { BASE_URL, serverResponseSchema } from "@src/api/helpers";

const loginRequestSchema = z.object({
    email: z.string().email(),
    password: z.string(),
});

const loginResponseSchema = serverResponseSchema.extend({
    user: userSchema.optional(),
});

type LoginRequest = z.infer<typeof loginRequestSchema>;

const login = async (req: LoginRequest) => {
    try {
        const parsedReq = loginRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/login`, {
            method: "POST",
            body: JSON.stringify(parsedReq),
        });

        const parsedRes = loginResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useLogin = () => {
    return useMutation({
        mutationFn: login,
    });
};

const signupRequestSchema = z.object({
    email: z.string().email(),
    password: z.string(),
});

type SignupRequest = z.infer<typeof signupRequestSchema>;

const signupResponseSchema = serverResponseSchema.extend({
    user: userSchema.optional(),
});

const signup = async (req: SignupRequest) => {
    try {
        const parsedReq = signupRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/signup`, {
            method: "POST",
            body: JSON.stringify(parsedReq),
        });

        const parsedRes = signupResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
        }
    } catch (err) {
        throw err as Error;
    }
};

const useSignup = () => {
    return useMutation({
        mutationFn: signup,
    });
};

const getUserResponseSchema = serverResponseSchema.extend({
    user: userSchema.optional(),
    bookmarks: z.array(bookmarkSchema).optional(),
});

const getUser = async () => {
    try {
        const res = await fetch(`${BASE_URL}/me`, {
            method: "GET",
        });

        const parsedRes = getUserResponseSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return {
                user: parsedRes.user,
                bookmarks: parsedRes.bookmarks,
            };
        }
    } catch (err) {
        throw err as Error;
    }
};

const useGetUser = () => {
    return useQuery({
        queryKey: ["user"],
        queryFn: getUser,
    });
};

export { useLogin, useSignup, useGetUser };
