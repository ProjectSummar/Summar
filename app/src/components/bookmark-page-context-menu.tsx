import { Ionicons } from "@expo/vector-icons";
import {
    Menu,
    MenuOption,
    MenuOptions,
    MenuTrigger,
} from "react-native-popup-menu";
import { Share, Text } from "react-native";
import { Bookmark, IconName } from "@src/types";
import { useToast } from "@src/contexts/toast-context";
import { Dispatch, SetStateAction } from "react";
import { useSummariseBookmark } from "@src/api/bookmark";

const BookmarkPageContextMenu = (
    { bookmark, setSummaryView }: {
        bookmark: Bookmark;
        setSummaryView: Dispatch<SetStateAction<boolean>>;
    },
) => {
    const { errorToast, successToast } = useToast();

    const { mutate: summariseBookmark } = useSummariseBookmark();

    const summariseBookmarkOnPress = () => {
        if (bookmark.summary.length !== 0) {
            return;
        }

        summariseBookmark(
            { id: bookmark.id },
            {
                onSuccess: () =>
                    successToast("Summarised bookmark successfully"),
                onError: () => errorToast("Error summarising bookmark"),
            },
        );
    };

    const toggleSummaryView = () => {
        if (bookmark.summary.length === 0) {
            errorToast("No summary available");
            return;
        }
        setSummaryView((summaryView) => !summaryView);
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
                    text="Summarise Bookmark"
                    onSelect={summariseBookmarkOnPress}
                    icon="flash-outline"
                />
                <ContextMenuOption
                    text="Toggle Summary View"
                    onSelect={toggleSummaryView}
                    icon="book-outline"
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

export default BookmarkPageContextMenu;
