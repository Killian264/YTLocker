import axios from "axios";
import { useContext } from "react";
import { AlertContext } from "../AlertContext";
import { useBearer } from "../useBearer";

export const useUserSession = (): (() => void) => {
	const { pushAlert } = useContext(AlertContext);
	const [, setBearer] = useBearer("");

	return () => {
		axios
			.post("/user/session/create", {}, {})
			.then((response) => {
				setBearer(response.data.Data.Bearer);
			})
			.catch(() => {
				pushAlert({
					message: "The provided email and password were incorrect. User may not exist.",
					type: "failure",
				});
			});
	};
};
