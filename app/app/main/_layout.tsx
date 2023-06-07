import { Drawer } from "expo-router/drawer";

const Layout = () => {
    return (
        <Drawer initialRouteName="bookmarks">
            <Drawer.Screen
                name="bookmarks"
                options={{ title: "Bookmarks", headerShown: false }}
            />
            <Drawer.Screen
                name="settings"
                options={{ title: "Settings" }}
            />
        </Drawer>
    );
};

export default Layout;
