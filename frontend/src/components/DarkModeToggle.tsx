import { Switch } from '@mui/joy';
import React from 'react';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import useToggleColorScheme from '../hooks/useToggleColorScheme';
export const DarkModeToggle: React.FC = () => {
  const [mode, toggleColorScheme] = useToggleColorScheme();
  const initialSetChecked = mode === 'dark';
  const [checked, setChecked] = React.useState<boolean>(initialSetChecked);

  return (
    <Switch
      size="md"
      variant="soft"
      checked={checked}
      onChange={(event) => {
        event.stopPropagation();
        setChecked(event.target.checked);
        toggleColorScheme();
      }}
      slotProps={{
        input: { 'aria-label': 'Dark mode' },
        thumb: {
          children: <DarkModeIcon />,
        },
      }}
      sx={{
        '--Switch-thumbSize': '16px',
      }}
    />
  );
};
