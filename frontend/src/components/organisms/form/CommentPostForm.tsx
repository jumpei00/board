import React, { ChangeEvent, useState } from "react";
import { useDispatch } from "react-redux";
import { Box, Text, Input, Flex, Spacer } from "@chakra-ui/react";
import { ImageButton } from "../../atoms/button/ImageButton";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { createCommentPayload } from "../../../pages/threadContent/redux/comments/type";
import { createComment } from "../../../pages/threadContent/redux/comments";

type CommentPostFormProps = {
    loginUsername: string;
    threadKey: string;
}

export const CommentPostform: React.FC<CommentPostFormProps> = (props) => {
    const dispatch = useDispatch();
    const [comment, setComment] = useState("");
    const [buttonDisable, setButtonDisable] = useState(true);

    const createCommentPayload: createCommentPayload = {
        threadKey: props.threadKey,
        comment,
        contributer: props.loginUsername,
    };

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        const currentValue = event.target.value;
        currentValue === "" ? setButtonDisable(true) : setButtonDisable(false);
        setComment(currentValue);
    };

    const postCommentOnClick = () => {
        dispatch(createComment(createCommentPayload));
        setComment("");
    };

    return (
        <Box w="70%" m="auto">
            <Text mb="10px">コメント</Text>
            <Input
                mb="20px"
                variant="flushed"
                placeholder="コメントしよう"
                value={comment}
                onChange={handleChange}
            ></Input>
            <Flex>
                <ImageButton>画像</ImageButton>
                <Spacer></Spacer>
                <PrimaryButton colorScheme="teal" isDisabled={buttonDisable} onClick={postCommentOnClick}>
                    投稿
                </PrimaryButton>
            </Flex>
        </Box>
    );
};
