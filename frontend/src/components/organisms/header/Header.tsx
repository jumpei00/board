import React from "react";
import { useSelector } from "react-redux";
import { Outlet } from "react-router-dom";
import { Box, Heading, Flex, Spacer } from "@chakra-ui/react";
import { HeaderNavigation } from "../navigation/headerNavigation";
import { RootState } from "../../../store/store";

export const Header: React.FC = () => {
    const user = useSelector((state: RootState) => state.user);

    return (
        <>
            <Box p="30px" bg="gray.300">
                <Flex>
                    <Heading>6ちゃんねる</Heading>
                    <Spacer></Spacer>
                    <HeaderNavigation
                        isSignupText={user.username === ""}
                        isSigninText={user.username === ""}
                        isSignoutText={user.username !== ""}
                        username={user.username}
                    ></HeaderNavigation>
                </Flex>
            </Box>
            <Outlet></Outlet>
        </>
    );
};
