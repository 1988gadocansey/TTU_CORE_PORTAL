import axios from 'axios';
import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from 'react-router-dom';
import GoogleLogin from 'react-google-login';
const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await axios.post('login', {
            email,
            password
        });

        setRedirect(true);
    }
    const responseGoogle = async (response:any) => {
      //console.log(JSON.stringify(response));
      const data = JSON.stringify(response)
      const profile = response.getBasicProfile();
      setName(profile.getName());
      setEmail(profile.getEmail());
      console.log("profile" + profile.getName())
     // await axios.get('/oauth/auth/google');
      //console.log(JSON.stringify(response));
    }


  const [gapi, setGapi] = useState();
  const [googleAuth, setGoogleAuth] = useState();
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [name, setName] = useState('');

  const [imageUrl, setImageUrl] = useState();

  const onSuccess = (googleUser: { getBasicProfile: () => any; }) => {
    setIsLoggedIn(true);
    const profile = googleUser.getBasicProfile();
    setName(profile.getName());
    setEmail(profile.getEmail());
    setImageUrl(profile.getImageUrl());
  };

  const onFailure = () => {
    setIsLoggedIn(false);
  }



    if (redirect) {
        return <Redirect to={'/'}/>;
    }

    return (
        <main className="form-signin">
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

                <input type="email" className="form-control" placeholder="Email" required
                       onChange={e => setEmail(e.target.value)}
                />

                <input type="password" className="form-control" placeholder="Password" required
                       onChange={e => setPassword(e.target.value)}
                />

                <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
              <GoogleLogin
                clientId="523793815073-33qjc07d518jpoh5hcrv4ar12aa6atdc.apps.googleusercontent.com"
                buttonText="Login"
                onSuccess={responseGoogle}
                onFailure={responseGoogle}
                cookiePolicy={'single_host_origin'}
              />,

            </form>
        </main>
    )
};

export default Login;
