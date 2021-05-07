import React from 'react';
import { Login } from './Login';
import { Register } from './Register';


export const LoginPage = ({ className }) => {

	const [page, setPage] = React.useState("login")

	const login = (user) => {

	}

	const register = (user) => {

	}

	return (
		<div className={className}>
			{ page == "login" &&
				<Login
					onSubmit={login}
					onClickRegister={() => {setPage("register")}}
				/>
			}
			{ page == "register" &&
				<Register
					onSubmit={login}
					onClickLogin={() => {setPage("login")}}
				/>
			}
		</div>
	);
};

LoginPage.propTypes = {};