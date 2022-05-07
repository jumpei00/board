import React from "react";
import { useNavigate } from "react-router-dom";
import { Text, Box, Flex, Spacer, Heading, HStack } from "@chakra-ui/react";
import { Thread } from "../../../models/Thread";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { ThreadViewButton } from "../../atoms/button/ThreadViewButton";

interface ThreadBoadState extends Thread {
    isStatic?: boolean;
}

export const ThreadBoard: React.FC<ThreadBoadState> = (props) => {
    const navigate = useNavigate();

    return (
        <>
            <Box
                p="20px"
                border="1px"
                bg="red.100"
                borderRadius="lg"
                boxShadow="dark-lg"
            >
                <Flex>
                    <Heading>{props.title}</Heading>
                    <Spacer></Spacer>
                    <HStack>
                        {props.isStatic || (
                            <ThreadViewButton
                                onClick={() =>
                                    navigate(`thread/${props.threadKey}`)
                                }
                            >
                                Look!
                            </ThreadViewButton>
                        )}
                        <MenuIconButton
                            onOpen={() => undefined}
                        ></MenuIconButton>
                    </HStack>
                </Flex>
                <Text textAlign="right">投稿者: {props.contributer}</Text>
                <Text textAlign="right">投稿日: {props.postDate}</Text>
                <Flex>
                    <Text>閲覧数: {props.views}人</Text>
                    <Spacer></Spacer>
                    <Text w="600px">コメント数: {props.sumComment}人</Text>
                    <Spacer></Spacer>
                    <Text>更新日: {props.updateDate}</Text>
                </Flex>
            </Box>
        </>
    );
};
