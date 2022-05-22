import React from "react";
import { Link } from "react-router-dom";
import { Box, HStack, Text } from "@chakra-ui/react";
import { useDispatch } from "react-redux";
import { signout } from "../../../state/user/redux";

type HeaderNavigationProps = {
    isSignupText: boolean;
    isSigninText: boolean;
    isSignoutText: boolean;
    username: string;
};

export const HeaderNavigation: React.FC<HeaderNavigationProps> = (props) => {
    const dispatch = useDispatch();

    return (
        <HStack spacing={4}>
            <Text>
                <Link to="/">Home</Link>
            </Text>
            {props.isSigninText && (
                <Text>
                    <Link to="/signin">ログイン</Link>
                </Text>
            )}
            {props.isSignoutText && (
                <Text>
                    <Link to="/" onClick={() => dispatch(signout())}>
                        ログアウト
                    </Link>
                </Text>
            )}
            {props.isSignupText && (
                <Text>
                    <Link to="/signup">ユーザー登録</Link>
                </Text>
            )}
            <Box textAlign="center">
                <Text>ようこそ</Text>
                <Text>{props.username !== "" ? `${props.username}さん` : "ゲストさん"}</Text>
            </Box>
        </HStack>
    );
};
