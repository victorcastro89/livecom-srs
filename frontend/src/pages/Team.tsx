import React from 'react';

import {
  Avatar,
  Box,
  Button,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  IconButton,
  Modal,
  ModalDialog,
  Table,
  Tooltip,
  Typography,
  useTheme,
} from '@mui/joy';
import { useNavigate } from 'react-router-dom';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import DeleteIcon from '@mui/icons-material/Delete';
import useMediaQuery from '@mui/material/useMediaQuery';
import WarningRoundedIcon from '@mui/icons-material/WarningRounded';
import { InviteUserModal } from '../components/Modal/InviteUser';
export const TeamPage: React.FC = () => {
  function createData(
    email: string,
    role: string,
    pendingAccept: boolean,
    action: string
  ) {
    return { email, role, pendingAccept, action };
  }
  const rows = [
    createData('asd@asd.com', 'Owner', false, ''),
    createData('Ice@mix.com', 'Admin', true, ''),
    createData('Elviz@elviz.com', 'Cohost', false, ''),
    createData('Cupcake@gmail.com', 'Cohost', true, ''),
    createData('Dinda@uol.com.br', 'Admin', false, ''),
  ];
  const [openTooltips, setOpenTooltips] = React.useState(rows.map(() => false));
  // Refs to store timeout IDs for each tooltip
  const timeoutRefs = React.useRef<(number | null)[]>(rows.map(() => null));
  const handleTooltipClick = (index: number) => {
    // Open the tooltip immediately
    setOpenTooltips((prevOpenTooltips) =>
      prevOpenTooltips.map((isOpen, i) => i === index)
    );

    // Clear any existing timeout to prevent multiple timeouts from being set
    const currentTimeout = timeoutRefs.current[index];
    if (typeof currentTimeout === 'number') {
      clearTimeout(currentTimeout);
    }

    // Set a new timeout to close the tooltip
    timeoutRefs.current[index] = window.setTimeout(() => {
      setOpenTooltips((prevOpenTooltips) =>
        prevOpenTooltips.map((isOpen, i) => (i === index ? false : isOpen))
      );
      // Clear the timeout id from the refs
      timeoutRefs.current[index] = null;
    }, 4000);
  };
  // Clear all timeouts when the component unmounts
  React.useEffect(() => {
    return () => {
      timeoutRefs.current.forEach((timeoutId) => {
        if (timeoutId !== null) {
          clearTimeout(timeoutId);
        }
      });
    };
  }, []);

  const [open, setOpen] = React.useState<boolean>(false);
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  const navigate = useNavigate();

  return (
    <Box sx={{ px: isMobile ? 1 : 20 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
        <Box>
          <Typography level="h1">Team</Typography>
        </Box>
        <Box>
          <InviteUserModal
            label="Invite"
            onFormSubmit={(name, description) => {
              console.log('Project Name:', name);
              console.log('Project Description:', description);
              // You can add more logic here, like sending data to an API
            }}
          />
        </Box>
      </Box>
      <Table
        borderAxis="xBetween"
        sx={{ mt: 2 }}
        variant={'plain'}
        color={'neutral'}
      >
        <tbody>
          {rows.map((row, index) => (
            <tr key={row.email}>
              <td>
                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                  {isMobile ? (
                    <></>
                  ) : (
                    <>
                      <Avatar sx={{ marginRight: 2 }}>
                        {row.email[0].toUpperCase()}
                      </Avatar>
                    </>
                  )}
                  <Box
                    sx={{
                      display: 'flex',
                      flexDirection: 'column',
                      width: isMobile ? '100px' : '100%',
                    }}
                  >
                    <Typography noWrap level="body-sm">
                      {row.email}
                    </Typography>
                    <Typography noWrap level="body-sm">
                      {row.role}
                    </Typography>
                  </Box>
                </Box>
              </td>
              <td>
                {row.pendingAccept ? (
                  <>
                    <Box
                      sx={{
                        display: 'flex',
                        alignItems: 'center',
                      }}
                    >
                      <Typography noWrap level="body-sm">
                        Pending Invite
                      </Typography>
                      <Box
                        sx={{ display: 'flex', flexDirection: 'column', mx: 2 }}
                      >
                        <Tooltip
                          title="O link foi copiado e também foi enviado um e-mail para o usuário."
                          arrow
                          open={openTooltips[index]}
                          placement="bottom"
                          size="sm"
                          disableFocusListener
                          disableHoverListener
                          disableTouchListener
                          sx={{ maxWidth: '150px' }}
                        >
                          <IconButton onClick={() => handleTooltipClick(index)}>
                            <ContentCopyIcon />
                          </IconButton>
                        </Tooltip>
                      </Box>
                    </Box>
                  </>
                ) : (
                  ''
                )}
              </td>
              <td>
                <Box
                  sx={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'flex-end',
                  }}
                  onClick={() => setOpen(true)}
                >
                  <DeleteIcon />
                </Box>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
      <Modal open={open} onClose={() => setOpen(false)}>
        <ModalDialog variant="outlined" role="alertdialog">
          <DialogTitle>
            <WarningRoundedIcon />
            Confirmation
          </DialogTitle>
          <Divider />
          <DialogContent>Are you sure you want to remove user?</DialogContent>
          <DialogActions>
            <Button
              variant="solid"
              color="danger"
              onClick={() => setOpen(false)}
            >
              Remove
            </Button>
            <Button
              variant="plain"
              color="neutral"
              onClick={() => setOpen(false)}
            >
              Cancel
            </Button>
          </DialogActions>
        </ModalDialog>
      </Modal>
    </Box>
  );
};
