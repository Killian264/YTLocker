import React from 'react';
import { Login } from './Login';
import { Register } from './Register';



export const LoginPage = ({ }) => {

	const [page, setPage] = React.useState("login")

	const login = (user) => {

	}

	const register = (user) => {

	}

	return (
		<div>
			<div>
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
		</div>
	);
};

LoginPage.propTypes = {};