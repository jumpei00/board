import React from "react";
import { Box, Stack, Flex, Divider, Text, Spacer } from "@chakra-ui/react";
import { MenuIconButton } from "../../atoms/button/MenuIconButton";
import { GoogButton } from "../../atoms/button/GoogButton";
import { Picture } from "../../atoms/picture/Picture";

export const CommentBoard: React.FC = () => {
    return (
        <Box
            p="20px"
            border="1px"
            bg="blue.100"
            borderRadius="lg"
            boxShadow="dark-lg"
        >
            <Stack ml="15px" spacing={3}>
                <Flex>
                    <Text m="auto">投稿者: ゲスト</Text>
                    <Spacer></Spacer>
                    <MenuIconButton></MenuIconButton>
                </Flex>
                <Divider></Divider>
                <Text>テストです。</Text>
                <Picture url={"https://bit.ly/dan-abramov"}></Picture>
                <Flex>
                    <GoogButton></GoogButton>
                    <Spacer></Spacer>
                    <Text m="auto">更新日時: 2022/1/1 15:00</Text>
                </Flex>
            </Stack>
        </Box>
    );
};
