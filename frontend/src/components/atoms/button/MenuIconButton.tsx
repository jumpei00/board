import React from "react";
import {
    Menu,
    MenuButton,
    MenuList,
    MenuItem,
    IconButton,
} from "@chakra-ui/react";
import { HamburgerIcon, EditIcon, DeleteIcon } from "@chakra-ui/icons";

type MenuIconButtonProps = {
    onOpen: () => void;
};

export const MenuIconButton: React.FC<MenuIconButtonProps> = (props) => {
    const { onOpen } = props;

    return (
        <Menu>
            <MenuButton
                as={IconButton}
                icon={<HamburgerIcon></HamburgerIcon>}
                color="black"
            ></MenuButton>
            <MenuList>
                <MenuItem icon={<EditIcon></EditIcon>} onClick={onOpen}>編集</MenuItem>
                <MenuItem icon={<DeleteIcon></DeleteIcon>}>削除</MenuItem>
            </MenuList>
        </Menu>
    );
};
