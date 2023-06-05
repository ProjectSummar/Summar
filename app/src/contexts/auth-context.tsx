import { useGetUser } from "@src/api";
import { useRouter, useSegments } from "expo-router";
import { ReactNode, useEffect } from "react";

const AuthContextProvider = ({ children }: { children: ReactNode }) => {
    const {
        data,
        isLoading: getUserLoading,
        isError: getUserError,
    } = useGetUser();

    const user = data?.user;

    const segments = useSegments();
    const router = useRouter();

    useEffect(() => {
        const inAuthRoute = segments[0] === "auth";
        const isLoggedIn = !getUserLoading && !getUserError && user;

        if (!inAuthRoute && !isLoggedIn) {
            // redirect to login page
            router.replace("/auth/login");
        } else if (inAuthRoute && isLoggedIn) {
            // redirect to bookmarks page
            router.replace("/bookmarks");
        }
    }, [user, getUserLoading, getUserError, segments]);

    return <>{children}</>;
};

export { AuthContextProvider };
