import { Drawer } from "expo-router/drawer";

const Layout = () => {
    return (
        <Drawer initialRouteName="bookmark">
            <Drawer.Screen
                name="bookmark"
                options={{ title: "Bookmarks", headerShown: false }}
            />
            <Drawer.Screen name="settings" options={{ title: "Settings" }} />
        </Drawer>
    );
};

export default Layout;
