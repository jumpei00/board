import React, { ReactNode } from "react";
import { Button } from "@chakra-ui/react";

type PrimaryButtonProps = {
    children: ReactNode;
    colorScheme: string;
};

export const PrimaryButton: React.FC<PrimaryButtonProps> = (props) => {
    const { children, colorScheme } = props;
    return (
        <Button colorScheme={colorScheme} size="md" _hover={{ opacity: 0.8 }}>
            {children}
        </Button>
    );
};
