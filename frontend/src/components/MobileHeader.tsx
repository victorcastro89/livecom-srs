import {
  AspectRatio,
  Avatar,
  Box,
  DialogTitle,
  Drawer,
  Dropdown,
  List,
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
import { NavLink } from 'react-router-dom';
import { styled, useColorScheme, useTheme } from '@mui/joy/styles';
import Button from '@mui/joy/Button';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import MenuIcon from '@mui/icons-material/Menu';
import { Divider } from '@mui/material';
import LightModeIcon from '@mui/icons-material/LightMode';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import useToggleColorScheme from '../hooks/useToggleColorScheme';
import LogoutIcon from '@mui/icons-material/Logout';
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
            VC
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
            <Avatar variant="outlined">VC</Avatar>

            <Typography level="title-sm">Your name</Typography>
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
            <ListItemButton sx={{ fontWeight: 'lg' }}>Home</ListItemButton>
            <ListItemButton>Team</ListItemButton>

            <Divider sx={{ mt: 8 }}></Divider>
            <ListItemButton onClick={toggleColorScheme}>
              {mode === 'dark' ? (
                <>
                  {' '}
                  <LightModeIcon />
                  <Typography level="title-sm">Ligth Mode</Typography>{' '}
                </>
              ) : (
                <>
                  {' '}
                  <DarkModeIcon />
                  <Typography level="title-sm">Dark Mode</Typography>{' '}
                </>
              )}
            </ListItemButton>
            <ListItemButton>
              <LogoutIcon />
              <Typography level="title-sm">Log out</Typography>{' '}
            </ListItemButton>
          </List>
        </Drawer>
      </Box>
    </Box>
  );
};
