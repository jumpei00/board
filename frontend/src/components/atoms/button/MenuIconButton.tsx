import React from "react";
import { Menu, MenuButton, MenuList, MenuItem, IconButton } from "@chakra-ui/react";
import { HamburgerIcon, EditIcon, DeleteIcon } from "@chakra-ui/icons";

type MenuIconButtonProps = {
    onOpen: () => void;
    setIsEdit: React.Dispatch<React.SetStateAction<boolean>>;
};

export const MenuIconButton: React.FC<MenuIconButtonProps> = (props) => {
    const editMenuItemOnClick = () => {
        props.setIsEdit(true);
        props.onOpen();
    };

    const deleteMenuItemOnClick = () => {
        props.setIsEdit(false);
        props.onOpen();
    };

    return (
        <Menu>
            <MenuButton as={IconButton} icon={<HamburgerIcon></HamburgerIcon>} color="black"></MenuButton>
            <MenuList>
                <MenuItem icon={<EditIcon></EditIcon>} onClick={editMenuItemOnClick}>
                    編集
                </MenuItem>
                <MenuItem icon={<DeleteIcon></DeleteIcon>} onClick={deleteMenuItemOnClick}>
                    削除
                </MenuItem>
            </MenuList>
        </Menu>
    );
};
