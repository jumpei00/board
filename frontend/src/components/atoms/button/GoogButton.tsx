import React from "react";
import { Button } from "@chakra-ui/react";
import { FcLike } from "react-icons/fc";

export const GoogButton: React.FC = () => {
    return (
        <Button
            leftIcon={<FcLike></FcLike>}
            colorScheme="yellow"
            _hover={{ opacity: 0.8 }}
        >
            10
        </Button>
    );
};
