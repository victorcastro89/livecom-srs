import React, { Fragment } from 'react';
import Counter from '../components/counter/Counter';
import { Box } from '@mui/joy';
import { Outlet } from 'react-router-dom';
import { DesktopLeftNav } from '../components/DesktopLeftNav';

import { DesktopHeader } from '../components/DesktopHeader';

export const MainDesktop: React.FC = () => {
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
      {/* Left Nav */}
      <Box>
        <DesktopLeftNav />
      </Box>

      {/* Rigth content */}
      <Box id="teste" sx={{ flex: 1 }}>
        {/* Header */}
        <DesktopHeader />
        {/* Main Content */}
        <Box
          sx={{
            p: '2vh',
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
