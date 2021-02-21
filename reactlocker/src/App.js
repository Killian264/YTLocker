import logo from "./logo.svg";
import "./App.css";
import { createMuiTheme, withStyles, makeStyles, ThemeProvider } from '@material-ui/core/styles';
// import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import { green, purple, orange } from '@material-ui/core/colors';
import React, { useState } from 'react';

import Switch from '@material-ui/core/Switch';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Avatar from '@material-ui/core/Avatar';
import { Container } from '@material-ui/core';
import Box from '@material-ui/core/Box';

import { Card } from '@material-ui/core';
import { CardHeader } from '@material-ui/core';

import PropTypes from 'prop-types';

const darkTheme = createMuiTheme({
	palette: {
		type: "dark",
	}
});

const lightTheme = createMuiTheme({
	palette: {
		type: "light",
	}
});



const themes = {
	"light": lightTheme,
	"dark": darkTheme,
};


function App() {

	const [checked, setChecked] = useState(false);

	const [theme, setTheme] = useState(darkTheme);

	const toggleChecked = () => {
		setChecked((prev) => !prev);
	}

	return (
		<ThemeProvider theme={theme}>
			<header className="App-header">
				<Container fixed>
					<Box display="flex" justifyContent="space-between" pt={2}>
						<Box display="flex">
							<Box>
								<Button>Dashboard</Button>
							</Box>
							<Box ml={3}>
								<Button>Playlists</Button>
							</Box>
							<Box ml={3}>
								<Button m={3}>All Videos</Button>
							</Box>
						</Box>
						<Box display="flex">
							<FormGroup>
								<FormControlLabel
									control={
										<Switch checked={checked} onChange={ () => toggleChecked() } />
									}
									label="this is cool2"
								/>
							</FormGroup>
							<Avatar>K</Avatar>
						</Box>
					</Box>
				</Container>
			</header>
			<main className="app">
			</main>
		</ThemeProvider>
	);
}

export default App;
