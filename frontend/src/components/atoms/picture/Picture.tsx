import React from "react";
import { Box, Image } from "@chakra-ui/react";

type PictureProps = {
    url: string;
};

export const Picture: React.FC<PictureProps> = (props) => {
    const { url } = props;

    return (
        <Box p="5px">
            <Image src={url} boxSize="300px" m="auto"></Image>
        </Box>
    );
};
