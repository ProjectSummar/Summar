import { Stack } from "expo-router";

const Layout = () => {
    return (
        <Stack
            initialRouteName="index"
            screenOptions={{
                title: "Bookmarks",
            }}
        />
    );
};

export default Layout;
