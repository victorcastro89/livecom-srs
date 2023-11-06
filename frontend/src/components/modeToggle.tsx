import React from 'react';
import { useColorScheme, Button } from '@mui/joy';

export const ModeToggle: React.FC = () => {
  const { mode, setMode } = useColorScheme();
  return (
    <Button onClick={() => setMode(mode === 'dark' ? 'light' : 'dark')}>
      {mode === 'dark' ? 'Turn light' : 'Turn dark'}
    </Button>
  );
};
