import React, { useState, useEffect } from "react";
import ReactDOM from 'react-dom';
import styled from 'styled-components';

import {NavBar} from './NavBar';
import { useNavDropMenu } from './Hooks/useNavDropMenu.js';
import { NavDropMenu } from './NavDropMenu.js';

function App () {
    const ndm = useNavDropMenu();

    const doLogOut = () => {
        console.log ('log out func')
    };

    return (
        <>
            <NavBar {...ndm}/>
            <NavDropMenu {...ndm} doLogOut={doLogOut}  />

        </>
    )
}

if (document.getElementById('react_root')) {
    ReactDOM.render(<App />, document.getElementById('react_root')); 
}