import { Ionicons } from "@expo/vector-icons";
import { Drawer } from "expo-router/drawer";

const Layout = () => {
    return (
        <Drawer
            initialRouteName="bookmark"
            detachInactiveScreens
            screenOptions={({ navigation }) => ({
                headerLeft: () => (
                    <Ionicons
                        style={{ marginLeft: 20 }}
                        name="menu-sharp"
                        size={30}
                        onPress={navigation.toggleDrawer}
                    />
                ),
            })}
        >
            <Drawer.Screen
                name="bookmark"
                options={{ title: "Bookmarks", headerShown: false }}
            />
            <Drawer.Screen name="settings" options={{ title: "Settings" }} />
        </Drawer>
    );
};

export default Layout;
