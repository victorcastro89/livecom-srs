import React, { useEffect, useRef } from 'react';
// Import OvenPlayer if it's installed via npm
import OvenPlayer from 'ovenplayer';
import { current } from '@reduxjs/toolkit';

export function VideoPlayer({ streamUrl }: { streamUrl: string }) {
  const playerRef = useRef<HTMLDivElement>(null);

  // Function to load a script dynamically
  const loadScript = (src: string, onLoad: any) => {
    // Check if the script is already included
    if (!document.querySelector(`script[src="${src}"]`)) {
      const script = document.createElement('script');
      script.src = src;
      script.onload = onLoad;
      script.onerror = () => console.error(`Error loading script: ${src}`);
      document.head.appendChild(script);
    } else {
      // If the script is already included, call the onLoad callback
      onLoad();
    }
  };

  useEffect(() => {
    loadScript(
      'https://cdn.jsdelivr.net/npm/hls.js@latest/dist/hls.min.js',
      () => {
        if (playerRef.current?.id) {
          // Initialize OvenPlayer
          OvenPlayer.create(playerRef.current?.id, {
            sources: [{ type: 'hls', file: streamUrl }],
            // other configuration options...
          });
        }
      }
    );
    return () => {
      // Cleanup

      if (playerRef.current) {
        playerRef.current.innerHTML = '';
      }
    };
  }, [streamUrl]);

  return <div id="player_id" ref={playerRef} />;
}
