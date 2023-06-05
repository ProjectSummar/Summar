import { AuthContextProvider } from "@src/contexts/auth-context";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Stack } from "expo-router";
import { SafeAreaProvider } from "react-native-safe-area-context";

const AUTH_CONTEXT = false;

const Layout = () => {
    const queryClient = new QueryClient();

    if (AUTH_CONTEXT) {
        return (
            <SafeAreaProvider>
                <QueryClientProvider client={queryClient}>
                    <AuthContextProvider>
                        <Stack screenOptions={{ headerShown: false }} />
                    </AuthContextProvider>
                </QueryClientProvider>
            </SafeAreaProvider>
        );
    } else {
        return (
            <SafeAreaProvider>
                <QueryClientProvider client={queryClient}>
                    <Stack screenOptions={{ headerShown: false }} />
                </QueryClientProvider>
            </SafeAreaProvider>
        );
    }
};

export default Layout;
