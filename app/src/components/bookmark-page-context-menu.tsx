import { Share } from "react-native";
import { Bookmark } from "@src/types";
import { useToast } from "@src/contexts/toast-context";
import { Dispatch, SetStateAction } from "react";
import { useSummariseBookmark } from "@src/api/bookmark";
import { ContextMenu, ContextMenuOption } from "./context-menu";

const BookmarkPageContextMenu = ({
    bookmark,
    setSummaryView,
}: {
    bookmark: Bookmark;
    setSummaryView: Dispatch<SetStateAction<boolean>>;
}) => {
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
        <ContextMenu>
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
        </ContextMenu>
    );
};

export default BookmarkPageContextMenu;
