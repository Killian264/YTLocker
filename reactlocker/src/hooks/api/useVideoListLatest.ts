import axios from "axios";
import { useQuery } from "react-query";

export const useVideoListLatest = (): [boolean, number[]] => {
	const { isSuccess, data } = useQuery(["latest"], () => fetchLatest());

	if (!isSuccess || data === undefined) {
		return [true, []];
	}

	if (data === null) {
		return [true, []];
	}

	return [false, data];
};

const fetchLatest = async (): Promise<number[]> => {
	return axios.get("/playlist/videos/latest").then((response) => {
		let { Data: data } = response.data;

		return data;
	});
};
