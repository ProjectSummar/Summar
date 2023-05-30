import { Stack } from "expo-router";

function Layout() {
    return (
        <Stack
            initialRouteName="signup"
            screenOptions={{ headerShown: false }}
        />
    );
}

export default Layout;
