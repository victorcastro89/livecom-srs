import React, { useEffect } from 'react';

import { Button, Stack } from '@mui/joy';
import { Auth } from 'firebase/auth';

import { useLoaderData } from 'react-router-dom';
import SwaggerUI from 'swagger-ui-react';
import 'swagger-ui-react/swagger-ui.css';

export const ApidDocs = () => {

    return (
        <div>
          <h1>API Documentation</h1>
          <SwaggerUI url="/swagger.json" />
        </div>
      );
};
