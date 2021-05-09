import React, { useState } from 'react';
import { Alert } from '../components/Alert';
import { Login } from '../components/Login';
import { Register } from '../components/Register';
import { LogoBar } from '../components/LogoBar';


export const LoginPage = ({ className, history }) => {

	const [page, setPage] = React.useState("login")

	let [alert, setAlert] = useState("")

	const login = (user) => {
		history.push("/")
	}

	const register = (user) => {
		setPage("login")
		setAlert("Successfully created user account.")
		setTimeout(() => { setAlert(""); }, 5000);
	}

	return (
		<>
			<LogoBar className="absolute top-1 left-5" ></LogoBar>
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
		</>
	);
};

LoginPage.propTypes = {};