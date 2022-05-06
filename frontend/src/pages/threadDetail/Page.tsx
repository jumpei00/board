import React from "react";
import { Box } from "@chakra-ui/react";
import { ThreadBoard } from "../../components/organisms/thread/ThreadBoard";
import { CommentPostform } from "../../components/organisms/form/CommentPostForm";
import { Thread } from "../../models/Thread";
import { CommentBoardList } from "../../components/templates/comments/CommentBoardList";

export const ThreadDetail: React.FC<Thread> = (props) => {
    const {
        hashID,
        title,
        contributer,
        postDate,
        updateDate,
        views,
        sumComment,
    } = props;

    return (
        <>
            <Box w="70%" m="50px auto">
                <ThreadBoard
                    hashID={hashID}
                    title={title}
                    contributer={contributer}
                    postDate={postDate}
                    updateDate={updateDate}
                    views={views}
                    sumComment={sumComment}
                ></ThreadBoard>
            </Box>
            <CommentPostform></CommentPostform>
            <CommentBoardList></CommentBoardList>
        </>
    );
};
