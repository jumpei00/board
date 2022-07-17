import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { Text, Box, Flex, Spacer, Heading, useDisclosure } from "@chakra-ui/react";
import { Thread } from "../../../models/Thread";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { ThreadViewButton } from "../../atoms/button/ThreadViewButton";
import { GeneralModal } from "../modal/GeneralModal";
import { UpdateThreadPayload, threadSagaActions } from "../../../state/threads/modules";

interface ThreadBoadProps {
    isStatic?: boolean;
    thread: Thread
}

export const ThreadBoard: React.FC<ThreadBoadProps> = (props) => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [isEdit, setIsEdit] = useState(true);

    const updateButtonOnClick = (title: string) => {
        const updateThreadPayload: UpdateThreadPayload = {
            threadKey: props.thread.threadKey,
            body: {
                title,
                contributor: props.thread.contributor,
            }
        };
        dispatch(threadSagaActions.update(updateThreadPayload));
        onClose();
    };

    const deleteButtonOnClick = () => {
        dispatch(threadSagaActions.delete(props.thread.threadKey));
        onClose();
    };

    return (
        <>
            <Box p="20px" border="1px" bg="red.100" borderRadius="lg" boxShadow="dark-lg">
                <Flex>
                    <Heading>{props.thread.title}</Heading>
                    <Spacer></Spacer>
                    {props.isStatic || (
                        <ThreadViewButton onClick={() => navigate(`thread/${props.thread.threadKey}`)}>Look!</ThreadViewButton>
                    )}
                    <MenuIconButton onOpen={onOpen} setIsEdit={setIsEdit}></MenuIconButton>
                </Flex>
                <Text textAlign="right">投稿者: {props.thread.contributor}</Text>
                <Text textAlign="right">投稿日: {props.thread.createDate}</Text>
                <Flex>
                    <Text>閲覧数: {props.thread.views}人</Text>
                    <Spacer></Spacer>
                    <Text w="500px">コメント数: {props.thread.commentSum}人</Text>
                    <Spacer></Spacer>
                    <Text>更新日: {props.thread.updateDate}</Text>
                </Flex>
            </Box>
            <GeneralModal
                content={props.thread.title}
                isEdit={isEdit}
                isOpen={isOpen}
                onClose={onClose}
                updateOnClick={updateButtonOnClick}
                deleteOnClick={deleteButtonOnClick}
            ></GeneralModal>
        </>
    );
};
