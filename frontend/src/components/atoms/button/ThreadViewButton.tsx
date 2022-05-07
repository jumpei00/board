import React, { ReactNode } from "react";
import { Button } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";

type ThreadViewButtonProps = {
    children: ReactNode;
    onClick: () => void;
};

export const ThreadViewButton: React.FC<ThreadViewButtonProps> = (props) => {
    const { children, onClick } = props;

    return (
        <Button
            rightIcon={<ExternalLinkIcon></ExternalLinkIcon>}
            colorScheme="yellow"
            size="md"
            _hover={{ opacity: 0.8 }}
            onClick={onClick}
        >
            {children}
        </Button>
    );
};
