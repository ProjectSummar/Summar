import { Ionicons } from "@expo/vector-icons";
import {
    Menu,
    MenuOption,
    MenuOptions,
    MenuTrigger,
} from "react-native-popup-menu";
import { Text } from "react-native";
import { Bookmark } from "@src/types";
import { useDeleteBookmark } from "@src/api/bookmark";
import { useRouter } from "expo-router";

const BookmarkContextMenu = ({ bookmark }: { bookmark: Bookmark }) => {
    const { mutate: deleteBookmark } = useDeleteBookmark();

    const deleteBookmarkOnSelect = () => {
        console.log("deleting bookmark", bookmark.title);
        deleteBookmark({ id: bookmark.id });
    };

    const router = useRouter();

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
                    onSelect={deleteBookmarkOnSelect}
                    icon="trash"
                />
                <ContextMenuOption
                    text="Update Bookmark"
                    onSelect={() =>
                        router.push({
                            pathname: "/main/bookmark/update",
                            params: { id: bookmark.id },
                        })}
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
