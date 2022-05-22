import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { useParams } from "react-router-dom";
import { Box } from "@chakra-ui/react";
import { ThreadBoard } from "../../components/organisms/thread/ThreadBoard";
import { CommentPostform } from "../../components/organisms/form/CommentPostForm";
import { CommentBoardList } from "../../components/templates/comments/CommentBoardList";
import { RootState } from "../../store/store";
import { getAllCommentByThreadKey } from "./redux/comments";
import { getThreadByThreadKeyPayload } from "./redux/thread/type";
import { getThreadByThreadKey } from "./redux/thread";
import { getAllCommentPayload } from "./redux/comments/type";

export const ThreadContent: React.FC = () => {
    const urlParams = useParams();

    const user = useSelector((state: RootState) => state.user);
    const thread = useSelector((state: RootState) => state.thread);
    const threads = useSelector((state: RootState) => state.threads.threads);
    const comments = useSelector((state: RootState) => state.comments.comments);
    const dispatch = useDispatch();

    const getAllCommentPayload: getAllCommentPayload = {
        threadKey: urlParams.threadKey as string,
    };

    const getThreadByThreadKeyPayload: getThreadByThreadKeyPayload = {
        threads,
        threadKey: urlParams.threadKey as string,
    };

    useEffect(() => {
        dispatch(getThreadByThreadKey(getThreadByThreadKeyPayload));
    }, [threads]);

    useEffect(() => {
        dispatch(getAllCommentByThreadKey(getAllCommentPayload));
    }, []);

    return (
        <>
            <Box w="70%" m="50px auto">
                <ThreadBoard
                    threadKey={thread.threadKey}
                    isStatic
                    title={thread.title}
                    contributer={thread.contributer}
                    postDate={thread.postDate}
                    updateDate={thread.updateDate}
                    views={thread.views}
                    sumComment={thread.sumComment}
                ></ThreadBoard>
            </Box>
            {user.username === "" || (
                <CommentPostform
                    loginUsername={user.username}
                    threadKey={urlParams.threadKey as string}
                ></CommentPostform>
            )}
            <CommentBoardList comments={comments}></CommentBoardList>
        </>
    );
};
