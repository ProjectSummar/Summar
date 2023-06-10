import { Ionicons } from "@expo/vector-icons";
import {
    Menu,
    MenuOption,
    MenuOptions,
    MenuTrigger,
} from "react-native-popup-menu";
import { Text } from "react-native";

const BookmarkContextMenu = () => {
    return (
        <Menu>
            <MenuTrigger>
                <Ionicons
                    name="ellipsis-vertical"
                    size={20}
                    style={{ padding: 10 }}
                />
            </MenuTrigger>

            <MenuOptions
                customStyles={{ optionsContainer: { borderRadius: 10 } }}
            >
                <ContextMenuOption
                    text="Delete Bookmark"
                    onSelect={() => console.log("bookmark deleted")}
                    icon="trash"
                />
                <ContextMenuOption
                    text="Update Bookmark"
                    onSelect={() => console.log("bookmark updated")}
                    icon="md-pencil-sharp"
                />
            </MenuOptions>
        </Menu>
    );
};

const ContextMenuOption = (
    { text, onSelect, icon }: {
        text: string;
        onSelect: () => any;
        icon: keyof typeof Ionicons.glyphMap;
    },
) => {
    return (
        <MenuOption
            onSelect={onSelect}
            customStyles={{
                optionWrapper: {
                    flexDirection: "row",
                    alignItems: "center",
                    justifyContent: "space-between",
                    padding: 10,
                },
            }}
        >
            <Text style={{ fontSize: 15 }}>{text}</Text>
            <Ionicons name={icon} size={20} />
        </MenuOption>
    );
};

export default BookmarkContextMenu;
