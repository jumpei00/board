import React from "react";
import { Stack } from "@chakra-ui/react";
import { CommentBoard } from "../../organisms/comment/CommentBoard";

export const CommentBoardList: React.FC = () => {
    return (
        <Stack w="70%" m="50px auto" spacing={6}>
            <CommentBoard></CommentBoard>
        </Stack>
    );
};
