import axios from "axios";
import { useContext } from "react";
import { useHistory } from "react-router-dom";
import { AlertContext } from "../AlertContext";
import { useBearer } from "../useBearer";

export const useUserRefreshSession = (): ((bearer: string) => void) => {
	const { pushAlert } = useContext(AlertContext);
	const [, setBearer] = useBearer("");
	const history = useHistory();

	return (bearer: string) => {
		axios
			.post("/user/session/refresh", {}, {
				headers: {
				  'Authorization': bearer 
				}
			  })
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
};
