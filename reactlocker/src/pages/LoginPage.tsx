import React from "react";
import { Login } from "../components/Login";
import { Register } from "../components/Register";
import { LogoBar } from "../components/LogoBar";
import { RouteComponentProps } from "react-router-dom";
import { useUserLogin } from "../hooks/api/useUserLogin";
import { useUserRegister, UserRegister } from "../hooks/api/useUserRegister";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const [page, setPage] = React.useState("login");

	const postLogin = useUserLogin();
	const postRegister = useUserRegister();

	const register = async (user: UserRegister) => {
		postRegister(user).then(() => {
			setPage("login");
		});
	};

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				<div className="m-auto">
					{page === "login" && (
						<Login
							onSubmit={postLogin}
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
