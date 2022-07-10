import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { Text, Box, Flex, Spacer, Heading, useDisclosure } from "@chakra-ui/react";
import { Thread } from "../../../models/Thread";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { ThreadViewButton } from "../../atoms/button/ThreadViewButton";
import { GeneralModal } from "../modal/GeneralModal";
import { deleteThreadPayload, editThreadPayload } from "../../../pages/threads/redux/threads/type";
import { deleteThread, editThreadTitle } from "../../../state/threads";

interface ThreadBoadProps extends Thread {
    isStatic?: boolean;
}

export const ThreadBoard: React.FC<ThreadBoadProps> = (props) => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [isEdit, setIsEdit] = useState(true);

    const updateButtonOnClick = (title: string) => {
        const editThreadPayload: editThreadPayload = {
            threadKey: props.threadKey,
            title,
            contributer: props.contributer,
        };
        dispatch(editThreadTitle(editThreadPayload));
        onClose();
    };

    const deleteButtonOnClick = () => {
        const deleteThreadPayload: deleteThreadPayload = {
            threadKey: props.threadKey,
            contributer: props.contributer,
        };
        dispatch(deleteThread(deleteThreadPayload));
        onClose();
    };

    return (
        <>
            <Box p="20px" border="1px" bg="red.100" borderRadius="lg" boxShadow="dark-lg">
                <Flex>
                    <Heading>{props.title}</Heading>
                    <Spacer></Spacer>
                    {props.isStatic || (
                        <ThreadViewButton onClick={() => navigate(`thread/${props.threadKey}`)}>Look!</ThreadViewButton>
                    )}
                    <MenuIconButton onOpen={onOpen} setIsEdit={setIsEdit}></MenuIconButton>
                </Flex>
                <Text textAlign="right">投稿者: {props.contributer}</Text>
                <Text textAlign="right">投稿日: {props.postDate}</Text>
                <Flex>
                    <Text>閲覧数: {props.views}人</Text>
                    <Spacer></Spacer>
                    <Text w="500px">コメント数: {props.sumComment}人</Text>
                    <Spacer></Spacer>
                    <Text>更新日: {props.updateDate}</Text>
                </Flex>
            </Box>
            <GeneralModal
                content={props.title}
                isEdit={isEdit}
                isOpen={isOpen}
                onClose={onClose}
                updateOnClick={updateButtonOnClick}
                deleteOnClick={deleteButtonOnClick}
            ></GeneralModal>
        </>
    );
};
