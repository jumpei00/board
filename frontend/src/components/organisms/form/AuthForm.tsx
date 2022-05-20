import React from "react";
import { Box, Stack, FormControl, FormLabel, Input, Heading } from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";

type AuthFormProps = {
    formName: string;
    buttonName: string;
};

export const AuthForm: React.FC<AuthFormProps> = (props) => {
    const { formName, buttonName } = props;

    return (
        <Box w="50%" m="200px auto" bg="gray.300" boxShadow="dark-lg" p="50px">
            <FormControl>
                <Stack spacing={8}>
                    <Heading>{formName}</Heading>
                    <Box>
                        <FormLabel>ユーザー名</FormLabel>
                        <Input type="text" variant="flushed" placeholder="username"></Input>
                    </Box>
                    <Box>
                        <FormLabel>パスワード</FormLabel>
                        <Input type="password" variant="flushed" placeholder="password"></Input>
                    </Box>
                    <Box textAlign="right">
                        <PrimaryButton colorScheme="teal" onClick={() => undefined}>
                            {buttonName}
                        </PrimaryButton>
                    </Box>
                </Stack>
            </FormControl>
        </Box>
    );
};
