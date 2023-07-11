import { createDrawerNavigator } from "@react-navigation/drawer";
import { NavigationContainer } from "@react-navigation/native";
import Settings from "@app/bookmarks/settings";

const Drawer = createDrawerNavigator();

export default function DrawerMenu() {
    return (
        <NavigationContainer>
            <Drawer.Navigator initialRouteName="Home">
                <Drawer.Screen name="Settings" component={Settings} />
            </Drawer.Navigator>
        </NavigationContainer>
    );
}
