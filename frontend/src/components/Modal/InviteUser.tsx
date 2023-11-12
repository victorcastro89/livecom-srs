import React, { useState, FormEvent } from 'react';
import Button from '@mui/joy/Button';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import Input from '@mui/joy/Input';
import Modal from '@mui/joy/Modal';
import ModalDialog from '@mui/joy/ModalDialog';
import DialogTitle from '@mui/joy/DialogTitle';
import DialogContent from '@mui/joy/DialogContent';
import Stack from '@mui/joy/Stack';
import Add from '@mui/icons-material/Add';
import { Select } from '@mui/joy';
import Option from '@mui/joy/Option';
// Define the props that you want to accept for the component
interface CustomModalDialogProps {
  onFormSubmit: (email: string, role: string) => void; // Callback for form submission
  label?: string;
}

// The component definition with destructuring of props
export const InviteUserModal: React.FC<CustomModalDialogProps> = ({
  label,
  onFormSubmit,
}) => {
  const [open, setOpen] = useState<boolean>(false);
  const [email, setName] = useState<string>('');
  const [role, setRole] = useState<string>('');

  // Function to handle form submission
  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onFormSubmit(email, role); // Call the provided callback function
    setOpen(false); // Close the modal
  };

  return (
    <>
      <Button
        variant="outlined"
        color="neutral"
        startDecorator={<Add fontSize="small" />}
        onClick={() => setOpen(true)}
      >
        {label}
      </Button>
      <Modal open={open} onClose={() => setOpen(false)}>
        <ModalDialog>
          <DialogTitle>Adicionar usuário</DialogTitle>
          <DialogContent>
            Preencha o endereço de e-mail, o enviaremos um convite
          </DialogContent>
          <form onSubmit={handleSubmit}>
            <Stack spacing={2}>
              <FormControl>
                <FormLabel>E-mail</FormLabel>
                <Input
                  type="email"
                  autoFocus
                  required
                  value={email}
                  onChange={(e) => setName(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel>Role</FormLabel>
                <Select
                  onChange={(_event, newValue) => setRole(newValue as string)}
                  placeholder="Admin"
                  variant="outlined"
                >
                  <Option value="Admin">Admin</Option>
                  <Option value="Cohost">Cohost</Option>
                </Select>
              </FormControl>
              <Button type="submit">Submit</Button>
            </Stack>
          </form>
        </ModalDialog>
      </Modal>
    </>
  );
};
