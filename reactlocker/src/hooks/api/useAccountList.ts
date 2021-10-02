import axios from "axios";
import { useQuery } from "react-query";

export const useAccountList = (): [boolean, number[]] => {
	const { isSuccess, data } = useQuery(["accounts"], () => fetchLatest());

	if (!isSuccess || data === undefined) {
		return [true, []];
	}

	if (data === null) {
		return [true, []];
	}

	return [false, data];
};

const fetchLatest = async (): Promise<number[]> => {
	return axios.get("/account/list").then((response) => {
		let { Data: data } = response.data;

		return data;
	});
};
