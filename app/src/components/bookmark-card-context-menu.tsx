import { Ionicons } from "@expo/vector-icons";
import {
    Menu,
    MenuOption,
    MenuOptions,
    MenuTrigger,
} from "react-native-popup-menu";
import { Share, Text } from "react-native";
import { Bookmark, IconName } from "@src/types";
import { useDeleteBookmark } from "@src/api/bookmark";
import { useRouter } from "expo-router";
import { useToast } from "@src/contexts/toast-context";

const BookmarkCardContextMenu = ({ bookmark }: { bookmark: Bookmark }) => {
    const { errorToast, successToast } = useToast();

    const router = useRouter();

    const { mutate: deleteBookmark } = useDeleteBookmark();

    const deleteBookmarkOnSelect = () => {
        console.log("deleting bookmark", bookmark.title);
        deleteBookmark(
            { id: bookmark.id },
            {
                onSuccess: () => successToast("Bookmark deleted successfully"),
                onError: () => errorToast("Error deleting bookmark"),
            }
        );
    };

    const shareBookmark = async () => {
        return await Share.share({ url: bookmark.url });
    };

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
                        })
                    }
                    icon="md-pencil-sharp"
                />
                <ContextMenuOption
                    text="Share Bookmark"
                    onSelect={shareBookmark}
                    icon="share-outline"
                />
            </MenuOptions>
        </Menu>
    );
};

const ContextMenuOption = ({
    text,
    onSelect,
    icon,
}: {
    text: string;
    onSelect: () => any;
    icon: IconName;
}) => {
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

export default BookmarkCardContextMenu;
