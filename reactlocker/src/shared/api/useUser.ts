import { useQuery } from "react-query";
import { useHistory } from "react-router-dom";
import { DROPLET_BASE } from "../env";
import { useBearer } from "../hooks/useBearer";
import { User } from "../types";
import { ApiResponse2 } from "./api";

export interface UserInformationResponse extends ApiResponse2 {
	user: User | null;
}

export const useUser = (): [boolean, User | null] => {
	const [bearer] = useBearer("");
	let history = useHistory();

	const { isLoading, isError, data } = useQuery(["user"], () => UserInformation(bearer));

	if (isLoading || isError || data === undefined || data.user == null) {
		return [true, null];
	}

	if (data.status === 401) {
		history.push("/login");
	}

	return [false, data.user];
};

const UserInformation = async (bearer: string) => {
	const url = DROPLET_BASE + "/user/information";

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	let user: UserInformationResponse = {
		status: res.Status,
		message: res.Message,
		user: null,
	};

	if (user.status == 200) {
		user.user = {
			username: res.Data.Username,
			email: res.Data.Email,
			joined: new Date(Date.parse(res.Data.CreatedAt)),
		};
	}

	return user;
};
