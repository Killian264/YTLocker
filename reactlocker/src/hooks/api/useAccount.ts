import axios from "axios";
import { useQuery } from "react-query";
import { Account } from "../../shared/types";

export const useAccount = (id: number, enabled = true): [boolean, Account | null] => {
	const { isSuccess, data } = useQuery(["account", id], () => fetchAccount(id), {
		staleTime: Infinity,
		enabled: enabled,
	});

	if (!isSuccess || data === undefined) {
		return [true, null];
	}

	return [false, data];
};

const fetchAccount = async (id: number): Promise<Account> => {
	return axios.get(`/account/${id}`).then((response) => {
		let { Data: data } = response.data;

		let parsed: Account = {
			id: data.ID,
			username: data.Username,
			email: data.Email,
			picture: data.Picture,
			permissionLevel: data.PermissionLevel,
		};

		return parsed;
	});
};
