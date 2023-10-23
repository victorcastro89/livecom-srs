import React, { useEffect } from 'react';

import { Button, Stack } from '@mui/joy';
import { Auth } from 'firebase/auth';

import { useLoaderData } from 'react-router-dom';

export const SignedIn = () => {
  const { auth } = useLoaderData() as { auth: Auth };
  useEffect(() => {
    console.log('auth', auth);
  }, []);

  return (
    <>
      <Stack>
        <h1>Authenticated</h1>
        <h2>{auth.currentUser?.email ? auth.currentUser.email : null} </h2>
        <Button onClick={() => auth.signOut()}>Logout</Button>
      </Stack>
    </>
  );
};
