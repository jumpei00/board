import React from "react";
import { Stack } from "@chakra-ui/react";
import { CommentBoard } from "../../../../components/organisms/comment/CommentBoard";
import { Comment } from "../../../../models/Comment";

type CommentListProps = {
    threadKey: string;
    comments: Array<Comment>;
}

export const CommentBoardList: React.FC<CommentListProps> = (props) => {
    return (
        <Stack w="70%" m="50px auto" spacing={6}>
            {props.comments.map((comment) => (
                <CommentBoard
                    key={comment.commentKey}
                    threadKey={props.threadKey}
                    comment={comment}
                ></CommentBoard>
            ))}
        </Stack>
    );
};
