import { useMutation, useQuery } from "@tanstack/react-query";
import { z } from "zod";
import { bookmarkSchema, userSchema } from "./types";

const BASE_URL = process.env.BASE_URL;

const serverResponseSchema = z.object({
    ok: z.boolean(),
    msg: z.string(),
});

type LoginRequest = {
    email: string;
    password: string;
};

const login = async (req: LoginRequest) => {
    const raw = await fetch(`${BASE_URL}/login`, {
        method: "POST",
        body: JSON.stringify(req),
    });

    const res = serverResponseSchema.parse(await raw.json());

    if (!res.ok) {
        throw new Error(res.msg);
    } else {
        return res;
    }
};

const useLogin = () => {
    return useMutation({
        mutationFn: login,
    });
};

type SignupRequest = {
    email: string;
    password: string;
};

const signupResponseSchema = serverResponseSchema.extend({
    user: userSchema,
});

const signup = async (req: SignupRequest) => {
    const raw = await fetch(`${BASE_URL}/signup`, {
        method: "POST",
        body: JSON.stringify(req),
    });

    const res = signupResponseSchema.parse(await raw.json());

    if (!res.ok) {
        throw new Error(res.msg);
    } else {
        return res;
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
    const raw = await fetch(`${BASE_URL}/me`, {
        method: "GET",
    });

    const res = getUserSchema.parse(await raw.json());

    if (!res.ok) {
        throw new Error(res.msg);
    } else {
        return res;
    }
};

const useGetUser = () => {
    return useQuery({
        queryKey: ["user"],
        queryFn: getUser,
    });
};

export { useLogin, useSignup, useGetUser };
