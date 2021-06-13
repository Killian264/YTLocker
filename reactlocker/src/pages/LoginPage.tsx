import React from "react";
import { Alert } from "../components/Alert";
import { Login } from "../components/Login";
import { Register } from "../components/Register";
import { LogoBar } from "../components/LogoBar";
import { UserLogin, UserRegister } from "../shared/types";
import { RouteComponentProps } from "react-router-dom";
import { useBearer } from "../shared/hooks/useBearer";
import { useAlert } from "../shared/hooks/useAlert";
import axios from "axios";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const [page, setPage] = React.useState("login");
	const [, setBearer] = useBearer("");
	const [alert, setAlert] = useAlert({
		message: "",
		type: "success",
	});

	const login = (user: UserLogin) => {
		axios
			.post("/user/login", user)
			.then((response) => {
				let { Bearer } = response.data.Data;
				setBearer(Bearer);
				history.push("/");
			})
			.catch(() => {
				setAlert({
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
				setAlert({
					message: "Successfully created user account.",
					type: "success",
				});
			})
			.catch(() => {
				setAlert({
					message: "Failed to create account",
					type: "failure",
				});
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
