import React, { useState } from 'react';
import { Alert } from './Alert';
import { Login } from './Login';
import { Register } from './Register';


export const LoginPage = ({ className }) => {

	const [page, setPage] = React.useState("login")

	let [alert, setAlert] = useState("")

	const login = (user) => {

	}

	const register = (user) => {
		setPage("login")
		setAlert("Successfully created user account.")
		setTimeout(() => { setAlert(""); }, 5000);
	}

	return (
		<div className="flex h-screen">
			{alert !== "" && 
				<Alert>{alert}</Alert>
			}
			<div className="m-auto">
				{ page === "login" &&
					<Login
						onSubmit={login}
						onClickRegister={() => {setPage("register")}}
					/>
				}
				{ page === "register" &&
					<Register
						onSubmit={register}
						onClickLogin={() => {setPage("login")}}
					/>
				}
			</div>
		</div>
	);
};

LoginPage.propTypes = {};