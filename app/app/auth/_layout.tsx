import { Stack } from "expo-router";

const Layout = () => {
    return (
        <Stack
            initialRouteName="login"
            screenOptions={{ headerShown: false }}
        />
    );
};

export default Layout;
