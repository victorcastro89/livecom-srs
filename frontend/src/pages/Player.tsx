import React, { useEffect, useRef } from 'react';

import { Button, Stack, Box, Typography } from '@mui/joy';
import { VideoPlayer } from '../components/VideoPlayer';

// interface CustomWindow extends Window {
//   OvenPlayer: typeof OvenPlayer;
// }

// declare let window: CustomWindow;
// const options: OvenPlayer.OvenPlayerConfig = {
//   autoStart: true,
//   autoFallback: true,
//   mute: false,
//   sources: [
//     {
//       type: 'hls',
//       file: 'https://bitdash-a.akamaihd.net/content/MI201109210084_1/m3u8s/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.m3u8',
//     },
//   ],
// };

export const Player = () => {
  return (
    <>
      <Stack m={0} p={0}>
        <Typography>Oi</Typography>
        <Box className="player-wrapper">
          <VideoPlayer streamUrl="http://localhost:8080/live/live_360.m3u8" />
        </Box>
      </Stack>
    </>
  );
};
