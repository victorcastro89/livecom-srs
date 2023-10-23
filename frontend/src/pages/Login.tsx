/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/naming-convention */
import React, { useEffect } from 'react';

import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth';

import Sheet from '@mui/joy/Sheet';
import { AuthStateType, setUser } from '../features/auth/authSlice';
import store from '../store';
import { auth } from '../App';
import { EmailAuthProvider, GoogleAuthProvider } from 'firebase/auth';

// Configure FirebaseUI.
const uiConfig = {
  callbacks: {
    // Avoid redirects after sign-in.
    signInSuccessWithAuthResult: (
      res: { additionalUserInfo: any; user: any },
      _redirectUrl: string
    ) => {
      const { additionalUserInfo, user } = res;
      console.log('additionalUserInfo', additionalUserInfo);
      let ustate: AuthStateType = {
        email: user?.email ? user?.email : null,
        emailVerified: user?.emailVerified ? user?.emailVerified : null,
        isAnonymous: user?.isAnonymous ? user?.isAnonymous : null,
        displayName: user?.displayName ? user?.displayName : null,
        phoneNumber: user?.phoneNumber ? user?.phoneNumber : null,
        photoURL: user?.photoURL ? user?.photoURL : null,
        uid: user?.uid ? user?.uid : null,
        refreshToken: user?.refuserhToken ? user?.refuserhToken : null,
      };
      // console.log('redirectUrl', window.location.href); To get path parameter
      store.dispatch(setUser(ustate)); // Assuming you have access to your Redux store here

      return true; // Avoids redirects after sign-in
    },
  },
  // Popup signin flow rather than redirect flow.
  signInFlow: 'redirect',
  // Redirect to /signedIn after sign in is successful. Alternatively you can provide a callbacks.signInSuccess function.
  signInSuccessUrl: '/signedIn',
  // We will display Google and Facebook as auth providers.
  signInOptions: [
    {
      provider:EmailAuthProvider.PROVIDER_ID,
      requireDisplayName: false,
      signInMethod:
        EmailAuthProvider.EMAIL_PASSWORD_SIGN_IN_METHOD,
    },
   GoogleAuthProvider.PROVIDER_ID,
  ],
};

export const Login = () => {
  useEffect(() => {}, []);

  return (
    <div>
      <h1>Auth</h1>
      <Sheet sx={{ mx: 'auto' }} variant="soft">
        Welcome!
      </Sheet>
      <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />
    </div>
  );
};
