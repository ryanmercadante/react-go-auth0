import React from 'react'
import { useAuth0 } from '../react-auth0-spa'

const Home = () => {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0()
  return (
    <>
      <div className='container'>
        <div className='jumbotron text-center mt-5'>
          <h1>We R VR</h1>
          <p>Provide valuable feedback to VR experience developers.</p>
          {!isAuthenticated && (
            <button
              className='btn btn-primary btn-lg btn-login btn-block'
              onClick={() => loginWithRedirect({})}
            >
              Sign in
            </button>
          )}
        </div>
      </div>
    </>
  )
}

export default Home
