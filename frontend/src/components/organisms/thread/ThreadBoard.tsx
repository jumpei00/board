import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import { Text, Box, Flex, Spacer, Heading, useDisclosure } from "@chakra-ui/react";
import { Thread } from "../../../models/Thread";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { ThreadViewButton } from "../../atoms/button/ThreadViewButton";
import { GeneralModal } from "../modal/GeneralModal";
import { UpdateThreadPayload, threadSagaActions } from "../../../state/threads/modules";
import { RootState } from "../../../store/store";

interface ThreadBoadProps {
    isStatic?: boolean;
    thread: Thread;
}

export const ThreadBoard: React.FC<ThreadBoadProps> = (props) => {
    const userState = useSelector((state: RootState) => state.user.userState);
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [isEdit, setIsEdit] = useState(true);

    const updateButtonOnClick = (title: string) => {
        const updateThreadPayload: UpdateThreadPayload = {
            threadKey: props.thread.threadKey,
            body: {
                title,
            },
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
                        <ThreadViewButton onClick={() => navigate(`thread/${props.thread.threadKey}`)}>
                            Look!
                        </ThreadViewButton>
                    )}
                    {props.isStatic || userState.username === props.thread.contributor && (
                        <MenuIconButton onOpen={onOpen} setIsEdit={setIsEdit}></MenuIconButton>
                    )}
                </Flex>
                <Text textAlign="right">?????????: {props.thread.contributor}</Text>
                <Text textAlign="right">?????????: {props.thread.createDate}</Text>
                <Flex>
                    <Text>?????????: {props.thread.views}???</Text>
                    <Spacer></Spacer>
                    <Text w="500px">???????????????: {props.thread.commentSum}???</Text>
                    <Spacer></Spacer>
                    <Text>?????????: {props.thread.updateDate}</Text>
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
