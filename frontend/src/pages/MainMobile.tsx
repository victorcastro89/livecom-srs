import React, { Fragment } from 'react';
import Counter from '../components/counter/Counter';
import { Box } from '@mui/joy';
import { Outlet } from 'react-router-dom';

import { MobileHeader } from '../components/MobileHeader';

export const MainMobile: React.FC = () => {
  return (
    <Box
      sx={{
        display: 'flex',
        height: '100vh',
        overflow: 'hidden',
        backgroundColor: 'background.surface',
        color: 'text.primary',
      }}
    >
      {/* Main content */}
      <Box id="teste" sx={{ flex: 1 }}>
        <MobileHeader />

        <Box
          sx={{
            p: '1vh',
            mt: '1vh',
            height: '99vh',
            overflowY: 'auto',
          }}
        >
          <Outlet></Outlet>
        </Box>
      </Box>
    </Box>
  );
};
