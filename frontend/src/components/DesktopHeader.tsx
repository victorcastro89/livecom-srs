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
  Switch,
  Typography,
} from '@mui/joy';
import React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import Button from '@mui/joy/Button';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import useToggleColorScheme from '../hooks/useToggleColorScheme';
import LightModeIcon from '@mui/icons-material/LightMode';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/auth/authSlice';
import { getUserInitials } from '../helpers/helpers';
import { useAppDispatch } from '../hooks';
import { logOut } from '../features/auth/authThunk';
import { DarkModeToggle } from './DarkModeToggle';
export const DesktopHeader: React.FC = () => {
  const [mode, toggleColorScheme] = useToggleColorScheme();

  const menuButtonRef = React.useRef<any>(null);
  const [menuWidth, setMenuWidth] = React.useState<number | null>(null);
  const navigate = useNavigate();
  const user = useSelector(selectUser);
  const dispatch = useAppDispatch();
  React.useEffect(() => {
    if (menuButtonRef.current) {
      setMenuWidth(menuButtonRef.current.getBoundingClientRect().width);

      const resizeObserver = new ResizeObserver((entries) => {
        for (let entry of entries) {
          setMenuWidth(menuButtonRef.current.getBoundingClientRect().width);
          setMenuWidth(entry.contentRect.width);
        }
      });

      resizeObserver.observe(menuButtonRef.current);

      // Clean up observer on component unmount
      return () => {
        resizeObserver.disconnect();
      };
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
            <Avatar>{getUserInitials(user.data)}</Avatar>
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
            <MenuItem
              onClick={() => {
                toggleColorScheme();
              }}
            >
              <DarkModeToggle />
              {/* <DarkModeIcon /> */}
              <Typography>Dark Mode</Typography>
            </MenuItem>

            <Divider />
            <MenuItem
              onClick={() => {
                dispatch(logOut());
              }}
            >
              <Typography>Log out</Typography>
            </MenuItem>
          </Menu>
        </Dropdown>
      </Box>
    </Box>
  );
};
