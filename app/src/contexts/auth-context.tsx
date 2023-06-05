import { useGetUser } from "@src/api/auth";
import { useRouter, useSegments } from "expo-router";
import { ReactNode, useEffect } from "react";

const AuthContextProvider = ({ children }: { children: ReactNode }) => {
    const { data, isLoading } = useGetUser();

    const user = data?.user;

    const segments = useSegments();
    const router = useRouter();

    // TODO: if loading, show splash screen? (but need to only show on startup, not between screen transitions)

    useEffect(() => {
        const inAuthRoute = segments[0] === "auth";
        const isLoggedIn = !isLoading && user;

        console.log("auth context", {
            inAuthRoute,
            segments,
            user,
        });

        if (!inAuthRoute && !isLoggedIn) {
            // redirect to login page
            router.replace("/auth/login");
        } else if (inAuthRoute && isLoggedIn) {
            // redirect to bookmarks page
            router.replace("/bookmarks");
        }
    }, [isLoading, segments, user]);

    return <>{children}</>;
};

export { AuthContextProvider };
