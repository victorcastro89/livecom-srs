import {
  AspectRatio,
  Avatar,
  Box,
  Divider,
  Dropdown,
  Menu,
  MenuButton,
  MenuItem,
  Sheet,
  Stack,
  Typography,
} from '@mui/joy';
import React from 'react';
import { NavLink } from 'react-router-dom';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import Button from '@mui/joy/Button';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import useToggleColorScheme from '../hooks/useToggleColorScheme';
import LightModeIcon from '@mui/icons-material/LightMode';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/auth/authSlice';
export const DesktopHeader: React.FC = () => {
  const [mode, toggleColorScheme] = useToggleColorScheme();
  const menuButtonRef = React.useRef<any>(null);
  const [menuWidth, setMenuWidth] = React.useState<number | null>(null);
  const user = useSelector(selectUser);
  // Effect to update the width of the menu
  React.useEffect(() => {
    if (menuButtonRef.current) {
      setMenuWidth(menuButtonRef.current.offsetWidth);
    }
  }, []);

  return (
    <Box
      sx={{
        width: '100%',
        display: 'flex',
        justifyContent: 'flex-end',
        padding: 1,
        boxSizing: 'border-box',
      }}
    >
      <Box sx={{ flexGrow: 0 }}>
        <Dropdown>
          <MenuButton
            variant="plain"
            sx={{ display: 'flex', alignItems: 'center', maxWidth: '250px' }}
            ref={menuButtonRef} // Set the ref to measure width
          >
            <Avatar>
              {user.data.DisplayName
                ? user.data.DisplayName.substring(0, 1)
                : user.data.Email.split('@')[0].substring(0, 2)}
            </Avatar>
            <Typography
              noWrap
              sx={{
                overflow: 'hidden',
                textOverflow: 'ellipsis',
                whiteSpace: 'nowrap',
              }}
              level="title-md"
            >
              {user.data.Email}
            </Typography>
            <ArrowDropDownIcon></ArrowDropDownIcon>
          </MenuButton>
          <Menu
            variant="outlined"
            sx={{
              minWidth: menuWidth, // Set the width of the menu
              width: menuWidth, // Set the width of the menu
            }}
          >
            <MenuItem onClick={toggleColorScheme}>
              {mode === 'dark' ? (
                <>
                  {' '}
                  <LightModeIcon />
                  <Typography>Ligth Mode</Typography>{' '}
                </>
              ) : (
                <>
                  {' '}
                  <DarkModeIcon />
                  <Typography>Dark Mode</Typography>
                </>
              )}
            </MenuItem>
            <Divider />
            <MenuItem>
              <Typography>Log out</Typography>
            </MenuItem>
          </Menu>
        </Dropdown>
      </Box>
    </Box>
  );
};
