import React from "react";
import { Box, Text, Textarea, Flex, Spacer } from "@chakra-ui/react";
import { ImageButton } from "../../atoms/button/ImageButton";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";

export const CommentPostform: React.FC = () => {
    return (
        <Box w="70%" m="auto">
            <Text mb="10px">コメント</Text>
            <Textarea h="120px" mb="10px" variant="flushed" placeholder="コメントしよう"></Textarea>
            <Flex>
                <ImageButton>画像</ImageButton>
                <Spacer></Spacer>
                <PrimaryButton colorScheme="teal">投稿</PrimaryButton>
            </Flex>
        </Box>
    );
};
