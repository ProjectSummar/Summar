import { Stack } from "expo-router";

const Layout = () => {
    return (
        <Stack
            initialRouteName="signup"
            screenOptions={{ headerShown: false }}
        />
    );
};

export default Layout;
