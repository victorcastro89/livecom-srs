import { AspectRatio, Box, Sheet, Stack, Typography } from '@mui/joy';
import React from 'react';
import { NavLink } from 'react-router-dom';
import { useColorScheme } from '@mui/joy/styles';
import Button from '@mui/joy/Button';

function ModeToggle() {
  const { mode, setMode } = useColorScheme();
  return (
    <Button
      variant="outlined"
      color="neutral"
      onClick={() => setMode(mode === 'dark' ? 'light' : 'dark')}
    >
      {mode === 'dark' ? 'Turn light' : 'Turn dark'}
    </Button>
  );
}
export const Navbar: React.FC = () => (
  <Sheet
    component={'nav'}
    sx={{
      width: '100%',
      top: 0,
      zIndex: 1000, // Adjust zIndex value as per your requirement
    }}
  >
    <Stack
      direction="row"
      alignItems="center"
      justifyContent={'space-between'}
      spacing={2}
    >
      <Box sx={{ width: 8 / 10, padding: 1, paddingBottom: 0 }}>
        <Box
          component={'img'}
          src={`${process.env.PUBLIC_URL}/logolivecom.png`}
          alt="logo"
          sx={{
            width: 70,
            objectFit: 'cover',
          }}
        ></Box>
      </Box>

      <Stack
        direction="row"
        alignItems="center"
        justifyContent={'flex-start'}
        spacing={2}
        sx={{ width: 2 / 10 }}
      >
        <Typography>
          <NavLink to="/">Home</NavLink>
        </Typography>
        <Typography>
          <NavLink to="/about">About</NavLink>
        </Typography>
        <ModeToggle></ModeToggle>
      </Stack>
    </Stack>
  </Sheet>
);
