import React from "react";
import { Text, Input, Box } from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";

export const ThreadPostForm: React.FC = () => {
    return (
        <Box w="70%" m="auto">
            <Text mb="10px">スレッドタイトル</Text>
            <Input
                type="text"
                size="lg"
                mb="10px"
                variant="flushed"
                placeholder="話題を投稿しましょう"
            ></Input>
            <PrimaryButton colorScheme="teal">スレッドを投稿</PrimaryButton>
        </Box>
    );
};
