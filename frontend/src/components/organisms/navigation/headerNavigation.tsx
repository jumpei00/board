import React from "react";
import { Box, HStack, Text } from "@chakra-ui/react";

type HeaderNavigationProps = {
    isSignupText?: boolean;
    isSigninText?: boolean;
    isSignoutText?: boolean;
    username?: string;
}

export const HeaderNavigation: React.FC<HeaderNavigationProps> = (props) => {
    const {isSignupText, isSigninText, isSignoutText, username} = props

    return (
        <HStack spacing={4}>
            <Text>Home</Text>
            {isSigninText && <Text>ログイン</Text>}
            {isSignoutText && <Text>ログアウト</Text>}
            {isSignupText && <Text>ユーザー登録</Text>}
            <Box textAlign="center">
                <Text>ようこそ</Text>
                <Text>{username ? `${username}さん`: "ゲストさん"}</Text>
            </Box>
        </HStack>
    );
};
