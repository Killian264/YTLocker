import React, { useContext } from "react";
import { Login } from "../components/Login";
import { Register } from "../components/Register";
import { LogoBar } from "../components/LogoBar";
import { UserLogin, UserRegister } from "../shared/types";
import { RouteComponentProps } from "react-router-dom";
import { useBearer } from "../hooks/useBearer";
import axios from "axios";
import { AlertContext } from "../hooks/AlertContext";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const [page, setPage] = React.useState("login");
	const [, setBearer] = useBearer("");
	const { pushAlert } = useContext(AlertContext);

	const login = (user: UserLogin) => {
		axios
			.post("/user/login", user)
			.then((response) => {
				let { Bearer } = response.data.Data;
				setBearer(Bearer);
				history.push("/");
			})
			.catch(() => {
				pushAlert({
					message: "The provided email and password were incorrect. User may not exist.",
					type: "failure",
				});
			});
	};

	const register = async (user: UserRegister) => {
		axios
			.post("/user/register", user)
			.then(() => {
				setPage("login");
				pushAlert({
					message: "Successfully created user account.",
					type: "success",
				});
			})
			.catch(() => {
				pushAlert({
					message: "Failed to create account",
					type: "failure",
				});
			});
	};

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
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
