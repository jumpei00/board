import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { useParams } from "react-router-dom";
import { Box } from "@chakra-ui/react";
import { ThreadBoard } from "../../../components/organisms/thread/ThreadBoard";
import { CommentPostform } from "../../../components/organisms/form/CommentPostForm";
import { CommentBoardList } from "./organisms/CommentBoardList";
import { RootState } from "../../../store/store";
import { commentSagaActions } from "../../../state/comments/modules";
import { userSagaActions } from "../../../state/user/modules";

export const ThreadContent: React.FC = () => {
    const urlParams = useParams();
    const userState = useSelector((state: RootState) => state.user.userState);
    const commentState = useSelector((state: RootState) => state.comment.commentState);
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(userSagaActions.getme());
        if (urlParams.threadKey) {
            dispatch(commentSagaActions.getAll(urlParams.threadKey));
        }
    }, []);

    return (
        <>
            <Box w="70%" m="50px auto">
                <ThreadBoard isStatic thread={commentState.thread}></ThreadBoard>
            </Box>
            {userState.username === "" || <CommentPostform threadKey={urlParams.threadKey}></CommentPostform>}
            <CommentBoardList
                threadKey={commentState.thread.threadKey}
                comments={commentState.comments}
            ></CommentBoardList>
        </>
    );
};
