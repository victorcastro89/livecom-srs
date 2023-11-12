import {
  Box,
  List,
  ListItem,
  ListItemButton,
  ListItemDecorator,
  Tooltip,
  styled,
} from '@mui/joy';
import React from 'react';
import Home from '@mui/icons-material/Home';
import SmartDisplay from '@mui/icons-material/SmartDisplay';
import makeStyles from '@mui/material/styles/makeStyles';
import { ListItemText } from '@mui/material';
import PeopleIcon from '@mui/icons-material/People';
import { useNavigate } from 'react-router-dom';
const NoHoverListItem = styled(ListItem)(({ theme }) => ({
  '&:hover': {
    backgroundColor: 'transparent !important', // Use !important to override MUI styles if necessary
    color: 'inherit !important',
  },
  '& .MuiListItemButton-root:hover': {
    backgroundColor: 'transparent !important',
    color: 'inherit !important',
  },
  '& .MuiListItemButton-root': {
    display: 'flex', // Make the ListItemButton a flex container
    justifyContent: 'center', // Center the children horizontally
    alignItems: 'center', // Center the children vertically
    width: '100%', // Take the full width to center within the parent
    // Add any other styles for ListItemButton here...
  },
  '&': {
    display: 'flex', // Make the ListItemButton a flex container
    justifyContent: 'center', // Center the children horizontally
    alignItems: 'center', // Center the children vertically
    width: '100%', // Take the full width to center within the parent
    // Add any other styles for ListItemButton here...
  },
}));

const MyListItem = styled(ListItem)(({ theme }) => ({
  '&:hover': {
    backgroundColor:
      theme.palette.mode === 'dark'
        ? theme.palette.grey[900]
        : '#be43cf !important', // Remove the hover background color
    color: 'inherit !important', // Remove the hover text color change (optional)
  },
  '& .MuiListItemButton-root:hover': {
    // Target the ListItemButton hover within CustomListItem
    backgroundColor:
      theme.palette.mode === 'dark'
        ? theme.palette.grey[900]
        : '#be43cf !important',
    color: 'inherit !important', // Remove the hover text color change (optional)
  },
  '& .MuiListItemButton-root': {
    display: 'flex', // Make the ListItemButton a flex container
    justifyContent: 'center', // Center the children horizontally
    alignItems: 'center', // Center the children vertically
    width: '100%', // Take the full width to center within the parent
    // Add any other styles for ListItemButton here...
  },
  '&': {
    display: 'flex', // Make the ListItemButton a flex container
    justifyContent: 'center', // Center the children horizontally
    alignItems: 'center', // Center the children vertically
    width: '100%', // Take the full width to center within the parent
    // Add any other styles for ListItemButton here...
  },
}));
export const DesktopLeftNav: React.FC = () => {
  const navigate = useNavigate();
  return (
    <Box
      sx={(theme) => ({
        backgroundColor:
          theme.palette.mode === 'dark'
            ? 'background.surface'
            : theme.palette.primary[200],
        height: '100vh',
        boxShadow:
          theme.palette.mode === 'dark'
            ? '10px 10px 12px -13px rgba(100,100,100,0.5)'
            : 'none',

        // borderRight: '1px solid rgba(255, 255, 255, 0.02)',
      })}
    >
      <List
        sx={{
          '--ListItem-paddingY': '25px',
        }}
      >
        <NoHoverListItem>
          {/* <ListItemButton> */}
          <Box
            component={'img'}
            src={`logowhite.png`}
            alt="logo"
            sx={{
              width: 30,
              objectFit: 'cover',
            }}
          ></Box>
          {/* </ListItemButton> */}
        </NoHoverListItem>
        <MyListItem onClick={() => navigate('/home')}>
          <Tooltip title="Home" placement="right" variant="soft">
            <ListItemButton>
              <ListItemDecorator>
                <Home sx={{ color: 'white', fontSize: 30 }} />
              </ListItemDecorator>
            </ListItemButton>
          </Tooltip>
        </MyListItem>
        <MyListItem onClick={() => navigate('/team')}>
          <Tooltip title="Team" placement="right" variant="soft">
            <ListItemButton>
              <ListItemDecorator>
                <PeopleIcon sx={{ color: 'white', fontSize: 30 }} />
              </ListItemDecorator>
            </ListItemButton>
          </Tooltip>
        </MyListItem>
      </List>
    </Box>
  );
};
