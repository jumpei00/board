import React from "react";
import {
    Menu,
    MenuButton,
    MenuList,
    MenuItem,
    IconButton,
} from "@chakra-ui/react";
import { HamburgerIcon, EditIcon, DeleteIcon } from "@chakra-ui/icons";

export const MenuIconButton: React.FC = () => {
    return (
        <Menu>
            <MenuButton
                as={IconButton}
                icon={<HamburgerIcon></HamburgerIcon>}
                color="black"
            ></MenuButton>
            <MenuList>
                <MenuItem icon={<EditIcon></EditIcon>}>編集</MenuItem>
                <MenuItem icon={<DeleteIcon></DeleteIcon>}>削除</MenuItem>
            </MenuList>
        </Menu>
    );
};
