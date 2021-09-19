import axios from "axios";
import { useContext } from "react";
import { AlertContext } from "../AlertContext";

export interface UserRegister {
	username: string;
	email: string;
	password: string;
	password2: string;
}

export const useUserRegister = (): ((user: UserRegister) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);

	return (user: UserRegister) => {
		return axios
			.post("/user/register", user)
			.then(() => {
				pushAlert({
					message: "Successfully created user account.",
					type: "success",
				});
				return;
			})
			.catch((error) => {
				pushAlert({
					message: `Failed to create account: '${error.response.data.Message}'.`,
					type: "failure",
				});
			});
	};
};
