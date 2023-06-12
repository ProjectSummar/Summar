import { Ionicons } from "@expo/vector-icons";
import { Stack, useNavigation } from "expo-router";

const Layout = () => {
    const drawer = useNavigation("/main") as any;

    return (
        <Stack
            initialRouteName="index"
            screenOptions={{
                headerLeft: () => (
                    <Ionicons
                        name="menu-sharp"
                        size={30}
                        onPress={drawer.toggleDrawer}
                    />
                ),
            }}
        >
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
