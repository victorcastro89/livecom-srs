/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/naming-convention */
import React, { useEffect } from 'react';
import { AuthErrorMap } from 'firebase/auth';
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth';

import Sheet from '@mui/joy/Sheet';
import {
  AuthStateType,
  FirebaseAuthStateType,
  setFirebaseUser,
  setFirebaseUserError,
} from '../features/auth/authSlice';
import store from '../store';
import { auth } from '../App';
import { EmailAuthProvider, GoogleAuthProvider } from 'firebase/auth';
import '../css/firebase_styling.global.css';
import { Box, Typography } from '@mui/joy';
import { createOrGetUser } from '../features/auth/authThunk';
import { Roles } from '../factory/apiFactory';
import { useNavigate } from 'react-router-dom';

export const Login = () => {
  const navigate = useNavigate();
  const signInSuccessWithAuthResult = (
    res: { additionalUserInfo: any; user: any },
    _redirectUrl: string
  ) => {
    const { additionalUserInfo, user } = res;
    console.log('user', user);
    console.log('additionalUserInfo', additionalUserInfo);
    let ustate: FirebaseAuthStateType = {
      email: user?.email ? user?.email : null,
      emailVerified: user?.emailVerified ? user?.emailVerified : null,
      isAnonymous: user?.isAnonymous ? user?.isAnonymous : null,
      displayName: user?.displayName ? user?.displayName : null,
      phoneNumber: user?.phoneNumber ? user?.phoneNumber : null,
      photoURL: user?.photoURL ? user?.photoURL : null,
      uid: user?.uid ? user?.uid : null,
      refreshToken: user?.refuserhToken ? user?.refuserhToken : null,
      accessToken: user?.accessToken ? user?.accessToken : null,
    };
    // console.log('redirectUrl', window.location.href); To get path parameter
    store.dispatch(setFirebaseUser(ustate)); // Assuming you have access to your Redux store here

    store
      .dispatch(
        createOrGetUser({
          first_name: ustate.displayName || '',
          last_name: '',
          phone_number: ustate.phoneNumber || '',
          photo_url: ustate.photoURL || '',
          account_name: ustate.email?.split('@')[0] || '',
          role: Roles.OWNER,
        })
      )
      .unwrap()
      .then(() => {
        navigate('/signedIn');
      })
      //TOTO : handle error
      .catch((err) => {
        console.log('error on Dispatch createOrGetUser', err);
        store.dispatch(
          setFirebaseUserError((err as string) || 'Failed to login')
        );
      });
    return false;
  };

  // Configure FirebaseUI.
  const uiConfig = {
    callbacks: {
      signInFailure: (error: any) => {
        console.log('Firebase  signInFailure', error);
        store.dispatch(setFirebaseUserError(error.code || 'Failed to login'));
        return false;
      },
      // Avoid redirects after sign-in.
      signInSuccessWithAuthResult: signInSuccessWithAuthResult,
    },
    // Popup signin flow rather than redirect flow.
    signInFlow: 'redirect',
    // Redirect to /signedIn after sign in is successful. Alternatively you can provide a callbacks.signInSuccess function.
    signInSuccessUrl: '/signedIn',
    // We will display Google and Facebook as auth providers.
    signInOptions: [
      {
        provider: EmailAuthProvider.PROVIDER_ID,
        requireDisplayName: false,
        signInMethod: EmailAuthProvider.EMAIL_PASSWORD_SIGN_IN_METHOD,
      },
      {
        provider: GoogleAuthProvider.PROVIDER_ID,
        clientId:
          '226538007405-7dnimnlrkv6f6b3ifbn9p9in5c012cbk.apps.googleusercontent.com',
      },
    ],
    credentialHelper: 'googleyolo',
  };
  return (
    <>
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          alignContent: 'center',
          height: '100vh', // Full viewport height
          margin: 0, // Reset default margin
          backgroundColor: 'primary.400', // Set the background color to gray
          width: '100vw', // Set the width to full viewport width
        }}
      >
        <Box
          sx={{
            margin: 2,
            padding: 0.5,
            borderRadius: 12,
            boxShadow: 1,
            width: '100%',
            maxWidth: '400px',
            minHeight: '320px',
            boxSizing: 'border-box', // To include padding in the width calculation
            backgroundColor: 'background.body',
            textAlign: 'center',
          }}
        >
          <Box
            component={'img'}
            src={`/Logo.png`}
            alt="logo"
            sx={{
              width: 70,
              paddingTop: 2,
              maxWidth: '100%',
              objectFit: 'cover',
            }}
          ></Box>
          <Typography
            level="h1"
            sx={{
              textAlign: 'center', // Center the text
              p: 2,
            }}
          >
            Livee
          </Typography>
          <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />
        </Box>
      </Box>
    </>
  );
};
