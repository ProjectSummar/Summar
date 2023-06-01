import { Stack } from "expo-router";
import { StatusBar } from "expo-status-bar";

const Layout = () => {
    return (
        <>
            <StatusBar style="dark" />
            <Stack
                initialRouteName="login"
                screenOptions={{ headerShown: false }}
            />
        </>
    );
};

export default Layout;
