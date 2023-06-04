import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Stack } from "expo-router";
import { SafeAreaProvider } from "react-native-safe-area-context";

const Layout = () => {
    const queryClient = new QueryClient();

    return (
        <SafeAreaProvider>
            <QueryClientProvider client={queryClient}>
                <Stack screenOptions={{ headerShown: false }} />
            </QueryClientProvider>
        </SafeAreaProvider>
    );
};

export default Layout;
