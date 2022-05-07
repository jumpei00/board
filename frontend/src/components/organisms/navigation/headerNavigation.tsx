import React from "react";
import { Link } from "react-router-dom";
import { Box, HStack, Text } from "@chakra-ui/react";

type HeaderNavigationProps = {
    isSignupText?: boolean;
    isSigninText?: boolean;
    isSignoutText?: boolean;
    username?: string;
};

export const HeaderNavigation: React.FC<HeaderNavigationProps> = (props) => {
    const { isSignupText, isSigninText, isSignoutText, username } = props;

    return (
        <HStack spacing={4}>
            <Text><Link to="/">Home</Link></Text>
            {isSigninText && <Text><Link to="/signin">ログイン</Link></Text>}
            {isSignoutText && <Text><Link to="/">ログアウト</Link></Text>}
            {isSignupText && <Text><Link to="/signup">ユーザー登録</Link></Text>}
            <Box textAlign="center">
                <Text>ようこそ</Text>
                <Text>{username ? `${username}さん` : "ゲストさん"}</Text>
            </Box>
        </HStack>
    );
};
