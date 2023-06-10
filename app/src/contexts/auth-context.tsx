import { useGetUser } from "@src/api/auth";
import { useRouter, useSegments } from "expo-router";
import { ReactNode, useEffect } from "react";

const AuthContextProvider = ({ children }: { children: ReactNode }) => {
    const { data: user, isLoading } = useGetUser();

    const segments = useSegments();
    const router = useRouter();

    useEffect(() => {
        const inAuthRoute = segments[0] === "auth";
        const isLoggedIn = !isLoading && user !== undefined;

        console.log("auth context", {
            inAuthRoute,
            segments,
            user,
            isLoggedIn,
        });

        if (!inAuthRoute && !isLoggedIn) {
            // redirect to login page
            router.replace("/auth/login");
        } else if (inAuthRoute && isLoggedIn) {
            // redirect to main page
            router.replace("/main");
        }
    }, [isLoading, segments, user]);

    return <>{children}</>;
};

export { AuthContextProvider };
