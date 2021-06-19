import axios from "axios";
import { useContext } from "react";
import { useHistory } from "react-router-dom";
import { Channel } from "../../shared/types";
import { AlertContext } from "../AlertContext";
import { useBearer } from "../useBearer";

export const useLogin = (): ((email: string, password: string) => void) => {
	const { pushAlert } = useContext(AlertContext);
	const [, setBearer] = useBearer("");
	const history = useHistory();

	return (email: string, password: string) => {
		axios
			.post("/user/login", { email, password })
			.then((response) => {
				let { Bearer } = response.data.Data;
				setBearer(Bearer);
				history.push("/");
				return null;
			})
			.catch(() => {
				pushAlert({
					message: "The provided email and password were incorrect. User may not exist.",
					type: "failure",
				});
			});
	};
};
