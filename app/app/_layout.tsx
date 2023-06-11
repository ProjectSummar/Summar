import { AuthContextProvider } from "@src/contexts/auth-context";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Slot } from "expo-router";
import { useEffect, useState } from "react";
import { SafeAreaProvider } from "react-native-safe-area-context";
import * as Font from "expo-font";
import * as SplashScreen from "expo-splash-screen";
import Ionicons from "@expo/vector-icons/Ionicons";
import { MenuProvider } from "react-native-popup-menu";
import { ToastProvider } from "@src/contexts/toast-context";

const AUTH_CONTEXT = true;

const Layout = () => {
    const [ready, setReady] = useState(false);

    const queryClient = new QueryClient();

    useEffect(() => {
        (async () => {
            try {
                SplashScreen.preventAutoHideAsync();

                const fonts = cacheFonts([Ionicons.font]);

                await Promise.all(fonts);
            } catch (err) {
                console.warn(err);
            } finally {
                setReady(true);
                SplashScreen.hideAsync();
            }
        })();
    }, []);

    if (!ready) {
        return null;
    }

    if (AUTH_CONTEXT) {
        return (
            <SafeAreaProvider>
                <QueryClientProvider client={queryClient}>
                    <AuthContextProvider>
                        <MenuProvider>
                            <ToastProvider>
                                <Slot />
                            </ToastProvider>
                        </MenuProvider>
                    </AuthContextProvider>
                </QueryClientProvider>
            </SafeAreaProvider>
        );
    } else {
        return (
            <SafeAreaProvider>
                <QueryClientProvider client={queryClient}>
                    <MenuProvider>
                        <ToastProvider>
                            <Slot />
                        </ToastProvider>
                    </MenuProvider>
                </QueryClientProvider>
            </SafeAreaProvider>
        );
    }
};

const cacheFonts = (fonts: any[]) => {
    return fonts.map((font) => Font.loadAsync(font));
};

export default Layout;
