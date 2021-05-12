import React, { useState } from "react";
import { Alert, AlertProps } from "../components/Alert";
import { Login } from "../components/Login";
import { Register } from "../components/Register";
import { LogoBar } from "../components/LogoBar";
import { UserLogin, UserRegister } from "../shared/types";
import { RouteComponentProps } from "react-router-dom";
import { API } from "../shared/api";
import { useBearer } from "../shared/hooks/useBearer";
import { useAlert } from "../shared/hooks/useAlert";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const [page, setPage] = React.useState("login");
	const [bearer, setBearer] = useBearer("success");
	const [alert, setAlert] = useAlert({
		message: "",
		type: "success",
	});

	const login = (user: UserLogin) => {
		API.Login(user).then((res) => {
			if (res.success) {
				setBearer(res.bearer);
				history.push("/");
				return;
			}

			const alert: AlertProps = {
				message:
					"The provided email and password were incorrect. User may not exist.",
				type: "failure",
			};

			setAlert(alert);
		});
	};

	const register = (user: UserRegister) => {
		setPage("login");
		setAlert({
			message: "Successfully created user account.",
			type: "success",
		});
	};

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				{alert.message !== "" && <Alert className="mt-20" {...alert} />}
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
