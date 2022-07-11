import React from "react";
import { Stack } from "@chakra-ui/react";
import { CommentBoard } from "../../../../components/organisms/comment/CommentBoard";
import { Comments } from "../../../../models/comment";

export const CommentBoardList: React.FC<Comments> = (props) => {
    return (
        <Stack w="70%" m="50px auto" spacing={6}>
            {props.comments.map((comment) => (
                <CommentBoard
                    key={comment.commentKey}
                    threadKey={comment.threadKey}
                    commentKey={comment.commentKey}
                    contributer={comment.contributer}
                    comment={comment.comment}
                    updateDate={comment.updateDate}
                ></CommentBoard>
            ))}
        </Stack>
    );
};
