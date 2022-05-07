import React from "react";
import {
    Text,
    Box,
    Flex,
    Spacer,
    Heading,
} from "@chakra-ui/react";
import { Thread } from "../../../models/Thread";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";

export const ThreadBoard: React.FC<Thread> = (props) => {
    const { title, contributer, postDate, updateDate, views, sumComment } =
        props;

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
                    <Heading>{title}</Heading>
                    <Spacer></Spacer>
                    <MenuIconButton onOpen={() => undefined}></MenuIconButton>
                </Flex>
                <Text textAlign="right">投稿者: {contributer}</Text>
                <Text textAlign="right">投稿日: {postDate}</Text>
                <Flex>
                    <Text>閲覧数: {views}人</Text>
                    <Spacer></Spacer>
                    <Text w="600px">コメント数: {sumComment}人</Text>
                    <Spacer></Spacer>
                    <Text>更新日: {updateDate}</Text>
                </Flex>
            </Box>
        </>
    );
};
