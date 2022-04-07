import axios from 'axios';
import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from 'react-router-dom';
import { GoogleLogin } from 'react-google-login';

const OauthClient = () => async (googleData: { tokenId: any }) => {
  const res = await fetch("/api/oauth/auth/google", {
    method: "POST",
    body: JSON.stringify({
      token: googleData.tokenId
    }),
    headers: {
      "Content-Type": "application/json"
    }
  })
  const data = await res.json()
  // store returned user somehow
}
