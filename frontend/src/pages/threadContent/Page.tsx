import React from "react";
import { Box } from "@chakra-ui/react";
import { ThreadBoard } from "../../components/organisms/thread/ThreadBoard";
import { CommentPostform } from "../../components/organisms/form/CommentPostForm";
import { CommentBoardList } from "../../components/templates/comments/CommentBoardList";
import { Thread } from "../../models/Thread";

export const ThreadContent: React.FC = () => {
    const damyThread: Thread = {
        threadKey: "1",
        title: "test",
        contributer: "motohashi",
        postDate: "2022/1/1 12:00",
        updateDate: "2022/1/1 13:00",
        views: 10,
        sumComment: 20,
    };

    return (
        <>
            <Box w="70%" m="50px auto">
                <ThreadBoard
                    threadKey={damyThread.threadKey}
                    isStatic
                    title={damyThread.title}
                    contributer={damyThread.contributer}
                    postDate={damyThread.postDate}
                    updateDate={damyThread.updateDate}
                    views={damyThread.views}
                    sumComment={damyThread.sumComment}
                ></ThreadBoard>
            </Box>
            <CommentPostform></CommentPostform>
            <CommentBoardList></CommentBoardList>
        </>
    );
};
