import axios from "axios";
import { useQuery } from "react-query";
import { Account } from "../../shared/types";

export const useAccountList = (): [boolean, Account[]] => {
	const { isSuccess, data } = useQuery(["accounts"], () => fetchLatest());

	if (!isSuccess || data === undefined) {
		return [true, []];
	}

	if (data === null) {
		return [true, []];
	}

	return [false, data];
};

const fetchLatest = async (): Promise<Account[]> => {
	return axios.get("/account/list").then((response) => {
		let { Data: data } = response.data;

		let accounts = data.map((account: any): Account => {
			return {
				id: account.ID,
				username: account.Username,
				email: account.Email,
				picture: account.Picture,
				permissionLevel: account.PermissionLevel,
			};
		});

		return accounts;
	});
};
