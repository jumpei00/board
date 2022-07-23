import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Box, Stack, Flex, Divider, Text, Spacer, useDisclosure } from "@chakra-ui/react";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { Comment } from "../../../models/Comment";
import { GeneralModal } from "../modal/GeneralModal";
import { commentSagaActions } from "../../../state/comments/modules";
import { RootState } from "../../../store/store";

type CommentBoardProps = {
    threadKey: string;
    comment: Comment;
};

export const CommentBoard: React.FC<CommentBoardProps> = (props) => {
    const userState = useSelector((state: RootState) => state.user.userState)
    const dispatch = useDispatch();
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [isEdit, setIsEdit] = useState(true);

    const updateButtonOnClick = (comment: string) => {
        dispatch(
            commentSagaActions.update({
                threadKey: props.threadKey,
                commentKey: props.comment.commentKey,
                body: {
                    comment,
                },
            })
        );
        onClose();
    };

    const deleteButtonOnClick = () => {
        dispatch(
            commentSagaActions.delete({
                threadKey: props.threadKey,
                commentKey: props.comment.commentKey,
            })
        );
        onClose();
    };

    return (
        <>
            <Box p="20px" border="1px" bg="blue.100" borderRadius="lg" boxShadow="dark-lg">
                <Stack ml="15px" spacing={3}>
                    <Flex>
                        <Text m="auto">投稿者: {props.comment.contributor}</Text>
                        <Spacer></Spacer>
                        {userState.username === props.comment.contributor && (
                            <MenuIconButton onOpen={onOpen} setIsEdit={setIsEdit}></MenuIconButton>
                        )}
                    </Flex>
                    <Divider></Divider>
                    <Text>{props.comment.comment}</Text>
                    {/* <Picture url={"https://bit.ly/dan-abramov"}></Picture> */}
                    <Flex>
                        {/* <GoogButton></GoogButton> */}
                        <Spacer></Spacer>
                        <Text m="auto">更新日時: {props.comment.updateDate}</Text>
                    </Flex>
                </Stack>
            </Box>
            <GeneralModal
                content={props.comment.comment}
                isEdit={isEdit}
                isOpen={isOpen}
                onClose={onClose}
                updateOnClick={updateButtonOnClick}
                deleteOnClick={deleteButtonOnClick}
            ></GeneralModal>
        </>
    );
};
