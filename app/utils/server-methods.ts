import { useMutation, useQuery } from "@tanstack/react-query";
import { z } from "zod";
import { bookmarkSchema, userSchema } from "./types";
import Constants from "expo-constants";

const BASE_URL = Constants.expoConfig?.extra?.baseUrl;

const serverResponseSchema = z.object({
    ok: z.boolean(),
    msg: z.string(),
});

const loginRequestSchema = z.object({
    email: z.string().email(),
    password: z.string(),
});

type LoginRequest = z.infer<typeof loginRequestSchema>;

const login = async (req: LoginRequest) => {
    try {
        const parsedReq = loginRequestSchema.parse(req);

        const res = await fetch(`${BASE_URL}/login`, {
            method: "POST",
            body: JSON.stringify(parsedReq),
        });

        const parsedRes = serverResponseSchema.parse(await res.json());

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
    user: userSchema,
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

const getUserSchema = serverResponseSchema.extend({
    user: userSchema,
    bookmarks: z.array(bookmarkSchema),
});

const getUser = async () => {
    try {
        const res = await fetch(`${BASE_URL}/me`, {
            method: "GET",
        });

        const parsedRes = getUserSchema.parse(await res.json());

        if (!parsedRes.ok) {
            throw new Error(parsedRes.msg);
        } else {
            return parsedRes;
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
