import React from "react";
import { useSelector } from "react-redux";
import { Outlet } from "react-router-dom";
import { Box, Heading, Flex, Spacer } from "@chakra-ui/react";
import { HeaderNavigation } from "../navigation/headerNavigation";
import { RootState } from "../../../store/store";

export const Header: React.FC = () => {
    const userState = useSelector((state: RootState) => state.user.userState);

    return (
        <>
            <Box p="30px" bg="gray.300">
                <Flex>
                    <Heading>6ちゃんねる</Heading>
                    <Spacer></Spacer>
                    <HeaderNavigation
                        isSignupText={userState.username === ""}
                        isSigninText={userState.username === ""}
                        isSignoutText={userState.username !== ""}
                        username={userState.username}
                    ></HeaderNavigation>
                </Flex>
            </Box>
            <Outlet></Outlet>
        </>
    );
};
