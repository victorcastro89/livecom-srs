import {
  AspectRatio,
  Avatar,
  Box,
  DialogTitle,
  Drawer,
  Dropdown,
  List,
  ListItem,
  ListItemButton,
  Menu,
  MenuButton,
  MenuItem,
  ModalClose,
  Sheet,
  Stack,
  Typography,
} from '@mui/joy';
import React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import { styled, useColorScheme, useTheme } from '@mui/joy/styles';
import Button from '@mui/joy/Button';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import MenuIcon from '@mui/icons-material/Menu';
import { Divider } from '@mui/material';
import LightModeIcon from '@mui/icons-material/LightMode';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import useToggleColorScheme from '../hooks/useToggleColorScheme';
import LogoutIcon from '@mui/icons-material/Logout';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/auth/authSlice';
import { getUserInitials } from '../helpers/helpers';
import { useAppDispatch } from '../hooks';
import { logOut } from '../features/auth/authThunk';
import { DarkModeToggle } from './DarkModeToggle';
function ModeToggle() {
  const { mode, setMode } = useColorScheme();
  return (
    <Button onClick={() => setMode(mode === 'dark' ? 'light' : 'dark')}>
      {mode === 'dark' ? 'Turn light' : 'Turn dark'}
    </Button>
  );
}

// Create a styled button with custom focus styles
const CustomFocusButton = styled(Button)(({ theme }) => ({
  // // Default styles
  // backgroundColor:
  //   theme.palette.mode === 'dark' ? theme.palette.primary[700] : 'transparent',
  // color: theme.palette.mode === 'dark' ? 'white' : 'black',
  // // Override hover styles
  // '&:hover': {
  //   backgroundColor:
  //     theme.palette.mode === 'dark'
  //       ? theme.palette.primary[600]
  //       : 'transparent',
  // },
  // // Override focus styles
  // '&:focus': {
  //   outline: 'none',
  //   backgroundColor:
  //     theme.palette.mode === 'dark'
  //       ? theme.palette.primary[600]
  //       : 'transparent',
  // },
}));
export const MobileHeader: React.FC = () => {
  const [open, setOpen] = React.useState(false);
  const [mode, toggleColorScheme] = useToggleColorScheme();
  const theme = useTheme();
  const user = useSelector(selectUser);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
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
        <CustomFocusButton variant="plain" onClick={() => setOpen(true)}>
          <Avatar sx={{ mr: 1 }} variant="outlined">
            {getUserInitials(user.data)}
          </Avatar>
          <MenuIcon
            htmlColor={theme.palette.mode === 'dark' ? 'white' : 'black'}
          ></MenuIcon>
        </CustomFocusButton>
        <Drawer
          open={open}
          onClose={() => setOpen(false)}
          sx={{
            width: '100vw', // Set the width to 100% of the viewport width
            '& .MuiDrawer-content': {
              // Target the inner paper element
              width: '100vw', // Set the width to 100% of the viewport width
              maxWidth: '100vw', // Ensure it doesn't exceed the viewport width
            },
          }}
        >
          <ModalClose />
          <Stack
            direction="column"
            spacing={1}
            justifyContent="center"
            alignItems={'center'}
            mt={8}
          >
            <Avatar variant="outlined">{getUserInitials(user.data)}</Avatar>

            <Typography level="title-sm">{user.data.Email}</Typography>
          </Stack>
          <Divider sx={{ mt: 8 }}></Divider>
          <List
            size="lg"
            component="nav"
            sx={{
              flex: 'none',
              fontSize: 'xl',
              '& > div': { justifyContent: 'center' },
            }}
          >
            <ListItemButton
              onClick={() => navigate('/home')}
              sx={{ fontWeight: 'lg' }}
            >
              Home
            </ListItemButton>
            <ListItemButton onClick={() => navigate('/team')}>
              Team
            </ListItemButton>
            <Divider sx={{ mt: 8 }}></Divider>
            <ListItem>
              <DarkModeToggle />
              <Typography level="title-sm">Dark Mode</Typography>{' '}
            </ListItem>
            <ListItemButton onClick={() => dispatch(logOut())}>
              <LogoutIcon />
              <Typography level="title-sm">Log out</Typography>{' '}
            </ListItemButton>
          </List>
        </Drawer>
      </Box>
    </Box>
  );
};
