import axios from "axios";
import { useQuery } from "react-query";
import { User } from "../../shared/types";

export const useUser = (): [boolean, User | null] => {
	const { isSuccess, data } = useQuery(["user"], UserInformation);

	if (!isSuccess || data === undefined) {
		return [true, null];
	}

	return [false, data];
};

const UserInformation = (): Promise<User> => {
	return axios.get("/user/information").then((response) => {
		let { Username, Email, CreatedAt } = response.data.Data;

		return {
			username: Username,
			email: Email,
			joined: new Date(Date.parse(CreatedAt)),
		};
	});
};
