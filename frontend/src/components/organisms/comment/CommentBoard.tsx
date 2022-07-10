import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { Box, Stack, Flex, Divider, Text, Spacer, useDisclosure } from "@chakra-ui/react";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { GoogButton } from "../../atoms/button/GoogButton";
import { Picture } from "../../atoms/picture/Picture";
import { Comment } from "../../../models/Comment";
import { GeneralModal } from "../modal/GeneralModal";
import { deleteCommentPayload, editCommentPayload } from "../../../pages/threadContent/redux/comments/type";
import { deleteComment, editComment } from "../../../state/comments";

export const CommentBoard: React.FC<Comment> = (props) => {
    const dispatch = useDispatch();
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [isEdit, setIsEdit] = useState(true);

    const updateButtonOnClick = (comment: string) => {
        const editCommentPayload: editCommentPayload = {
            commentKey: props.commentKey,
            comment,
            contributer: props.contributer,
        };
        dispatch(editComment(editCommentPayload));
        onClose();
    };

    const deleteButtonOnClick = () => {
        const deleteCommentPayload: deleteCommentPayload = {
            commentKey: props.commentKey,
            contributer: props.contributer,
        };
        dispatch(deleteComment(deleteCommentPayload));
        onClose();
    };

    return (
        <>
            <Box p="20px" border="1px" bg="blue.100" borderRadius="lg" boxShadow="dark-lg">
                <Stack ml="15px" spacing={3}>
                    <Flex>
                        <Text m="auto">投稿者: {props.contributer}</Text>
                        <Spacer></Spacer>
                        <MenuIconButton onOpen={onOpen} setIsEdit={setIsEdit}></MenuIconButton>
                    </Flex>
                    <Divider></Divider>
                    <Text>{props.comment}</Text>
                    <Picture url={"https://bit.ly/dan-abramov"}></Picture>
                    <Flex>
                        <GoogButton></GoogButton>
                        <Spacer></Spacer>
                        <Text m="auto">更新日時: {props.updateDate}</Text>
                    </Flex>
                </Stack>
            </Box>
            <GeneralModal
                content={props.comment}
                isEdit={isEdit}
                isOpen={isOpen}
                onClose={onClose}
                updateOnClick={updateButtonOnClick}
                deleteOnClick={deleteButtonOnClick}
            ></GeneralModal>
        </>
    );
};
