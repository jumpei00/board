import React, { ReactNode } from "react";
import { Button } from "@chakra-ui/react";
import { AttachmentIcon } from "@chakra-ui/icons";

type ImageButtonProps = {
    children: ReactNode;
};

export const ImageButton: React.FC<ImageButtonProps> = (props) => {
    const { children } = props;

    return (
        <Button leftIcon={<AttachmentIcon></AttachmentIcon>} colorScheme="cyan" size="md" _hover={{ opacity: 0.8 }}>
            {children}
        </Button>
    );
};
