import React from "react";
import { Outlet } from "react-router-dom";
import { Box, Heading, Flex, Spacer } from "@chakra-ui/react";
import { HeaderNavigation } from "../../organisms/navigation/HeaderNavigation";

export const Header: React.FC = () => {
    return (
        <>
            <Box p="30px" bg="gray.300">
                <Flex>
                    <Heading>6ちゃんねる</Heading>
                    <Spacer></Spacer>
                    <HeaderNavigation isSigninText isSignoutText isSignupText></HeaderNavigation>
                </Flex>
            </Box>
            <Outlet></Outlet>
        </>
    );
};
