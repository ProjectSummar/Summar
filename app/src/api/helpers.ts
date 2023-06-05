import { z } from "zod";
import Constants from "expo-constants";

const BASE_URL = Constants.expoConfig?.extra?.baseUrl;

const serverResponseSchema = z.object({
    ok: z.boolean(),
    msg: z.string(),
});

export { BASE_URL, serverResponseSchema };
