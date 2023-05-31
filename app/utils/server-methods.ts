import { useMutation } from "@tanstack/react-query";
import { z } from "zod";
import { userSchema } from "./types";

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
    const res = await fetch(`${BASE_URL}/login`, {
        method: "POST",
        body: JSON.stringify(req),
    });

    return serverResponseSchema.parse(await res.json());
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
    const res = await fetch(`${BASE_URL}/signup`, {
        method: "POST",
        body: JSON.stringify(req),
    });

    return signupResponseSchema.parse(await res.json());
};

const useSignup = () => {
    return useMutation({
        mutationFn: signup,
    });
};

export { useLogin, useSignup };
