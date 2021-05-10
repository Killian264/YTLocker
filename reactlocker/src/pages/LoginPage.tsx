import React, { useState } from "react";
import { Alert } from "../components/Alert";
import { Login } from "../components/Login";
import { Register } from "../components/Register";
import { LogoBar } from "../components/LogoBar";
import { UserLogin, UserRegister } from "../shared/types";
import { RouteComponentProps } from "react-router-dom";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const [page, setPage] = React.useState("login");

	let [alert, setAlert] = useState("");

	const login = (user: UserLogin) => {
		history.push("/");
	};

	const register = (user: UserRegister) => {
		setPage("login");
		setAlert("Successfully created user account.");
		setTimeout(() => {
			setAlert("");
		}, 5000);
	};

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				{alert !== "" && <Alert type="success">{alert}</Alert>}
				<div className="m-auto">
					{page === "login" && (
						<Login
							onSubmit={login}
							onClickRegister={() => {
								setPage("register");
							}}
						/>
					)}
					{page === "register" && (
						<Register
							onSubmit={register}
							onClickLogin={() => {
								setPage("login");
							}}
						/>
					)}
				</div>
			</div>
		</>
	);
};

LoginPage.propTypes = {};
