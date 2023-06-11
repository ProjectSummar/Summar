import { Stack } from "expo-router";

const Layout = () => {
    return (
        <Stack initialRouteName="index">
            <Stack.Screen
                name="create"
                options={{
                    title: "Create Bookmark",
                    presentation: "modal",
                }}
            />
            <Stack.Screen
                name="update"
                options={{
                    title: "Update Bookmark Title",
                    presentation: "modal",
                }}
            />
        </Stack>
    );
};

export default Layout;
