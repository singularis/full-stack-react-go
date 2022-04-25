import React, { useState, useEffect } from "react";
import ReactDOM from 'react-dom';
import styled from 'styled-components';

import {NavBar} from './NavBar';
import { useNavDropMenu } from './Hooks/useNavDropMenu.js';
import { NavDropMenu } from './NavDropMenu.js';
import { useModal } from './Hooks/useModal';
import { Modal } from './Modal.js';

function App () {
    const modal = useModal();
    const ndm = useNavDropMenu();

    const doLogOut = () => {
        console.log ('log out func')
    };

    return (
        <>
            <NavBar {...ndm} {...modal}/>
            <NavDropMenu {...ndm} doLogOut={doLogOut}  />
            { modal.modalShowing && <Modal {...modal}/>
            }
        </>
    )
}

if (document.getElementById('react_root')) {
    ReactDOM.render(<App />, document.getElementById('react_root')); 
}